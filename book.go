package main

import (
	"fmt"

	"github.com/nowakf/tview"
	"github.com/nowakf/ubcell"
)

func Book() *book {
	return &book{
		names:  tview.NewList(),
		pages:  tview.NewPages(),
		plain:  tview.NewTextView(),
		table:  tview.NewTable(),
		spells: make(map[string]*spell),
	}
}

type book struct {
	names  *tview.List
	pages  *tview.Pages
	plain  *tview.TextView
	table  *tview.Table
	spells map[string]*spell
}

var testSpells = map[string]*spell{
	"the crooked path": &spell{
		handle:            "the crooked path",
		short_description: "a spell to conceal the caster",
		long_description:  `The crooked path is invaluable to anyone who wishes to pass unnoticed. It does not grant true invisibility, but rather acts to make its subject forgettable, causes paper or electronic records of their presence to be misplaced or deleted, and allows exceptional acts to be overlooked.`,
		keyValue: map[string]int{
			"Cost": 25,
			"XP":   125,
			"HP":   15,
		},
	},
	"four broken arrows": &spell{
		handle:            "four broken arrows",
		short_description: "a spell to strike from afar",
		long_description:  `Four Broken Arrows is a spell that will inflict harm from a distance. The target will lose all sense of caution, and engage in increasingly risky behaviour until they perish.`,
		keyValue: map[string]int{
			"Cost": 15,
			"XP":   43,
			"HP":   4,
		},
	},
	"the unbroken pot": &spell{
		handle:            "the unbroken pot",
		short_description: "a spell to bring luck",
		long_description:  `The unbroken pot is a spell that will bring luck in all endeavours. Unfortunately, it also vastly increases the appetite of the user, forcing them to gorge themselves almost constantly, and giving them an irresistible taste for raw meat.`,
		keyValue: map[string]int{
			"Cost": 15,
			"XP":   13,
			"HP":   40,
		},
	},
}

func (b *book) UI(nextMode func()) (title string, content tview.Primitive) {

	b.names.
		SetBorder(true)

	b.plain.SetWordWrap(true).SetBorder(true)

	b.table.SetCellSimple(0, 0, "names").
		SetCellSimple(0, 1, "XP").
		SetCellSimple(0, 2, "Cost").
		SetCellSimple(0, 3, "HP").
		SetBorders(true)

	b.pages = tview.NewPages().
		AddPage("text", b.plain, true, false).
		AddPage("table", b.table, true, true)

	prim := tview.NewFlex().
		AddItem(b.names, 0, 1, true).
		AddItem(b.pages, 0, 2, false)

	prim.SetDrawFunc(func(screen ubcell.Screen, x, y, width, height int) (int, int, int, int) {
		app.SetFocus(b.names)
		return x, y, width, height
	})
	b.spells = testSpells

	b.names.AddItem("overview", "", 'a', func() {
		b.pages.ShowPage("table")
	}).ShowSecondaryText(false)

	b.Update()

	return "Book", prim
}

func (b *book) Count() int {
	return 0
}

func (b *book) Update() {

	i := 0
	for _, s := range b.spells {
		i++
		b.table.SetCellSimple(i, 0, s.handle).
			SetCellSimple(i, 1, fmt.Sprintf("%d", s.keyValue["XP"])).
			SetCellSimple(i, 2, fmt.Sprintf("%d", s.keyValue["Cost"])).
			SetCellSimple(i, 3, fmt.Sprintf("%d", s.keyValue["HP"]))

		selectedFunc := func(s *spell) func() {
			return func() {
				b.pages.SwitchToPage("text")
				b.plain.SetText(b.formatSpell(s))
			}
		}(s)

		b.names.AddItem(s.handle, s.short_description, rune(i+97), selectedFunc)
	}

}
func (b *book) formatSpell(s *spell) string {
	text := "\n"

	text += s.long_description

	text += "\n\n"

	text += chapterBreak(b.pages.Box)

	text += "\n\n"

	for key, value := range s.keyValue {
		text += fmt.Sprintf("%s : %d \n", key, value)
	}
	return text
}

//get either returns the specified element, or one at random
func (b *book) Get(s string) *spell {
	return b.spells[s]
}

type spell struct {
	handle            string
	short_description string
	long_description  string
	keyValue          map[string]int
}
