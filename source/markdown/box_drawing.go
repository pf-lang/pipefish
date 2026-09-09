package markdown

import (
	"strings"

	"github.com/tim-hardcastle/pipefish/source/dtypes"
)

func BoxDrawing(s string) string {
	raw := strings.Split(s, "\n")
	rawRunes := [][]rune{}
	for _, line := range raw {
		rawRunes = append(rawRunes, []rune(line))
	}
	cookedStrings := []string{}
	for y, line := range rawRunes {
		cookedLine := []rune{}
		for x, r := range line {
			q := getLocus(locus{
					r,
					safeLook(rawRunes, x, y-1),
					safeLook(rawRunes, x, y+1),
					safeLook(rawRunes, x-1, y),
					safeLook(rawRunes, x+1, y),
				},
			)
			cookedLine = append(cookedLine, q)
		}
		cookedStrings = append(cookedStrings, string(cookedLine))
	}
	return strings.Join(cookedStrings, "\n")
}

func safeLook(rawRunes [][]rune, x, y int) rune {
	if y < 0 || len(rawRunes) <= y || x < 0 || len(rawRunes[y]) <= x {
		return ' '
	}
	return rawRunes[y][x]
}

type locus struct{center, up, down, left, right rune}

func getLocus(L locus) rune {
	switch L.center {
	case '|':
		if dtypes.SetOf('|', '^', '-', '+').Contains(L.up) {
			if dtypes.SetOf('|', 'v', '-', '+').Contains(L.down) {
				if dtypes.SetOf('-', '+').Contains(L.left) {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '┼'
					} else {
						return '┤'
					}
				} else {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '├'
					} else {
						return '│'
					}
				}
			} else {
				if dtypes.SetOf('-', '+').Contains(L.left) {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '┴'
					} else {
						return '╯'
					}
				} else {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '╰'
					} else {
						return '│'
					}
				}
			}
		} else {
			if dtypes.SetOf('|', 'v', '-', '+').Contains(L.down) {
				if dtypes.SetOf('-', '+').Contains(L.left) {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '┬'
					} else {
						return '╮'
					}
				} else {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '╭'
					} else {
						return '|'
					}
				}
			} else {
				if dtypes.SetOf('-', '+').Contains(L.left) {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '┼'
					} else {
						return '╮'
					}
				} else {
					if dtypes.SetOf('-', '+').Contains(L.right) {
						return '╭'
					} else {
						return '|'
					}
				}
			}
		}
	case '-':
		if dtypes.SetOf('-', '<', '|', '+').Contains(L.left) {
			if dtypes.SetOf('-', '>', '|', '+').Contains(L.right) {
				if dtypes.SetOf('|', '+').Contains(L.up) {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '┼'
					} else {
						return '┴'
					}
				} else {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '┬'
					} else {
						return '─'
					}
				}
			} else {
				if dtypes.SetOf('|', '+').Contains(L.up) {
					if '|' == L.down {
						return '┤'
					} else {
						return '╯'
					}
				} else {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '╮'
					} else {
						return '─'
					}
				}
			}
		} else {
			if dtypes.SetOf('-', '>', '|', '+').Contains(L.right) {
				if dtypes.SetOf('|', '+').Contains(L.up) {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '├'
					} else {
						return '╰'
					}
				} else {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '╭'
					} else {
						return '─'
					}
				}
			} else {
				if dtypes.SetOf('|', '+').Contains(L.up) {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '┼'
					} else {
						return '┴'
					}
				} else {
					if dtypes.SetOf('|', '+').Contains(L.down) {
						return '┬'
					} else {
						return '-'
					}
				}
			}
		}
	case '^':
		if dtypes.SetOf('|', '+').Contains(L.down) {
			return '🢑'
		} else {
			return '^'
		}
	case 'v':
		if dtypes.SetOf('|', '+').Contains(L.up) {
			return '🢓'
		} else {
			return 'v'
		}
	case '<':
		if dtypes.SetOf('-', '+').Contains(L.right) {
			return '🢐'
		} else {
			return '<'
		}
	case '>':
		if dtypes.SetOf('-', '+').Contains(L.left) {
			return '🢒'
		} else {
			return '>'
		}
	case '+':
		if dtypes.SetOf('-', '|', '^', '+').Contains(L.up) ||
		dtypes.SetOf('-', '|', 'v', '+').Contains(L.down) ||
		dtypes.SetOf('-', '|', '<', '+').Contains(L.left) ||
		dtypes.SetOf('-', '|', '>', '+').Contains(L.right) {
			return '┼'
		} else {
			return '+'
		}
			

	default:
		return L.center
	}
}

