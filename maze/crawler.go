package maze

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/nowakf/undefined_behaviour/maze/cell"
	"golang.org/x/image/colornames"
)

type crawler struct {
	loc        vector
	arms       []arm
	prim       int
	directions [8]dir

	desired     cell.Key
	feared      cell.Key
	traversable cell.Key

	stats
	duration int

	state crawlerState

	enviro *maze
}

type stats struct {
	armMaxLength int
	resilience   int
	health       int
	intelligence int
	vision       int
}

type dir struct {
	path   vlist
	desire float64
}

type bstate int

func Crawler(m *maze, x, y int) (c crawler) {
	return crawler{
		loc:        vector{x, y},
		arms:       make([]arm, 5),
		prim:       -1,
		directions: [8]dir{},

		desired:     cell.Food,
		feared:      cell.Body,
		traversable: cell.Food | cell.Tendril | cell.Body,

		stats: stats{
			armMaxLength: 10,
			resilience:   10,
			health:       10,
			intelligence: 10,
			vision:       10,
		},
		duration: -1,
		state:    Wandering,
		enviro:   m,
	}

}

func (c *crawler) Update() {
	switch c.state {
	case Unconscious:
		if rand.Int()%c.resilience == 0 {
			c.state = Stunned
		}
	case Stunned:
		if c.duration > 0 {
			c.duration--
			return
		}
		if c.duration == 0 {
			c.duration = -1
			c.state = Wandering
			return
		}
		if c.duration < 0 {
			c.duration = c.resilience
			return
		}

	case Wandering:
		c.assign()
		c.prim = c.poll()
		for i := range c.arms {
			c.arms[i].updateLength()
		}
		if c.prim != -1 {
			c.state = Reaching
		}

	case Reaching:
		for i := range c.arms {
			c.arms[i].updateLength()
		}

		if c.arms[c.prim].segments.Len() > c.armMaxLength || c.arms[c.prim].segments.Len() == c.arms[c.prim].path.Len() {
			c.state = Chasing
		}

	case Chasing:
		for i := range c.arms {
			if i != c.prim {
				c.arms[i].state = retracting
			}
			c.arms[i].updateLength()
		}
		if c.arms[c.prim].segments.Len() == 0 || c.arms[c.prim].path.Len() == 0 {
			c.state = Eating
			return
		}
		c.loc = c.arms[c.prim].path.PopFront()
		c.arms[c.prim].segments.PopFront()
		c.move(c.loc)

	case Eating:
		c.consume()
		c.prim = -1
		c.state = Wandering
	}

}

func (c *crawler) Orient() {

}

func (c *crawler) Vision() {
	//you need to implement some kind of complex FOV later.

	topLeft := (c.loc.y+c.stats.vision)*c.enviro.width + (c.loc.x + c.stats.vision)
	topRight := (c.loc.y-c.stats.vision)*c.enviro.width + (c.loc.x - c.stats.vision)

	if topLeft < 0 {
		topLeft = 0
	}
	if topRight >= len(c.enviro.old) {
		topRight = len(c.enviro.old) - 1
	}
	//01234
	//56789
	//
	for i := topLeft; i < topRight; i++ {
		relativeX := (i % c.enviro.width) - c.loc.x
		relativeY := (i / c.enviro.width) - c.loc.y
		//there is definitely a more elegant way to do this

		switch {
		case relativeY >= 2*relativeX && -relativeX <= 2*relativeY:
			//top
		case relativeY <= 2*relativeX && relativeX <= 2*relativeY:
			//topright
		case relativeY >= 2*relativeX && relativeX >= -2*relativeY:
			//right
		case -relativeY >= 2*relativeX && relativeX <= -2*relativeY:
			//bottom right
		case -relativeY <= 2*relativeX && relativeX <= -2*relativeY:
			//bottom

		}

	}
}

func (c *crawler) getPaths(count int) []vlist {
	nearbyFood := c.enviro.Nearby(c.loc, count, c.desired, c.feared)
	out := make([]vlist, len(nearbyFood))
	for i, v := range nearbyFood {
		out[i] = c.enviro.Path(c.loc, v, c.traversable)
	}
	return out
}

func (c *crawler) Draw() {
	for i := range c.arms {
		c.arms[i].draw(c.enviro)
	}
	c.enviro.Set(c.loc, cell.Body, int(c.state))
}

//TODO make it assign multiple paths
func (c *crawler) assign() {

	var list []int

	for i := range c.arms {
		path := c.arms[i].path
		if path.Len() <= 0 || !c.enviro.Read(*path.I(path.Len() - 1)).Has(cell.Food) {
			list = append(list, i)

		}
	}

	paths := c.getPaths(len(list))

	for i, path := range paths {
		c.arms[list[i]].assign(path)
	}
}

func (c *crawler) move(newLoc vector) {
	for i := range c.arms {
		if i != c.prim {
			c.arms[i].move(newLoc)
		}
	}
}

func (c *crawler) consume() {
	c.enviro.Erase(c.loc, cell.Food)
}

func (c *crawler) poll() int {
	pathLength := math.MaxInt64
	prim := -1
	for i := range c.arms {
		plen := c.arms[i].path.Len()
		if pathLength > plen && rand.Int()%3 == 0 {
			prim = i
			pathLength = plen
		}
	}
	return prim
}

const (
	Wandering crawlerState = iota
	Stunned
	Unconscious
	Chasing
	Reaching
	Eating
)

type crawlerState int

type bodyState struct {
	name  string
	sigil rune
	color color.RGBA
}

var BodyStates = [...]bodyState{
	Wandering: bodyState{
		name:  "Wandering",
		sigil: 'w',
		color: colornames.Grey,
	},
	Stunned: bodyState{
		name:  "Stunned",
		sigil: 'x',
		color: colornames.Cornflowerblue,
	},
	Unconscious: bodyState{
		name:  "Unconscious",
		sigil: 'x',
		color: colornames.Grey,
	},
	Chasing: bodyState{
		name:  "Chasing",
		sigil: 'x',
		color: colornames.Red,
	},
	Reaching: bodyState{
		name:  "Reaching",
		sigil: 'x',
		color: colornames.Palevioletred,
	},
	Eating: bodyState{
		name:  "Eating",
		sigil: 'x',
		color: colornames.Lemonchiffon,
	},
}
