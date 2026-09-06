class PipefishHighlighter extends HTMLElement {
    constructor() {
    super();
        this.ready = this.loadWasm();
    }

    highlight(source) {
        return window.pipefishHighlight(source);
    }

    async loadWasm() {
        const wasmURL = "assets/pipefish.wasm";
        const execURL = "assets/wasm_exec.js";

        if (!window.Go) {
            await this.loadScript(execURL);
        }

        this.go = new Go();

        const result = await WebAssembly.instantiateStreaming(
            fetch(wasmURL),
            this.go.importObject
        );

        this.go.run(result.instance);
    }

    loadScript(url) {
        return new Promise((resolve, reject) => {
            const script = document.createElement("script");

            script.src = url;
            script.onload = resolve;
            script.onerror = reject;

            document.head.appendChild(script);
        });
    }
}

customElements.define("pf-highlighter", PipefishHighlighter);
