//go:build !js && !wasm

package text_test

import (
	"strconv"
	"testing"

	"github.com/tim-hardcastle/pipefish/source/markdown"
	"github.com/tim-hardcastle/pipefish/source/test_helper"
	"github.com/tim-hardcastle/pipefish/source/text"
)

func TestColors(t *testing.T) {
	if !(text.Red("foo") == "\x1b[31mfoo\x1b[39m") {
		t.Fatal("Can't make things red.", strconv.Quote(text.Red("foo")))
	}
	if !(text.Cyan("foo") == "\x1b[36mfoo\x1b[39m") {
		t.Fatal("Can't make things cyan.")
	}
	if !(text.Green("foo") == "\x1b[32mfoo\x1b[39m") {
		t.Fatal("Can't make things green.")
	}
	if !(text.Yellow("foo") == "\x1b[33mfoo\x1b[39m") {
		t.Fatalf("Can't make things green.")
	}
	if !(text.Emph("foo") == "`foo`") {
		t.Fatal("Can't emphasize things.")
	}
	if !(text.ErrorFont("foo") == "\x1b[38;2;244;71;71m\x1b[4mfoo\x1b[0m") {
		t.Fatal("Can't make error font.")
	}
}

func TestMarkdown(t *testing.T) {
	tests := []test_helper.TestItem{
		{`Hello`, `Hello`},
		{`Hello *darkness* my **old** friend.`, "Hello \x1b[3mdarkness\x1b[23m my \x1b[1mold\x1b[22m friend."},
		{`<R>red</>.`, "\x1b[31mred\x1b[39m."},
		{"inline `code` looks like this.", "inline \x1b[48;2;0;0;64m\x1b[97mcode\x1b[49m\x1b[39m looks like this."},
		{`## Heading`, "════ Heading ═══════════════════════════════════════════════════════════════════════════════"},
		{"- Bullet point", "  ▪ Bullet point"},
	}
	render := markdown.GetTuiRenderer(92)
	for _, test := range tests {
		got := render(test.Input)
		println(got)
		if !(test.Want == got) {
			t.Fatalf("Test failed with input %s \nExp :\n%s\nGot :\n%s", test.Input, strconv.Quote(test.Want), strconv.Quote(got))
		}
	}
}

func TestTextUtils(t *testing.T) {
	if !(text.Flatten("foo/bar.troz") == "foo_bar_troz") {
		t.Fatalf("Flatten failed")
	}
	if !(text.Head("foolish", "foo")) {
		t.Fatalf("Head failed")
	}
	if text.Head("aardvark", "foo") {
		t.Fatalf("Head failed")
	}
	if text.Head("aa", "foo") {
		t.Fatalf("Head failed")
	}
	if !(text.Tail("proof", "oof")) {
		t.Fatalf("Tail failed")
	}
	if text.Tail("aardvark", "oof") {
		t.Fatalf("Tail failed")
	}
	if text.Tail("aa", "oof") {
		t.Fatalf("Tail failed")
	}
}
