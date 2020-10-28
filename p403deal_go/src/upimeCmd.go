package main

import (
	"encoding/json"
	"fmt"
	"logger"
)

type CmdInfos struct {
	cmds  []UpimeCmd
	index int
}

func (this *CmdInfos) UnmarshalJSON(c []byte) error {
	var value []interface{}
	err := json.Unmarshal(c, &value)
	if err != nil {
		return err
	}
	this.cmds = make([]UpimeCmd, len(value))
	var j = 0
	for _, one := range value {
		cmd := convertToCmd(one.([]interface{}))
		if cmd != nil {
			this.cmds[j] = cmd
			j++
		}
	}
	this.cmds = this.cmds[:j]
	this.index = 0
	return nil
}

func convertToCmd(value []interface{}) UpimeCmd {
	// defer func() {
	// err := recover()
	// logger.Error("error:", err)
	// }()
	// logger.Debug("deal command %v\n", value)
	fmt.Println("convert To Cmd ...")
	t := uint(value[1].(float64))
	var one UpimeCmd
	switch t {
	case LINE:
		//logger.Debug("line get")
		one = &Line{}
	case ERASER:
		one = &eraser{}
	case INSERTIMAGE:
		one = &insertImage{}
	case CLEAR:
		one = &clearCmd{}
	case UPDATEIMAGE:
		one = &updateImage{}
	case POSITION:
		one = &position{}
	case SIZE:
		one = &resize{}
	case SCROLL:
		one = &scroll{}
	case DEACTIVE:
		one = &deactive{}
	case DELETE:
		one = &drop{}
	case TEXT:
		one = &text{}
	case FOCUSLINE:
		one = &focusPoint{}
	case BKPAPER:
		one = &bgCmd{}
	case PAGE:
		one = &switchPage{}
	case SINGLECMD:
		one = &singleCmd{}
	default:
		logger.Warn("unknown command %d\n", t)
		return nil
	}
	one.init(value) //初始化命令
	return one
}
