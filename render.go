package main

import (
	"io/ioutil"
	"math"
	"os"

	c "ub/common"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"unicode"
)

type render struct {
	atlas          *text.Atlas
	window         *pixelgl.Window
	face           font.Face
	text           *text.Text
	glyphH, glyphW float64
}

var (
	backgroundColour = pixel.RGB(0.3, 0.3, 0.3)
)

func newRender(w *pixelgl.Window) *render {
	r := new(render)
	r.window = w

	face, err := r.loadTTF("./assets/fonts/DejaVuSansMono.ttf", 20)
	if err != nil {
		panic(err)
	}
	r.face = face

	r.atlas = text.NewAtlas(r.face, text.RangeTable(unicode.Latin), text.RangeTable(unicode.Space), text.RangeTable(unicode.Po), text.RangeTable(unicode.S), text.ASCII)

	r.text = text.New(pixel.V(0, 0), r.atlas)

	r.text.Color = pixel.RGB(0.8, 0.8, 0.8)

	r.glyphW, r.glyphH = r.getIncrement()

	return r
}

//draws cells to the screen
func (r *render) update(cells []c.Cell) {
	r.clear()
	for _, cell := range cells {
		//	println(string(cell.Content), cell.Content)
		r.drawCell(cell.Content, float64(cell.X)*r.glyphW, float64(cell.Y)*r.glyphH)

	}
	r.text.Draw(r.window, pixel.IM)
}

//clears the text
func (r *render) clear() {
	r.text.Clear()
	r.window.Clear(backgroundColour)
}

// gets the height and width of the window
func (r *render) Stats() (int, int) {
	return r.getCellCount(r.window)
}

//draws a letter at a specific position
func (r *render) drawCell(content rune, xpos, ypos float64) {
	r.text.Dot = pixel.V(xpos, r.window.Bounds().H()-ypos)
	r.text.WriteRune(content)
}

//gets the step size
func (r *render) getIncrement() (float64, float64) {

	wi := r.text.BoundsOf("S").W()
	hi := r.text.BoundsOf("S").H()

	r.glyphW, r.glyphH = wi, hi

	return wi, hi
}

//gets the height and width of the window, in cells
func (r *render) getCellCount(w *pixelgl.Window) (int, int) {

	bounds := w.Bounds()

	glyphwidth, glyphheight := r.getIncrement()

	h := bounds.H() / glyphheight
	wi := bounds.W() / glyphwidth

	return int(math.Floor(h)), int(math.Floor(wi))
}

func (r *render) loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
