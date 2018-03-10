package world

type Record struct {
	origin *person
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

func (e Record) Origin() *person {
	return e.origin
}

type option struct {
	Title string
	Url   string
}

func (o *option) Do(p *person) {
}

type Options []option

func (o *Options) Insert(title, url string) {
	(*o) = append(*o, option{title, url})
}
