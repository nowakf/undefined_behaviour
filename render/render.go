package render

import (
	"fmt"
	"math"

	c "ub/common"
	"ub/data"
	"ub/ui"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type render struct {
	Window         *pixelgl.Window
	fonts          map[pixel.RGBA]*text.Text
	glyphH, glyphW float64
}

func New(d *data.FontLoader, fontSize float64) *render {
	r := new(render)
	r.Window = newWindow()

	r.fonts = d.Fonts(fontSize, c.Colors...)

	r.glyphW, r.glyphH = r.getIncrement()

	return r
}

//draws cells to the screen
//change: [][][]rune-
// [] refers to a layer. so you send all the background layers, then the foreground layers
func (r *render) Update(stack []ui.Layer) {

	r.clear()

	for _, layer := range stack {

		color := layer.Color()
		_, ok := r.fonts[color]
		if !ok {
			println(fmt.Sprintf("there is no %v in the fonts", layer.Color()))
			color = c.White
		}

		for _, cell := range layer.Content() {
			r.fonts[color].Dot = pixel.V(float64(cell.X)*r.glyphW, r.Window.Bounds().H()-float64(cell.Y)*r.glyphH)
			r.fonts[color].WriteRune(cell.Letter)
		}
		r.fonts[color].Draw(r.Window, pixel.IM)
	}

}

//clears the text
func (r *render) clear() {
	screenColor := pixel.RGB(0.3, 0.3, 0.3)
	for _, font := range r.fonts {
		font.Clear()
	}
	r.Window.Clear(screenColor)
}

// gets the height and width of the Window
func (r *render) Stats() (int, int) {
	return r.getCellCount(r.Window)
}

//gets the step size
func (r *render) getIncrement() (float64, float64) {

	_, ok := r.fonts[c.White]
	if !ok {
		for _, color := range c.Colors {
			_, ok = r.fonts[color]
			if ok {
				break
			} else {
				fmt.Printf("%v is not defined in colors", color)
			}
		}
	}

	wi := r.fonts[c.White].BoundsOf("S").W()
	hi := r.fonts[c.White].BoundsOf("S").H()

	r.glyphW, r.glyphH = wi, hi

	return wi, hi
}

//gets the height and width of the Window, in cells
func (r *render) getCellCount(w *pixelgl.Window) (int, int) {

	bounds := w.Bounds()

	glyphwidth, glyphheight := r.getIncrement()

	h := bounds.H() / glyphheight
	wi := bounds.W() / glyphwidth

	return int(math.Floor(h)), int(math.Floor(wi))
}
