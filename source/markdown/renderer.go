package markdown

import (
	"reflect"
	"strings"

	"github.com/tim-hardcastle/pipefish/source/dtypes"
	"github.com/tim-hardcastle/pipefish/source/text"
)

type Renderer struct {
	RenderAst func(mdNode) string
	Width     int
	LineBreak string
}

func NewAstRenderer() Renderer {
	return NewRenderer(PpAst, -1, "")
}

func Iota(s string) string {
	return s
}

func NewRenderer(fn func(mdNode) string, width int, lineBreak string) Renderer {
	return Renderer{fn, width, lineBreak}
}

func (rnd Renderer) Render(raw string) string {
	justified := Justify(rnd.Width, raw)
	ast := Parse(justified)
	textWithoutLineBreaks := rnd.RenderAst(ast)
	return strings.TrimRight(strings.ReplaceAll(textWithoutLineBreaks, "⏎", rnd.LineBreak), "\n")
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
				sep = "\n\n"
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
			if dtypes.SetOf("pf-coder", "pf-editor", "pf-ide", "pf-reader", "pf-service", "pf-tui").Contains(n.tag) {
				sb.WriteString("<")
				sb.WriteString(n.tag)
				sb.WriteString(">")
				sb.WriteString(result)
				sb.WriteString("</")
				sb.WriteString(n.tag)
				sb.WriteString(">")
			} else {
				sb.WriteString(textWrapper[stCodeBlock](result))
			}
		case mdCliBlock:
			result := ""
			for _, line := range n.lines {
				lineOut := ""
				if ix := strings.Index(line, "→"); ix != -1 {
					if ix > 0 {
						lineOut = "<span class=\"service\">" + (line[:ix-1]) + "</span> "
					}
					lineOut = lineOut + "→" + codeHighlighter(line[ix+3:]) // +3 because that's how many bytes there are in the rune.
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
	stIde
	stCodeBlock
	stTuiBlock
	stH1
	stH2
	stH3
	stH4
	stRed
	stYellow
	stGreen
	stCyan
	stBlue
	stPurple
	stLineBreak
)

var defaultSettings = map[mdStyle]func(string) string{
	stParagraph:text.Paragraph,
	stBold:     text.Bold,
	stItalic:   text.Italic,
	stInline:   text.InlineCode,
	stRed:		text.Red,
	stYellow:	text.Yellow,
	stGreen:	text.Green,
	stCyan:		text.Cyan,
	stBlue:		text.Blue,
	stPurple:	text.Purple,
	stList:     text.List,
	stListItem: text.ListItem,
}

func getTerminalRenderer(width int) map[mdStyle]func(string) string {
	terminalSettings := map[mdStyle]func(s string) string{
		stH1:        text.H1WithWidth(width),
		stH2:        text.H2WithWidth(width),
		stH3:        text.H3WithWidth(width),
		stH4:        text.H4WithWidth(width), 
	}
	return merge(defaultSettings, terminalSettings)
}


func merge(maps ...map[mdStyle]func(string) string) map[mdStyle]func(string) string {
	merged := map[mdStyle]func(string) string{}
	for _, M := range maps {
		for k, v := range M {
			merged[k] = v
		}
	}
	return merged
}
