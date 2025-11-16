# Installing GoFlow CLI

## Option 1: Via `go install` (Recommended)

The simplest and recommended way to install the GoFlow CLI:

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

### Setup PATH

Make sure `$GOPATH/bin` is in your PATH:

```bash
# Add to ~/.zshrc or ~/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin

# Apply changes
source ~/.zshrc  # or source ~/.bashrc
```

### Verify Installation

```bash
goflow version
# Output: GoFlow CLI v0.1.0

goflow help
# Shows help information
```

## Option 2: Build from Source

Clone the repository and build:

```bash
# Clone
git clone https://github.com/base-go/GoFlow.git
cd GoFlow/cmd/goflow

# Build with optimization
go build -ldflags="-s -w" -o goflow

# Install globally (optional)
sudo mv goflow /usr/local/bin/
```

## Quick Test

Create and run a test project:

```bash
# Create project
goflow create hello --template=minimal --platforms=macos

# Run it
cd hello/macos
go run main.go
```

## Binary Size

The optimized GoFlow CLI binary is approximately **2.6MB** with embedded templates.

## Updating

To update to the latest version:

```bash
go install github.com/base-go/GoFlow/cmd/goflow@latest
```

## Uninstalling

To remove the CLI:

```bash
rm $(which goflow)
```

## Troubleshooting

### Command not found

If you get "command not found" after installation:

1. Check if `$GOPATH/bin` is in your PATH:
   ```bash
   echo $PATH | grep "$(go env GOPATH)/bin"
   ```

2. If not, add it:
   ```bash
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
   source ~/.zshrc
   ```

3. Verify Go is installed and GOPATH is set:
   ```bash
   go env GOPATH
   ```

### Permission denied

On macOS/Linux, if you get permission errors when running `goflow`:

```bash
chmod +x $(which goflow)
```

## See Also

- [CLI README](./README.md) - Full CLI documentation
- [CLI Reference](../../docs/CLI.md) - Command reference
- [Main README](../../README.md) - GoFlow framework overview
