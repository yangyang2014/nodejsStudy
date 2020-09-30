package main

import (
	"fmt"
	"strings"
	"sort"
	)

func getMinMinute(dataStrArray []string){
	sort.Strings(dataStrArray)
	for i,value:=range dataStrArray{
		arr:=strings.Split(value,":")
		fmt.Println(i,value,arr)
	}
	
}

func getMinusByTwoTime(time1 string, time2 string) int {
	
}

func main() {
	dataStrArray := []string{"11:30","15:45","12:00"}
	getMinMinute(dataStrArray)
}