package buityrutime

import (
	"fmt"
	"time"
)

type BuityRuTime string

var months = []string{
	"января", "февраля", "марта", "апреля", "мая", "июня",
	"июля", "августа", "сентября", "октября", "ноября", "декабря",
}

func (b BuityRuTime) String() string {
	return string(b)
}

func Now() BuityRuTime {
	t := time.Now()

	return formatRu(t)
}

func formatRu(t time.Time) BuityRuTime {
	return BuityRuTime(fmt.Sprintf("%d %s %d", t.Day(), months[t.Month()-1], t.Year()))
}
