package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//para struct
type para struct {
	one    time.Time
	two    time.Time
	symbol string
}

//ans struct
type ans struct {
	one bool
}

//question struct
type question struct {
	p para
	a ans
}

//Test_OK ...
func Test_OK(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			p: para{
				one:    SunnyStr2Time(`2018-03-21 12:00:00`),
				two:    SunnyStr2Time(`2018-03-21 12:00:00`),
				symbol: `eq`,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one:    SunnyStr2Time(`2018-03-21 12:00:00`),
				two:    SunnyStr2Time(`2018-03-21 13:00:00`),
				symbol: `eq`,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one:    SunnyStr2Time(`2018-03-21 12:00:00`),
				two:    SunnyStr2Time(`2018-03-21 13:00:00`),
				symbol: `lt`,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one:    SunnyStr2Time(`2018-03-21 14:20:00`),
				two:    SunnyStr2Time(`2018-03-21 13:00:00`),
				symbol: `lt`,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one:    SunnyStr2Time(`2018-03-21 14:20:00`),
				two:    SunnyStr2Time(`2018-03-21 13:00:00`),
				symbol: `gt`,
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, SunnyCompareTime(p.one, p.two, p.symbol), "输入:%v", p)
	}
}
