package main
import "fmt"

func main(){

	var countryCapitalMap map[string]string 
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] ="巴黎"
	
    for k,v := range countryCapitalMap{
		fmt.Println(k,v)
	}
	
}
