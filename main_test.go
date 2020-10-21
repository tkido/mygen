package main

import (
	"fmt"
	"image/color"
	"testing"
)

func TestColor(t *testing.T) {
	c1 := color.RGBA{1, 22, 100, 99}

	fmt.Printf("%#v, %#v, %#v, %#v\n", c1.R, c1.G, c1.B, c1.A)
	fmt.Printf("%#v\n", c1)
	c2 := color.RGBAModel.Convert(c1).(color.RGBA)
	r, g, b, a := c2.RGBA()
	fmt.Printf("%#v, %#v, %#v, %#v\n", r, g, b, a)
	fmt.Printf("%v, %v, %v, %v\n", r, g, b, a)

}
