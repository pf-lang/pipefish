package markdown_test

import (
	"strconv"
	"testing"

	"github.com/tim-hardcastle/pipefish/source/markdown"
)

type mdTest struct{
	input string
	output string
}

func TestParsing(t *testing.T) {
	tests := []mdTest{
		{"foo *bar* qux", "doc(par(text(foo ), format(italic: text(bar)), text( qux)))"},
		{"foo **bar** qux", "doc(par(text(foo ), format(bold: text(bar)), text( qux)))"},
		{"foo `bar` qux", "doc(par(text(foo ), code(bar), text( qux)))"},
		{"hello\ndarkness", "doc(par(text(hello darkness)))"},
		{"> hello\n> darkness", "doc(quote(text(hello darkness)))"},
		{"* hello\n* darkness", "doc(list(item(text(hello)), item(text(darkness))))"},
		{"```\nhello\ndarkness\n```", "doc(codeblock(hello, darkness))"},
		{"##hello darkness", "doc(heading(2, text(hello darkness)))"},
	}
	runTests(t, tests, markdown.NewAstRenderer())
}

func TestHtml(t *testing.T) {
	tests := []mdTest{
		{"foo *bar* qux", "<p>foo <i>bar</i> qux</p>\n"},
		{"foo **bar** qux", "<p>foo <b>bar</b> qux</p>\n"},
		{"foo `bar` qux", "<p>foo <code>bar</code> qux</p>\n"},
		{"* hello\n* darkness", "<ul>\n  <li>hello</li>\n  <li>darkness</li>\n</ul>\n"},
		{"# Title", "<h1>Title</h1>\n"},
		{"## Title", "<h2 id=\"Title\">Title</h2>\n"},
		{"### Title", "<h3 id=\"Title\">Title</h3>\n"},
		{"#### Title", "<h4>Title</h4>\n"},
	}
	runTests(t, tests, markdown.NewHtmlRenderer())
}

func runTests(t *testing.T, tests []mdTest, rnd markdown.Renderer) {
	for _, test := range tests {
		result := rnd.Render(test.input)
		if result != test.output {
			t.Fatalf("expected \n%s\n and got \n%s\n", strconv.Quote(test.output), strconv.Quote(result))
		}
	}
}