package main

import (
	"image"
	"logger"
)

var _ = image.NewUniform

const (
	IMAGEELE = iota
	TEXTELE
)

type kindCmd struct {
	IDCmd
	kind uint
}

type position struct {
	changeZindex
	pos Point
}

type resize struct {
	position
}

type drop struct {
	kindCmd
}

type changeZindex struct {
	kindCmd
}

type deactive struct {
	changeZindex
}

func (p *kindCmd) init(value []interface{}) {
	p.IDCmd.init(value)
	p.kind = uint(value[3].(float64))
}
func (p *changeZindex) init(value []interface{}) {
	p.kindCmd.init(value)
}
func (p *position) init(value []interface{}) {
	p.changeZindex.init(value)
	p1 := value[4].([]interface{})
	p.pos = Point{p1[0].(float64), p1[1].(float64)}
}
func (this *position) Merge(p *Page) {
	var one shape
	switch this.kind {
	case IMAGEELE:
		one = p.Images[this.id]
	case TEXTELE:
		one = p.texts[this.id]
	default:
		return
	}
	if one != nil {
		one.move(this.pos)
		this.changeZindex.Merge(p)
	} else {
		logger.Warn("can't find element %v", this.id)
	}

}
func (this *resize) Merge(p *Page) {
	var one shape
	switch this.kind {
	case IMAGEELE:
		one = p.Images[this.id]
	case TEXTELE:
		one = p.texts[this.id]
	default:
		return
	}
	if one != nil {
		one.resize(this.pos)
		this.changeZindex.Merge(p)
	} else {
		logger.Warn("can't find element %v", this.id)
	}
}
func (this *changeZindex) Merge(p *Page) {
	var a *[]string
	switch this.kind {
	case IMAGEELE:
		a = &p.imageIndex
	case TEXTELE:
		a = &p.textIndex
	default:
		return
	}
	l := len(*a)
	if l > 1 && (*a)[l-1] != this.id {
		for i := 0; i < l; i++ {
			if (*a)[i] == this.id {
				*a = append((*a)[:i], (*a)[i+1:]...)
				*a = append(*a, this.id)
				break
			}
		}
	}
}

func (this *drop) Merge(p *Page) {
	var a *[]string
	switch this.kind {
	case IMAGEELE:
		delete(p.Images, this.id)
		a = &p.imageIndex
	case TEXTELE:
		delete(p.texts, this.id)
		a = &p.textIndex
	default:
		return
	}
	for i, l := 0, len(*a); i < l; i++ {
		if (*a)[i] == this.id {
			*a = append((*a)[:i], (*a)[i+1:]...)
			break
		}
	}
}
