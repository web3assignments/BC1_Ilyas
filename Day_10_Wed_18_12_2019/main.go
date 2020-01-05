package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	shell "github.com/ipfs/go-ipfs-api"
)

const (
	DefaultHttpAddress = "localhost"
	DefaultHttpPort    = 8081
)

func main() {
	httpAddress := DefaultHttpAddress
	httpPort := DefaultHttpPort

	sh := shell.NewShell("localhost:5001")

	r := mux.NewRouter()
	r.HandleFunc("/upload_text", uploadText(sh)).Methods("POST")
	r.HandleFunc("/retrieve_text", retrieveText(sh)).Methods("GET")

	log.Printf("Server ready for connections at %v:%v\n", httpAddress, httpPort)
	if err := http.ListenAndServe(fmt.Sprint(httpAddress, ":", httpPort), r); err != nil {
		log.Println(err)
	}
}

func uploadText(sh *shell.Shell) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)

			return
		}

		fileHash, err := sh.Add(strings.NewReader(string(requestBytes)))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)

			return
		}

		if _, err := rw.Write([]byte(fileHash)); err != nil {
			log.Println(err)

			return
		}
	}
}

func retrieveText(sh *shell.Shell) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println(err)

			return
		}

		fileHash := string(requestBytes)

		blockBytes, err := sh.BlockGet(fileHash)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Println(err)

			return
		}

		if _, err := rw.Write(blockBytes); err != nil {
			log.Println(err)

			return
		}
	}
}
