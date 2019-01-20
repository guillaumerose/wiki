package server

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) show(w http.ResponseWriter, r *http.Request) {
	relative := strings.Replace(string(s.getPageName(r)), ".md", "", -1)

	bin, err := ioutil.ReadFile(filepath.Join(s.basePath, relative+".md"))
	text := bin
	if err != nil || len(bin) == 0 {
		text = emptyPageText
	}

	files, err := ioutil.ReadDir(s.basePath)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var links []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".md" {
			links = append(links, file.Name())
		}
	}

	show.Execute(w, data{
		"Title": relative,
		"Path":  "/" + relative + "/edit",
		"Text":  bytesAsHTML(parsedMarkdown(text)),
		"Files": links,
	})
}
