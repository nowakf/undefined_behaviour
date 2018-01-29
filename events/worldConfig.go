package events

type worldConfig struct {
	//some configuration figures
	madness int
	year    int
}

func NewWorldConfig() *worldConfig {
	w := new(worldConfig)
	w.year = 1961
	w.madness = 0
	return w
}
func (w *worldConfig) SetYear(year int) {
	w.year = year
}

func (w *worldConfig) SetMadness(mad int) {
	w.madness = mad
}
