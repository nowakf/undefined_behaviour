package components

import (
	c "cthu3/common"
	"strings"
	"unicode"
)

type Text struct {
	content  []c.Cell
	asString string
	*Rect
}

func (t *Text) Content() string {
	return t.asString
}

func (t *Text) WithOutline(x, y int) []c.Cell {
	transformed := make([]c.Cell, len(t.content))

	for i, cell := range t.content {
		if x == cell.X && y == cell.Y && x == cell.X+t.W() && y == cell.Y+t.H() {

			highlighted := unicode.ToUpper(cell.Content)

			transformed[i] = c.Cell{X: cell.X + x, Y: cell.Y + y, Content: highlighted}

		} else {
			transformed[i] = c.Cell{X: cell.X + x, Y: cell.Y + y, Content: cell.Content}
		}

	}
	return transformed
}

func (t *Text) Draw(x, y int) []c.Cell {

	transformed := make([]c.Cell, len(t.content))

	for i, cell := range t.content {
		transformed[i] = c.Cell{X: cell.X + x, Y: cell.Y + y, Content: cell.Content}
	}

	return transformed
}

//returns a formatted body text, and a height
func NewBodyText(width int, content string) (*Text, int) {
	t := new(Text)
	t.asString = content
	paragraphs := strings.Split(content, "\n")

	text := ""
	height := 0
	for _, paragraph := range paragraphs {
		t, h := t.wrap(width, paragraph)
		text += "\n" + t
		height += h + 1
	}

	t.content = t.toCellArray(text, false)
	t.Rect = NewRect(height, width)
	return t, height
}

//returns a formatted title
func NewTitleText(width int, content string) *Text {

	t := new(Text)
	t.asString = content
	t.Rect = NewRect(1, width)
	t.content = t.toCellArray(t.horizTruncate(content, width), true)
	return t

}

func (t *Text) horizTruncate(s string, width int) string {
	if len(s) < width {
		return s
	} else {
		return s[0:width]
	}
}

func (t *Text) vertTruncate(Text string) string {

	lines := strings.Fields(strings.Trim(Text, "\n"))

	if len(lines) <= t.H() {
		return Text
	} else {
		trimmed := ""

		for i := 0; i < t.H(); i++ {
			trimmed += lines[i] + "\n"
		}
		return trimmed + "\n" + "..."
	}

}

func (t *Text) wrap(width int, content string) (string, int) {

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

func (t *Text) toCellArray(s string, isCentered bool) []c.Cell {
	lines := strings.Split(s, "\n")
	output := make([]c.Cell, 0)
	if isCentered {
		for y, line := range lines {
			offset := (t.W() - len(line)) / 2
			for x, letter := range line {
				c := c.Cell{X: x + offset, Y: y, Content: letter}
				output = append(output, c)
			}

		}
	} else {
		for y, line := range lines {
			for x, letter := range line {
				c := c.Cell{X: x, Y: y, Content: letter}
				output = append(output, c)
			}

		}
	}
	return output
}
