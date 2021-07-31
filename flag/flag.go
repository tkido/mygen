package flag

import (
	"flag"
	"testing"

	"github.com/tkido/mygen/base"
)

var (
	Id   int
	Base base.Type
	Name string
)

func init() {
	testing.Init()
	var btInt int

	flag.IntVar(&Id, "id", 1, "id Actor ID")
	flag.IntVar(&btInt, "base", 1, "base Actor Base")
	flag.StringVar(&Name, "name", "", "name Actor Name")
	flag.Parse()

	Base = base.Type(btInt)
}
