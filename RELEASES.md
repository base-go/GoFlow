# GoFlow Release Process

This document describes how to create and manage releases for the GoFlow project.

## Version Numbering

GoFlow follows [Semantic Versioning](https://semver.org/):

- **MAJOR.MINOR.PATCH** (e.g., 1.2.3)
  - **MAJOR**: Breaking changes, incompatible API changes
  - **MINOR**: New features, backwards-compatible
  - **PATCH**: Bug fixes, backwards-compatible

### Current Version

The current CLI version is defined in `cmd/goflow/main.go`:

```go
const cliVersion = "0.2.0"
```

## Creating a Release

### 1. Prepare the Release

Before creating a release, ensure:

- All tests pass: `go test ./...`
- Code is properly formatted: `goflow format --check`
- No linting issues: `goflow analyze`
- Documentation is up to date
- CHANGELOG.md is updated with release notes

### 2. Update Version Number

Edit `cmd/goflow/main.go` and update the version constant:

```go
const cliVersion = "0.3.0"  // Update this
```

Commit the version change:

```bash
git add cmd/goflow/main.go
git commit -m "Bump version to 0.3.0"
git push origin dev
```

### 3. Create a Git Tag

Tags mark specific points in Git history as releases:

```bash
# Create an annotated tag
git tag -a v0.3.0 -m "Release v0.3.0: Description of changes"

# Push the tag to remote
git push origin v0.3.0
```

**Note**: Always prefix version tags with `v` (e.g., `v0.3.0`, not `0.3.0`)

### 4. Create a GitHub Release

#### Via GitHub Web Interface

1. Go to https://github.com/base-go/GoFlow/releases
2. Click "Draft a new release"
3. Select the tag you just created (v0.3.0)
4. Set the release title: "GoFlow v0.3.0"
5. Add release notes (see template below)
6. Attach release binaries (see Building Release Binaries section)
7. Click "Publish release"

#### Via GitHub CLI (gh)

```bash
# Create release with notes from file
gh release create v0.3.0 \
  --title "GoFlow v0.3.0" \
  --notes-file RELEASE_NOTES.md \
  ./build/goflow-macos-arm64 \
  ./build/goflow-macos-amd64 \
  ./build/goflow-linux-amd64 \
  ./build/goflow-windows-amd64.exe
```

### Release Notes Template

```markdown
## GoFlow v0.3.0

### What's New

- Feature 1: Description
- Feature 2: Description
- Enhancement: Description

### Bug Fixes

- Fix 1: Description
- Fix 2: Description

### Breaking Changes

- Breaking change 1: Migration guide
- Breaking change 2: Migration guide

### Dependencies

- Updated dependency X to v1.2.3
- Added dependency Y

### Contributors

Thanks to all contributors who made this release possible!

### Installation

\`\`\`bash
# macOS (ARM64)
curl -L https://github.com/base-go/GoFlow/releases/download/v0.3.0/goflow-macos-arm64 -o goflow
chmod +x goflow

# Linux (AMD64)
curl -L https://github.com/base-go/GoFlow/releases/download/v0.3.0/goflow-linux-amd64 -o goflow
chmod +x goflow
\`\`\`

**Full Changelog**: https://github.com/base-go/GoFlow/compare/v0.2.0...v0.3.0
```

## Building Release Binaries

### Manual Build

Build for multiple platforms using Go cross-compilation:

```bash
# Create build directory
mkdir -p build/releases/v0.3.0

# macOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" \
  -o build/releases/v0.3.0/goflow-macos-arm64 \
  ./cmd/goflow

# macOS AMD64 (Intel)
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" \
  -o build/releases/v0.3.0/goflow-macos-amd64 \
  ./cmd/goflow

# Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" \
  -o build/releases/v0.3.0/goflow-linux-amd64 \
  ./cmd/goflow

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" \
  -o build/releases/v0.3.0/goflow-windows-amd64.exe \
  ./cmd/goflow
```

**Build flags explained:**
- `-ldflags="-s -w"`: Strip debug symbols (smaller binaries)
- `-o`: Output file path

### Using GoFlow CLI

```bash
# Build using the GoFlow CLI (once implemented)
goflow build --release --bundle --output ./build/releases/v0.3.0
```

### Automated Release Build Script

Create `scripts/build-release.sh`:

```bash
#!/bin/bash
set -e

VERSION=${1:-"0.3.0"}
BUILD_DIR="build/releases/v${VERSION}"

echo "Building GoFlow v${VERSION}..."
mkdir -p "${BUILD_DIR}"

# Build for all platforms
for OS in darwin linux windows; do
  for ARCH in amd64 arm64; do
    OUTPUT="${BUILD_DIR}/goflow-${OS}-${ARCH}"

    # Skip invalid combinations
    if [[ "$OS" == "windows" && "$ARCH" == "arm64" ]]; then
      continue
    fi

    # Add .exe extension for Windows
    if [[ "$OS" == "windows" ]]; then
      OUTPUT="${OUTPUT}.exe"
    fi

    echo "Building $OS/$ARCH..."
    GOOS=$OS GOARCH=$ARCH go build -ldflags="-s -w" \
      -o "$OUTPUT" ./cmd/goflow
  done
done

echo "✓ Build complete: ${BUILD_DIR}"
ls -lh "${BUILD_DIR}"
```

Usage:

```bash
chmod +x scripts/build-release.sh
./scripts/build-release.sh 0.3.0
```

## Release Checklist

Use this checklist before every release:

- [ ] All tests pass (`go test ./...`)
- [ ] Code formatted (`goflow format --check`)
- [ ] No lint issues (`goflow analyze`)
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
- [ ] Version bumped in `cmd/goflow/main.go`
- [ ] Version committed and pushed
- [ ] Git tag created and pushed
- [ ] Release binaries built for all platforms
- [ ] GitHub release created with binaries
- [ ] Release notes published
- [ ] Announcement posted (if applicable)

## Release Cadence

GoFlow follows this release schedule:

- **Major releases**: As needed for breaking changes
- **Minor releases**: Monthly or when significant features are ready
- **Patch releases**: As needed for critical bug fixes

## Pre-releases and Beta Versions

For pre-release versions, use suffixes:

- **Alpha**: `v0.3.0-alpha.1`
- **Beta**: `v0.3.0-beta.1`
- **Release Candidate**: `v0.3.0-rc.1`

Example:

```bash
git tag -a v0.3.0-beta.1 -m "Beta release for v0.3.0"
git push origin v0.3.0-beta.1

gh release create v0.3.0-beta.1 \
  --title "GoFlow v0.3.0-beta.1" \
  --prerelease \
  --notes "Beta release for testing. Not recommended for production."
```

## Hotfix Releases

For urgent bug fixes in production:

1. Create a hotfix branch from the release tag:
   ```bash
   git checkout -b hotfix/v0.2.1 v0.2.0
   ```

2. Fix the bug and commit:
   ```bash
   git commit -m "Fix critical bug in ..."
   ```

3. Update version to patch level (0.2.0 → 0.2.1)

4. Create tag and release:
   ```bash
   git tag -a v0.2.1 -m "Hotfix: Critical bug fix"
   git push origin v0.2.1
   ```

5. Merge back to dev and main:
   ```bash
   git checkout dev
   git merge hotfix/v0.2.1
   git push origin dev
   ```

## Rolling Back a Release

If a release has critical issues:

1. Delete the GitHub release (keep the tag)
2. Yank the release in package managers (if applicable)
3. Create a hotfix release immediately
4. Document the issue in release notes

```bash
# Delete release (GitHub CLI)
gh release delete v0.3.0

# Delete tag (if needed)
git tag -d v0.3.0
git push origin :refs/tags/v0.3.0
```

## Changelog Guidelines

Keep CHANGELOG.md updated with each change:

```markdown
## [Unreleased]

### Added
- New feature X

### Changed
- Modified behavior of Y

### Fixed
- Bug in Z

### Deprecated
- Feature A (will be removed in v1.0.0)

### Removed
- Feature B

### Security
- Security fix for CVE-XXXX

## [0.2.0] - 2025-01-15

### Added
- CLI migration to Mamba framework
- Phase 2 commands: run, test, doctor, build
...
```

## Versioning Best Practices

1. **Breaking Changes**: Always increment MAJOR version
2. **New Features**: Increment MINOR version
3. **Bug Fixes**: Increment PATCH version
4. **Pre-1.0.0**: Use 0.x.x for initial development (breaking changes allowed in MINOR)
5. **Documentation**: Version docs alongside code
6. **Dependencies**: List all dependency changes in release notes

## Automation

Consider setting up GitHub Actions for automated releases:

`.github/workflows/release.yml`:

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.25'

      - name: Build binaries
        run: ./scripts/build-release.sh ${GITHUB_REF#refs/tags/v}

      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: build/releases/**/*
          generate_release_notes: true
```

## Questions?

For questions about the release process, please:

- Open an issue: https://github.com/base-go/GoFlow/issues
- Check documentation: https://goflow.dev/docs
- Ask in discussions: https://github.com/base-go/GoFlow/discussions
