package testing

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

// GoldenTester provides golden file testing (screenshot comparison)
type GoldenTester struct {
	t              *testing.T
	goldenDir      string
	updateGoldens  bool
	threshold      float64 // Pixel difference threshold (0.0 - 1.0)
	diffDir        string  // Directory to save diff images
}

// GoldenTestOptions configures golden testing
type GoldenTestOptions struct {
	GoldenDir     string  // Directory containing golden files
	UpdateGoldens bool    // If true, update golden files instead of comparing
	Threshold     float64 // Pixel difference threshold (default: 0.01)
	DiffDir       string  // Directory to save diff images (default: goldenDir/diffs)
}

// NewGoldenTester creates a new golden tester
func NewGoldenTester(t *testing.T, opts *GoldenTestOptions) *GoldenTester {
	if opts == nil {
		opts = &GoldenTestOptions{}
	}

	if opts.GoldenDir == "" {
		opts.GoldenDir = "testdata/golden"
	}

	if opts.Threshold == 0 {
		opts.Threshold = 0.01 // 1% difference allowed
	}

	if opts.DiffDir == "" {
		opts.DiffDir = filepath.Join(opts.GoldenDir, "diffs")
	}

	// Create directories if they don't exist
	os.MkdirAll(opts.GoldenDir, 0755)
	os.MkdirAll(opts.DiffDir, 0755)

	// Check for UPDATE_GOLDENS environment variable
	updateGoldens := opts.UpdateGoldens
	if os.Getenv("UPDATE_GOLDENS") == "true" || os.Getenv("UPDATE_GOLDENS") == "1" {
		updateGoldens = true
	}

	return &GoldenTester{
		t:             t,
		goldenDir:     opts.GoldenDir,
		updateGoldens: updateGoldens,
		threshold:     opts.Threshold,
		diffDir:       opts.DiffDir,
	}
}

// MatchesGolden compares an image against a golden file
func (gt *GoldenTester) MatchesGolden(name string, img image.Image) bool {
	goldenPath := filepath.Join(gt.goldenDir, name+".png")

	if gt.updateGoldens {
		// Update the golden file
		if err := gt.saveImage(goldenPath, img); err != nil {
			gt.t.Errorf("Failed to update golden file %s: %v", name, err)
			return false
		}
		gt.t.Logf("Updated golden file: %s", goldenPath)
		return true
	}

	// Load golden file
	goldenImg, err := gt.loadImage(goldenPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Golden file doesn't exist, create it
			if err := gt.saveImage(goldenPath, img); err != nil {
				gt.t.Errorf("Failed to create golden file %s: %v", name, err)
				return false
			}
			gt.t.Logf("Created golden file: %s", goldenPath)
			return true
		}
		gt.t.Errorf("Failed to load golden file %s: %v", name, err)
		return false
	}

	// Compare images
	diff, err := gt.compareImages(goldenImg, img)
	if err != nil {
		gt.t.Errorf("Failed to compare images for %s: %v", name, err)
		return false
	}

	if diff > gt.threshold {
		// Images differ, save diff image
		diffPath := filepath.Join(gt.diffDir, name+"_diff.png")
		diffImg := gt.createDiffImage(goldenImg, img)
		if err := gt.saveImage(diffPath, diffImg); err != nil {
			gt.t.Errorf("Failed to save diff image %s: %v", name, err)
		}

		gt.t.Errorf("Golden test failed for %s: difference %.4f exceeds threshold %.4f\nDiff saved to: %s",
			name, diff, gt.threshold, diffPath)
		return false
	}

	return true
}

// MatchesGoldenBytes compares image bytes against a golden file
func (gt *GoldenTester) MatchesGoldenBytes(name string, imgBytes []byte) bool {
	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		gt.t.Errorf("Failed to decode image bytes for %s: %v", name, err)
		return false
	}
	return gt.MatchesGolden(name, img)
}

// MatchesGoldenHash compares a hash against a golden hash
func (gt *GoldenTester) MatchesGoldenHash(name string, data []byte) bool {
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])

	goldenPath := filepath.Join(gt.goldenDir, name+".hash")

	if gt.updateGoldens {
		// Update the golden hash
		if err := os.WriteFile(goldenPath, []byte(hashStr), 0644); err != nil {
			gt.t.Errorf("Failed to update golden hash %s: %v", name, err)
			return false
		}
		gt.t.Logf("Updated golden hash: %s", goldenPath)
		return true
	}

	// Load golden hash
	goldenHash, err := os.ReadFile(goldenPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Golden hash doesn't exist, create it
			if err := os.WriteFile(goldenPath, []byte(hashStr), 0644); err != nil {
				gt.t.Errorf("Failed to create golden hash %s: %v", name, err)
				return false
			}
			gt.t.Logf("Created golden hash: %s", goldenPath)
			return true
		}
		gt.t.Errorf("Failed to load golden hash %s: %v", name, err)
		return false
	}

	goldenHashStr := string(goldenHash)
	if hashStr != goldenHashStr {
		gt.t.Errorf("Golden hash test failed for %s:\nExpected: %s\nGot:      %s",
			name, goldenHashStr, hashStr)
		return false
	}

	return true
}

// saveImage saves an image to a file
func (gt *GoldenTester) saveImage(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

// loadImage loads an image from a file
func (gt *GoldenTester) loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}

// compareImages compares two images and returns a difference score (0.0 - 1.0)
func (gt *GoldenTester) compareImages(img1, img2 image.Image) (float64, error) {
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	// Images must have the same size
	if bounds1.Dx() != bounds2.Dx() || bounds1.Dy() != bounds2.Dy() {
		return 1.0, fmt.Errorf("image sizes differ: %v vs %v", bounds1.Size(), bounds2.Size())
	}

	var totalDiff uint64
	var maxDiff uint64 = uint64(bounds1.Dx() * bounds1.Dy() * 255 * 4) // RGBA

	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()

			// Convert from 16-bit to 8-bit
			r1, g1, b1, a1 = r1>>8, g1>>8, b1>>8, a1>>8
			r2, g2, b2, a2 = r2>>8, g2>>8, b2>>8, a2>>8

			// Calculate absolute differences
			totalDiff += abs(int64(r1) - int64(r2))
			totalDiff += abs(int64(g1) - int64(g2))
			totalDiff += abs(int64(b1) - int64(b2))
			totalDiff += abs(int64(a1) - int64(a2))
		}
	}

	return float64(totalDiff) / float64(maxDiff), nil
}

// createDiffImage creates a diff image highlighting differences
func (gt *GoldenTester) createDiffImage(img1, img2 image.Image) image.Image {
	bounds := img1.Bounds()
	diffImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, a1 := img1.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()

			// Convert from 16-bit to 8-bit
			r1, g1, b1, a1 = r1>>8, g1>>8, b1>>8, a1>>8
			r2, g2, b2, a2 = r2>>8, g2>>8, b2>>8, a2>>8

			// Calculate differences
			dr := abs(int64(r1) - int64(r2))
			dg := abs(int64(g1) - int64(g2))
			db := abs(int64(b1) - int64(b2))
			da := abs(int64(a1) - int64(a2))

			// Highlight differences in red
			var r, g, b, a uint8
			if dr+dg+db+da > 10 { // Threshold for visible difference
				r = 255
				g = 0
				b = 0
				a = 255
			} else {
				// Show original pixel (grayscale)
				gray := uint8((r1 + g1 + b1) / 3)
				r, g, b, a = gray, gray, gray, uint8(a1)
			}

			diffImg.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	return diffImg
}

func abs(x int64) uint64 {
	if x < 0 {
		return uint64(-x)
	}
	return uint64(x)
}
