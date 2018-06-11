package object

//this holds the descriptions of specific portions of memory.
var descriptions = map[category]map[index]description{
	body: map[index]description{
		strength:     description{"strength", ""},
		intelligence: description{"intelligence", ""},
		willpower:    description{"willpower", ""},
	},
	mind:     map[index]description{},
	social:   map[index]description{},
	identity: map[index]description{},
}

type description struct {
	title string
	blurb string
}

func (o description) Title() string {
	return o.title
}
func (o description) Blurb() string {
	return o.blurb
}
