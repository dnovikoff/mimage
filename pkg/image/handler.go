package image

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

type Handler struct {
	Sprites Images
	MaxLen  int
}

var _ http.Handler = &Handler{}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	base := strings.TrimPrefix(req.URL.Path, "/")
	ext := path.Ext(base)
	if ext != ".png" {
		http.Error(rw, "only png supported", http.StatusBadRequest)
		return
	}
	input := strings.TrimSuffix(base, ext)
	if len(input) > h.MaxLen {
		http.Error(rw, "max len exceded", http.StatusBadRequest)
		return
	}
	output, err := h.Sprites.Parse(input)
	if err != nil {
		http.Error(rw, "error parsing data", http.StatusBadRequest)
		return
	}
	rec := &RecordDrawer{}
	w := &Writer{
		Drawer: rec,
	}
	w.WriteImages(output)
	sd := &SizeDrawer{}
	rec.Repeat(sd)
	dest := sd.NewRGBA()
	id := &ImageDrawer{Dest: dest}
	rec.Repeat(id)
	data, err := EncodePNG(dest)
	if err != nil {
		http.Error(rw, "error encoding image", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	rw.Header().Set("Content-Type", "image/png")
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}
