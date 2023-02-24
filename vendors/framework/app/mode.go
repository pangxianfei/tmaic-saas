package app

var mode Mode

type Mode byte

const (
	ModeProduction Mode = iota
	ModeDevelop
	ModeTest
)

func init() {
	SetMode(ModeDevelop)
}

func SetMode(m Mode) {
	switch m {
	case ModeTest:
	case ModeDevelop:
	default:
		fallthrough
	case ModeProduction:
	}
	mode = m
}
func GetMode() Mode {
	return mode
}
