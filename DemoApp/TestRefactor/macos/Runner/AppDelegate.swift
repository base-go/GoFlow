import Cocoa

@main
class AppDelegate: NSObject, NSApplicationDelegate {
    var mainWindow: MainWindow?

    func applicationDidFinishLaunching(_ notification: Notification) {
        // Initialize GoFlow runtime
        print("GoFlow: Initializing TestRefactor")

        // Create main window
        mainWindow = MainWindow()
        mainWindow?.makeKeyAndOrderFront(nil)

        NSApp.activate(ignoringOtherApps: true)
    }

    func applicationWillTerminate(_ notification: Notification) {
        print("GoFlow: Terminating TestRefactor")
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        return true
    }
}
