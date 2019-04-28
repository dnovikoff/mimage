package image

import (
	"image"

	"github.com/dnovikoff/tempai-core/tile"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type Images []image.Image

func (i Images) InitBlank() {
	i[int(BlankTile-1)] = BlankFromImage(i[int(BlankTile-1)])
}

func (i Images) Tile(t tile.Tile) image.Image {
	return i[int(t-tile.TileBegin)]
}

func DefaultSprites(src image.Image) Images {
	x := Sprites(src, 7, 6)
	x.InitBlank()
	return x
}

func Sprites(src image.Image, cols, rows int) Images {
	sub := src.(SubImager)
	size := src.Bounds().Size()
	w := size.X / cols
	h := size.Y / rows
	x := 0
	y := 0
	mp := make(Images, rows*cols)
	i := 0
	for r := 0; r < rows; r++ {
		x = 0
		for c := 0; c < cols; c++ {
			mp[i] = sub.SubImage(image.Rectangle{
				Min: image.Point{X: x, Y: y},
				Max: image.Point{X: x + w, Y: y + h},
			})
			x += w
			i++
		}
		y += h
	}
	return mp
}
