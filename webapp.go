package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate">
		<meta http-equiv="Pragma" content="no-cache">
		<meta http-equiv="Expires" content="0">
		<meta charset="UTF-8">
		<title>{{.Region}}</title>
	</head>
	<body>
		<h1>Hello {{.Region}}</h1>
		<img src={{.Img}}></img>
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Region string
		Img    string
	}{
		Region: os.Getenv("REGION"),
		Img:    os.Getenv("IMG"),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = t.Execute(w, data)
		check(err)
	})

	panic(http.ListenAndServe(":8000", nil))
}
