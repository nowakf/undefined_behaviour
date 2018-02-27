package data

import "testing"

func TestLoad(t *testing.T) {
	s := load()
	for _, e := range s {

		if e.Event_url == "" {
			t.Fatalf("got no event url")
		}
		t.Log(e.Event_url)

		if e.Title == "" {
			t.Fatalf("got no title, %v", e)
		}
		t.Log(e.Title)

		if e.Content == "" {
			t.Fatalf("got no content, %v", e)
		}
		if e.Consequences == nil {
			t.Fatalf("got no consequences")
		}
	}

}
