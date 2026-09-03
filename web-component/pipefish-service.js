class PipefishService extends HTMLElement {
    constructor() {
        super();

        this.ready = Promise.all([
            this.loadWasm(),
            this.loadFont()
        ]);

        this.history = [];
        this.historyIndex = 0;

        const shadow = this.attachShadow({ mode: "open" });

        const style = document.createElement("link");
        style.rel = "stylesheet";
        style.href = "pipefish-service.css";

        const syntax = document.createElement("link");
        syntax.rel = "stylesheet";
        syntax.href = "syntax.css";

        const wrapper = document.createElement("div");
        wrapper.classList.add("service");

        const heading = document.createElement("h2");
        heading.textContent = "Pipefish service";

        const editor = document.createElement("div");
        editor.classList.add("code-editor");

        const highlighted = document.createElement("pre");
        highlighted.classList.add("highlighted");

        const code = document.createElement("textarea");
        code.placeholder = "Type Pipefish code here...";
        code.classList.add("code-input");
        code.spellcheck = false;

        editor.append(highlighted, code);

        code.addEventListener("input", () => { this.updateHighlighting(); });

        code.addEventListener("scroll", () => {
            highlighted.scrollTop = code.scrollTop;
            highlighted.scrollLeft = code.scrollLeft;
        });

        this.code = code;
        this.highlighted = highlighted;

        const compileButton = document.createElement("button");
        compileButton.textContent = "Compile";

        const repl = document.createElement("div");
        repl.classList.add("repl");

        const transcript = document.createElement("div");
        transcript.classList.add("transcript");
        this.transcript = transcript;

        const inputLine = document.createElement("div");
        inputLine.classList.add("input-line");

        const prompt = document.createElement("span");
        prompt.classList.add("prompt");
        prompt.textContent = "→ ";

        const input = document.createElement("input");
        input.classList.add("input");
        input.type = "text";
        input.autocomplete = "off";
        input.spellcheck = false;

        inputLine.append(prompt, input);
        repl.append(transcript, inputLine);

        wrapper.append(
            heading,
            editor,
            compileButton,
            repl
        );

        shadow.append(style, syntax, wrapper);

        console.log("shadow styles:", [...shadow.querySelectorAll("link")].map(x => x.href));

        compileButton.addEventListener("click", async () => {
            try {
                await this.compile(code.value);
                this.write("Compiled successfully.");
                input.focus();
            } catch (error) {
                this.write("Error: " + error);
            }
        });

        input.addEventListener("keydown", async (event) => {
            if (event.key === "Enter") {
                event.preventDefault();
                await this.submit(input);
            }
        });
    }

    syncEditorScroll(input, highlighted) {
        highlighted.scrollTop = input.scrollTop;
        highlighted.scrollLeft = input.scrollLeft;
    }

    async updateHighlighting() {
        await this.ready;

        this.highlighted.innerHTML =
            window.pipefishHighlight(this.code.value);

        this.syncEditorScroll(this.code, this.highlighted);
    }

    async submit(input) {
        const line = input.value;

        if (!line.trim()) {
            return;
        }

        this.history.push(line);
        this.historyIndex = this.history.length;

        this.write("→ " + line);
        input.value = "";

        try {
            const result = await this.do(line);

            if (result !== undefined && result !== "") {
                this.write(result);
            }
        } catch (error) {
            this.write("Error: " + error);
        }
    }

    write(text) {
        const line = document.createElement("div");
        line.textContent = text;
        this.transcript.appendChild(line);
    }

    async loadWasm() {
        const wasmURL = "pipefish.wasm";
        const execURL = "wasm_exec.js";

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

    async loadFont() {
    const fontURL = new URL(
        "code-font/GoogleSansCode-VariableFont_MONO,wght.woff2",
        import.meta.url
    );

    const font = new FontFace(
        "Google Sans Code",
        `url("${fontURL}")`,
            {
                weight: "100 900",
                style: "normal"
            }
        );

        await font.load();
        document.fonts.add(font);
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

    async compile(source) {
        await this.ready;
        return window.pipefishCompile(source);
    }

    async do(line) {
        await this.ready;
        return window.pipefishDo(line);
    }
}

customElements.define("pipefish-service", PipefishService);