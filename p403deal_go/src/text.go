package main

import (
	// "gg"
	"image/color"
	"logger"
)

type text struct {
	IDCmd
	text  string
	size  float64
	color *color.RGBA
	upimeRectangle
}

func (this *text) init(raw []interface{}) {

	this.IDCmd.init(raw)
	this.text = raw[3].(string)
	switch len(raw) {
	case 9, 8:
		a := raw[6:8]
		this.upimeRectangle.init(a)
		fallthrough
	case 6:
		cs := raw[5].([]interface{})
		var colorv [3]uint8
		for i, l := 0, len(cs); i < 3 && i < l; i++ {
			colorv[i] = uint8(cs[i].(float64))
		}
		this.color = &color.RGBA{colorv[0], colorv[1], colorv[2], 0xff}
		fallthrough
	case 5:
		sf := raw[4].([]interface{})
		this.size = sf[0].(float64)
	}
}

func (this *text) Merge(p *Page) {
	txt := p.texts[this.id]
	if txt == nil {
		p.texts[this.id] = this
		p.textIndex = append(p.textIndex, this.id)
		return
	}
	txt1 := txt.(*text)
	txt1.text = this.text
	//update position is not realized yet
}

var fontPath string

func (this *text) Draw(p *Page, offset float64) {
	gc := p.linePen
	gc.SetColor(*this.color)
	var x, y, width int
	x = this.Min.X + 4
	y = this.Min.Y + 20
	width = this.Max.X - this.Min.X - 4 - 10
	err := gc.LoadFontFace(fontPath, this.size)
	if err != nil {
		logger.Info("load font %s failed", fontPath)
	}
	logger.Info("will paint x=%d y=%d width=%d", x, y, width)
	gc.DrawStringWrapped(this.text, float64(x), float64(y)-offset, 0, 0, float64(width), (21/this.size)+1, gg.AlignLeft)
}
