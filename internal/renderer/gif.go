package renderer

import (
	"image"
	"image/gif"
	"os"
)

func CreateGif(frames []*image.Paletted, path string) {
	delays := make([]int, len(frames))
	f, _ := os.Create(path)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{Image: frames, Delay: delays})
}
