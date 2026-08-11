/******************************************************************************
 *
 * Pipefish Website
 *
 * site.js
 *
 ******************************************************************************/

(function () {
    "use strict";


    /*
     * ------------------------------------------------------------------------
     * Theme
     * ------------------------------------------------------------------------
     */

    const themeKey = "pipefish-theme";

    function preferredTheme() {
        const saved = localStorage.getItem(themeKey);

        if (saved === "light" || saved === "dark") {
            return saved;
        }

        return window.matchMedia("(prefers-color-scheme: dark)").matches
            ? "dark"
            : "light";
    }

    function setTheme(theme) {
        document.documentElement.dataset.theme = theme;

        const button = document.querySelector(".theme-toggle");

        if (button) {
            button.setAttribute(
                "aria-label",
                theme === "dark"
                    ? "Switch to light mode"
                    : "Switch to dark mode"
            );

            button.textContent =
                theme === "dark"
                    ? "Light"
                    : "Dark";
        }
    }

    function toggleTheme() {
        const current =
            document.documentElement.dataset.theme;

        const next =
            current === "dark"
                ? "light"
                : "dark";

        localStorage.setItem(themeKey, next);

        setTheme(next);
    }

    setTheme(preferredTheme());

    const themeButton =
        document.querySelector(".theme-toggle");

    if (themeButton) {
        themeButton.addEventListener(
            "click",
            toggleTheme
        );
    }


    /*
     * ------------------------------------------------------------------------
     * Code-block copy buttons
     * ------------------------------------------------------------------------
     */

    document
        .querySelectorAll(".code-copy")
        .forEach(function (button) {

            button.addEventListener(
                "click",
                async function () {

                    const codeBlock =
                        button.closest(".code-block");

                    if (!codeBlock) {
                        return;
                    }

                    const code =
                        codeBlock.querySelector("code");

                    if (!code) {
                        return;
                    }

                    try {
                        await navigator.clipboard.writeText(
                            code.textContent
                        );

                        const original =
                            button.textContent;

                        button.textContent = "Copied";

                        setTimeout(function () {
                            button.textContent = original;
                        }, 1500);

                    } catch (error) {
                        console.error(
                            "Could not copy code:",
                            error
                        );
                    }
                }
            );
        });

})();