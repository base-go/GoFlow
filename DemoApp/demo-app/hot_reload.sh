#!/bin/bash

# Simple hot reload script for GoFlow demo
echo "🚀 Starting GoFlow Hot Reload Demo"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "⚡ Hot reload enabled"
echo "📁 Watching: lib/main_darwin.go"
echo "🔨 Build command: go run -tags darwin lib/main_darwin.go"
echo ""
echo "💡 Commands:"
echo "   Ctrl+C - Stop watching"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Kill any existing process
pkill -f "go run -tags darwin lib/main_darwin.go" 2>/dev/null

# Start the app initially
echo "🔄 Initial build..."
go run -tags darwin lib/main_darwin.go &
APP_PID=$!
echo "✅ App started (PID: $APP_PID)"
echo ""

# Watch for file changes
LAST_MODIFIED=$(stat -f %m lib/main_darwin.go 2>/dev/null || echo 0)

while true; do
    sleep 1

    # Check if app is still running
    if ! kill -0 $APP_PID 2>/dev/null; then
        echo "👋 App exited"
        break
    fi

    # Check if file was modified
    CURRENT_MODIFIED=$(stat -f %m lib/main_darwin.go 2>/dev/null || echo 0)

    if [ "$CURRENT_MODIFIED" != "$LAST_MODIFIED" ] && [ "$CURRENT_MODIFIED" != "0" ]; then
        echo "🔥 File change detected - Hot reloading..."
        RELOAD_START=$(date +%s%3N)

        # Kill current app
        kill $APP_PID 2>/dev/null
        wait $APP_PID 2>/dev/null

        # Start new instance
        go run -tags darwin lib/main_darwin.go &
        APP_PID=$!

        RELOAD_END=$(date +%s%3N)
        RELOAD_TIME=$((RELOAD_END - RELOAD_START))

        echo "🔄 Hot reload completed in ${RELOAD_TIME}ms"
        echo "✅ App restarted (PID: $APP_PID)"
        echo ""

        LAST_MODIFIED=$CURRENT_MODIFIED
    fi
done

echo "🛑 Hot reload stopped"