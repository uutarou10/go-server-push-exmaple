package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(getFilePath(r.URL.Path))
		if err != nil {
			w.WriteHeader(404)
			return
		}
		defer file.Close()

		header := w.Header()
		header.Add("Content-Type", getMimeTypeByPath(r.URL.Path))
		if r.URL.Path == "/" {
			pusher, ok := w.(http.Pusher)
			if ok {
				for i := 1; i <= 100; i++ {
					err := pusher.Push("/sushi_chutoro_"+strconv.Itoa(i)+".png", nil)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}

		data, err := ioutil.ReadFile(getFilePath(r.URL.Path))
		w.Write(data)
	})

	http.ListenAndServeTLS(":3000", "localhost.pem", "localhost-key.pem", nil)
}

func getFilePath(fname string) string {
	if fname == "/" {
		fname = "index.html"
	}

	basePath := "./www"
	fmt.Println(path.Join(basePath, fname))
	return path.Join(basePath, fname)
}

func getMimeTypeByPath(path string) string {
	if path == "/" {
		path = "index.html"
	}
	return mime.TypeByExtension(filepath.Ext(path))
}
