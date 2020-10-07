package main 
import (
	"fmt"
	"log"
	"net/http"
)

type MyServeMux struct {
	http.ServeMux
	id int
}



type MessageHandler struct {
    msg string 
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, m.msg)
}

/*
http://localhost:8080/first -> handler1
http://localhost:8080/second -> handler2
http://localhost:8080/second/item1 -> handler21
http://localhost:8080/second/item2 -> 404 page not found
*/
func main() {
    mux := new (MyServeMux)

    handler1 := &MessageHandler{"The first handler."}
    handler2 := &MessageHandler{"The second handler."}
    handler21 := &MessageHandler{"The second handler, and item is 1."}

    mux.Handle("/first",  handler1)
    mux.Handle("/second", handler2)
    mux.Handle("/second/item1", handler21)

    log.Println("Listening...")
    http.ListenAndServe(":8080", mux)
}