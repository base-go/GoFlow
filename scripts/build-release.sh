#!/bin/bash
set -e

VERSION=${1:-"0.2.0"}
BUILD_DIR="build/releases/v${VERSION}"

echo "🔨 Building GoFlow v${VERSION}..."
mkdir -p "${BUILD_DIR}"

# Build for all platforms
PLATFORMS=(
  "darwin/amd64"
  "darwin/arm64"
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
)

for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | cut -d'/' -f1)
  ARCH=$(echo $PLATFORM | cut -d'/' -f2)
  OUTPUT="${BUILD_DIR}/goflow-${OS}-${ARCH}"

  # Add .exe extension for Windows
  if [[ "$OS" == "windows" ]]; then
    OUTPUT="${OUTPUT}.exe"
  fi

  echo "Building $OS/$ARCH..."
  GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w" \
    -o "$OUTPUT" ./cmd/goflow

  # Make binaries executable (except Windows)
  if [[ "$OS" != "windows" ]]; then
    chmod +x "$OUTPUT"
  fi
done

echo ""
echo "✓ Build complete: ${BUILD_DIR}"
echo ""
ls -lh "${BUILD_DIR}"

# Calculate checksums
echo ""
echo "📝 Generating checksums..."
cd "${BUILD_DIR}"
shasum -a 256 * > SHA256SUMS
cat SHA256SUMS
