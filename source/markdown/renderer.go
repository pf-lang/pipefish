package markdown

import (
	"reflect"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/text"
)

type Renderer struct {
	RenderAst func(mdNode) string
}

func NewAstRenderer() Renderer {
	return NewRenderer(PpAst)
}

func NewHtmlRenderer() Renderer {
	return NewRenderer(MakeRenderFunction(html, htmlHighlighter))
}

func NewTerminalRenderer(highlighter func(string) string) Renderer {
	return NewRenderer(MakeRenderFunction(getTerminalRenderer(92), highlighter))
}

func NewRenderer(fn func(mdNode) string) Renderer {
	return Renderer{fn}
}

func (rnd Renderer) Render(raw string) string {
	ast := rnd.Parse(raw)
	return rnd.RenderAst(ast)
}

type ContentsItem struct {
	Heading    string
	Subheading []string
}

func (rnd Renderer) ExtractHeadings(doc mdDocument) []ContentsItem {
	result := []ContentsItem{}
	for _, bit := range doc.nodes {
		if heading, ok := bit.(mdHeading); ok {
			if !(heading.level == 2 || heading.level == 3) {
				continue
			}
			hText := ""
			for _, hBit := range heading.nodes {
				hText = hText + rnd.RenderAst(hBit)
			}
			if heading.level == 2 {
				result = append(result, ContentsItem{hText, nil})
			} else {
				result[len(result)-1].Subheading = append(result[len(result)-1].Subheading, hText)
			}
		}
	}
	return result
}

func MakeRenderFunction(textWrapper map[mdStyle]func(s string) string, codeHighlighter func(s string) string) func(mdNode) string {

	var render func(mdNode) string

	render = func(n mdNode) string {
		var builder strings.Builder
		sb := &builder
		sep := ""
		switch n := n.(type) {
		case mdDocument:
			for _, block := range n.nodes {
				sb.WriteString(sep)
				sb.WriteString(render(block))
				sep = "\n"
			}
		case mdParagraph:
			result := ""
			for _, text := range n.nodes {
				result = result + render(text)
			}
			sb.WriteString(textWrapper[stParagraph](result))
		case mdText:
			return n.text
		case mdFormat:
			for _, text := range n.nodes {
				if plain, ok := text.(mdText); ok {
					sb.WriteString(textWrapper[n.style](plain.text))
				} else {
					sb.WriteString(render(text))
				}
			}
		case mdHeading:
			text := ""
			for _, bit := range n.nodes {
				text = text + (render(bit))
			}
			sb.WriteString(textWrapper[mdStyle(int(stH1)+n.level-1)](text))
		case mdInlineCode:
			sb.WriteString(textWrapper[stInline](n.text))
		case mdList:
			list := ""
			for _, item := range n.nodes {
				list = list + sep
				itemText := ""
				for _, text := range item.children() {
					itemText = itemText + render(text)
				}
				list = list + textWrapper[stListItem](itemText)
				sep = "\n"
			}
			sb.WriteString(textWrapper[stList](list))
		case mdCodeBlock:
			result := ""
			for _, line := range n.lines {
				result = result + sep + codeHighlighter(line)
				sep = "\n"
			}
			sb.WriteString(textWrapper[stCodeBlock](result))
		case mdCliBlock:
			result := ""
			for _, line := range n.lines {
				lineOut := ""
				if ix := strings.Index(line, "→"); ix != -1 {
					if ix > 0 {
						lineOut = "<span class=\"service\">" + (line[:ix-1]) + "</span> "
					}
					lineOut = lineOut + "→" + codeHighlighter(line[ix+3:])    // +3 because that's how many bytes there are in the rune.
				} else {
					lineOut = line
				}
				result = result + sep + lineOut
				sep = "\n"
			}
			sb.WriteString(textWrapper[stTuiBlock](result))
		default:
			panic("unhandled case " + reflect.TypeOf(n).String())
		}
		return sb.String()
	}
	return render
}

// An enum for styles.
type mdStyle int

const (
	stParagraph mdStyle = iota
	stBold
	stItalic
	stInline
	stList
	stListItem
	stCodeBlock
	stTuiBlock
	stH1
	stH2
	stH3
	stH4
)

var html = map[mdStyle]func(string) string{
	stParagraph: func(s string) string { return "<p>" + s + "</p>\n" },
	stBold:      func(s string) string { return "<b>" + s + "</b>" },
	stItalic:    func(s string) string { return "<i>" + s + "</i>" },
	stInline:    func(s string) string { return "<code>" + s + "</code>" },
	stList:      func(s string) string { return "<ul>\n" + s + "\n</ul>\n" },
	stListItem:  func(s string) string { return "  <li>" + s + "</li>" },
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
	// By convention the major divisions in the document are h2 and the lower divisions are h3
	// and nothing else counts for the purposes of making the table of contents.
	stH1: func(s string) string { return "<h1>" + s + "</h1>\n" },
	stH2: func(s string) string { return "<h2 id=\"" + text.Hyphenate(s) + "\">" + s + "</h2>\n" },
	stH3: func(s string) string { return "<h3 id=\"" + text.Hyphenate(s) + "\">" + s + "</h3>\n" },
	stH4: func(s string) string { return "<h4>" + s + "</h4>\n" },
}

func getTerminalRenderer(width int) map[mdStyle]func(string) string {
	return map[mdStyle]func(s string) string{
		stBold:   func(s string) string { return text.BOLD + s + text.RESET_BOLD },
		stItalic: func(s string) string { return text.ITALIC + s + text.RESET_ITALIC },
		stInline: func(s string) string {
			return text.INLINE_CODE_BACKGROUND + text.WHITE + s +
				text.RESET_BACKGROUND + text.RESET_FOREGROUND
		},
		stList:     func(s string) string { return "<ul>\n" + s + "\n</ul>" },
		stListItem: func(s string) string { return "  <li>" + s + "</li>" },
		stH1:       func(s string) string { return "<h1>" + s + "</h1>" },
		stH2:       func(s string) string { return "<h2>" + s + "</h2>" },
		stH3:       func(s string) string { return "<h3>" + s + "</h3>" },
		stH4:       func(s string) string { return "<h4>" + s + "</h4>" },
	}
}
