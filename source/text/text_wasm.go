//go:build js && wasm
package text

import (
	"strings"
	"unicode/utf8"
)

func Red(s string) string {
	return color(s, "red")
}

func Yellow(s string) string {
	return color(s, "yellow")
}

func Green(s string) string {
	return color(s, "green")
}

func Cyan(s string) string {
	return color(s, "cyan")
}

func Blue(s string) string {
	return color(s, "blue")
}

func Purple(s string) string {
	return color(s, "purple")
}

func Bold(s string) string {
	return "<b>" + s + "</b>"
}

func Italic(s string) string {
	return "<i>" + s + "</i>"
}

func Paragraph(s string) string {
	return "<p>" + s + "</p>\n"
}

func InlineCode(s string) string {
	return color(s, "code")
}

func color(s, c string) string {
	return("<span class = \"" + c + "\">" + s + "</span>")
}

func List(s string) string { return "<ul>\n" + s + "\n</ul>\n" }
func ListItem(s string) string { return "  <li>" + s + "</li>" }

func H1(s string) string { return "<h1>" + s + "</h1>\n" }
func H2(s string) string { return "<h2 id=\"" + Hyphenate(s) + "\">" + s + "</h2>\n" }
func H3(s string) string { return "<h3 id=\"" + Hyphenate(s) + "\">" + s + "</h3>\n" }
func H4(s string) string { return "<h4>" + s + "</h4>\n" }

func H1WithWidth(width int) (func(string) string) { return func(s string) string{return "≡≡≡≡ " + s + " " + strings.Repeat("≡", width - 6 - utf8.RuneCountInString(s))}}
func H2WithWidth(width int) (func(string) string) { return func(s string) string{return "════ " + s + " " + strings.Repeat("═", width - 6 - utf8.RuneCountInString(s))}}
func H3WithWidth(width int) (func(string) string) { return func(s string) string{return "―――― " + s + " " + strings.Repeat("―", width - 6 - utf8.RuneCountInString(s))}}
func H4WithWidth(width int) (func(string) string) { return func(s string) string{return "┈┈┈┈ " + s + " " + strings.Repeat("┈", width - 6 - utf8.RuneCountInString(s))}}

var BookRenderer = GetBookRenderer()

func GetTuiRenderer(width int) func(string) string {
	return NewRenderer(MakeRenderFunction(makeTuiHtml(width), htmlHighlighter), width, "<br>").Render
}

func RenderMdAsBookHtml(text string) string {
	return BookRenderer.Render(text)
}

func GetBookRenderer() Renderer {
	return NewRenderer(MakeRenderFunction(getBookHtml(), htmlHighlighter), -1, "")
}

func getBookHtml() map[mdStyle]func(string) string {
	bookSettings := map[mdStyle]func(string) string{
		stIde: func(s string) string { return "<pf-ide>" + s + "</pf-ide>" },
		stCodeBlock: func(s string) string {
			return "" +
				`<div class="code-block">
	<div class="code-header">
	<span class="code-language">Pipefish</span>
	<button class="code-copy" type="button">Copy</button>
	</div>
	<pre><code class="language-pipefish">` + s + `</code></pre>
	</div>
	`
		},
		stTuiBlock: func(s string) string {
			return "" +
				`<div class="code-block">
	<div class="code-header">
	<span class="code-language">TUI</span>
	</div>
	<pre><code class="language-pipefish">` + s + `</code></pre>
	</div>
	`
		},
		stH1:        H1,
		stH2:        H2,
		stH3:        H3,
		stH4:        H4,
		stLineBreak: func(s string) string { return "This shouldn't happen." },
	}
	return merge(defaultSettings, bookSettings)
}

func makeTuiHtml(width int) map[mdStyle]func(string) string {
	tuiSettings := map[mdStyle]func(string) string{
		stIde:       func(s string) string { return `</div><pre><code>` + s + `</code></pre></div>` },
		stCodeBlock: func(s string) string { return `</div><pre><code>` + s + `</code></pre></div>` },
		stTuiBlock:  func(s string) string { return `</div><pre><code>` + s + `</code></pre></div>` },
		stH1:        H1WithWidth(width),
		stH2:        H2WithWidth(width),
		stH3:        H3WithWidth(width),
		stH4:        H4WithWidth(width),
		stLineBreak: func(s string) string { return "<br>" },
	}
	return merge(defaultSettings, tuiSettings)
}
