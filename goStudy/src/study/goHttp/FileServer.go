package main
import (
	"net/http"
	"fmt"
)

func main () {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	fmt.Println(fs)
	// mux.Handle("/",fs)
	http.ListenAndServe(":8080",mux)
}