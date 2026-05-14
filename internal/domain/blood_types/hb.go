package bloodtypes

const hbName = "Белок в эритроцитах, переносит кислород"

type HB struct {
	val  int
	name string
}

func (h *HB) Value() int {
	return h.val
}

func (h *HB) Name() string {
	return h.name
}

func NewHb(val int) *HB {
	return &HB{val: val, name: hbName}
}
