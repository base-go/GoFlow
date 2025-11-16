#!/bin/bash
# Build script for input test

set -e

echo "Building GoFlow Input Test..."

# Build the binary
go build -o input-test main.go

# Create app bundle structure
echo "Creating app bundle..."
mkdir -p InputTest.app/Contents/MacOS
mkdir -p InputTest.app/Contents/Resources

# Copy binary
cp input-test InputTest.app/Contents/MacOS/InputTest
chmod +x InputTest.app/Contents/MacOS/InputTest

# Create Info.plist
cat > InputTest.app/Contents/Info.plist << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>InputTest</string>
    <key>CFBundleIdentifier</key>
    <string>com.goflow.inputtest</string>
    <key>CFBundleName</key>
    <string>GoFlow Input Test</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>1.0</string>
    <key>CFBundleVersion</key>
    <string>1</string>
    <key>LSMinimumSystemVersion</key>
    <string>10.12</string>
    <key>NSHighResolutionCapable</key>
    <true/>
</dict>
</plist>
EOF

echo "✅ Build complete!"
echo ""
echo "To run the test:"
echo "  open InputTest.app"
echo ""
echo "Or double-click InputTest.app in Finder"
