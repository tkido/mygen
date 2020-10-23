package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/part"
	"github.com/tkido/mygen/status"
)

type PartMenu struct {
	MenuBase
	Part part.Type
	Data []*ebiten.Image
}

func NewPartMenu(w, h, col, row int) *PartMenu {
	canvas, _ := ebiten.NewImage(w*col, h*row, ebiten.FilterDefault)
	cursorImg, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)

	menu := &PartMenu{
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
		Data: []*ebiten.Image{},
	}
	menu.Self = menu
	return menu
}

func (m *PartMenu) SetCursor(index int) {
	m.MenuBase.SetCursor(index)
	g.Character.StatusMap[status.Human].Parts[m.Part] = part.Index(index - 1)
	m.Reflesh()
}

func (m *PartMenu) Update() {
	ps, ok := g.VariationManager.Map[g.Character.Base][m.Part]
	if !ok {
		log.Fatalf("not found")
	}
	m.Data = []*ebiten.Image{}
	m.Data = append(m.Data, g.View.Bg)
	for _, p := range ps {
		m.Data = append(m.Data, g.ImageManager.LoadImage(p.file))
	}
	m.Limit = len(m.Data) - 1

	m.SetCursor(int(g.Character.StatusMap[status.Human].Parts[m.Part] + 1))
}

func (m *PartMenu) Reflesh() {
	log.Println("PartMenu.Reflesh")
	m.Canvas.Fill(color.Transparent)

	for i, img := range m.Data {
		x := i % m.Col
		y := i / m.Col
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
		m.Canvas.DrawImage(g.View.Bg, op)
		m.Canvas.DrawImage(img, op)
		if m.Cursor == i {
			m.CursorImg.Fill(g.View.GetFocusColor(m))
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*m.W), float64(y*m.H))
			m.Canvas.DrawImage(m.CursorImg, op)
		}
	}
}
