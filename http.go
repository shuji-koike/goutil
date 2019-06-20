package goutil

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// HTTPStaticFile ...
type HTTPStaticFile string

func (s HTTPStaticFile) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if len(s) == 0 {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, err := os.Open(string(s))
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	defer file.Close()
	io.Copy(res, file)
}

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

// HTTPhandleGetStatus ...
func HTTPhandleGetStatus(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	MemStats := runtime.MemStats{}
	runtime.ReadMemStats(&MemStats)
	HTTPServeJSON(res, req, map[string]interface{}{
		"rev":      os.Getenv("CIRCLE_SHA1"),
		"MemStats": MemStats,
	})
}
