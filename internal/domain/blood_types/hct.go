package bloodtypes

const hctName = "Доля объема крови, занимаемая эритроцитами"

type HCT struct {
	val  int
	name string
}

func NewHCT(val int) *HCT {
	return &HCT{val: val, name: hctName}
}

func (h HCT) Name() string {
	return h.name
}

func (h HCT) Value() int {
	return h.val
}
