package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/tkido/mygen/part"

	"github.com/hajimehoshi/ebiten"
	"github.com/tkido/mygen/sprite"
	"github.com/tkido/mygen/status"
)

type ExportManager struct {
	Root string
	Face *ebiten.Image
	Tv   *ebiten.Image
}

func NewExportManager(root string) ExportManager {
	return ExportManager{
		Root: root,
	}
}

func (m *ExportManager) Export() {
	// Face
	m.Face, _ = ebiten.NewImage(144*4, 144*2, ebiten.FilterDefault)
	for i, st := range status.Types {
		g.StatusMenu.Status = st
		g.Sprites.reflesh(sprite.Face)
		op := &ebiten.DrawImageOptions{}
		x, y := i/2, i%2
		op.GeoM.Translate(float64(x)*144, float64(y)*144)
		m.Face.DrawImage(g.Sprites.Face, op)
	}
	facePath := filepath.Join(".", m.Root, "faces", fmt.Sprintf("%04d.png", g.Character.Id))
	m.SaveImage(facePath, m.Face)

	// SV
	for _, st := range status.Types {
		g.StatusMenu.Status = st
		g.Sprites.reflesh(sprite.Sv)
		svPath := filepath.Join(".", m.Root, "sv_actors", fmt.Sprintf("%04d_%02d.png", g.Character.Id, st))
		m.SaveImage(svPath, g.Sprites.Sv)
	}

	// TV, TVD
	for i, st := range status.Types {
		if int(st)%4 == 0 {
			m.Tv, _ = ebiten.NewImage(48*3*4, 48*4*2, ebiten.FilterDefault)
		}
		g.StatusMenu.Status = st
		g.Sprites.reflesh(sprite.Tv)
		g.Sprites.reflesh(sprite.Tvd)

		op := &ebiten.DrawImageOptions{}
		x, y := i%4%2, i%4/2
		op.GeoM.Translate(float64(x)*48*3*2, float64(y)*48*4)
		m.Tv.DrawImage(g.Sprites.Tv, op)
		// sleeping TV
		sleepingHeadIndex := part.Index(6)
		g.Character.StatusMap[st].Parts[part.Head] = sleepingHeadIndex
		g.Character.StatusMap[st].Parts[part.Glasses] = part.Null
		g.Sprites.reflesh(sprite.Tv)
		op.GeoM.Translate(48*3, 0)
		m.Tv.DrawImage(g.Sprites.Tv.SubImage(image.Rect(0, 0, 48*3, 48*3)).(*ebiten.Image), op)
		// TVD
		op.GeoM.Translate(0, 48*3)
		m.Tv.DrawImage(g.Sprites.Tvd, op)

		if int(st)%4 == 3 {
			tvPath := filepath.Join(".", m.Root, "characters", fmt.Sprintf("%04d_%02d.png", g.Character.Id, int(st)/4))
			m.SaveImage(tvPath, m.Tv)
		}
	}
	fmt.Println("Exported!!")
}

func (m *ExportManager) SaveImage(path string, img *ebiten.Image) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
