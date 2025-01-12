package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"

	datastar "github.com/starfederation/datastar/sdk/go"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("index.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		panic("uhoh")
	}

}

func feed(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)

	for {

		id := uuid.New()

		sse.MergeFragments(
			fmt.Sprintf("<div id='feed'> <p id='feed_uuid' style='font-size:30px;color:#0074D9'> %s </p> <p id='feed_time'> %s </p> </div>", id.String(), time.Now().UTC()),
			// to replace the entire parent div, use withSelector. Otherwise it will simply replace the span with id feed that is calling this endpoint
			// datastar.WithSelector("div"),
		)

		// you can do send a uuid (or message or notification) once, or at a specified number of times or every second
		// in this example the feed is called on load, but you could also do it on click, etc
		time.Sleep(1 * time.Second)
	}

}

func main() {
	// port to serve on
	port := "3000"

	http.HandleFunc("GET /", Index)
	http.HandleFunc("GET /feed", feed)

	fmt.Printf("Server starting: http://127.0.0.1:%s \n", port)

	http.ListenAndServe(":"+port, nil)
}
