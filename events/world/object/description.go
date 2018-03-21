package object

//this holds the descriptions of specific portions of memory.
var objects = map[category]map[index]*object{
	body: map[index]*object{
		strength:  &object{"strength", ""},
		willpower: &object{"willpower", ""},
	},
	mind:     map[index]*object{},
	social:   map[index]*object{},
	identity: map[index]*object{},
}

type object struct {
	title       string
	description string
}

func (o *object) String() string {
	return o.title
}
func (o *object) Describe() string {
	return o.description
}
