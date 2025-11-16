#ifndef WINDOW_BRIDGE_H
#define WINDOW_BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif

// Window management
typedef void* WindowHandle;

// Application initialization
void initApp(void);

WindowHandle createWindow(int width, int height, const char* title);
void destroyWindow(WindowHandle window);
void showWindow(WindowHandle window);
void hideWindow(WindowHandle window);
int windowShouldClose(WindowHandle window);
void setWindowShouldClose(WindowHandle window, int shouldClose);

// Event handling
void pollEvents(WindowHandle window);
void waitEvents(WindowHandle window);

// Drawing
void* getWindowGraphicsContext(WindowHandle window);
void presentWindow(WindowHandle window);
void setWindowNeedsDisplay(WindowHandle window);

// Window properties
void getWindowSize(WindowHandle window, int* width, int* height);
void setWindowSize(WindowHandle window, int width, int height);
void getWindowPosition(WindowHandle window, int* x, int* y);
void setWindowPosition(WindowHandle window, int x, int y);
void setWindowTitle(WindowHandle window, const char* title);

// Callbacks (to be called from Go)
typedef void (*DrawCallback)(WindowHandle window, void* userData);
typedef void (*ResizeCallback)(WindowHandle window, int width, int height, void* userData);
typedef void (*MouseCallback)(WindowHandle window, int button, int action, double x, double y, void* userData);
typedef void (*KeyCallback)(WindowHandle window, int key, int action, void* userData);

void setDrawCallback(WindowHandle window, DrawCallback callback, void* userData);
void setResizeCallback(WindowHandle window, ResizeCallback callback, void* userData);
void setMouseCallback(WindowHandle window, MouseCallback callback, void* userData);
void setKeyCallback(WindowHandle window, KeyCallback callback, void* userData);

#ifdef __cplusplus
}
#endif

#endif // WINDOW_BRIDGE_H
