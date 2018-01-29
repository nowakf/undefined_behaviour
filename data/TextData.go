//data unmarshalls textual data from files.
//as it stands, this is all loaded upfront, since it's all plaintext.
package data

type TextData struct {
	ActorData   map[string]DeadActor
	EventsData  map[string]DeadEvent
	ActionsData map[string]DeadAction
	SpellsData  map[string]DeadAction
}

func NewData() *TextData {
	return new(TextData)
}
