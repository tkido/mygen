package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/gradient"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
)

type ColorMenu struct {
	MenuBase
	Part part.Type
	Data []gradient.Row
}

func NewColorMenu(w, h, col, row int) *ColorMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursorImg, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)

	menu := &ColorMenu{
		MenuBase: MenuBase{
			W:         w,
			H:         h,
			Col:       col,
			Row:       row,
			Cursor:    0,
			Canvas:    canvas,
			CursorImg: cursorImg,
			Dirty:     true,
		},
		Part: part.Face,
		Data: []gradient.Row{},
	}
	menu.Self = menu
	return menu
}

func (m *ColorMenu) SetCursor(index int) {
	m.MenuBase.SetCursor(index)
	g.Character.StatusMap[status.Human].Parts[m.Part] = part.Index(index - 1)
	m.Reflesh()
	g.Logic.UpdateFace()
}

func (m *ColorMenu) Update() {
	// ps, ok := g.VariationManager.Map[g.Character.Base][m.Part]
	// if !ok {
	// 	log.Fatalf("not found")
	// }
	m.Data = []gradient.Row{-1}
	if gt, ok := part.GradientMap[g.PartMenu.Part]; ok {
		g := gradient.Map[gt]
		for row := g.Start; row <= g.Start+g.Number; row++ {
			m.Data = append(m.Data, row)
		}
	}
	m.Limit = len(m.Data) - 1

	m.SetCursor(int(g.Character.StatusMap[status.Human].Parts[m.Part] + 1))
}

func (m *ColorMenu) Reflesh() {
	log.Println("ColorMenu.Reflesh")
	m.Canvas.Fill(color.Transparent)
	for i, row := range m.Data {
		x := i % m.Col
		y := i / m.Col
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.25, 8)
		op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
		r := int(row)
		rect := image.Rect(64, r*4, 256-64, (r+1)*4)
		m.Canvas.DrawImage(g.ImageManager.Gradient.SubImage(rect).(*ebiten.Image), op)
		if m.Cursor == i {
			m.CursorImg.Fill(g.View.GetFocusColor(m))
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
			m.Canvas.DrawImage(m.CursorImg, op)
		}
	}
}
