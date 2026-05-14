package bloodtypes

const mchName = "Среднее содержание гемоглобина в эритроците"

type MCH struct {
	val  int
	name string
}

func NewMCH(val int) *MCH {
	return &MCH{val: val, name: mchName}
}

func (m MCH) Name() string {
	return m.name
}

func (m MCH) Value() int {
	return m.val
}
