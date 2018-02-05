package events

import "testing"

func TestLoad(t *testing.T) {
	s := load()

	if s.event_url == "" {
		t.Fatalf("got no event url")
	}
	t.Log(s.event_url)

	if s.title == "" {
		t.Fatalf("got no title, %v", s)
	}
	t.Log(s.title)

	if s.content == "" {
		t.Fatalf("got no content, %v", s)
	}
	if s.consequences == nil {
		t.Fatalf("got no consequences")
	}

}
