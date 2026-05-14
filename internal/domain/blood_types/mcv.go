package bloodtypes

const mcvName = "Средний объем эритроцита"

type MCV struct {
	val  int
	name string
}

func NewMCV(val int) *MCV {
	return &MCV{val: val, name: mcvName}
}

func (m MCV) Name() string {
	return m.name
}

func (m MCV) Value() int {
	return m.val
}
