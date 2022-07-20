package handler

import (
	"html/template"
	"net/http"

	"portfolio/svcUtils/logging"

	"github.com/gorilla/csrf"
)

type IndexTempData struct {
	CSRFField   template.HTML
	FormAction  string
	FormErrors  map[string]string
	FormMessage map[string]string
	FormName    string
}

func (s *Server) indexHander(w http.ResponseWriter, r *http.Request) {
	logging.FromContext(r.Context()).WithField("method", "indexHander")
	data := IndexTempData{
		CSRFField: csrf.TemplateField(r),
		FormName:  "index.html",
	}
	s.loadIndexTemplate(w, r, data)
}

func (s *Server) loadIndexTemplate(w http.ResponseWriter, r *http.Request, data IndexTempData) {
	log := logging.FromContext(r.Context()).WithField("method", "loadIndexTemplate")
	tmpl := s.lookupTemplate(data.FormName)
	if tmpl == nil {
		log.Error("unable to load template")
		http.Redirect(w, r, ErrorPath, http.StatusSeeOther)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Redirect(w, r, ErrorPath, http.StatusSeeOther)
		return
	}
}
