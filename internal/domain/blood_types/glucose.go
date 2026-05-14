package bloodtypes

const glucoseName = "Уровень сахара в крови"

type GLUCOSE struct {
	val  float64
	name string
}

func (g GLUCOSE) Value() float64 {
	return g.val
}

func (g GLUCOSE) Name() string {
	return g.name
}

func NewGLUCOSE(val float64) *GLUCOSE {
	return &GLUCOSE{val: val, name: glucoseName}
}
