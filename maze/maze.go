package maze

import (
	"errors"
	"image/color"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/nowakf/tview"
	"github.com/nowakf/undefined_behaviour/maze/cell"
	termbox "github.com/nsf/termbox-go"
)

type maze struct {
	width          int
	fresh          []cell.Cell
	old            []cell.Cell
	activeEntities []active
	info           string
	stopped        bool
}

//Accuracy sets how many slots should be considered.
//a higher value means better, slower results
const Accuracy = 8

type active interface {
	Update()
	Draw()
}

const (
	left = iota
	right
	top
	bottom
)

func (m *maze) Area() int {
	return len(m.fresh)
}

func (m *maze) Nearby(source vector, maxDist int, desired key, feared key) []vector {

	x, y := source.x, source.y
	bounds := [4]int{
		left:   y*m.width + (x - maxDist),
		right:  y*m.width + (x + maxDist),
		top:    (y-maxDist)*m.width + (x + maxDist),
		bottom: (y+maxDist)*m.width + x,
	}

	for i := range bounds {
		if bounds[i] < 0 {
			bounds[i] = 0
		} else if bounds[i] >= len(m.old) {
			bounds[i] = len(m.old) - 1
		}
	}

	type dir struct {
		float64 //desirableness
		vector  //where the object is
	}

	getNearby := func(obs object) (results [Accuracy]dir) {

		for x := bounds[left]; x < bounds[right]; x++ {
			for y := bounds[top]; y < bounds[bottom]; y++ {

				thingsOfInterest := m.old[y*m.width+x].Some(desired)

				diff := vector{x, y}.sub(s)

				interestQuotient := float64(maxDist) / math.Sqrt(float64(diff.abs().x*diff.abs().x+diff.abs().y*diff.abs().y)) * float64(thingsOfInterest.Count())
				//treat the slots as a 3x3 array

				//slot := 5 + diff.unit().y*3 + diff.unit.x

				slot := (y*m.width + x) % Accuracy

				if interestQuotient > results[slot].float64 {
					results[slot] = dir{interestQuotient, loc}
				}
			}
		}
	}

	wanted := getNearby(desired)
	hated := getNearby(feared)

	for i := range wanted {
		wanted[i] -= hated[i]
	}
	blurred := wanted
	for i := range blurred {
		averaged[i] = (wanted[(Accuracy-(i-1))%Accuracy] + wanted[(Accuracy+(i+1))%Accuracy]) / 2
	}

	out := make([]vector, count)

	sort.Slice(blurred, func(i, j int) bool {
		return blurred[i] < blurred[j]
	})

	for i := 0; i < count; i++ {
		out[i] = distances[i].vector
	}
	return out

}

//pathfinding from source to sink - obstacle is a superset of all non-passable objects
func (m *maze) Path(source vector, sink vector, obstacle key) (path vlist) {

	visited := make(map[vector]vector)
	frontier := []vector{source}
	increment := 0

	var current vector

	for current != sink && len(frontier)-increment > 0 {
		m.Set(current, cloud)

		current = frontier[increment]
		increment++

		for _, neighbor := range m.neighbors(current, passable) {
			if _, ok := visited[neighbor]; !ok {

				frontier = append(frontier, neighbor)
				//mark as touched
				visited[neighbor] = current

				if neighbor == sink {
					current = neighbor
					break
				}
			}
		}

	}

	if current == sink {
		for current != source {
			path.PushFront(current)
			current = visited[current]
		}
		path.PushFront(current)

		return path
	}

	return vlist{}

	//panic("no path")

}

func (m *maze) neighbors(of vector, passable key) []vector {
	neighbors := make([]vector, 0)

	check := func(at vector) bool {
		return at.x < m.width && at.x >= 0 && at.y < len(m.old)/m.width && at.y >= 0 && !m.old[at.y*m.width+at.x].Has(passable)
	}
	var n vector
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if check(n) {
				neighbors = append(neighbors, vector{x, y})
			}
		}
	}
	return neighbors

}

func (m *maze) Convert(from string) error {

	lines := strings.Split(from, "\n")

	rand.Seed(time.Now().UnixNano())

	if len(lines) < 1 {
		return errors.New("0 length input!")
	}

	m.width = len(lines[0])

	for y, line := range lines {
		if len(line) != m.width {
			return errors.New("uneven widths!")
		}

		m.old = append(m.old, make([]cell.Cell, len(line))...)
		m.fresh = append(m.fresh, make([]cell.Cell, len(line))...)

		for x, letter := range line {
			switch letter {
			case ' ':
			case '#':
				m.fresh[y*m.width+x].Set(wall)
			case 'X':
				m.fresh[y*m.width+x].Set(body)
				m.activeEntities = append(m.activeEntities, Crawler(m, x, y))
			case '!':
				m.fresh[y*m.width+y].Set(food)
			default:
				return errors.New("unknown symbol " + string(letter))

			}

		}
	}

	return nil

}

func (m *maze) Update() {
	for i := range m.Active {
		m.Active[i].Update()
	}
}

func (m *maze) CelltoRune(c cell.Cell, offset int) rune {
	switch {
	case c.Has(Wall):
		return block
	case c.Has(Food):
		return fmark
	case c.Has(Body):
		return crmark
	case c.Has(Tendril):
		return tendrils[m.orient(Tendril|Body, offset)]
	case c.Has(Cloud):
		return space
	case c.Some(All) == 0:
		return space
	default:
		return '?'
	}
}

func (m *maze) orient(q key, offset int) uint8 {
	//   x
	// x   x
	//   x
	//16 possible combinations:
	var key uint8

	dirs := []uint{
		uint(offset - m.width), //up
		uint(offset + m.width), //down
		uint(offset + 1),       //right
		uint(offset - 1),       //left
	}

	for i, dir := range dirs {
		if dir >= 0 && dir < uint(len(m.old)) {
			if m.old[dir].Some(q) != 0 {
				key |= 1 << uint(i)
			}
		}
	}
	return key

}

func (m *maze) Draw(canvas *tview.Canvas) {

	for i := range m.activeEntities {
		m.activeEntities[i].Draw()
	}

	var empty cell.Cell

	for i, cel := range m.old {

		r, f, b := m.DrawCell(cel, i)

		canvas.SetCell(i%m.width, i/m.width, r, f, b)

		if cel.Some(Wall|Food|Cloud) != 0 {
			m.fresh[i] = cel
		}

	}
	for i := range m.old {
		m.old[i] = empty
	}

	y := 0
	x := m.width
	for _, letter := range m.info {
		if letter == '\n' {
			x = m.width
			y++
		}
		x++
		canvas.SetCell(x, y, letter, termbox.ColorDefault, termbox.ColorBlack)
	}
	m.fresh, m.old = m.old, m.fresh

}

//TODO change
func (m *maze) DrawCell(c cell.Cell, offset int) (letter rune, fg, bg color.RGBA) {
	var obj int
	if obj = c.Read(Wall); obj >= 0 {
		letter = Walls[obj].sigil
		fg = Walls[obj].color
	} else if obj = c.Read(Body); obj >= 0 {
		letter = BodyStates[obj].sigil
		fg = BodyStates[obj].color
	} else if obj = c.Read(Tendril); obj >= 0 {
		letter = Tendrils[m.orient(Body|Tendril, offset)].sigil
		fg = Tendrils[m.orient(Body|Tendril, offset)].color
	} else if obj = c.Read(Food); obj >= 0 {
		letter = FoodTypes[obj].sigil
		fg = FoodTypes[obj].color
	} else {
		letter = space
	}
	if obj = c.Read(Cloud); obj >= 0 {
		bg = Clouds[obj].color
	} else {
		bg = Colors.DefaultBackground
	}

}

func (m *maze) Query(x, y int) {
	if y*m.width+x < len(m.fresh) {
		m.info = m.Read(vector{x, y}).Enumerate()
	}
}
func (m *maze) updateFood() {
	if rand.Int()%5 == 0 {
		m.ScatterFood(0)
		//if rand.Int()%10 == 0 {
		//	m.ScatterCrawlers(0)
		//}
	}
	temp := make([]vector, 0, len(m.food))
	for _, f := range m.food {
		if m.old[f.y*m.width+f.x].Has(food) {
			temp = append(temp, f)
		}
	}
	m.food = temp

}
func (m *maze) Set(at vector, o key, val int) {
	m.fresh[at.y*m.width+at.x].Set(o, val)
}
func (m *maze) Read(at vector) cell.Cell {
	return m.old[at.y*m.width+at.x]
}
func (m *maze) set(x, y int, r rune) {
	if x >= m.width || x < 0 || y >= len(m.fresh)/m.width || y < 0 {
		return
	}
	//TODO
	//termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorBlack)
}
