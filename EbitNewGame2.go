package main

import (
	//E "EbitNew"
	E2 "EbitNew2"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&E2.Game{}); err != nil {
		log.Fatal(err)
	}
}

var Images []Image2

func init() {

	var err error
	image1, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	//hmm := entity1{}
	//hmm.Position = Position{5, 5}
	NewEntity1(image1)
	//println(int(hmm.X))
	//println(int(hmm.Y))
}

type entity1 struct {
	*E2.Position
	*E2.Image
}

var draw []Draw

func NewEntity1(image *ebiten.Image) *entity1 {
	new := entity1{}
	new.Position = E2.NewPosition(200, 200)
	new.Image = E2.NewImage(image)
	Images = append(Images, &new)
	E2.Draws = append(E2.Draws, &new)
	return &new
}

type Image2 interface {
	Draw()
}

type Image struct {
	Image *ebiten.Image
	Op    ebiten.DrawImageOptions
}

func (s *Image) getImage() *ebiten.Image {
	return s.Image
}

func (s *Image) getOp() *ebiten.DrawImageOptions {
	return &s.Op
}

type Draw interface {
	getPosition() (float64, float64)
	getImage() *ebiten.Image
	getOp() *ebiten.DrawImageOptions
}

func (s *Image) Draw() {
	println("hej")
}

func NewImage(image *ebiten.Image) *Image {
	new := Image{}
	//new.Entity = s
	new.Image = image
	Images = append(Images, &new)
	return &new
}

type Position struct {
	X float64
	Y float64
}

func (s *Position) getPosition() (float64, float64) {
	return s.X, s.Y
}

func NewPosition(x, y float64) *Position {
	new := Position{}
	new.X = x
	new.Y = y
	return &new
}
