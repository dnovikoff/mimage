package image

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
)

func DecodePNG(data []byte) (image.Image, error) {
	img, err := png.Decode(bytes.NewReader(data))
	return img, err
}

func LoadPNG(filename string) (image.Image, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DecodePNG(data)
}

func EncodePNG(src image.Image) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := png.Encode(buf, src)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func WritePNG(filename string, src image.Image) error {
	data, err := EncodePNG(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
