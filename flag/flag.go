package flag

import (
	"flag"
	"testing"

	"github.com/tkido/mygen/base"
)

var (
	Id   int
	Base base.Type
)

func init() {
	testing.Init()
	var btInt int

	flag.IntVar(&Id, "id", 1, "id Actor ID")
	flag.IntVar(&btInt, "base", 1, "base Actor Base")
	flag.Parse()

	Base = base.Type(btInt)
}
