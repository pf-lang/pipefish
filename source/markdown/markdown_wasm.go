//go:build js && wasm

package markdown

import (
	"github.com/tim-hardcastle/pipefish/source/text"
)

var BookRenderer = GetBookRenderer()

func GetTuiRenderer(width int) func(string)string {
	return NewRenderer(MakeRenderFunction(makeTuiHtml(width), htmlHighlighter)).Render
}

func RenderMdAsBookHtml(text string) string {
	return BookRenderer.Render(text)
}

func GetBookRenderer() Renderer {
	return NewRenderer(MakeRenderFunction(getBookHtml(), htmlHighlighter))
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
		stH1: text.H1,
		stH2: text.H2,
		stH3: text.H3,
		stH4: text.H4,
	}
	return merge(defaultSettings, bookSettings)
}

func makeTuiHtml(width int) map[mdStyle]func(string) string {
	tuiSettings := map[mdStyle]func(string) string{
		stIde:       func(s string) string { return `</div><pre><code>` + s + `</code></pre></div>` },
		stCodeBlock: func(s string) string { return `</div><pre><code>` + s + `</code></pre></div>` },
		stTuiBlock:  func(s string) string { return `</div><pre><code>` + s + `</code></pre></div>` },
		stH1:        text.H1WithWidth(width),
		stH2:        text.H2WithWidth(width),
		stH3:        text.H3WithWidth(width),
		stH4:        text.H4WithWidth(width),
	}
	return merge(defaultSettings, tuiSettings)
}
