//go:build !js && !wasm

package text

import (
	"strings"
	"unicode/utf8"
)

func Red(s string) string {
	return RED + s + RESET_FOREGROUND
}

func Yellow(s string) string {
	return YELLOW + s + RESET_FOREGROUND
}

func Green(s string) string {
	return GREEN + s + RESET_FOREGROUND
}

func Cyan(s string) string {
	return CYAN + s + RESET_FOREGROUND
}

func Blue(s string) string {
	return BLUE + s + RESET_FOREGROUND
}

func Purple(s string) string {
	return PURPLE + s + RESET_FOREGROUND
}

func Bold(s string) string {
	return BOLD + s + RESET_BOLD
}

func Italic(s string) string {
	return ITALIC + s + RESET_ITALIC
}

func Paragraph(s string) string {
	return s
}

func InlineCode(s string) string {
	return INLINE_CODE_BACKGROUND + WHITE + s +
		RESET_BACKGROUND + RESET_FOREGROUND
}

func List(s string) string {
	return s
}

func ListItem(s string) string {
	return BULLET + s
}

func H1WithWidth(width int) (func(string) string) { return func(s string) string{return "≡≡≡≡ " + s + " " + strings.Repeat("≡", width - 6 - utf8.RuneCountInString(s))}}
func H2WithWidth(width int) (func(string) string) { return func(s string) string{return "════ " + s + " " + strings.Repeat("═", width - 6 - utf8.RuneCountInString(s))}}
func H3WithWidth(width int) (func(string) string) { return func(s string) string{return "―――― " + s + " " + strings.Repeat("―", width - 6 - utf8.RuneCountInString(s))}}
func H4WithWidth(width int) (func(string) string) { return func(s string) string{return "┈┈┈┈ " + s + " " + strings.Repeat("┈", width - 6 - utf8.RuneCountInString(s))}}

func GetTuiRenderer(width int) func(string) string {
	return NewTerminalRenderer(Iota, width).Render
}

func NewTerminalRenderer(highlighter func(string) string, width int) Renderer {
	return NewRenderer(MakeRenderFunction(getTerminalRenderer(width), highlighter), width, "\n")
}