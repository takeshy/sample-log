package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const rootTemplateHTML = `
<html>
	<body>
		<form action="/sample" method="post">
			code: <input type="text" name="code" value="" >
			<input type="submit" value="send">
		</form>
	</body>
</html>
`

type SampleForm struct {
	Code string `json:"code"`
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rootTemplate.Execute(w, nil)
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(500)
			return
		}
		code := r.Form.Get("code")
		f := SampleForm{Code: code}
		t, err := json.Marshal(f)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		fmt.Printf("%s\n", string(t))
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("location", "/sample")
		w.WriteHeader(http.StatusMovedPermanently) // 301 Moved Permanently
	default:
		w.WriteHeader(404)
	}
}

func main() {
	http.HandleFunc("/sample", sampleHandler)
	http.ListenAndServe(":80", nil)
}
