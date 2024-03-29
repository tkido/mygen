package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Element is ebiten UI element
type Element interface {
	Id() int
	isDecendantOf(child Element) bool

	Show()
	Hide()
	IsVisible() bool
	draw(screen *ebiten.Image, clip image.Rectangle)
	Reflesh()
	Dirty()

	Move(x, y int)
	Position() (x, y int)
	Resize(w, h int)
	Size() (w, h int)
	SetDrawImageOptions(op *ebiten.DrawImageOptions)
	SetBackgroundColor(c color.Color)
	SetAnimation(anime Animation)
	StopAnimation()

	Add(x, y int, el Element)
	Clear()
	setParent(el Element)

	SetMouseCallback(e MouseEvent, c Callback)
	handleMouseEvent(ev mouseEvents, origin image.Point, clip image.Rectangle) (handled bool)
	onMouseEvent(MouseEvent)

	SetKeyCallback(key ebiten.Key, cb Callback)
	handleKeyEvent(k ebiten.Key) bool
	SetUiCallback(e UiEvent, cb Callback)
	handleUiEvent(e UiEvent) bool
	SetFocus()
	IsFocused() bool

	String() string
}

// Texter has internal text as string
type Texter interface {
	SetText(text string)
	Text() (text string)
}

// Imager has internal image source as image.Image
type Imager interface {
	SetImage(srcImage image.Image)
	Source() (srcImage image.Image)
}
