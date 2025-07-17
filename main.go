package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"rc2/bug/pages"

	"github.com/a-h/templ"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed assets
	embededFiles embed.FS
	playerList   = []string{"dave", "jason", "mark", "dan", "sam"}
)

func main() {

	r := chi.NewRouter()

	assets, err := fs.Sub(embededFiles, "assets")
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/*", http.FileServer(http.FS(assets)))

	r.Handle("/qr/*", http.StripPrefix("/qr/", http.FileServer(http.Dir("./qr"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		templ.Handler(pages.Admin(playerList)).ServeHTTPStreamed(w, r)
	})

	http.ListenAndServe(":3000", r)

	fmt.Println("listening on port :3000")
}
