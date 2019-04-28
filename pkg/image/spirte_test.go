package image

import (
	"fmt"
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpriting(t *testing.T) {
	s := testSprites(t)
	require.Equal(t, 42, len(s))
	for i, v := range s {
		name := fmt.Sprintf("sprites/s%02d", i)
		t.Run(name, func(t *testing.T) {
			testPNG(t, name, v)
		})
	}
}

type testCase struct {
	Name   string
	Images []image.Image
}

func TestDrawing(t *testing.T) {
	s := testSprites(t)
	for _, v := range []testCase{
		{
			Name:   "writer",
			Images: []image.Image{s[1], s[6], s[18], s[22]},
		},
		{
			Name: "flip",
			Images: []image.Image{
				&Flip{Image: s[4], X: false, Y: false},
				&Flip{Image: s[4], X: false, Y: true},
				&Flip{Image: s[4], X: true, Y: false},
				&Flip{Image: s[4], X: true, Y: true},
				&Flip{Image: s[14], X: false, Y: false},
				&Flip{Image: s[14], X: false, Y: true},
				&Flip{Image: s[14], X: true, Y: false},
				&Flip{Image: s[14], X: true, Y: true},
			},
		},
		{
			Name: "rotate",
			Images: []image.Image{
				&Rotate{Image: s[4], Count: 0},
				&Rotate{Image: s[4], Count: 1},
				&Rotate{Image: s[4], Count: 2},
				&Rotate{Image: s[4], Count: 3},
			},
		},
		{
			Name: "group",
			Images: []image.Image{
				&Rotate{Image: NewGroup(s[4], s[18]), Count: 0},
				&Rotate{Image: NewGroup(s[4], s[18]), Count: 1},
				&Rotate{Image: NewGroup(s[4], s[18]), Count: 2},
				&Rotate{Image: NewGroup(s[4], s[18]), Count: 3},
			},
		},
	} {
		t.Run(v.Name, func(t *testing.T) {
			rec := &RecordDrawer{}
			w := &Writer{Drawer: rec}
			for _, v := range v.Images {
				w.Write(v)
			}
			sd := &SizeDrawer{}
			rec.Repeat(sd)
			dest := image.NewRGBA(sd.Bounds)
			id := &ImageDrawer{Dest: dest}
			rec.Repeat(id)
			testPNG(t, v.Name, dest)
		})
	}
}
