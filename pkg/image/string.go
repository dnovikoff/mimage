package image

import (
	"image"

	"github.com/dnovikoff/tempai-core/tile"
	"github.com/facebookgo/stackerr"
)

const (
	RedMan = tile.TileEnd + iota
	RedPin
	RedSou
	BackTile
	BlankTile
)

func (s Images) Parse(str string) ([]image.Image, error) {
	if len(str) == 0 {
		return nil, nil
	}
	tmp := make([]image.Image, 0, len(str))
	index := 0
	t := tile.TileEnd
	max := '0'
	for k, v := range str {
		r := rune(v)
		switch r {
		case 's':
			t = tile.Sou1
		case 'm':
			t = tile.Man1
		case 'p':
			t = tile.Pin1
		case 'z':
			t = tile.East
			if max > '7' {
				return nil, stackerr.Newf("Unexpected value '%s' for type '%s'", string(max), string(v))
			}
		case '_', '-', 'I', '0':

		default:
			if r < '1' || r > '9' {
				return nil, stackerr.Newf("Unexpected symbol '%s' at position %v", string(v), k)
			}
			if r > max {
				max = r
			}
		}
		if t != tile.TileEnd {
			if index == k {
				return nil, stackerr.Newf("Empty range at %v", index)
			}
			rotate := false
			for _, val := range str[index:k] {
				var i image.Image
				switch val {
				case '-':
					rotate = true
					continue
				case '_':
					i = s.Tile(BlankTile)
				case 'I':
					i = s.Tile(BackTile)
				case '0':
					switch t.Type() {
					case tile.TypeMan:
						i = s.Tile(RedMan)
					case tile.TypePin:
						i = s.Tile(RedPin)
					case tile.TypeSou:
						i = s.Tile(RedSou)
					default:
						return nil, stackerr.Newf("Unexpected red five symbol")
					}
				default:
					i = s.Tile(tile.Tile(int(rune(val)-'1')) + t)
				}
				if rotate {
					var prev *Rotate
					if len(tmp) > 0 {
						if rot, ok := tmp[len(tmp)-1].(*Rotate); ok {
							if _, ok := rot.Image.(*Group); !ok {
								prev = rot
							}
						}
					}
					if prev != nil {
						i = NewGroup(prev.Image, i)
						tmp = tmp[:len(tmp)-1]
					}
					i = &Rotate{Image: i, Count: 1}
				}
				tmp = append(tmp, i)
				rotate = false
			}
			if rotate {
				return nil, stackerr.Newf("Incorrect rotate symbol at position %v", k-1)
			}
			index = k + 1
			max = '0'
			t = tile.TileEnd
		}
	}
	if index != len(str) {
		return nil, stackerr.Newf("Expected to end with a letter")
	}
	return tmp, nil
}
