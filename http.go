package goutil

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// HTTPListenAndServe ...
func HTTPListenAndServe(mux http.Handler, port int) error {
	log.Printf("HTTPListenAndServe: port=%d\n", port)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
	return err
}

// HTTPServeJSON ...
func HTTPServeJSON(res http.ResponseWriter, req *http.Request, data interface{}) {
	var writer io.Writer = res
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		res.Header().Set("Content-Encoding", "gzip")
		w := gzip.NewWriter(res)
		defer w.Close()
		writer = w
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}
