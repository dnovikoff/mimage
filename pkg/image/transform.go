package image

import (
	"image"
	"image/color"
)

type Translate struct {
	image.Image
	Point image.Point
}

var _ image.Image = &Translate{}

func (t *Translate) Bounds() image.Rectangle {
	return t.Image.Bounds().Add(t.Point)
}

func (t *Translate) At(x, y int) color.Color {
	return t.Image.At(x-t.Point.X, y-t.Point.Y)
}

type Flip struct {
	image.Image
	X bool
	Y bool
}

var _ image.Image = &Flip{}

func (t *Flip) At(x, y int) color.Color {
	min := t.Bounds().Min
	max := t.Bounds().Max
	if t.X {
		x = max.X - (x - min.X) - 1
	}
	if t.Y {
		y = max.Y - (y - min.Y) - 1
	}
	return t.Image.At(x, y)
}

type Rotate struct {
	image.Image
	Count int
}

var _ image.Image = &Rotate{}

func (t *Rotate) Bounds() image.Rectangle {
	if t.Count%2 == 0 {
		return t.Image.Bounds()
	}
	b := t.Image.Bounds()
	b.Max.X, b.Max.Y = b.Max.Y-b.Min.Y+b.Min.X, b.Max.X-b.Min.X+b.Min.Y
	return b
}

func (t *Rotate) At(x, y int) color.Color {
	min := t.Image.Bounds().Min
	max := t.Image.Bounds().Max
	switch t.Count % 4 {
	case 0:
	case 1:
		x, y = y-min.Y+min.X, max.Y-1-(x-min.X)
	case 2:
		x, y = max.X-1-(x-min.X), max.Y-1-(y-min.Y)
	case 3:
		x, y = max.X-1-(y-min.Y), (x-min.X)+min.Y
	}
	return t.Image.At(x, y)
}

type Group struct {
	images []image.Image
	bounds image.Rectangle
	model  color.Model
}

type Blank struct {
	bounds image.Rectangle
	model  color.Model
}

var _ image.Image = &Blank{}

func BlankFromImage(img image.Image) *Blank {
	return &Blank{
		bounds: img.Bounds(),
		model:  img.ColorModel(),
	}
}

func (b *Blank) ColorModel() color.Model {
	return b.model
}

func (b *Blank) Bounds() image.Rectangle {
	return b.bounds
}

func (b *Blank) At(x, y int) color.Color {
	return color.Transparent
}

var _ image.Image = &Group{}

func NewGroup(x ...image.Image) *Group {
	if len(x) == 0 {
		return nil
	}
	g := &Group{
		images: x,
		model:  x[0].ColorModel(),
	}
	size := image.Point{}
	for _, v := range x {
		s := v.Bounds().Size()
		if size.Y < s.Y {
			size.Y = s.Y
		}
		size.X += s.X
	}
	g.bounds.Max = size
	return g
}

func (g *Group) ColorModel() color.Model {
	return g.model
}

func (g *Group) Bounds() image.Rectangle {
	return g.bounds
}

func (g *Group) At(x, y int) color.Color {
	for _, v := range g.images {
		s := v.Bounds().Size()
		if x < s.X {
			min := v.Bounds().Min
			return v.At(min.X+x, min.Y+y)
		}
		x -= s.X
	}
	return nil
}

func Transform(img image.Image, rotate int, flip bool) image.Image {
	if flip {
		img = &Flip{Image: img, Y: true}
	}
	if rotate > 0 {
		img = &Rotate{Image: img, Count: rotate}
	}
	return img
}
