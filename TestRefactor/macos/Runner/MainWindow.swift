import Cocoa

class MainWindow: NSWindow {
    init() {
        super.init(
            contentRect: NSRect(x: 0, y: 0, width: 1200, height: 800),
            styleMask: [.titled, .closable, .miniaturizable, .resizable],
            backing: .buffered,
            defer: false
        )

        self.title = "TestRefactor"
        self.center()

        // TODO: Load GoFlow shared library and initialize engine
        setupGoFlowEngine()
    }

    private func setupGoFlowEngine() {
        // This will load the compiled Go shared library (libapp.dylib)
        // and initialize the GoFlow rendering engine

        print("GoFlow: Setting up rendering engine for TestRefactor")

        // Future implementation:
        // 1. Load libapp.dylib from build output
        // 2. Call Go initialization function
        // 3. Set up rendering surface
        // 4. Start event loop

        // For now, show placeholder content
        let textField = NSTextField(frame: NSRect(x: 20, y: 20, width: 760, height: 60))
        textField.stringValue = "GoFlow is initializing..."
        textField.isEditable = false
        textField.isBordered = false
        textField.backgroundColor = .clear
        contentView?.addSubview(textField)
    }
}
