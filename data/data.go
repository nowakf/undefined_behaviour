package data

//data handles loading files. Will probably want something more sophisticated later.
import (
	"io/ioutil"
	"os"
	"unicode"

	"github.com/golang/freetype/truetype"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font"
)

type data struct {
}
type Save struct{}

func NewData() *data {
	return new(data)
}

func (d *data) LoadSave() *Save {
	return nil
}

type FontLoader struct {
}

//returns as many fonts as you give it colors. Useful, I know.
func (d FontLoader) Fonts(fontSize float64, colors ...pixel.RGBA) map[pixel.RGBA]*text.Text {

	face, err := d.loadTTF("./assets/fonts/DejaVuSansMono.ttf", fontSize)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face,
		text.RangeTable(unicode.Space), //for spaces
		//text.RangeTable(unicode.Po),    //for the '‾' mark
		text.RangeTable(unicode.S), //for the '█' mark,
		text.ASCII,                 //for misc stuff
	)

	fonts := make(map[pixel.RGBA]*text.Text)
	for _, color := range colors {
		font := text.New(pixel.V(0, 0), atlas)
		font.Color = color
		fonts[color] = font
	}

	return fonts

}

func (d FontLoader) loadTTF(path string, size float64) (font.Face, error) {
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
