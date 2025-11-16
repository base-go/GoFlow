#!/bin/bash
# Build script for macOS rendering demo

set -e

echo "Building GoFlow Rendering Demo..."

# Build the binary
go build -o rendering-demo main.go

# Create app bundle structure
echo "Creating app bundle..."
mkdir -p RenderingDemo.app/Contents/MacOS
mkdir -p RenderingDemo.app/Contents/Resources

# Copy binary
cp rendering-demo RenderingDemo.app/Contents/MacOS/RenderingDemo
chmod +x RenderingDemo.app/Contents/MacOS/RenderingDemo

# Info.plist already exists

echo "✅ Build complete!"
echo ""
echo "To run the app:"
echo "  open RenderingDemo.app"
echo ""
echo "Or double-click RenderingDemo.app in Finder"
