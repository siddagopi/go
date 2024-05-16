package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Mapper struct {
	Mapping map[string]string
	lock    sync.Mutex
}

var urlMapper Mapper

func init() {
	urlMapper = Mapper{
		Mapping: make(map[string]string),
	}
}
func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running...ji"))
	})

	createShortURLHandler()
	fmt.Println("hi it's coming here after short url handler")
	r.Get("/short/{key}", redirectHandler)

	// fmt.Println("hi it's coming here after redirect url handler")
	http.ListenAndServe(":3000", r)
}

func createShortURLHandler() {

	fmt.Println("coming to create short url handler function")
	u := "http://google.com"

	//generate key
	key := "1"

	//insert key
	insertMapping(key, u)

}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("coming inside of redirect handler")
	key := chi.URLParam(r, "key")

	fmt.Println("key here", key)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key field is empty"))
		return
	}
	//fetch mapping

	u := fetchMapping(key)

	if u == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url field is empty"))
		return
	}

	fmt.Println("url here", u)

	http.Redirect(w, r, u, http.StatusFound)
}

func insertMapping(key string, u string) {
	urlMapper.lock.Lock()

	defer urlMapper.lock.Unlock()

	urlMapper.Mapping[key] = u
}

func fetchMapping(key string) string {
	urlMapper.lock.Lock()
	defer urlMapper.lock.Unlock()

	return urlMapper.Mapping[key]
}
