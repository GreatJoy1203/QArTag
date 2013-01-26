package main

import (
					"fmt"
					"net/http"
					
					"qr/web"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main(){
	http.HandleFunc("/qr/draw", web.Draw)
	 http.HandleFunc("/test", handler)
	fmt.Println("handle")
    http.ListenAndServe(":8080", nil)
}

