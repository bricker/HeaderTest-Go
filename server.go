package server

import (
	"net/http"

	"fmt"

	"io/ioutil"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/curl", func(w http.ResponseWriter, r *http.Request) {
		context := appengine.NewContext(r)
		client := urlfetch.Client(context)
		req, _ := http.NewRequest("GET", "http://localhost:8080/http", nil)
		req.Header.Add("X-Test", "Hello1")
		req.Header.Add("X-Test", "Hello2")
		req.Header.Add("X-Test", "Hello3")

		resp, _ := client.Do(req)
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprint(w, string(bytes))
	})

	http.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%+#v", r.Header)
	})
}
