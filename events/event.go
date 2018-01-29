package events

type event struct {
	event_url    string
	etype        string
	content      string
	consequences []string
}
