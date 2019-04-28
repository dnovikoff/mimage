package image

import (
	"image"
	"image/draw"
)

type Drawer interface {
	Draw(r image.Rectangle, src image.Image, sp image.Point)
}

type Writer struct {
	Drawer Drawer
	Point  image.Point
}

func (w *Writer) SkipX(x int) {
	w.Point.X += x
}

func (w *Writer) WriteImages(img []image.Image) {
	for _, v := range img {
		w.Write(v)
	}
}

func (w *Writer) Write(img image.Image) {
	nb := normalize(img.Bounds())
	bounds := nb.Add(w.Point).Sub(image.Point{Y: nb.Max.Y})
	w.Drawer.Draw(bounds, img, img.Bounds().Min)
	w.Point.X = bounds.Max.X
}

func normalize(r image.Rectangle) image.Rectangle {
	return image.Rectangle{
		Max: r.Max.Sub(r.Min),
	}
}

func getShift(r image.Rectangle) image.Point {
	shift := image.Point{}
	if r.Min.X < 0 {
		shift.X = r.Min.X
	}
	if r.Min.Y < 0 {
		shift.Y = r.Min.Y
	}
	return shift
}

func fix(r image.Rectangle) image.Rectangle {
	return r.Sub(getShift(r))
}

type RecordDrawer struct {
	records []func(Drawer)
}

func (d *RecordDrawer) Repeat(x Drawer) {
	for _, v := range d.records {
		v(x)
	}
}

func (d *RecordDrawer) Draw(r image.Rectangle, src image.Image, sp image.Point) {
	d.records = append(d.records, func(d Drawer) {
		d.Draw(r, src, sp)
	})
}

type ImageDrawer struct {
	Dest draw.Image
}

func (d *ImageDrawer) Draw(r image.Rectangle, src image.Image, sp image.Point) {
	draw.Over.Draw(d.Dest, r, src, sp)
}

type SizeDrawer struct {
	Bounds image.Rectangle
}

func (d *SizeDrawer) NewRGBA() draw.Image {
	return image.NewRGBA(d.Bounds)
}

func (d *SizeDrawer) Draw(r image.Rectangle, src image.Image, sp image.Point) {
	d.Bounds.Min = minPoint(r.Bounds().Min, d.Bounds.Min)
	d.Bounds.Max = maxPoint(r.Bounds().Max, d.Bounds.Max)
}

func min(a1, a2 int) int {
	if a1 < a2 {
		return a1
	}
	return a2
}

func minPoint(a1, a2 image.Point) image.Point {
	return image.Point{
		X: min(a1.X, a2.X),
		Y: min(a1.Y, a2.Y),
	}
}

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}

func maxPoint(a1, a2 image.Point) image.Point {
	return image.Point{
		X: max(a1.X, a2.X),
		Y: max(a1.Y, a2.Y),
	}
}
