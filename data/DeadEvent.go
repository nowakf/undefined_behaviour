package data

type DeadEvent struct {
	event_url    string
	etype        string
	content      string
	consequences []string
}

func (d *DeadEvent) Event_Url() string {
	return d.event_url
}
func (d *DeadEvent) Etype() string {
	return d.etype
}
func (d *DeadEvent) Content() string {
	return d.content
}
func (d *DeadEvent) Consequences() []string {
	return d.consequences
}
