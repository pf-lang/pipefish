class PipefishService extends HTMLElement {
    constructor() {
        super();

        this.ready = Promise.all([
            this.loadWasm(),
            this.loadFont()
        ]);

        this.history = [];
        this.historyIndex = 0;
        this.multiline = false;

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

        // ------------------------------------------------------------
        // Main code editor
        // ------------------------------------------------------------

        const editor = document.createElement("div");
        editor.classList.add("code-editor");

        const highlighted = document.createElement("pre");
        highlighted.classList.add("highlighted");

        const code = document.createElement("textarea");
        code.placeholder = "Type Pipefish code here...";
        code.classList.add("code-input");
        code.spellcheck = false;

        editor.append(highlighted, code);

        code.addEventListener("input", async () => {
            await this.ready;
            highlighted.innerHTML =
                window.pipefishHighlight(code.value);
        });

        code.addEventListener("scroll", () => {
            highlighted.scrollTop = code.scrollTop;
            highlighted.scrollLeft = code.scrollLeft;
        });

        code.addEventListener("keydown", event => {
            if (event.key === "Tab") {
                event.preventDefault();

                code.setRangeText(
                    "\t",
                    code.selectionStart,
                    code.selectionEnd,
                    "end"
                );

                code.dispatchEvent(new Event("input"));
                return;
            }

            if (event.key !== "Enter") {
                return;
            }

            event.preventDefault();

            const start = code.selectionStart;
            const before = code.value.slice(0, start);
            const line = before.split("\n").pop();

            const indent = line.match(/^[\t ]*/)[0];
            const extraIndent =
                /(:\s*|--\s*)$/.test(line) ? "\t" : "";

            code.setRangeText(
                "\n" + indent + extraIndent,
                start,
                code.selectionEnd,
                "end"
            );

            code.dispatchEvent(new Event("input"));
        });

        this.code = code;
        this.highlighted = highlighted;

        // ------------------------------------------------------------
        // Compile button
        // ------------------------------------------------------------

        const compileButton = document.createElement("button");
        compileButton.textContent = "Compile";

        // ------------------------------------------------------------
        // REPL
        // ------------------------------------------------------------

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

        const inputEditor = document.createElement("div");
        inputEditor.classList.add("input-editor");

        const highlightedInput = document.createElement("pre");
        highlightedInput.classList.add("highlighted-input");

        const input = document.createElement("textarea");
        input.classList.add("input");
        input.autocomplete = "off";
        input.spellcheck = false;
        input.rows = 1;

        inputEditor.append(highlightedInput, input);
        inputLine.append(prompt, inputEditor);

        this.input = input;
        this.highlightedInput = highlightedInput;

        // ------------------------------------------------------------
        // REPL live highlighting and sizing
        // ------------------------------------------------------------

        input.addEventListener("input", async () => {
            await this.ready;

            highlightedInput.innerHTML =
                window.pipefishHighlight(input.value);

            this.resizeReplInput();
        });

        input.addEventListener("scroll", () => {
            this.syncEditorScroll(input, highlightedInput);
        });

        // ------------------------------------------------------------
        // REPL keyboard handling
        // ------------------------------------------------------------

        input.addEventListener("keydown", async event => {
            // Tab inserts a literal tab.
            if (event.key === "Tab") {
                event.preventDefault();

                input.setRangeText(
                    "\t",
                    input.selectionStart,
                    input.selectionEnd,
                    "end"
                );

                input.dispatchEvent(new Event("input"));
                return;
            }

            if (event.key !== "Enter") {
                return;
            }

            const start = input.selectionStart;
            const before = input.value.slice(0, start);
            const currentLine = before.split("\n").pop();

            // --------------------------------------------------------
            // Ordinary one-line REPL input
            // --------------------------------------------------------

            if (!this.multiline) {
                if (!input.value.trim()) {
                    event.preventDefault();
                    return;
                }

                // A line ending in : or -- starts multiline mode.
                if (/(:\s*|--\s*)$/.test(currentLine)) {
                    event.preventDefault();

                    input.setRangeText(
                        "\n" + this.replIndent(currentLine),
                        input.selectionStart,
                        input.selectionEnd,
                        "end"
                    );

                    this.multiline = true;

                    input.dispatchEvent(new Event("input"));
                    return;
                }

                // Ordinary command: submit it.
                event.preventDefault();
                await this.submitReplInput();
                return;
            }

            // --------------------------------------------------------
            // Multiline REPL input
            // --------------------------------------------------------

            // Enter on an otherwise-empty line terminates the command.
            // The empty line itself is not part of the command.
            if (!currentLine.trim()) {
                event.preventDefault();
                await this.submitReplInput();
                return;
            }

            // Otherwise insert a newline, preserving indentation and
            // adding one tab after : or --.
            event.preventDefault();

            input.setRangeText(
                "\n" + this.replIndent(currentLine),
                input.selectionStart,
                input.selectionEnd,
                "end"
            );

            input.dispatchEvent(new Event("input"));
        });

        repl.append(transcript, inputLine);

        wrapper.append(
            heading,
            editor,
            compileButton,
            repl
        );

        shadow.append(style, syntax, wrapper);

        // ------------------------------------------------------------
        // Compile
        // ------------------------------------------------------------

        compileButton.addEventListener("click", async () => {
            try {
                await this.compile(code.value);
                this.write("Compiled successfully.");
                input.focus();
            } catch (error) {
                this.write("Error: " + error);
            }
        });

        // Establish the initial one-line input height.
        this.resizeReplInput();
    }

    // ------------------------------------------------------------
    // REPL helpers
    // ------------------------------------------------------------

    replIndent(line) {
        const indent = line.match(/^[\t ]*/)[0];

        if (/(:\s*|--\s*)$/.test(line)) {
            return indent + "\t";
        }

        return indent;
    }

    async submitReplInput() {
        let command = this.input.value;

        // Remove the terminating empty/whitespace-only physical line.
        if (this.multiline) {
            const lines = command.split("\n");

            if (!lines[lines.length - 1].trim()) {
                lines.pop();
            }

            command = lines.join("\n");
        }

        if (!command.trim()) {
            this.multiline = false;
            this.input.value = "";
            this.highlightedInput.innerHTML = "";
            this.resizeReplInput();
            return;
        }

        this.multiline = false;

        // Capture the command before clearing the live editor.
        this.input.value = "";
        this.highlightedInput.innerHTML = "";
        this.resizeReplInput();

        await this.executeReplCommand(command);
    }

    async executeReplCommand(command) {
        this.history.push(command);
        this.historyIndex = this.history.length;

        const entry = document.createElement("pre");
        entry.classList.add("transcript-input");

        entry.innerHTML =
            `<span class="prompt">→ </span>` +
            window.pipefishHighlight(command);

        this.transcript.appendChild(entry);

        try {
            const result = await this.do(command);

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

    resizeReplInput() {
        const input = this.input;
        const editor = input.parentElement;

        input.style.height = "auto";

        const lineHeight =
            parseFloat(getComputedStyle(input).lineHeight);

        const maxHeight = lineHeight * 12;
        const height = Math.min(input.scrollHeight, maxHeight);

        input.style.height = height + "px";
        editor.style.height = height + "px";
    }

    syncEditorScroll(input, highlighted) {
        highlighted.scrollTop = input.scrollTop;
        highlighted.scrollLeft = input.scrollLeft;
    }

    // ------------------------------------------------------------
    // Main editor
    // ------------------------------------------------------------

    async updateHighlighting() {
        await this.ready;

        this.highlighted.innerHTML =
            window.pipefishHighlight(this.code.value);

        this.syncEditorScroll(this.code, this.highlighted);
    }

    // ------------------------------------------------------------
    // Pipefish/WASM
    // ------------------------------------------------------------

    async compile(source) {
        await this.ready;
        return window.pipefishCompile(source);
    }

    async do(line) {
        await this.ready;
        return window.pipefishDo(line);
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
}

customElements.define("pipefish-service", PipefishService);

