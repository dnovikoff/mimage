package image

import (
	"bytes"
	"flag"
	"image"
	"io"
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	update = flag.Bool("update", false, "update test data")
)

func testSprites(t *testing.T) Images {
	src, err := LoadPNG("test_data/sprite.png")
	require.NoError(t, err)
	return DefaultSprites(src)
}

func testPNG(t *testing.T, name string, img image.Image) {
	data, err := EncodePNG(img)
	require.NoError(t, err)
	testPNGBytes(t, name, data)
}

func testPNGReader(t *testing.T, name string, r io.ReadCloser) {
	defer r.Close()
	data, err := ioutil.ReadAll(r)
	require.NoError(t, err)
	testPNGBytes(t, name, data)
}

func testPNGBytes(t *testing.T, name string, data []byte) {
	filename := path.Join("test_data", name+".png")
	if *update {
		require.NoError(t, ioutil.WriteFile(filename, data, 0644))
		return
	}
	expected, err := ioutil.ReadFile(filename)
	require.NoError(t, err)
	assert.True(t, bytes.Equal(expected, data))
}
