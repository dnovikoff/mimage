package image

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	src, err := LoadPNG("test_data/sprite.png")
	require.NoError(t, err)
	s := DefaultSprites(src)
	handler := &Handler{
		Sprites: s,
		MaxLen:  15,
	}
	server := httptest.NewServer(handler)
	defer server.Close()
	client := server.Client()

	t.Run("ok", func(t *testing.T) {
		resp, err := client.Get("http://" + server.Listener.Addr().String() + "/123z123p_-123s.png")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		testPNGReader(t, "handler", resp.Body)
	})
	t.Run("string to long", func(t *testing.T) {
		resp, err := client.Get("http://" + server.Listener.Addr().String() + "/123z123p_-123123123123s.png")
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("only png", func(t *testing.T) {
		resp, err := client.Get("http://" + server.Listener.Addr().String() + "/123z123p_-123s.jpeg")
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("subfolder", func(t *testing.T) {
		resp, err := client.Get("http://" + server.Listener.Addr().String() + "/subfolder/123z123p_-123s.png")
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("bad imput", func(t *testing.T) {
		resp, err := client.Get("http://" + server.Listener.Addr().String() + "/123z123p_-1239z.png")
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
