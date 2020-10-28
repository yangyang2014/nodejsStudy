package main

import (
	"image"
	"image/draw"
	"logger"

	"github.com/anthonynsimon/bild/transform"
)

type upimeRectangle struct {
	image.Rectangle
}

func (this *upimeRectangle) init(raw []interface{}) {
	p1 := raw[0].([]interface{})
	this.Min = image.Point{int(p1[0].(float64)), int(p1[1].(float64))}
	p2 := raw[1].([]interface{})
	this.Max = image.Point{int(p1[0].(float64) + p2[0].(float64)), int(p1[1].(float64) + p2[1].(float64))}
}
func (this *upimeRectangle) move(p Point) {
	a := this.Size()
	this.Min.X = int(p.X)
	this.Min.Y = int(p.Y)
	this.Max.X = this.Min.X + a.X
	this.Max.Y = this.Min.Y + a.Y
}
func (this *upimeRectangle) resize(p Point) {
	this.Max.X = this.Min.X + int(p.X)
	this.Max.Y = this.Min.Y + int(p.Y)
}

func (i *upimeRectangle) inRange(top, height float64) bool {
	if i.Max.Y > int(top) && i.Min.Y < int(top+height) {
		return true
	}
	return false
}

type insertImage struct {
	updateImage
	upimeRectangle
	left, top, width, height float64
	active                   bool
	content                  image.Image
}

type updateImage struct {
	IDCmd
	addr string
}

func (this *updateImage) init(raw []interface{}) {
	this.IDCmd.init(raw)
	this.addr = raw[3].(string)
}

func (i *updateImage) Merge(p *Page) {
	image := p.Images[i.id]
	if image == nil {
		logger.Warn("no image %s\n", i.id)
		return
	}
	image1 := image.(*insertImage)
	image1.addr = i.addr
	image1.content = nil
}

func (this *insertImage) init(raw []interface{}) {
	this.updateImage.init(raw)
	this.upimeRectangle.init(raw[4:6])
}

func (i *insertImage) Merge(p *Page) {
	image := p.Images[i.id]
	if image != nil {
		logger.Warn("duplicated image %s\n", i.id)
		return
	}
	p.Images[i.id] = i
	p.imageIndex = append(p.imageIndex, i.id)
}
func (this *insertImage) resize(p Point) {
	this.upimeRectangle.resize(p)
	this.content = nil
}
func (i *insertImage) Draw(p *Page, offset float64) {
	logger.Debug("image draw %v at %v within page %v\n", i.id, i.Rectangle, p.scroll)
	if i.content == nil {
		var err error
		i.content, err = p.getImage(i.addr)
		if err != nil {
			return
		}
		size := i.Size()
		i.content = transform.Resize(i.content, size.X, size.Y, transform.Linear)
	}
	drawRect := i.Rectangle
	drawRect.Min.Y -= int(offset)
	drawRect.Max.Y -= int(offset)
	draw.Draw(p.ImageCanvas, drawRect, i.content, image.ZP, draw.Over)
}
