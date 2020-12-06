package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Serve all static files from the public directory.
	router.ServeFiles("/assets/*filepath", http.Dir("web/public/assets"))

	// Top level files must be named explicitly. (No * matching)
	router.GET("/security.txt", file("web/public/standards/security.txt"))
	router.GET("/robots.txt", file("web/public/standards/robots.txt"))
	router.GET("/sitemap.xml", file("web/public/standards/sitemap.xml"))
	router.GET("/hall-of-fame", file("web/public/standards/hall-of-fame"))

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err)
	}
}

// Closure for static files.
func file(file string) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, file)
	}
}
