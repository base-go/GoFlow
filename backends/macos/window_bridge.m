#import <Cocoa/Cocoa.h>
#include "window_bridge.h"
#include "cgo_bridge.h"
#include <stdlib.h>

// Initialize the Cocoa application
void initApp(void) {
    @autoreleasepool {
        [NSApplication sharedApplication];
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
    }
}

// Run the Cocoa event loop (blocking)
void runApp(void) {
    @autoreleasepool {
        [NSApp run];
    }
}

// Custom view that handles drawing
@interface GoFlowView : NSView {
    void* graphicsContext;
    DrawCallback drawCallback;
    void* drawUserData;
}
@property (nonatomic, assign) void* graphicsContext;
@property (nonatomic, assign) DrawCallback drawCallback;
@property (nonatomic, assign) void* drawUserData;
@end

@implementation GoFlowView

- (id)initWithFrame:(NSRect)frame {
    self = [super initWithFrame:frame];
    if (self) {
        self.graphicsContext = NULL;
        self.drawCallback = NULL;
        self.drawUserData = NULL;
    }
    return self;
}

- (void)dealloc {
    if (self.graphicsContext) {
        destroyGraphicsContext(self.graphicsContext);
        self.graphicsContext = NULL;
    }
}

- (BOOL)isFlipped {
    return YES; // Use top-left origin
}

- (BOOL)acceptsFirstResponder {
    return YES; // Accept keyboard and mouse events
}

- (void)mouseDown:(NSEvent *)event {
    NSPoint location = [self convertPoint:[event locationInWindow] fromView:nil];
    NSWindow *window = [self window];
    GoFlowWindowDelegate *delegate = (GoFlowWindowDelegate*)[window delegate];

    if (delegate.mouseCallback && delegate.mouseUserData) {
        delegate.mouseCallback((__bridge WindowHandle)window, 0, 1, location.x, location.y, delegate.mouseUserData);
    }
}

- (void)mouseUp:(NSEvent *)event {
    NSPoint location = [self convertPoint:[event locationInWindow] fromView:nil];
    NSWindow *window = [self window];
    GoFlowWindowDelegate *delegate = (GoFlowWindowDelegate*)[window delegate];

    if (delegate.mouseCallback && delegate.mouseUserData) {
        delegate.mouseCallback((__bridge WindowHandle)window, 0, 0, location.x, location.y, delegate.mouseUserData);
    }
}

- (void)mouseMoved:(NSEvent *)event {
    NSPoint location = [self convertPoint:[event locationInWindow] fromView:nil];
    NSWindow *window = [self window];
    GoFlowWindowDelegate *delegate = (GoFlowWindowDelegate*)[window delegate];

    if (delegate.mouseCallback && delegate.mouseUserData) {
        delegate.mouseCallback((__bridge WindowHandle)window, -1, 2, location.x, location.y, delegate.mouseUserData);
    }
}

- (void)mouseDragged:(NSEvent *)event {
    [self mouseMoved:event]; // Treat as mouse move while dragging
}

- (void)keyDown:(NSEvent *)event {
    NSWindow *window = [self window];
    GoFlowWindowDelegate *delegate = (GoFlowWindowDelegate*)[window delegate];

    if (delegate.keyCallback && delegate.keyUserData) {
        delegate.keyCallback((__bridge WindowHandle)window, [event keyCode], 1, delegate.keyUserData);
    }
}

- (void)keyUp:(NSEvent *)event {
    NSWindow *window = [self window];
    GoFlowWindowDelegate *delegate = (GoFlowWindowDelegate*)[window delegate];

    if (delegate.keyCallback && delegate.keyUserData) {
        delegate.keyCallback((__bridge WindowHandle)window, [event keyCode], 0, delegate.keyUserData);
    }
}

- (void)drawRect:(NSRect)dirtyRect {
    [super drawRect:dirtyRect];

    // Get current graphics context
    NSGraphicsContext *nsContext = [NSGraphicsContext currentContext];
    CGContextRef cgContext = [nsContext CGContext];

    // Fill with white background
    [[NSColor whiteColor] setFill];
    NSRectFill(dirtyRect);

    // Call the draw callback if set
    if (self.drawCallback && self.drawUserData) {
        // Get window handle
        NSWindow *window = [self window];
        self.drawCallback((__bridge WindowHandle)window, self.drawUserData);
    }

    // If we have a graphics context, copy its contents to the window
    if (self.graphicsContext) {
        // Get pixel data from our graphics context
        int width = (int)self.bounds.size.width;
        int height = (int)self.bounds.size.height;
        int bufferSize = width * height * 4;
        unsigned char *buffer = malloc(bufferSize);

        getPixelData(self.graphicsContext, buffer, bufferSize);

        // Create CGImage from pixel data
        CGColorSpaceRef colorSpace = CGColorSpaceCreateDeviceRGB();
        CGContextRef bitmapContext = CGBitmapContextCreate(
            buffer,
            width,
            height,
            8,
            width * 4,
            colorSpace,
            kCGImageAlphaPremultipliedLast
        );

        if (bitmapContext) {
            CGImageRef image = CGBitmapContextCreateImage(bitmapContext);
            if (image) {
                // Draw the image
                CGContextDrawImage(cgContext, self.bounds, image);
                CGImageRelease(image);
            }
            CGContextRelease(bitmapContext);
        }

        CGColorSpaceRelease(colorSpace);
        free(buffer);
    }
}

@end

// Custom window delegate to handle events
@interface GoFlowWindowDelegate : NSObject <NSWindowDelegate> {
    ResizeCallback resizeCallback;
    void* resizeUserData;
    MouseCallback mouseCallback;
    void* mouseUserData;
    KeyCallback keyCallback;
    void* keyUserData;
}
@property (nonatomic, assign) ResizeCallback resizeCallback;
@property (nonatomic, assign) void* resizeUserData;
@property (nonatomic, assign) MouseCallback mouseCallback;
@property (nonatomic, assign) void* mouseUserData;
@property (nonatomic, assign) KeyCallback keyCallback;
@property (nonatomic, assign) void* keyUserData;
@end

@implementation GoFlowWindowDelegate

- (void)windowDidResize:(NSNotification *)notification {
    NSWindow *window = [notification object];
    NSRect frame = [[window contentView] bounds];

    if (self.resizeCallback && self.resizeUserData) {
        self.resizeCallback(
            (__bridge WindowHandle)window,
            (int)frame.size.width,
            (int)frame.size.height,
            self.resizeUserData
        );
    }

    // Recreate graphics context with new size
    GoFlowView *view = (GoFlowView*)[window contentView];
    if (view.graphicsContext) {
        destroyGraphicsContext(view.graphicsContext);
    }
    view.graphicsContext = createGraphicsContext(
        (int)frame.size.width,
        (int)frame.size.height
    );

    [view setNeedsDisplay:YES];
}

- (BOOL)windowShouldClose:(NSWindow *)sender {
    return YES;
}

@end

// Window management implementation
WindowHandle createWindow(int width, int height, const char* title) {
    @autoreleasepool {
        // Create window
        NSRect frame = NSMakeRect(0, 0, width, height);
        NSWindow *window = [[NSWindow alloc]
            initWithContentRect:frame
            styleMask:(NSWindowStyleMaskTitled |
                       NSWindowStyleMaskClosable |
                       NSWindowStyleMaskMiniaturizable |
                       NSWindowStyleMaskResizable)
            backing:NSBackingStoreBuffered
            defer:NO];

        if (!window) {
            return NULL;
        }

        // Set window properties
        NSString *titleString = title ? [NSString stringWithUTF8String:title] : @"GoFlow";
        [window setTitle:titleString];
        [window center];
        [window setAcceptsMouseMovedEvents:YES];

        // Create custom view
        GoFlowView *view = [[GoFlowView alloc] initWithFrame:frame];
        [window setContentView:view];
        [window makeFirstResponder:view]; // Make view the first responder for events

        // Create graphics context for the view
        view.graphicsContext = createGraphicsContext(width, height);

        // Create and set delegate
        GoFlowWindowDelegate *delegate = [[GoFlowWindowDelegate alloc] init];
        delegate.mouseCallback = NULL;
        delegate.mouseUserData = NULL;
        delegate.keyCallback = NULL;
        delegate.keyUserData = NULL;
        [window setDelegate:delegate];

        return (__bridge_retained WindowHandle)window;
    }
}

void destroyWindow(WindowHandle window) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge_transfer NSWindow*)window;
        GoFlowView *view = (GoFlowView*)[nsWindow contentView];

        if (view.graphicsContext) {
            destroyGraphicsContext(view.graphicsContext);
            view.graphicsContext = NULL;
        }

        [nsWindow close];
    }
}

void showWindow(WindowHandle window) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        [nsWindow makeKeyAndOrderFront:nil];
        [NSApp activateIgnoringOtherApps:YES];
        // Force a display update immediately
        [nsWindow display];
    }
}

void hideWindow(WindowHandle window) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        [nsWindow orderOut:nil];
    }
}

int windowShouldClose(WindowHandle window) {
    if (!window) return 1;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        return ![nsWindow isVisible];
    }
}

void setWindowShouldClose(WindowHandle window, int shouldClose) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        if (shouldClose) {
            [nsWindow close];
        }
    }
}

void pollEvents(WindowHandle window) {
    @autoreleasepool {
        // Ensure app is activated (important for first frame)
        if (![NSApp isActive]) {
            [NSApp activateIgnoringOtherApps:YES];
        }

        // Process all pending events
        NSEvent *event;
        NSDate *timeout = [NSDate dateWithTimeIntervalSinceNow:0.001]; // 1ms timeout

        while ((event = [NSApp nextEventMatchingMask:NSEventMaskAny
                                           untilDate:timeout
                                              inMode:NSDefaultRunLoopMode
                                             dequeue:YES])) {
            [NSApp sendEvent:event];
            timeout = [NSDate distantPast]; // After first event, process remaining immediately
        }

        // Run the run loop briefly to allow window server updates
        [[NSRunLoop currentRunLoop] runMode:NSDefaultRunLoopMode
                                 beforeDate:[NSDate dateWithTimeIntervalSinceNow:0.001]];

        [NSApp updateWindows];
    }
}

void waitEvents(WindowHandle window) {
    @autoreleasepool {
        NSEvent *event = [NSApp nextEventMatchingMask:NSEventMaskAny
                                           untilDate:[NSDate distantFuture]
                                              inMode:NSDefaultRunLoopMode
                                             dequeue:YES];
        if (event) {
            [NSApp sendEvent:event];
            [NSApp updateWindows];
        }
    }
}

void* getWindowGraphicsContext(WindowHandle window) {
    if (!window) return NULL;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        GoFlowView *view = (GoFlowView*)[nsWindow contentView];
        return view.graphicsContext;
    }
}

void presentWindow(WindowHandle window) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        GoFlowView *view = (GoFlowView*)[nsWindow contentView];
        [view setNeedsDisplay:YES];
        [view displayIfNeeded];
    }
}

void setWindowNeedsDisplay(WindowHandle window) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        GoFlowView *view = (GoFlowView*)[nsWindow contentView];
        [view setNeedsDisplay:YES];
    }
}

void getWindowSize(WindowHandle window, int* width, int* height) {
    if (!window || !width || !height) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        NSRect frame = [[nsWindow contentView] bounds];
        *width = (int)frame.size.width;
        *height = (int)frame.size.height;
    }
}

void setWindowSize(WindowHandle window, int width, int height) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        NSRect frame = [nsWindow frame];
        frame.size.width = width;
        frame.size.height = height;
        [nsWindow setFrame:frame display:YES];
    }
}

void getWindowPosition(WindowHandle window, int* x, int* y) {
    if (!window || !x || !y) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        NSRect frame = [nsWindow frame];
        *x = (int)frame.origin.x;
        *y = (int)frame.origin.y;
    }
}

void setWindowPosition(WindowHandle window, int x, int y) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        NSPoint point = NSMakePoint(x, y);
        [nsWindow setFrameOrigin:point];
    }
}

void setWindowTitle(WindowHandle window, const char* title) {
    if (!window || !title) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        NSString *titleString = [NSString stringWithUTF8String:title];
        [nsWindow setTitle:titleString];
    }
}

void setDrawCallback(WindowHandle window, DrawCallback callback, void* userData) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        GoFlowView *view = (GoFlowView*)[nsWindow contentView];
        view.drawCallback = callback;
        view.drawUserData = userData;
    }
}

void setResizeCallback(WindowHandle window, ResizeCallback callback, void* userData) {
    if (!window) return;

    @autoreleasepool {
        NSWindow *nsWindow = (__bridge NSWindow*)window;
        GoFlowWindowDelegate *delegate = (GoFlowWindowDelegate*)[nsWindow delegate];
        delegate.resizeCallback = callback;
        delegate.resizeUserData = userData;
    }
}

void setMouseCallback(WindowHandle window, MouseCallback callback, void* userData) {
    // TODO: Implement mouse event handling
}

void setKeyCallback(WindowHandle window, KeyCallback callback, void* userData) {
    // TODO: Implement keyboard event handling
}
