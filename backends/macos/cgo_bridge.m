#import <Cocoa/Cocoa.h>
#import <QuartzCore/QuartzCore.h>
#include "cgo_bridge.h"
#include <stdlib.h>
#include <string.h>

// GraphicsContext wraps the Core Graphics context and bitmap
typedef struct {
    CGContextRef context;
    NSBitmapImageRep *bitmap;
    int width;
    int height;
} GraphicsContext;

// Stack for save/restore operations
typedef struct SaveState {
    CGAffineTransform transform;
    struct SaveState *next;
} SaveState;

typedef struct {
    GraphicsContext *gc;
    SaveState *saveStack;
} CanvasContext;

// Create a new graphics context
void* createGraphicsContext(int width, int height) {
    @autoreleasepool {
        // Create bitmap context
        NSBitmapImageRep *bitmap = [[NSBitmapImageRep alloc]
            initWithBitmapDataPlanes:NULL
            pixelsWide:width
            pixelsHigh:height
            bitsPerSample:8
            samplesPerPixel:4
            hasAlpha:YES
            isPlanar:NO
            colorSpaceName:NSCalibratedRGBColorSpace
            bytesPerRow:width * 4
            bitsPerPixel:32];

        if (!bitmap) {
            return NULL;
        }

        // Create Core Graphics context from bitmap
        NSGraphicsContext *nsContext = [NSGraphicsContext graphicsContextWithBitmapImageRep:bitmap];
        CGContextRef cgContext = [nsContext CGContext];

        if (!cgContext) {
            return NULL;
        }

        // Retain the context
        CGContextRetain(cgContext);

        // Set default parameters
        CGContextSetAllowsAntialiasing(cgContext, true);
        CGContextSetShouldAntialias(cgContext, true);
        CGContextSetInterpolationQuality(cgContext, kCGInterpolationHigh);

        // Keep Core Graphics natural coordinate system (bottom-left origin)
        // We'll convert coordinates in drawing functions as needed

        // Allocate graphics context structure
        GraphicsContext *gc = malloc(sizeof(GraphicsContext));
        gc->context = cgContext;
        gc->bitmap = (__bridge_retained void*)bitmap;
        gc->width = width;
        gc->height = height;

        // Allocate canvas context
        CanvasContext *canvas = malloc(sizeof(CanvasContext));
        canvas->gc = gc;
        canvas->saveStack = NULL;

        return canvas;
    }
}

// Destroy graphics context
void destroyGraphicsContext(void* ctx) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        GraphicsContext *gc = canvas->gc;

        if (gc) {
            if (gc->context) {
                CGContextRelease(gc->context);
            }
            if (gc->bitmap) {
                CFBridgingRelease(gc->bitmap);
            }
            free(gc);
        }

        // Clean up save stack
        SaveState *state = canvas->saveStack;
        while (state) {
            SaveState *next = state->next;
            free(state);
            state = next;
        }

        free(canvas);
    }
}

// Clear canvas with color
void clearCanvas(void* ctx, double r, double g, double b, double a) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        CGContextSetRGBFillColor(context, r, g, b, a);
        CGRect rect = CGRectMake(0, 0, canvas->gc->width, canvas->gc->height);
        CGContextFillRect(context, rect);
    }
}

// Draw rectangle
void drawRect(void* ctx, double x, double y, double w, double h,
              double r, double g, double b, double a, int filled, double strokeWidth) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        // Convert from UI coordinates (top-left) to Core Graphics coordinates (bottom-left)
        double cgY = canvas->gc->height - y - h;
        CGRect rect = CGRectMake(x, cgY, w, h);

        if (filled) {
            CGContextSetRGBFillColor(context, r, g, b, a);
            CGContextFillRect(context, rect);
        } else {
            CGContextSetRGBStrokeColor(context, r, g, b, a);
            CGContextSetLineWidth(context, strokeWidth);
            CGContextStrokeRect(context, rect);
        }
    }
}

// Draw circle
void drawCircle(void* ctx, double cx, double cy, double radius,
                double r, double g, double b, double a, int filled, double strokeWidth) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        // Convert from UI coordinates (top-left) to Core Graphics coordinates (bottom-left)
        double cgY = canvas->gc->height - cy - radius;
        CGRect rect = CGRectMake(cx - radius, cgY, radius * 2, radius * 2);

        if (filled) {
            CGContextSetRGBFillColor(context, r, g, b, a);
            CGContextFillEllipseInRect(context, rect);
        } else {
            CGContextSetRGBStrokeColor(context, r, g, b, a);
            CGContextSetLineWidth(context, strokeWidth);
            CGContextStrokeEllipseInRect(context, rect);
        }
    }
}

// Draw line
void drawLine(void* ctx, double x1, double y1, double x2, double y2,
              double r, double g, double b, double a, double strokeWidth) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        // Convert from UI coordinates (top-left) to Core Graphics coordinates (bottom-left)
        double cgY1 = canvas->gc->height - y1;
        double cgY2 = canvas->gc->height - y2;

        CGContextSetRGBStrokeColor(context, r, g, b, a);
        CGContextSetLineWidth(context, strokeWidth);

        CGContextBeginPath(context);
        CGContextMoveToPoint(context, x1, cgY1);
        CGContextAddLineToPoint(context, x2, cgY2);
        CGContextStrokePath(context);
    }
}

// Draw text
void drawText(void* ctx, const char* text, double x, double y,
              const char* fontName, double fontSize,
              double r, double g, double b, double a, int fontWeight) {
    if (!ctx || !text) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        // Create NSString from C string
        NSString *string = [NSString stringWithUTF8String:text];
        NSString *font = fontName ? [NSString stringWithUTF8String:fontName] : @"Helvetica";

        // Determine font weight
        NSFontWeight weight = (fontWeight >= 700) ? NSFontWeightBold : NSFontWeightRegular;

        // Create font
        NSFont *nsFont = [NSFont systemFontOfSize:fontSize weight:weight];
        if (fontName && strlen(fontName) > 0) {
            NSFont *customFont = [NSFont fontWithName:font size:fontSize];
            if (customFont) {
                nsFont = customFont;
            }
        }

        // Create text color
        NSColor *textColor = [NSColor colorWithCalibratedRed:r green:g blue:b alpha:a];

        // Create attributes dictionary
        NSDictionary *attributes = @{
            NSFontAttributeName: nsFont,
            NSForegroundColorAttributeName: textColor
        };

        // Calculate text size and convert coordinates
        NSSize textSize = [string sizeWithAttributes:attributes];

        // Convert from UI coordinates (top-left) to Core Graphics coordinates (bottom-left)
        double cgY = canvas->gc->height - y - textSize.height;

        // Use standard Core Graphics coordinate system (bottom-left origin)
        NSGraphicsContext *nsContext = [NSGraphicsContext graphicsContextWithCGContext:context flipped:NO];
        [NSGraphicsContext saveGraphicsState];
        [NSGraphicsContext setCurrentContext:nsContext];

        // Draw the text at the converted position
        [string drawAtPoint:NSMakePoint(x, cgY) withAttributes:attributes];

        // Restore NSGraphicsContext
        [NSGraphicsContext restoreGraphicsState];
    }
}

// Save canvas state
void saveCanvas(void* ctx) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        CGContextSaveGState(context);

        // Also save to our custom stack for tracking
        SaveState *state = malloc(sizeof(SaveState));
        state->transform = CGContextGetCTM(context);
        state->next = canvas->saveStack;
        canvas->saveStack = state;
    }
}

// Restore canvas state
void restoreCanvas(void* ctx) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        CGContextRestoreGState(context);

        // Pop from our custom stack
        if (canvas->saveStack) {
            SaveState *state = canvas->saveStack;
            canvas->saveStack = state->next;
            free(state);
        }
    }
}

// Translate canvas
void translateCanvas(void* ctx, double dx, double dy) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        CGContextTranslateCTM(context, dx, dy);
    }
}

// Scale canvas
void scaleCanvas(void* ctx, double sx, double sy) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        CGContextScaleCTM(context, sx, sy);
    }
}

// Rotate canvas
void rotateCanvas(void* ctx, double radians) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        CGContextRotateCTM(context, radians);
    }
}

// Clip to rectangle
void clipRect(void* ctx, double x, double y, double w, double h) {
    if (!ctx) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        CGContextRef context = canvas->gc->context;

        // Convert from UI coordinates (top-left) to Core Graphics coordinates (bottom-left)
        double cgY = canvas->gc->height - y - h;
        CGRect rect = CGRectMake(x, cgY, w, h);
        CGContextClipToRect(context, rect);
    }
}

// Get pixel data (for debugging or saving)
void getPixelData(void* ctx, unsigned char* buffer, int bufferSize) {
    if (!ctx || !buffer) return;

    @autoreleasepool {
        CanvasContext *canvas = (CanvasContext*)ctx;
        NSBitmapImageRep *bitmap = (__bridge NSBitmapImageRep*)canvas->gc->bitmap;

        unsigned char *bitmapData = [bitmap bitmapData];
        int dataSize = canvas->gc->width * canvas->gc->height * 4;

        if (bufferSize >= dataSize) {
            memcpy(buffer, bitmapData, dataSize);
        }
    }
}

// Get canvas dimensions
void getCanvasDimensions(void* ctx, int* width, int* height) {
    if (!ctx || !width || !height) return;

    CanvasContext *canvas = (CanvasContext*)ctx;
    *width = canvas->gc->width;
    *height = canvas->gc->height;
}
