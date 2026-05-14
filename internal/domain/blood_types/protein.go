package bloodtypes

const proteinName = "Суммарная концентрация белков плазмы"

type PROTEIN struct {
	val  int
	name string
}

func NewProtein(val int) *PROTEIN {
	return &PROTEIN{val: val, name: proteinName}
}

func (p PROTEIN) Name() string {
	return p.name
}

func (p PROTEIN) Value() int {
	return p.val
}
