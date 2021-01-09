package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/makotia/kanki-time/api/config"
	"github.com/makotia/kanki-time/api/util"
)

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/", get)

	http.ListenAndServe(":8080", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	type reqJSON struct {
		Text string `json:"Text"`
		Type string `json:"Type"`
	}

	var (
		id        string
		err       error
		b         []byte
		returnMap map[string]string
		reqBody   reqJSON
		buf       bytes.Buffer
	)

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	io.Copy(&buf, r.Body)

	if err = json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id, err = util.GenImage(reqBody.Text, reqBody.Type); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	returnMap = map[string]string{
		"id":    id,
		"error": "",
	}

	if b, err = json.Marshal(returnMap); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(b))
}

func get(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		p   []string
	)
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		p = strings.Split(r.URL.Path, "/")
		if _, err = os.Stat(path.Join(config.GetConfig().Server.StaticDir, p[len(p)-1]+".png")); err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
	}
}
