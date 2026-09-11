import "./pf-reader.js";

class PipefishEditor extends HTMLElement {
    constructor() {
        super();

        const shadow = this.attachShadow({ mode: "open" });

        const style = document.createElement("link");
        style.rel = "stylesheet";
        style.href =
            new URL("./assets/pf-editor.css", import.meta.url);

        const editor = document.createElement("div");
        editor.classList.add("editor");

        const reader = document.createElement("pf-reader");

        const code = document.createElement("textarea");
        code.placeholder = "Type Pipefish code here...";
        code.classList.add("code-input");
        code.spellcheck = false;

        editor.append(reader, code);
        shadow.append(style, editor);

        this.reader = reader;
        this.code = code;
        this.ready = reader.ready;

        code.addEventListener("input", async () => {
            await this.reader.display(code.value);
        });

        code.addEventListener("scroll", () => {
            this.syncScroll();
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

            if (this.handleDelimiter(event)) {
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
    }

    get value() {
        return this.code.value;
    }

    set value(value) {
        this.code.value = value;
        this.display(value);
    }

    async initialize(source) {
        await this.ready;

        const normalized =
            await this.reader.initialize(source);

        this.code.value = normalized;
        this.syncScroll();
    }

    async display(source) {
        await this.ready;

        this.code.value = source;
        await this.reader.display(source);
        this.syncScroll();
    }

    syncScroll() {
        this.reader.scrollTop = this.code.scrollTop;
        this.reader.scrollLeft = this.code.scrollLeft;
    }

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

    handleDelimiter(event) {
        const element = this.code;

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

        if (closingDelimiters.has(event.key)) {
            const pos = element.selectionStart;

            if (
                pos === element.selectionEnd &&
                element.value[pos] === event.key
            ) {
                if (!isQuote) {
                    event.preventDefault();
                    element.setSelectionRange(pos + 1, pos + 1);
                    return true;
                }

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

        if (pairs[event.key] !== undefined) {
            const start = element.selectionStart;
            const end = element.selectionEnd;

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

            const selected =
                element.value.slice(start, end);

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
}

customElements.define("pf-editor", PipefishEditor);