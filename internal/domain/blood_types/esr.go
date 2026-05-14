package bloodtypes

const esrName = "Скорость оседания эритроцитов, косвенный маркер воспаления"

type ESR struct {
	val  int
	name string
}

func (e ESR) String() string {
	return e.name
}

func (e ESR) Value() int {
	return e.val
}

func NewESR(val int) *ESR {
	return &ESR{val: val, name: esrName}
}
