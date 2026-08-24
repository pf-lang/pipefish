package markdown

import (
	"fmt"
	"reflect"
	"strings"
)

type mdNode interface {
	children() []mdNode
}

// A code block.
type mdCodeBlock struct {
	lines []string
}

func (n mdCodeBlock) children() []mdNode { return nil }

// A CLI code block.
type mdCliBlock struct {
	lines []string
}

func (n mdCliBlock) children() []mdNode { return nil }

// The root node.
type mdDocument struct {
	nodes []mdNode
}

func (n mdDocument) children() []mdNode { return n.nodes }

// Formatted text.
type mdFormat struct {
	style mdStyle
	nodes []mdNode
}

func (n mdFormat) children() []mdNode { return n.nodes }

// Headings with level numbers.
type mdHeading struct {
	level int
	nodes []mdNode
}

func (n mdHeading) children() []mdNode { return n.nodes }

// Inline code.
type mdInlineCode struct {
	text string
}

func (n mdInlineCode) children() []mdNode { return nil }

// A list node, the children of which will be of type mdListItem.
type mdList struct {
	nodes []mdNode
}

func (n mdList) children() []mdNode { return n.nodes }

// One item in the `nodes` field of an `mdList`.
type mdListItem struct {
	nodes []mdNode
}

func (n mdListItem) children() []mdNode { return n.nodes }

// Paragraph.
type mdParagraph struct {
	nodes []mdNode
}

func (n mdParagraph) children() []mdNode { return n.nodes }

// A quotation block.
type mdQuote struct {
	nodes []mdNode
}

func (n mdQuote) children() []mdNode { return n.nodes }

// Plain text.
type mdText struct {
	text string
}

func (n mdText) children() []mdNode { return nil }

var styleMap = map[mdStyle]string{
	stBold:   "bold",
	stItalic: "italic",
}

// Prettyprinter.

var mdFlavors = map[reflect.Type]string{
	reflect.TypeFor[mdCodeBlock]():  "codeblock",
	reflect.TypeFor[mdDocument]():   "doc",
	reflect.TypeFor[mdFormat]():     "format",
	reflect.TypeFor[mdHeading]():    "heading",
	reflect.TypeFor[mdInlineCode](): "code",
	reflect.TypeFor[mdList]():       "list",
	reflect.TypeFor[mdListItem]():   "item",
	reflect.TypeFor[mdQuote]():      "quote",
	reflect.TypeFor[mdParagraph]():  "par",
	reflect.TypeFor[mdText]():       "text",
}

func PpAst(n mdNode) string {
	var StringBuilder strings.Builder
	sb := &StringBuilder
	flavor := mdFlavors[reflect.TypeOf(n)]
	fmt.Fprint(sb, flavor, "(")
	switch n := n.(type) {
	case mdCodeBlock:
		fmt.Fprint(sb, strings.Join(n.lines, ", "), ")")
	case mdFormat:
		fmt.Fprint(sb, styleMap[n.style], ": ", strings.Join(prettyChildren(n), ", "), ")")
	case mdHeading:
		fmt.Fprint(sb, n.level, ", ", strings.Join(prettyChildren(n), ", "), ")")
	case mdInlineCode:
		fmt.Fprint(sb, n.text, ")")
	case mdText:
		fmt.Fprint(sb, n.text, ")")
	default:
		fmt.Fprint(sb, strings.Join(prettyChildren(n), ", "), ")")
	}
	return sb.String()
}

func prettyChildren(n mdNode) []string {
	var result []string
	for _, child := range n.children() {
		result = append(result, PpAst(child))
	}
	return result
}
