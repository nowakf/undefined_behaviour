package events

type Record struct {
	origin *actor
	url    string
	title  string
	body   string
	Depth  int
	*Options
}

func (e *Record) Title() string {
	return e.title
}

func (e *Record) Body() string {
	return e.body
}

func (e Record) Origin() *actor {
	return e.origin
}

type option struct {
	Title string
	Url   string
}

func (o *option) Do(p *Player) {
}

type Options []option

func (o *Options) Insert(title, url string) {
	(*o) = append(*o, option{title, url})
}
