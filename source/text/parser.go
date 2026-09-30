package text

import (
	"strings"
	"unicode/utf8"

	"github.com/tim-hardcastle/pipefish/source/dtypes"
)

type mdMode int

const (
	mdUnassigned mdMode = iota
	mdGettingParagraph
	mdGettingCodeBlock
	mdGettingCliBlock
	mdGettingQuote
	mdGettingList
	mdGettingHeading
	mdEnd
)

// As a pre-processing step, this inserts '⏎' characters where we need to insert a line break.
func Justify(width int, text string) string {
	if width <= 0 {
		return text
	}
	word := ""
	line := ""
	justifiedText := ""
	textLength := utf8.RuneCount([]byte(text))
	lineLengthCount := 0
	previousRune := '\n'
	textAsRunes := []rune(text)
	consumingCodeBlock := false
	for i, r := range textAsRunes {
		if previousRune == '\n' && r == '`' && i+2 < textLength && textAsRunes[i+1] == '`' &&
			textAsRunes[i+2] == '`' {
			consumingCodeBlock = !consumingCodeBlock
		}
		if consumingCodeBlock {
			justifiedText = justifiedText + string(r)
			previousRune = r
			continue
		}
		if previousRune == '\n' {
			if dtypes.SetOf('+', '-', '*').Contains(r) {
				lineLengthCount = 3
			}
			if r == '>' {
				lineLengthCount = -2
			}
		}
		word = word + string(r)
		if dtypes.SetOf('/', '-', ' ', '\n', ',', '.', ')', ';', ':', '>').Contains(r) || i+1 == textLength {
			wordLength := utf8.RuneCount([]byte(word))
			switch {
			case dtypes.SetOf("<R>", "<Y>", "<G>", "<C>", "<B>", "<P>", "</>").Contains(word):
				line = line + word
			case lineLengthCount+wordLength <= width:
				line = line + word
				lineLengthCount = lineLengthCount + wordLength
			default:
				justifiedText = justifiedText + line + "⏎"
				line = word
				lineLengthCount = wordLength
			}
			word = ""
		}
		previousRune = r
	}
	justifiedText = justifiedText + line
	return justifiedText
}

// The block parser.
func Parse(raw string) mdDocument {
	lines := strings.Split(raw, "\n")
	docNodes := []mdNode{}
	accumulator := []string{}
	mode := mdUnassigned
	var codeTag string
mainloop:
	for i := range len(lines) + 1 {
		line := ""
		if i < len(lines) {
			line = lines[i]
		}
		newMode := mdGettingParagraph
		if mode == mdGettingCodeBlock || mode == mdGettingCliBlock {
			newMode = mode
		}
		switch {
		case line == "" && !(mode == mdGettingCodeBlock || mode == mdGettingCliBlock):
			newMode = mdUnassigned
		case Head(line, "#"):
			newMode = mdGettingHeading
		case Head(line, "* ") || Head(line, "- ") || Head(line, "+ "):
			newMode = mdGettingList
		case Head(line, "> "):
			newMode = mdGettingQuote
		case Head(line, "```tui"):
			newMode = mdGettingCliBlock
		case Head(line, "```"):
			if mode == mdGettingCodeBlock || mode == mdGettingCliBlock {
				newMode = mdUnassigned
			} else {
				codeTag = line[3:]
				newMode = mdGettingCodeBlock
			}
		}
		if newMode == mode {
			accumulator = append(accumulator, line)
			continue mainloop
		}
		// So if we're here, we've switched from one sort of block to another.
		if len(accumulator) > 0 {
			// Then we can emit the appropriate old block.
			switch mode {
			case mdGettingHeading:
				docNodes = append(docNodes, makeHeading(accumulator))
			case mdGettingParagraph:
				docNodes = append(docNodes, makeParagraph(accumulator))
			case mdGettingCodeBlock:
				docNodes = append(docNodes, makeCodeBlock(codeTag, accumulator))
				codeTag = ""
			case mdGettingCliBlock:
				docNodes = append(docNodes, makeCliBlock(accumulator))
			case mdGettingQuote:
				docNodes = append(docNodes, makeQuote(accumulator))
			case mdGettingList:
				docNodes = append(docNodes, makeList(accumulator))
			}
		}

		// And we start a new block
		if Head(line, "```") { // We discard code block fences.
			accumulator = []string{}
		} else { // Otherwise the line that marked the end of the old block is the start of the new one.
			accumulator = []string{line}
		}
		mode = newMode
	}
	return mdDocument{docNodes}
}

func makeParagraph(lines []string) mdParagraph {
	ip := newInlineParser(strings.Join(lines, " "))
	return mdParagraph{ip.parseAll()}
}

func makeHeading(lines []string) mdHeading {
	heading := strings.Join(lines, " ")
	level := 0
	var ch rune
	for _, ch = range heading {
		if ch != '#' {
			break
		}
		level = level + 1
	}
	i := level
	for ; i < len(heading) && heading[i] == ' '; i++ {
	}
	ip := newInlineParser(heading[i:])
	return mdHeading{level, ip.parseAll()}
}

func makeCodeBlock(codeTag string, lines []string) mdCodeBlock {
	return mdCodeBlock{codeTag, lines}
}

func makeCliBlock(lines []string) mdCliBlock {
	return mdCliBlock{lines}
}

func makeQuote(lines []string) mdQuote {
	quote := ""
	sep := ""
	for _, line := range lines {
		quote = quote + sep + line[2:]
		sep = " "
	}
	ip := newInlineParser(quote)
	return mdQuote{ip.parseAll()}
}

func makeList(lines []string) mdList {
	items := []mdNode{}
	for _, line := range lines {
		ip := newInlineParser(line[2:])
		lineItems := ip.parseAll()
		items = append(items, mdListItem{lineItems})
	}
	return mdList{items}
}

// The inline parser.

type inlineParser struct {
	line string
	pos  int
}

type parserMode int

const (
	pmNone parserMode = iota
	pmText
	pmBold
	pmItalic
	pmRed
	pmYellow
	pmGreen
	pmCyan
	pmBlue
	pmPurple
)

var stopAt = map[parserMode][]string{
	pmText:   {"**", "*", "`", "<R>", "<Y>", "<G>", "<C>", "<B>", "<P>"},
	pmBold:   {"**"},
	pmItalic: {"*"},
	pmRed:    {"</>"},
	pmYellow: {"</>"},
	pmGreen:  {"</>"},
	pmCyan:   {"</>"},
	pmBlue:   {"</>"},
	pmPurple: {"</>"},
}

func (ip *inlineParser) parseAll() []mdNode {
	result := []mdNode{}
	for !ip.done() {
		result = append(result, ip.parse(pmNone)...)
	}
	return result
}

func (ip *inlineParser) parse(pM parserMode) []mdNode {
	txt := []byte{}
	for {
		if ip.done() {
			return []mdNode{mdText{string(txt)}}
		}
		// Inline code is handled differently since it doesn't get bold or italics or colors
		// in it.
		if pM == pmNone && ip.char() == '`' {
			return append([]mdNode{ip.parseInlineCode()}, ip.parse(pmNone)...)
		}
		if (pM == pmNone || pM == pmItalic) && ip.headIs("**") {
			ip.skip(2)
			bolded := ip.parse(pmBold)
			ip.skip(2)
			return []mdNode{mdFormat{stBold, bolded}}
		}
		if (pM == pmNone && ip.headIs("*")) ||
			(pM == pmBold && ip.headIs("*") && !ip.headIs("**")) {
			ip.next()
			italicized := ip.parse(pmItalic)
			ip.next()
			return []mdNode{mdFormat{stItalic, italicized}}
		}
		if pM == pmNone && ip.headIs("<R>") {
			ip.next()
			ip.next()
			ip.next()
			coloredText := ip.parse(pmRed)
			if ip.headIs("</>") {
				ip.skip(3)
			}
			return []mdNode{mdFormat{stRed, coloredText}}
		}
		if pM == pmNone && ip.headIs("<Y>") {
			ip.next()
			ip.next()
			ip.next()
			coloredText := ip.parse(pmYellow)
			if ip.headIs("</>") {
				ip.skip(3)
			}
			return []mdNode{mdFormat{stYellow, coloredText}}
		}
		if pM == pmNone && ip.headIs("<G>") {
			ip.next()
			ip.next()
			ip.next()
			coloredText := ip.parse(pmGreen)
			if ip.headIs("</>") {
				ip.skip(3)
			}
			return []mdNode{mdFormat{stGreen, coloredText}}
		}
		if pM == pmNone && ip.headIs("<C>") {
			ip.next()
			ip.next()
			ip.next()
			coloredText := ip.parse(pmCyan)
			if ip.headIs("</>") {
				ip.skip(3)
			}
			return []mdNode{mdFormat{stCyan, coloredText}}
		}
		if pM == pmNone && ip.headIs("<B>") {
			ip.next()
			ip.next()
			ip.next()
			coloredText := ip.parse(pmBlue)
			if ip.headIs("</>") {
				ip.skip(3)
			}
			return []mdNode{mdFormat{stBlue, coloredText}}
		}
		if pM == pmNone && ip.headIs("<P>") {
			ip.next()
			ip.next()
			ip.next()
			coloredText := ip.parse(pmPurple)
			if ip.headIs("</>") {
				ip.skip(3)
			}
			return []mdNode{mdFormat{stPurple, coloredText}}
		}

		if pM == pmNone {
			pM = pmText
		}
		// If we've hit some closing item we were looking for, then we return the text we
		// accumulated as `mdText`, which may then be wrapped in something else by the caller.
		if upTo, ok := stopAt[pM]; ok {
			for _, limit := range upTo {
				if ip.headIs(limit) && !(pM == pmItalic && ip.headIs("**")) {
					return []mdNode{mdText{string(txt)}}
				}
			}
		}
		txt = append(txt, ip.char())
		ip.next()
	}
}

func (ip *inlineParser) parseInlineCode() mdInlineCode {
	code := []byte{}
	ip.next()
	for !ip.done() {
		if ip.char() == '`' {
			ip.next()
			break
		}
		code = append(code, ip.char())
		ip.next()
	}
	return mdInlineCode{string(code)}
}

func (ip *inlineParser) done() bool {
	return ip.pos >= len(ip.line)
}

func (ip *inlineParser) char() byte {
	return ip.line[ip.pos]
}

func (ip *inlineParser) next() {
	ip.pos = ip.pos + 1
}

func (ip *inlineParser) skip(i int) {
	ip.pos = ip.pos + i
}

func (ip *inlineParser) headIs(s string) bool {
	upperBound := ip.pos + len(s)
	return upperBound < len(ip.line) && ip.line[ip.pos:upperBound] == s
}

func newInlineParser(raw string) inlineParser {
	return inlineParser{raw, 0}
}
