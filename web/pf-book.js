const DEFAULT_LOGO = "/assets/rene.png";

function parseContents(text) {
    const sections = [];
    let section = null;
    let header = {
        links: [],
        logo: DEFAULT_LOGO
    };

    for (const line of text.split(/\r?\n/)) {
        const heading = line.match(/^##\s+(.+)$/);

        if (heading) {
            const title = heading[1];

            if (section && section.title !== "Header") {
                sections.push(section);
            }

            if (title === "Header") {
                section = null;
            } else {
                section = {
                    title,
                    pages: []
                };
            }

            continue;
        }

        const link = line.match(
            /^-\s+\[([^\]]+)\]\(([^)]+)\)$/
        );

        if (link) {
            const title = link[1];
            const url = link[2];

            // Links immediately under ## Header
            if (!section) {
                if (title.toLowerCase() === "logo") {
                    header.logo = url;
                } else {
                    header.links.push({
                        title,
                        url
                    });
                }

                continue;
            }
        }

        const page = line.match(/^-\s+(.+)$/);

        if (page && section) {
            const title = page[1];
            const filename =
                title.replace(/\s+/g, "-") + ".md";

            section.pages.push({
                title,
                filename,
                hash: filename.slice(0, -3).toLowerCase()
            });
        }
    }

    // Don't forget the final section.
    if (section) {
        sections.push(section);
    }

    return {
        header,
        sections
    };
}

class PfBook extends HTMLElement {
    async connectedCallback() {
        const src = this.getAttribute("src");
        const contentsSrc =
            src.replace(/\/?$/, "/Contents.md");

        const response = await fetch(contentsSrc);
        const text = await response.text();

        const parsed = parseContents(text);

        this.header = parsed.header;
        this.contents = parsed.sections;

        const theme =
                localStorage.getItem("pipefish-theme") ||
                (window.matchMedia("(prefers-color-scheme: dark)").matches
                ? "dark"
                : "light");

        document.documentElement.dataset.theme = theme;

        this.render();

        const hash = location.hash.slice(1);

        if (hash) {
            await this.showPage(hash);
        } else {
            const firstPage =
                this.contents[0]?.pages[0];

            if (firstPage) {
                await this.showPage(firstPage.hash);
            }
        }

        window.addEventListener("hashchange", () => {
            this.showPage(location.hash.slice(1));
        });
    }

    async showPage(hash) {
        const page = this.contents
            .flatMap(section => section.pages)
            .find(page => page.hash === hash);

        if (!page) {
            return;
        }

        this.querySelectorAll("nav a").forEach(item => {
            item.classList.toggle(
                "current",
                item.hash === "#" + hash
            );
        });

        const pageElement =
            this.querySelector("pf-page");

        pageElement.setAttribute(
            "title",
            page.title
        );

        pageElement.setAttribute(
            "src",
            this.pageSrc(page)
        );
    }

    pageSrc(page) {
        return this.getAttribute("src")
            .replace(/\/?$/, "/") +
            page.filename;
    }

    render() {
        this.innerHTML = "";

        // ------------------------------------------------------------
        // Header
        // ------------------------------------------------------------

        const header =
            document.createElement("header");

        header.classList.add("site-header");

        const logoLink =
            document.createElement("a");

        logoLink.classList.add("site-logo");
        logoLink.href = "/";

        const logo =
            document.createElement("img");

        logo.src = this.header.logo;
        logo.alt = "Pipefish";

        logoLink.appendChild(logo);
        header.appendChild(logoLink);

        const topNav =
            document.createElement("nav");

        topNav.classList.add("top-nav");

        for (const link of this.header.links) {
            const item =
                document.createElement("a");

            item.textContent = link.title;
            item.href = link.url;

            topNav.appendChild(item);
        }

        header.appendChild(topNav);

        const themeToggle =
            document.createElement("button");

        themeToggle.classList.add("theme-toggle");
        themeToggle.type = "button";
        
        const dark =
            document.documentElement.getAttribute("data-theme") ===
            "dark";

        themeToggle.textContent =
            dark ? "Light" : "Dark";

        themeToggle.setAttribute(
            "aria-label",
            dark
                ? "Switch to light mode"
                : "Switch to dark mode"
        );

        themeToggle.addEventListener(
            "click",
            () => {
                const dark =
                    document.documentElement
                        .getAttribute("data-theme") ===
                    "dark";

                const theme =
                    dark ? "light" : "dark";

                document.documentElement
                    .setAttribute(
                        "data-theme",
                        theme
                    );

                localStorage.setItem(
                    "pipefish-theme",
                    theme
                );

                themeToggle.textContent =
                    dark ? "Dark" : "Light";

                themeToggle.setAttribute(
                    "aria-label",
                    dark
                        ? "Switch to dark mode"
                        : "Switch to light mode"
                );
            }
        );

        header.appendChild(themeToggle);

        this.appendChild(header);

        // ------------------------------------------------------------
        // Main book
        // ------------------------------------------------------------

        const sidebar =
            document.createElement("nav");

        sidebar.classList.add("book-sidebar");

        for (const section of this.contents) {
            const heading =
                document.createElement("h2");

            heading.textContent = section.title;
            sidebar.appendChild(heading);

            for (const page of section.pages) {
                const item =
                    document.createElement("a");

                item.textContent = page.title;
                item.href = "#" + page.hash;

                sidebar.appendChild(item);
            }
        }

        const page = document.createElement("pf-page");

        const body =
            document.createElement("div");

        body.classList.add("book-body");

       
        body.appendChild(page);
        body.appendChild(sidebar);

        this.appendChild(body);

        // ------------------------------------------------------------
        // Footer
        // ------------------------------------------------------------

        const footer =
            document.createElement("footer");

        footer.classList.add("site-footer");

        const text =
            document.createElement("p");

        text.textContent =
            "This page is powered by Pipefish.";

        footer.appendChild(text);
        this.appendChild(footer);
    }
}

customElements.define("pf-book", PfBook);

export { parseContents };