package image

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringErrors(t *testing.T) {
	src, err := LoadPNG("test_data/sprite.png")
	require.NoError(t, err)
	s := DefaultSprites(src)
	for _, v := range []string{
		"12345",
		"123-z",
		"123z-",
	} {
		t.Run(v, func(t *testing.T) {
			_, err := s.Parse(v)
			require.Error(t, err)
		})
	}
}

func TestString(t *testing.T) {
	s := testSprites(t)
	for _, v := range []string{
		"123z123p123s",
		"-123z_1-23p_12-3s",
		"I11Iz__I23Ip",
		"50p50s50m",
		"1-2-3-4-5s",
		"5-0-55m",
		// Examples
		"1379m2568p23456s_6z",
		"123445679m_-123s_4-444s",
		"I50Im_5-5-05s",
	} {
		t.Run(v, func(t *testing.T) {
			rec := &RecordDrawer{}
			w := &Writer{Drawer: rec}
			images, err := s.Parse(v)
			require.NoError(t, err)
			w.WriteImages(images)
			sd := &SizeDrawer{}
			rec.Repeat(sd)
			dest := image.NewRGBA(sd.Bounds)
			id := &ImageDrawer{Dest: dest}
			rec.Repeat(id)
			testPNG(t, v, dest)
		})
	}
}
