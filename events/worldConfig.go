package events

type WorldConfig struct {
	//some configuration figures
	madness int
	year    int
}

func NewWorldConfig() *WorldConfig {
	w := new(WorldConfig)
	w.year = 1961
	w.madness = 0
	return w
}
func (w *WorldConfig) SetYear(year int) {
	w.year = year
}

func (w *WorldConfig) SetMadness(mad int) {
	w.madness = mad
}
