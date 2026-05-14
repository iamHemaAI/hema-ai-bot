package bloodtypes

const wbcName = "Клетки иммунной системы"

type WBC struct {
	val  float64
	name string
}

func NewWBC(val float64) *WBC {
	return &WBC{val: val, name: wbcName}
}

func (w WBC) Name() string {
	return w.name
}

func (w WBC) Value() float64 {
	return w.val
}
