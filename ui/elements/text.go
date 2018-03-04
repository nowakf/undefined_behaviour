package elements

import (
	"strings"
	c "ub/common"
)

type text struct {
	scrollOffset  int
	unformatted   string
	formatted     string
	height, width int
	wrapping      wrapType
	align         alignment
	*container
}

type wrapType int
type alignment int

const (
	horizontal wrapType = iota
	h_wrap_v_cut
	h_cut
)

const (
	center alignment = iota
	left
	right
)

func (t *text) Content() string {
	return t.unformatted
}

func (t *text) Draw(x, y int) []c.Cell {

	formatted := t.format(t.unformatted, t.H(), t.W())

	lines := strings.Split(formatted, "\n")
	hasHidden := false

	if len(lines) > t.H() {
		lines = t.scroll(t.scrollOffset, lines)
		hasHidden = true
	}

	content := t.toCellArray(hasHidden, lines)

	transformed := make([]c.Cell, len(content))

	for i, cell := range content {
		transformed[i] = c.Cell{
			X:          cell.X + x,
			Y:          cell.Y + y,
			Letter:     cell.Letter,
			Foreground: t.foreground,
			Background: t.background,
		}
	}

	return transformed
}

func (t *text) format(input string, height, width int) string {
	text := ""
	switch t.wrapping {
	case horizontal:
		paragraphs := strings.Split(input, "\n")
		text, height = t.wrappedParagraphs(width, paragraphs)
	case h_wrap_v_cut:
		paragraphs := strings.Split(input, "\n")
		text, _ = t.wrappedParagraphs(width, paragraphs)
		text = t.vertTruncate(text, height)
	case h_cut:
		text = t.horizTruncate(input, width)
	default:
		panic("there's no style defined for this text")
	}

	leftpad, rightpad := "", ""

	switch t.align {
	case left:
		rightpad = t.rightPad(width, text)
	case center:
		leftpad, rightpad = t.justifiedPad(width, text)
	default:
		panic("no alignment defined for this text!")
	}
	return leftpad + text + rightpad

}

//returns a formatted body text, and a height
func newbodytext(content string, box *container) *text {
	t := new(text)
	t.container = box
	t.unformatted = content
	t.wrapping = horizontal
	t.align = left

	return t
}

func (t *text) wrappedParagraphs(width int, paragraphs []string) (string, int) {
	text := ""
	height := 1
	for _, paragraph := range paragraphs {
		t, h := t.wrap(width, paragraph)
		text += "\n" + t
		height += h + 1
	}
	return text, height
}

//returns a formatted title
func newTitleText(content string, box *container) *text {
	t := new(text)
	t.wrapping = h_cut
	t.align = center
	t.unformatted = content
	t.container = box
	return t

}

func (t *text) horizTruncate(s string, width int) string {
	if len(s) < width {
		return s
	} else {
		return s[0:width]
	}
}

func (t *text) vertTruncate(text string, height int) string {

	lines := strings.Split(text, "\n")

	if len(lines) <= height {
		return text
	} else {
		trimmed := ""

		for i := 0; i < height; i++ {
			trimmed += lines[i] + "\n"
		}
		return trimmed + "..."
	}

}
func (t *text) rightPad(width int, input string) string {
	spaces := ""
	for width-len(input)-len(spaces) > 0 {
		spaces += " "
	}
	return spaces
}
func (t *text) justifiedPad(space int, input string) (string, string) {
	left := ""
	right := ""
	space -= len(input)

	for space > 0 {
		left += " "
		space--
		if space > 0 {
			right += " "
		}
		space--
	}
	return left, right
}

func (t *text) wrap(width int, content string) (string, int) {

	words := strings.Fields(strings.TrimSpace(content))

	wrapped := ""

	spaceLeft := width

	height := 0

	for _, word := range words {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = width - len(word)
			height++
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}

	}
	return wrapped, height

}

func (t *text) toCellArray(hasHidden bool, lines []string) []c.Cell {
	output := make([]c.Cell, 0)
	for y, line := range lines {
		for x, letter := range line {
			c := c.Cell{
				X:          x,
				Y:          y,
				Letter:     letter,
				Foreground: t.foreground,
				Background: t.background,
			}
			output = append(output, c)
		}

	}

	return output
}

func (t *text) scroll(offset int, lines []string) []string {
	start := 0 + offset
	end := t.H() + offset
	if start < 0 {
		start = 0
	}
	if end >= len(lines) {
		end = len(lines) - 1
	}
	return lines[start:end]

}
