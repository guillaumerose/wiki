package server

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func (s *Server) show(w http.ResponseWriter, r *http.Request) {
	relative := string(s.getPageName(r))

	bin, err := ioutil.ReadFile(filepath.Join(s.basePath, relative))
	text := bin
	if err != nil || len(bin) == 0 {
		text = emptyPageText
	}

	show.Execute(w, data{
		"Title": relative,
		"Path":  "/" + relative + "/edit",
		"Text":  bytesAsHTML(parsedMarkdown(text)),
	})
}
