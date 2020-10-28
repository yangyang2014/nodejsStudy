package main

import (
	"image"
	"image/color"
	"image/draw"
	"logger"
)

type rubeCmd struct {
	IDCmd
	next   int
	points []float64
	width  float64
}

type lineCmd struct {
	rubeCmd
	color []uint
}

//橡皮
type eraser struct {
	IDCmd
	HasNext    bool
	Points     []Point
	Width      float64
	miny, maxy float64 //??
}

type Line struct {
	eraser
	color *color.RGBA
}

func (a *eraser) init(cmd []interface{}) {
	a.IDCmd.init(cmd) //初始化id
	a.maxy = 0        //??
	u := -1
	a.miny = float64(uint(u)>>1 - 1) //？？ 最大数值

	ps := cmd[4].([]interface{})
	for i, l := 0, len(ps); i+1 < l; i += 2 {
		a.Points = append(a.Points, Point{ps[i].(float64), ps[i+1].(float64)}) //解析绘图的坐标
	}
	//logger.Debug("point len %d\n", len(a.Points))
	if len(cmd) >= 6 {
		a.Width = cmd[5].(float64) //设置宽度
	}
}

func (a *Line) init(cmd []interface{}) {
	a.eraser.init(cmd)
	if len(cmd) >= 7 { //设置线条颜色
		rawColor := cmd[6].([]interface{})
		a.color = &color.RGBA{uint8(rawColor[0].(float64)),
			uint8(rawColor[1].(float64)),
			uint8(rawColor[2].(float64)),
			0xff}
	}
}
func (this *Line) Draw(page *Page, offset float64) {
	logger.Debug("line draw %s", this.id)
	gc := page.linePen
	if this.color == nil {
		logger.Warn("line with no color %s,use black instead", this.id)
		this.color = &color.RGBA{0, 0, 0, 0xff}
	}
	gc.SetColor(*this.color)
	gc.SetLineWidth(this.Width)
	points := this.Points[2:]
	if len(points) == 0 {
		gc.DrawPoint(this.Points[0].X, this.Points[0].Y-offset, 1) // should always be called first for a new path
		gc.Stroke()
		return
	}
	gc.MoveTo(points[0].X, points[0].Y-offset) // should always be called first for a new path
	var i, l int
	for i, l = 1, len(points); i+1 < l; i += 2 {
		gc.QuadraticTo(points[i].X, points[i].Y-offset, points[i+1].X, points[i+1].Y-offset)
	}
	if i != l {
		gc.LineTo(points[i].X, points[i].Y-offset)
	}
	gc.Stroke()
}
func (this *eraser) Draw(page *Page, offset float64) {
	logger.Debug("eraser draw %s", this.id)
	gc := page.linePen
	gc.SetColor(color.Transparent)
	gc.SetLineWidth(this.Width)
	gc.MoveTo(this.Points[0].X, this.Points[0].Y-offset) // should always be called first for a new path
	var i, l int
	for i, l = 1, len(this.Points); i < l; i++ {
		gc.LineTo(this.Points[i].X, this.Points[i].Y-offset)
	}
	gc.Stroke()
}

func (this *Line) Merge(p *Page) {
	exist := p.Lines[this.id]
	ps := this.Points
	if len(ps) == 0 {
		return
	}
	var lp *Point
	var pre *Line
	if exist == nil {
		logger.Debug("add line %s", this.id)
		p.Lines[this.id] = this
		p.lineIndex = append(p.lineIndex, this.id)
		this.Points = make([]Point, 2, len(this.Points)*2)
		this.Points[0] = ps[0]
		this.Points[1] = ps[0]
		lp = &ps[0]
		pre = this
	} else {
		var ok bool
		pre, ok = exist.(*Line)
		if ok == false {
			logger.Error("wrong type of %s as Line\n", this.id)
			return
		}
		this.Width = pre.Width
		this.color = pre.color
		lp = &(pre.Points[len(pre.Points)-1])
	}
	for i, l := 0, len(ps); i < l; i++ {
		p := &ps[i]
		if p.Y < pre.miny {
			pre.miny = p.Y
		}
		if p.Y > pre.maxy {
			pre.maxy = p.Y
		}
		if *p == *lp {
			continue
		}
		np := Point{(lp.X + p.X) / 2, (lp.Y + p.Y) / 2} //猜想防止一下子出现线条的情况
		pre.Points = append(pre.Points, np, *p)
		lp = p
	}
}
func (this *eraser) Merge(p *Page) {
	exist := p.Lines[this.id]
	logger.Debug("deal %v", this.Points)
	var pre *eraser
	if exist == nil {
		p.Lines[this.id] = this
		p.lineIndex = append(p.lineIndex, this.id)
		pre = this
	} else {
		var ok bool
		pre, ok = exist.(*eraser)
		if ok {
			pre.Points = append(pre.Points, this.Points...)
			this.Width = pre.Width
		}
	}
	for _, p := range this.Points {
		if p.Y < pre.miny {
			pre.miny = p.Y
		}
		if p.Y > pre.maxy {
			pre.maxy = p.Y
		}
	}
}
func (this *eraser) inRange(top, height float64) bool {
	//logger.Debug("miny(%v),maxy(%v), top(%v) height(%v)", this.miny, this.maxy, top, height)
	if this.maxy > top && this.miny < (top+height) {
		return true
	}
	return false
}

type focusPoint struct {
	Line
	time int64
	x    float64
}

func (this *focusPoint) init(raw []interface{}) {
	this.IDCmd.init(raw)
	var p = raw[4].([]interface{})
	var length = len(p)
	if length < 2 {
		return
	}
	this.x = p[length-2].(float64)
	this.miny = p[length-1].(float64)
	this.maxy = this.miny
}

func (this *focusPoint) Merge(p *Page) {
	p.focusPoint = this
	this.time = p.currentTime
	p.setReshresh(p.currentTime + 500)
}

func (this *focusPoint) Draw(page *Page, offset float64) {
	if page.currentTime-this.time >= 500 {
		page.focusPoint = nil
		return
	}
	x := int(this.x)
	y := int(this.miny - offset)
	draw.Draw(page.ImageCanvas, image.Rect(x-6, y-6, x+6, y+6), staticImages["focus.png"], image.ZP, draw.Over)
}
