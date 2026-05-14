package bloodtypes

const pltName = "Отвечают за свертываемость крови"

type PLT struct {
	val  int
	name string
}

func NewPLT(val int) *PLT {
	return &PLT{val: val, name: pltName}
}

func (p PLT) Name() string {
	return p.name
}

func (p PLT) Value() int {
	return p.val
}
