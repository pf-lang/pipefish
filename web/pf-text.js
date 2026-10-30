import "./pf-highlighter.js";

class PfText extends HTMLElement {
    constructor() {
        super();

        const shadow = this.attachShadow({ mode: "open" });

        const style = document.createElement("link");
        style.rel = "stylesheet";
        style.href = new URL("./assets/pf-text.css", import.meta.url);

        const syntax = document.createElement("link");
        syntax.rel = "stylesheet";
        syntax.href = new URL("./assets/syntax.css", import.meta.url);

        const surface = document.createElement("div");
        surface.className = "surface";

        const highlighted = document.createElement("pre");
        highlighted.className = "highlighted";
        highlighted.setAttribute("aria-hidden", "true");

        const input = document.createElement("textarea");
        input.className = "input";
        input.spellcheck = false;
        input.autocomplete = "off";
        input.autocapitalize = "off";
        input.setAttribute("aria-label", "File contents");

        surface.append(highlighted, input);
        shadow.append(style, syntax, surface);

        this.surface = surface;
        this.highlighted = highlighted;
        this.input = input;

        this._value = "";
        this._highlighter = null;
        this._renderVersion = 0;

        input.addEventListener("input", () => {
            this._value = input.value;
            this.renderHighlight();

            // The native input event bubbles through the shadow boundary.
            // Consumers can listen for "input" on <pf-text>.
        });

        input.addEventListener("scroll", () => {
            highlighted.style.transform =
                `translate(${-input.scrollLeft}px, ${-input.scrollTop}px)`;
        });

        input.addEventListener("keydown", event => {
            if (event.key === "Tab") {
                event.preventDefault();

                input.setRangeText(
                    "\t",
                    input.selectionStart,
                    input.selectionEnd,
                    "end"
                );

                input.dispatchEvent(new Event("input", {
                    bubbles: true,
                    composed: true
                }));

                return;
            }

            if (this.handleDelimiter(event)) {
                return;
            }

            if (event.key !== "Enter") {
                return;
            }

            event.preventDefault();

            const start = input.selectionStart;
            const before = input.value.slice(0, start);
            const line = before.split("\n").pop();
            const indent = line.match(/^[\t ]*/)[0];

            const extraIndent =
                /(:\s*|--\s*)$/.test(line) ? "\t" : "";

            input.setRangeText(
                "\n" + indent + extraIndent,
                start,
                input.selectionEnd,
                "end"
            );

            input.dispatchEvent(new Event("input", {
                bubbles: true,
                composed: true
            }));
        });
    }

    get value() {
        return this.input.value;
    }

    set value(value) {
        this._value = String(value ?? "");

        if (this.input.value !== this._value) {
            this.input.value = this._value;
        }

        this.renderHighlight();
    }

    get editable() {
        return !this.input.readOnly;
    }

    set editable(value) {
        this.input.readOnly = !value;
        this.input.tabIndex = value ? 0 : -1;
        this.input.setAttribute(
            "aria-readonly",
            value ? "false" : "true"
        );

        this.toggleAttribute("editable", value);
    }

    get highlighter() {
        return this._highlighter;
    }

    set highlighter(value) {
        this._highlighter = typeof value === "function"
            ? value
            : null;

        this.renderHighlight();
    }

    async renderHighlight() {
        const version = ++this._renderVersion;
        const source = this._value;

        let html;

        try {
            if (this._highlighter) {
                html = await this._highlighter(source);
            } else {
                html = escapeHTML(source);
            }
        } catch (error) {
            console.error("Syntax highlighting failed:", error);
            html = escapeHTML(source);
        }

        // A newer edit may have happened while highlighting was running.
        if (version !== this._renderVersion) {
            return;
        }

        this.highlighted.innerHTML =
            html + (source.endsWith("\n") ? "\u00a0" : "");

        this.highlighted.style.transform =
            `translate(${-this.input.scrollLeft}px, ${-this.input.scrollTop}px)`;
    }

    focus() {
        this.input.focus();
    }

    get selectionStart() {
        return this.input.selectionStart;
    }

    set selectionStart(value) {
        this.input.selectionStart = value;
    }

    get selectionEnd() {
        return this.input.selectionEnd;
    }

    set selectionEnd(value) {
        this.input.selectionEnd = value;
    }

    setSelectionRange(start, end, direction) {
        this.input.setSelectionRange(start, end, direction);
    }

    get scrollTop() {
        return this.input.scrollTop;
    }

    set scrollTop(value) {
        this.input.scrollTop = value;
        this.syncHighlightScroll();
    }

    get scrollLeft() {
        return this.input.scrollLeft;
    }

    set scrollLeft(value) {
        this.input.scrollLeft = value;
        this.syncHighlightScroll();
    }

    syncHighlightScroll() {
        this.highlighted.style.transform =
            `translate(${-this.input.scrollLeft}px, ${-this.input.scrollTop}px)`;
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
        const input = this.input;

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
            const pos = input.selectionStart;

            if (
                pos === input.selectionEnd &&
                input.value[pos] === event.key
            ) {
                if (!isQuote) {
                    event.preventDefault();
                    input.setSelectionRange(pos + 1, pos + 1);
                    return true;
                }

                if (
                    this.quoteContext(input.value, pos, event.key) &&
                    !this.isEscaped(input.value, pos)
                ) {
                    event.preventDefault();
                    input.setSelectionRange(pos + 1, pos + 1);
                    return true;
                }
            }
        }

        if (pairs[event.key] !== undefined) {
            const start = input.selectionStart;
            const end = input.selectionEnd;

            if (
                isQuote &&
                this.quoteContext(input.value, start, event.key)
            ) {
                return false;
            }

            event.preventDefault();

            const selected = input.value.slice(start, end);

            input.setRangeText(
                event.key + selected + pairs[event.key],
                start,
                end,
                "select"
            );

            input.setSelectionRange(
                start + 1,
                start + 1 + selected.length
            );

            input.dispatchEvent(new Event("input", {
                bubbles: true,
                composed: true
            }));

            return true;
        }

        return false;
    }
}

function escapeHTML(text) {
    return text.replace(/[&<>"']/g, char => ({
        "&": "&amp;",
        "<": "&lt;",
        ">": "&gt;",
        '"': "&quot;",
        "'": "&#39;"
    })[char]);
}

customElements.define("pf-text", PfText);