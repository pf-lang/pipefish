import "./pf-service.js"
import "./pf-highlighter.js"

class PipefishIde extends HTMLElement {
    constructor() {
        super();

        const service = document.createElement("pf-service");
        const highlighter = document.createElement("pf-highlighter");

        this.service = service;
        this.highlighter = highlighter;

        this.ready = Promise.all([
            service.ready,
            highlighter.ready
        ]);

        this.history = [];
        this.historyIndex = 0;
        this.multiline = false;

        const shadow = this.attachShadow({ mode: "open" });

        const style = document.createElement("link");
        style.rel = "stylesheet";
        style.href = "assets/pf-page.css";

        const syntax = document.createElement("link");
        syntax.rel = "stylesheet";
        syntax.href = "assets/syntax.css";

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
                this.highlighter.Highlight(code.value);
        });

        code.addEventListener("scroll", () => {
            highlighted.scrollTop = code.scrollTop;
            highlighted.scrollLeft = code.scrollLeft;
        });

        code.addEventListener("keydown", event => {
            // Tab inserts a literal tab.
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

            // Handle brackets and quotes.
            if (this.handleDelimiter(event, code)) {
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
                this.highlighter.Highlight(input.value);

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

            // Handle brackets and quotes.
            if (this.handleDelimiter(event, input)) {
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
                await this.service.compile(code.value);
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
    // Delimiter helpers
    // ------------------------------------------------------------

    isEscaped(text, pos) {
        let backslashes = 0;

        for (
            let i = pos - 1;
            i >= 0 && text[i] === "\\";
            i--
        ) {
            backslashes++;
        }

        return backslashes % 2 === 1;
    }

    quoteContext(text, pos, quote) {
        let currentQuote = null;
        let escaped = false;

        for (let i = 0; i < pos; i++) {
            const char = text[i];

            if (escaped) {
                escaped = false;
                continue;
            }

            if (char === "\\") {
                escaped = true;
                continue;
            }

            if (currentQuote === null) {
                if (
                    char === '"' ||
                    char === "'" ||
                    char === "`"
                ) {
                    currentQuote = char;
                }
            } else if (char === currentQuote) {
                currentQuote = null;
            }
        }

        return currentQuote === quote;
    }

    handleDelimiter(event, element) {
        const pairs = {
            "(": ")",
            "[": "]",
            "{": "}",
            '"': '"',
            "'": "'",
            "`": "`"
        };

        const closingDelimiters = new Set([
            ")",
            "]",
            "}",
            '"',
            "'",
            "`"
        ]);

        const isQuote =
            event.key === '"' ||
            event.key === "'" ||
            event.key === "`";

        // --------------------------------------------------------
        // Skip over an existing closing delimiter.
        // --------------------------------------------------------

        if (closingDelimiters.has(event.key)) {
            const pos = element.selectionStart;

            if (
                pos === element.selectionEnd &&
                element.value[pos] === event.key
            ) {
                // Ordinary brackets always skip an existing closer.
                if (!isQuote) {
                    event.preventDefault();
                    element.setSelectionRange(pos + 1, pos + 1);
                    return true;
                }

                // Quotes only skip an existing closer when we are
                // actually inside a string of that type.
                //
                // An escaped quote is not a closing quote.
                if (
                    this.quoteContext(
                        element.value,
                        pos,
                        event.key
                    ) &&
                    !this.isEscaped(element.value, pos)
                ) {
                    event.preventDefault();
                    element.setSelectionRange(pos + 1, pos + 1);
                    return true;
                }
            }
        }

        // --------------------------------------------------------
        // Insert the matching closing delimiter.
        // --------------------------------------------------------

        if (pairs[event.key] !== undefined) {
            const start = element.selectionStart;
            const end = element.selectionEnd;

            // Quotes should not auto-pair when we're already inside
            // a quoted literal.
            if (
                isQuote &&
                this.quoteContext(
                    element.value,
                    start,
                    event.key
                )
            ) {
                return false;
            }

            event.preventDefault();

            const selected = element.value.slice(start, end);

            element.setRangeText(
                event.key + selected + pairs[event.key],
                start,
                end,
                "select"
            );

            element.setSelectionRange(
                start + 1,
                start + 1 + selected.length
            );

            element.dispatchEvent(new Event("input"));
            return true;
        }

        return false;
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
            this.highlighter.Highlight(command);

        this.transcript.appendChild(entry);

        try {
            const result = await this.service.do(command);

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
            this.highlighter.Highlight(this.code.value);

        this.syncEditorScroll(this.code, this.highlighted);
    }
}

customElements.define("pf-ide", PipefishIde);