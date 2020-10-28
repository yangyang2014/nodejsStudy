package main

type scroll struct {
	Cmd
	pos float64
}

type clearCmd struct {
	Cmd
}

func (this *clearCmd) init(raw []interface{}) {
	this.Cmd.init(raw)
}
func (this *clearCmd) Merge(p *Page) {
	p.Lines = make(map[string]element)
	p.lineIndex = p.lineIndex[:0]
}

func (this *scroll) init(raw []interface{}) {
	this.Cmd.init(raw)
	this.pos = raw[2].(float64)
}
func (this *scroll) Merge(p *Page) {
	p.scroll = this.pos
}

type bgCmd struct {
	Cmd
	img string
}

func (this *bgCmd) init(raw []interface{}) {
	this.Cmd.init(raw)
	if len(raw) >= 4 {
		this.img = raw[3].(string)
	}
}
func (this *bgCmd) Merge(p *Page) {
	p.bpImage = this.img
}

type singleCmd struct {
	Cmd
	subCmd int
}

func (this *singleCmd) init(raw []interface{}) {
	this.Cmd.init(raw)
	this.subCmd = int(raw[2].(float64))
}
func (this *singleCmd) Merge(p *Page) {
	switch this.subCmd {
	case S_RESTBEGIN:
		p.resting = true
	case S_RESTEND:
		p.resting = false
	}
}

type switchPage struct {
	Cmd
	pageNo int
}

func (this *switchPage) init(raw []interface{}) {
	this.Cmd.init(raw)
	this.pageNo = int(raw[2].(float64))
}
func (this *switchPage) Merge(p *Page) {
	p.container.initPages(this.pageNo, true)
	p.setReshresh(p.currentTime - 1) //这样保证翻页过去后立即刷新;
}

const (
	_ = iota
	S_RESTBEGIN
	S_RESTEND
)
