package stats

import . "github.com/nowakf/undefined_behaviour/events/world/object"

type Stats Object

func (s Stats) Get(a Stat) int {
	return Object(s).Get(Key(a))
}

func (s *Stats) Set(a Stat, newVal int) *Stats {
	(*Object)(s).Set(Key(a), newVal)
	return s
}

type Stat Key

const (
	MemberofCult     Stat = Stat(BOOL0)
	MemberofAntiCult      = Stat(BOOL1)
	LawAbiding            = Stat(BOOL2)
	MissingLimb           = Stat(BOOL3)
	MissingEye            = Stat(BOOL4)
	HeartProblems         = Stat(BOOL5)
	Schizophrenia         = Stat(BOOL6)
	Paranoia              = Stat(BOOL7)
	Agoraphobia           = Stat(BOOL8)
	Depression            = Stat(BOOL9)
	Plague                = Stat(BOOL10)
	STR                   = Stat(NIBBLE16)
	CHA                   = Stat(NIBBLE20)
	WIS                   = Stat(NIBBLE24)
	INT                   = Stat(NIBBLE28)
	WIL                   = Stat(NIBBLE32)
	SAN                   = Stat(NIBBLE36)
	Location              = Stat(NIBBLE40)
	ID                    = Stat(UINT8_48)
	Stress                = Stat(UINT8_54)
)

func (s Stat) Describe() (description string) {
	switch s {
	case MemberofCult:
		return "a follower of the old gods."
	case MemberofAntiCult:
		return "recognizes the danger followers of the old gods pose."
	case LawAbiding:
		return "law-abiding citizen."
	case MissingLimb:
		return "missing limb"
	case MissingEye:
		return "missing eye"
	case HeartProblems:
		return "heart problems"
	case Schizophrenia:
		return "schizophrenia"
	case Paranoia:
		return "paranoid"
	case Agoraphobia:
		return "agoraphobic"
	case Depression:
		return "depressed"
	case Plague:
		return "afflicted with the plague"
	case STR:
		return "STRENGTH: A crude measure of physical ability."
	case CHA:
		return "CHARISMA: How socially capable they are."
	case WIS:
		return "WISDOM: Their intuition, critical thinking skills, and ability to stay objective under pressure."
	case INT:
		return "INTELLIGENCE: Their ability to learn, and process new information."
	case WIL:
		return "WILLPOWER: The degree to which they follow through with their goals, in adversity."
	case SAN:
		return "SANITY: A raw measure of how sane they are."
	case Location:
		return "where they are."
	case ID:
		return "ID: Internal only."
	case Stress:
		return "Stress: how nervous they feel about the way things are going."
	default:
	}
	return
}
