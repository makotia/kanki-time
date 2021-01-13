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
	http.HandleFunc("/api/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		p   []string
	)

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		p = strings.Split(r.URL.Path, "/")
		if _, err = os.Stat(path.Join(config.GetConfig().Server.StaticDir, p[len(p)-1]+".png")); err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
	case "POST":
		type reqJSON struct {
			Text string `json:"Text"`
			Type string `json:"Type"`
		}

		var (
			id        string
			b         []byte
			returnMap map[string]string
			reqBody   reqJSON
			buf       bytes.Buffer
		)

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
	case "HEAD":
		w.WriteHeader(http.StatusOK)
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
	default:
		fmt.Println(r.Method)
		w.WriteHeader(http.StatusNotAcceptable)
	}
}
