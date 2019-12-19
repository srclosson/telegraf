package eval

import (
	"github.com/cosmos72/gomacro/fast"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

var sampleConfig = `

`

type Eval struct {
	Statement string

	interp      *fast.Interp
	expr        *fast.Expr
	initialized bool
}

func (p *Eval) SampleConfig() string {
	return sampleConfig
}

func (p *Eval) Description() string {
	return "Directly change your metrics"
}

func (p *Eval) Apply(in ...telegraf.Metric) []telegraf.Metric {
	if !p.initialized {
		p.interp = fast.New()
		p.initialized = true
		p.expr = p.interp.Compile(p.Statement)
	}

	//for _, metric := range in {
	//	fast.NewEnv()
	//}
	return in
}

func init() {
	processors.Add("eval", func() telegraf.Processor {
		return &Eval{
			initialized: false,
		}
	})
}