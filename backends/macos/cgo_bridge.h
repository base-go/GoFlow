#ifndef CGO_BRIDGE_H
#define CGO_BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif

// Graphics context management
void* createGraphicsContext(int width, int height);
void destroyGraphicsContext(void* ctx);

// Drawing operations
void clearCanvas(void* ctx, double r, double g, double b, double a);
void drawRect(void* ctx, double x, double y, double w, double h,
              double r, double g, double b, double a, int filled, double strokeWidth);
void drawCircle(void* ctx, double cx, double cy, double radius,
                double r, double g, double b, double a, int filled, double strokeWidth);
void drawLine(void* ctx, double x1, double y1, double x2, double y2,
              double r, double g, double b, double a, double strokeWidth);
void drawText(void* ctx, const char* text, double x, double y,
              const char* fontName, double fontSize,
              double r, double g, double b, double a, int fontWeight);

// Transform operations
void saveCanvas(void* ctx);
void restoreCanvas(void* ctx);
void translateCanvas(void* ctx, double dx, double dy);
void scaleCanvas(void* ctx, double sx, double sy);
void rotateCanvas(void* ctx, double radians);
void clipRect(void* ctx, double x, double y, double w, double h);

// Utility functions
void getPixelData(void* ctx, unsigned char* buffer, int bufferSize);
void getCanvasDimensions(void* ctx, int* width, int* height);

#ifdef __cplusplus
}
#endif

#endif // CGO_BRIDGE_H
