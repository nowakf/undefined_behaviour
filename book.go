package main

import (
	"fmt"
	"github.com/nowakf/tview"
)

func Book() *book {
	return &book{
		names:       tview.NewList(),
		description: tview.NewTextView(),
		newSpells:   make(map[string]*spell, 0),
		oldSpells:   make(map[string]*spell, 0),
	}
}

type book struct {
	names       *tview.List
	description *tview.TextView
	newSpells   map[string]*spell
	oldSpells   map[string]*spell
}

var testSpells = map[string]*spell{
	"the crooked path": &spell{
		handle: "the crooked path",
		content: map[string]string{
			"short_description": "a spell to conceal the caster",
			"long_description":  `The crooked path is invaluable to anyone who wishes to pass unnoticed. It does not grant true invisibility, but rather acts to make its subject forgettable, causes paper or electronic records of their presence to be misplaced or deleted, and allows exceptional acts to be overlooked.`,
		},
		keyValue: map[string]int{
			"Cost": 25,
			"XP":   125,
		},
	},
	"four broken arrows": &spell{
		handle: "four broken arrows",
		content: map[string]string{
			"short_description": "a spell to strike from afar",
			"long_description":  `Four Broken Arrows is a spell designed to inflict harm from a distance. The target will lose all sense of caution, and engage in increasingly risky behaviour until they perish.`,
		},
		keyValue: map[string]int{
			"Cost": 15,
			"XP":   43,
		},
	},
}

func (b *book) UI(nextMode func()) (title string, content tview.Primitive) {

	b.names.
		SetBorder(true)
	b.description.
		SetWordWrap(true).
		SetBorder(true)

	prim := tview.NewFlex().
		AddItem(b.names, 0, 1, true).
		AddItem(b.description, 0, 2, false)

	b.newSpells = testSpells

	return "Book", prim
}

func (b *book) Count() int {
	return len(b.newSpells)
}

func (b *book) Update() {
	if b.Count() > 0 {
		for i := len(b.oldSpells); i < len(b.oldSpells)+b.Count(); i++ {

			spell := b.Get("")
			if spell == nil {
				println("trying to read non-existent spell")
				continue
			}

			b.names.AddItem(spell.handle, spell.content["short_description"], rune(i+97), func() {

				text := "\n"

				text += spell.content["long_description"]

				text += "\n\n"

				text += chapterBreak(b.description.Box)

				text += "\n\n"

				for key, value := range spell.keyValue {
					text += fmt.Sprintf("%s : %d \n", key, value)
				}

				b.description.SetText(text)

			})

		}
	}
}

//get either returns the specified element, or one at random
func (b *book) Get(s string) *spell {

	//Update spells here
	if s == "" {
		for handle, spell := range b.newSpells {
			b.oldSpells[handle] = spell
			delete(b.newSpells, handle)
			return spell
		}
	} else {
		spell, ok := b.newSpells[s]
		if ok {
			b.oldSpells[s] = spell
			delete(b.newSpells, s)
			return spell
		}
		spell, ok = b.oldSpells[s]
		if ok {
			return spell
		}

		if !ok {
			return nil
		}

	}
	return nil

}

type spell struct {
	handle   string
	content  map[string]string
	keyValue map[string]int
}
