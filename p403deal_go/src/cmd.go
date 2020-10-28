package main

import (
	"strconv"
)

type cmdInfo struct {
	name string
	page int
	time int64
}

func (this *cmdInfo) init(rawCmd []interface{}) {
	if a, ok := rawCmd[0].([]interface{}); ok {
		this.name = a[1].(string)
		this.page = int(a[0].(float64))
	} else {
		this.time = int64(rawCmd[0].(float64))
	}
	return
}

type Cmd struct {
	cmdInfo
	CmdId uint
}

func (c *Cmd) init(raw []interface{}) {
	c.cmdInfo.init(raw)
	c.CmdId = uint(raw[1].(float64))
}

func (c *Cmd) Info() *Cmd {
	return c
}

type IDCmd struct {
	Cmd
	id string
}

func (this *IDCmd) init(rawCmd []interface{}) {
	this.Cmd.init(rawCmd)
	id := rawCmd[2]
	switch id.(type) {
	case float64:
		this.id = strconv.Itoa(int(id.(float64)))
	case string:
		this.id = id.(string)
	default:
		// logger.Debug("know type")
	}
	return
}
func (this *IDCmd) GetId() string {
	return this.id
}

type Point struct {
	X, Y float64
}

type Draw interface {
	// Draw(p *Page, offset float64)
}
type UpimeCmd interface {
	Info() *Cmd
	init(rawCmd []interface{})
	// Merge(p *Page)
}
type CmdOp interface {
	Draw
	UpimeCmd
}

type element interface {
	Draw
	inRange(top, height float64) bool
}
type mr interface {
	move(p Point)
	resize(p Point)
}
type shape interface {
	element
	mr
}

type rect struct {
}

const (
	_ = iota
	INSERTIMAGE
	UPDATEIMAGE
	ACTIVE
	DEACTIVE
	POSITION
	SCROLL
	CLEAR
	LINE
	PAGE
	DELETE
	TEXT
	SIZE
	EMPTYSCREEN
	BEZIER2
	KEYFRAME
	ERASER
	BEZIER1
	COVER
	VOICE
	FRAME
	FOCUSLINE
	BKPAPER
	SINGLECMD
)
