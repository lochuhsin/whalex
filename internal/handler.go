package internal

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")

	if paths[1] == "" {
		logrus.Info("root request")
		w.Write([]byte("hello world"))
		return
	}

	upstreamHost, ok := GetUpstreamRouter().Get(paths[1])
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("Unable to find desired upstream"))
		return
	}
	path := strings.Join(paths[2:], "/")

	var builder strings.Builder
	builder.WriteString("http://")
	builder.WriteString(upstreamHost)
	builder.WriteString("/")
	builder.WriteString(path)

	http.Redirect(w, r, builder.String(), http.StatusFound)
}
