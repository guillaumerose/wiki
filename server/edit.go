package server

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func (s *Server) edit(w http.ResponseWriter, r *http.Request) {
	relative := string(s.getPageName(r))
	bin, _ := ioutil.ReadFile(filepath.Join(s.basePath, relative+".md"))
	edit.Execute(w, data{
		"Title": relative,
		"Path":  "/" + relative,
		"Text":  string(bin),
	})
}
