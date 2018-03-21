package world

type Record struct {
	origin *Person
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

func (e *Record) Origin() *Person {
	return e.origin
}
func (e *Record) Do(index int) {
	e.origin.Queue((*e.Options)[index].url)
}

type option struct {
	title string
	url   string
	index int
}

func (o *option) Title() string { return o.title }
func (o *option) Index() int    { return o.index }

type Options []option

func (o *Options) Insert(title, url string) {
	(*o) = append(*o, option{title, url, len(*o)})
}
