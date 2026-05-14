package bloodtypes

const rbcName = "Количество красных кровяных клеток"

type RBC struct {
	val  float64
	name string
}

func (r *RBC) Name() string {
	return r.name
}

func (r *RBC) Value() float64 {
	return r.val
}

func NewRbc(val float64) *RBC {
	return &RBC{val: val, name: rbcName}
}
