package events

type event struct {
	event_url    string
	etype        etype
	content      string
	consequences []string
}

func (e *event) Event_Url() string {
	return e.event_url
}
func (e *event) Content() string {
	return e.content
}
func (e *event) Consequences() []string {
	return e.consequences
}
func (e *event) Etype() etype {
	return e.etype
}
