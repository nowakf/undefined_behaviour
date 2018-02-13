package ui

type email interface {
	Subject() string
	Sender() string
	Body() string
}
