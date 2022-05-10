package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"
)

type Dumper struct {
	N, M, W, H, Cnt int
	S               *svg.SVG
	I               *image.RGBA
	imgFile         *os.File
}

func NewDumper(name string, n, m, w, h int) *Dumper {
	var err error
	dumper := Dumper{N: n, M: m, W: w, H: h}

	dumper.imgFile, err = os.Create(name + ".jpeg")
	if err != nil {
		panic(err)
	}
	dumper.I = image.NewRGBA(image.Rect(0, 0, n*w, m*h))
	bg := image.NewUniform(color.RGBA{0xff, 0xff, 0xff, 0xff})
	draw.Draw(dumper.I, dumper.I.Bounds(), bg, image.Point{}, draw.Src)

	return &dumper
}
func (d *Dumper) Close() {
	jpeg.Encode(d.imgFile, d.I, nil)
	d.imgFile.Close()
}

func (d *Dumper) Plot(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt%d.N

	igr := imgg.AddTo(d.I, col*d.W, row*d.H, d.W, d.H, color.RGBA{0xff, 0xff, 0xff, 0xff}, nil, nil)
	c.Plot(igr)

	d.Cnt++
}
