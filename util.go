package main

import "github.com/nowakf/tview"

const chapter_break = "\u00a7"

func chapterBreak(p *tview.Box) string {
	s := ""
	_, _, width, _ := p.GetInnerRect()
	for i := 0; i < width/2; i++ {
		s += " "
	}
	s += chapter_break

	return s

}
