package world

const (
	HOME placeEnum = iota
	HOUSE
	BAR
	PARK
	WORK_OFFICE
	WORK_FACTORY
	STORE
	GUN_STORE
	HOSPITAL
	CITY_HALL
	POLICE_STATION
	PRISON
	HIDEOUT
	POWER_PLANT
	_
	_
	_
)

type placeEnum int

type place struct {
	name            string
	visitors        []ID
	restrictionFunc func(i ID)
}

func (p place) Name() string {
	return p.name
}
func (p *place) Arrive(a ID) {
	if p.restrictionFunc != nil {
		p.restrictionFunc(a)
	}
	//some kind of interaction check
	p.visitors = append(p.visitors, a)

}
func (p *place) checkAllowed(a ID) bool {
	return true
}

func (p *place) Leave(a ID) {
	//some kind of interaction check
	for i, id := range p.visitors {
		if id == a {
			p.visitors[i] = 0
		}
	}

}

var places = [16]place{
	HOME: place{
		name: "home",
	},
	HOUSE: place{
		name: "house",
		restrictionFunc: func(i ID) {
		},
	},
	BAR: place{
		name: "bar",
	},
	PARK: place{
		name: "park",
	},
	WORK_OFFICE: place{
		name:            "office",
		restrictionFunc: func(i ID) {},
	},
	WORK_FACTORY: place{
		name:            "factory",
		restrictionFunc: func(i ID) {},
	},
	STORE: place{
		name: "store",
	},
	GUN_STORE: place{
		name: "gun store",
	},
	HOSPITAL: place{
		name: "hospital",
	},
	CITY_HALL: place{
		name:            "city hall",
		restrictionFunc: func(i ID) {},
	},
	POLICE_STATION: place{
		name:            "police station",
		restrictionFunc: func(i ID) {},
	},
	PRISON: place{
		name:            "prison",
		restrictionFunc: func(i ID) {},
	},
	HIDEOUT: place{
		name:            "secret hideout",
		restrictionFunc: func(i ID) {},
	},
	POWER_PLANT: place{
		name:            "the local nuclear power plant",
		restrictionFunc: func(i ID) {},
	},
}

func (p placeEnum) Place() *place {
	return &places[p]
}
