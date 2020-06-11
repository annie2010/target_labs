package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gopherland/labs2/profiling/internal/fib"
)

const (
	contentTypeHDR = "Content-Type"
	contentType    = "application/json; charset=utf-8"
)

type Result struct {
	Number    int `json:"n"`
	Fibonacci int `json:"fib"`
}

type Results []Result

func FibHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	var res Results
	for i := 0; i <= n; i++ {
		res = append(res, Result{Number: i, Fibonacci: fib.Compute(i)})
	}
	buff, err := json.Marshal(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contentTypeHDR, contentType)
	if _, err := w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FibHandlerIter(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.ParseInt(r.FormValue("n"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	res := make(Results, 0, n)
	for i := 0; i <= int(n); i++ {
		res = append(res, Result{Number: i, Fibonacci: fib.ComputeIter(i)})
	}
	buff, err := json.Marshal(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contentTypeHDR, contentType)
	if _, err := w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
