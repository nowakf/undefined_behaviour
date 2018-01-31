package main

import (
	"fmt"
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
	fonts          map[pixel.RGBA]*text.Text
	glyphH, glyphW float64
}

func newRender(w *pixelgl.Window) *render {
	r := new(render)
	r.window = w

	face, err := r.loadTTF("./assets/fonts/DejaVuSansMono.ttf", 20)
	if err != nil {
		panic(err)
	}
	r.face = face

	r.atlas = text.NewAtlas(r.face, text.RangeTable(unicode.Latin), text.RangeTable(unicode.Space), text.RangeTable(unicode.Po), text.RangeTable(unicode.S), text.ASCII)

	r.fonts = make(map[pixel.RGBA]*text.Text)

	for _, color := range c.Colors {
		font := text.New(pixel.V(0, 0), r.atlas)
		font.Color = color
		r.fonts[color] = font
	}

	r.glyphW, r.glyphH = r.getIncrement()

	return r
}

//draws cells to the screen
func (r *render) update(cells []c.Cell) {
	r.clear()
	for _, cell := range cells {

		r.drawCell(cell.Letter, float64(cell.X)*r.glyphW, float64(cell.Y)*r.glyphH, cell.Background, cell.Foreground)

	}
	//TODO: merge this into one draw, if possible
	for _, font := range r.fonts {
		font.Draw(r.window, pixel.IM)
	}
}

//clears the text
func (r *render) clear() {
	screenColor := pixel.RGB(0.3, 0.3, 0.3)
	for _, font := range r.fonts {
		font.Clear()
	}
	r.window.Clear(screenColor)
}

// gets the height and width of the window
func (r *render) Stats() (int, int) {
	return r.getCellCount(r.window)
}

//draws a cell at a specific position
func (r *render) drawCell(letter rune, xpos, ypos float64, foreground pixel.RGBA, background pixel.RGBA) {
	r.drawBack(xpos, ypos, background)
	r.drawText(letter, xpos, ypos, foreground)
}

//draws the text
func (r *render) drawText(letter rune, xpos, ypos float64, color pixel.RGBA) {

	_, ok := r.fonts[color]
	if !ok {
		println(fmt.Sprintf("there is no %v in the fonts", color))
	}

	r.fonts[color].Dot = pixel.V(xpos, r.window.Bounds().H()-ypos)
	r.fonts[color].WriteRune(letter)
}

//TODO: draw the background
func (r *render) drawBack(xpos, ypos float64, color pixel.RGBA) {
}

//gets the step size
func (r *render) getIncrement() (float64, float64) {

	wi := r.fonts[c.White].BoundsOf("S").W()
	hi := r.fonts[c.White].BoundsOf("S").H()

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
