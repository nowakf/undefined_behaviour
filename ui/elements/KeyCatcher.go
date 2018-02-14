package elements

//this interface covers anything that captures keypresses -
//only one of these should be registering input events at
//any given time
type KeyCatcher interface {
	Escape() KeyCatcher
}
