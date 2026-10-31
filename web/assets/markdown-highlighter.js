function escapeHTML(text) {
    return text
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#39;");
}

function span(className, text) {
    return `<span class="${className}">${text}</span>`;
}

function highlightInline(source) {
    const protectedParts = [];

    function protect(html) {
        const index = protectedParts.length;
        protectedParts.push(html);
        return `\u0000${index}\u0000`;
    }

    // Inline code.
    source = source.replace(/(`+)(.+?)\1/g, (_, ticks, content) =>
        protect(
            span("md-inline-code", escapeHTML(ticks + content + ticks))
        )
    );

    // Images.
    source = source.replace(
        /!\[([^\]]*)\]\(([^)]+)\)/g,
        (_, alt, url) =>
            protect(
                span("md-link", escapeHTML(`![${alt}](${url})`))
            )
    );

    // Links.
    source = source.replace(
        /\[([^\]]+)\]\(([^)]+)\)/g,
        (_, text, url) =>
            protect(
                span("md-link", escapeHTML(`[${text}](${url})`))
            )
    );

    // Escape remaining source text.
    source = escapeHTML(source);

    // Bold.
    source = source.replace(
        /(\*\*|__)(.+?)\1/g,
        (_, delimiter, content) =>
            span("md-bold", delimiter + content + delimiter)
    );

    // Italic.
    source = source.replace(
        /(?<!\*)\*(?!\*)(.+?)(?<!\*)\*(?!\*)|(?<!_)_(?!_)(.+?)(?<!_)_(?!_)/g,
        (_, a, b) => {
            const delimiter = a !== undefined ? "*" : "_";
            const content = a ?? b;
            return span("md-italic", delimiter + content + delimiter);
        }
    );

    // Restore protected parts.
    source = source.replace(/\u0000(\d+)\u0000/g, (_, index) =>
        protectedParts[Number(index)]
    );

    return source;
}

export function highlightMarkdown(source) {
    const lines = source.replace(/\r\n/g, "\n").split("\n");
    const result = [];

    let inFence = false;
    let fenceMarker = "";
    let fenceLength = 0;

    for (const line of lines) {
        const fence = line.match(/^\s*(`{3,}|~{3,})/);

        // Fenced code blocks.
        if (fence) {
            const marker = fence[1][0];
            const length = fence[1].length;

            if (!inFence) {
                inFence = true;
                fenceMarker = marker;
                fenceLength = length;
            } else if (
                marker === fenceMarker &&
                length >= fenceLength
            ) {
                inFence = false;
            }

            result.push(
                span("md-inline-code", escapeHTML(line))
            );
            continue;
        }

        if (inFence) {
            result.push(
                span("md-inline-code", escapeHTML(line))
            );
            continue;
        }

        // Headings.
        const heading = line.match(/^(#{1,6}\s+)(.*)$/);

        if (heading) {
            result.push(
                span("md-heading", escapeHTML(line))
            );
            continue;
        }

        // Horizontal rules.
        if (/^\s{0,3}([-*_])(?:\s*\1){2,}\s*$/.test(line)) {
            result.push(
                span("md-quote", escapeHTML(line))
            );
            continue;
        }

        // Blockquotes.
        const quote = line.match(/^(\s{0,3}>\s?)(.*)$/);

        if (quote) {
            result.push(
                span(
                    "md-quote",
                    escapeHTML(quote[1]) + highlightInline(quote[2])
                )
            );
            continue;
        }

        // Ordered and unordered lists.
        const list = line.match(/^(\s*)([-+*]|\d+[.)])(\s+)(.*)$/);

        if (list) {
            result.push(
                escapeHTML(list[1]) +
                span("md-list", escapeHTML(list[2] + list[3])) +
                highlightInline(list[4])
            );
            continue;
        }

        // Ordinary text.
        result.push(
            span("md-body", highlightInline(line))
        );
    }

    return result.join("\n");
}