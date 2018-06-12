package main

import (
	"github.com/nowakf/tview"
	"math"
	"strings"
)

const column_count = 3

func News() *news {
	return &news{
		stories: make([]story, 0),
	}
}

type news struct {
	columns [column_count]column
	stories []story
}

var testStories = []story{
	story{
		title:   "cannibal horror",
		content: "The nation is reeling in shock from the revelations that President Hogarth is a CANNIBAL. Turn to page 4 for the full story!",
	},
	story{
		title:   "flight 401 mystery",
		content: "Rescuers continue the search for the missing jumbo jet, which disappeared three weeks ago while flying a pioneering route over the antarctic.",
	},
	story{
		title:   "my secret",
		content: "EXCLUSIVE from hearthrob Joni Sams: her way to stay fresh and vibrant, no matter the situation!",
	},
	story{
		title:   "Bob Dylan, folk singer, dies",
		content: "Beloved musician and composer Bob Dylan passed away on saturday, surrounded by his familly and friends.",
	},
	story{
		title:   "cannibal horror",
		content: "The nation is reeling in shock from the revelations that President Hogarth is a CANNIBAL. Turn to page 4 for the full story!",
	},
	story{
		title:   "flight 401 mystery",
		content: "Rescuers continue the search for the missing jumbo jet, which disappeared three weeks ago while flying a pioneering route over the antarctic.",
	},
	story{
		title:   "my secret",
		content: "EXCLUSIVE from hearthrob Joni Sams: her way to stay fresh and vibrant, no matter the situation!",
	},
	story{
		title:   "police scandal coverup",
		content: "",
	},
}

func (n *news) UI(nextMode func()) (title string, content tview.Primitive) {

	frame := tview.NewFlex()

	frame.SetDirection(tview.FlexColumn)

	n.stories = testStories

	for i := 0; i < column_count; i++ {
		n.columns[i] = column{0, tview.NewFlex().SetDirection(tview.FlexRow)}
		frame.AddItem(n.columns[i].Flex, 0, 1, false)
	}

	for _, story := range n.stories {
		t := tview.NewTextView().SetText(n.formatStory(story)).SetWordWrap(true)
		t.SetBorder(true)
		n.AddItem(t)

	}
	for _, col := range n.columns {
		println(col.count)
	}

	return "News", frame
}
func (n *news) Update() {}

func (n *news) Count() int {
	return 0
}
func (n *news) Get(s string) (handle string, content map[string]string, keyvalue map[string]int) {
	return "", nil, nil
}

type column struct {
	count int
	*tview.Flex
}

func (n *news) formatStory(s story) string {
	return strings.ToUpper(s.title) + "\n" + s.content
}
func (n *news) AddItem(p tview.Primitive) {
	smallestCount := math.MaxInt64
	var smallestColumn int
	for i, column := range n.columns {
		if column.count < smallestCount {
			smallestCount = column.count
			smallestColumn = i

		}

	}

	if n.columns[smallestColumn].Flex != nil {
		n.columns[smallestColumn].AddItem(p, 0, 1, false)
		n.columns[smallestColumn].count++
	}

}

type story struct {
	title   string
	content string
}
