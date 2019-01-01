package server

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) save(w http.ResponseWriter, r *http.Request) {
	relative := string(s.getPageName(r))
	absolute := filepath.Join(s.basePath, relative + ".md")
	os.MkdirAll(filepath.Dir(absolute), 0755)
	ioutil.WriteFile(absolute, []byte(strings.TrimSpace(r.FormValue("text"))), 0644)
	s.redirect(w, r)
}
