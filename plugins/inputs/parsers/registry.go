package parsers

import "github.com/srclosson/telegraf/plugins/parsers"

type Creator func() parsers.Parser

var Parsers = map[string]Creator{}

func Add(name string, creator Creator) {
	Parsers[name] = creator
}
