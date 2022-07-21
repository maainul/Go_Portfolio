package handler

import (
	"html/template"
	"net/http"
	"portfolio/svcUtils/logging"
	"time"

	gk "portfolio/api/gunk/v1/admin/getInTouch"

	"github.com/gorilla/csrf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetInTouch struct {
	ID        string
	Name      string
	Email     string
	Message   string
	CreatedAt time.Time
	CreatedBy string
}

type GetInTouchTempData struct {
	CSRFField   template.HTML
	Form        GetInTouch
	FormAction  string
	FormErrors  map[string]string
	FormMessage map[string]string
	FormName    string
}

func (s *Server) getInTouchCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logging.FromContext(ctx).WithField("method", "GetInTouchCreateHandler")
	if err := r.ParseForm(); err != nil {
		errMsg := "parsing form"
		log.WithError(err).Error(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var form GetInTouch
	if err := s.decoder.Decode(&form, r.PostForm); err != nil {
		logging.WithError(err, log).Error("decoding form")
		http.Redirect(w, r, ErrorPath, http.StatusSeeOther)
		return
	}
	_, err := s.getInTouch.CreateGetInTouch(ctx, &gk.CreateGetInTouchRequest{
		GetInTouch: &gk.GetInTouch{
			Name:      form.Name,
			Email:     form.Email,
			Message:   form.Message,
			CreatedAt: timestamppb.Now(),
			CreatedBy: "12345",
		},
	})
	if err != nil {
		logging.WithError(err, log).Error("create GetInTouch failed")
		http.Redirect(w, r, ErrorPath, http.StatusSeeOther)
		return
	}

	data := GetInTouchTempData{
		CSRFField:  csrf.TemplateField(r),
		FormName:   "index.html",
	}
	s.loadGetInTouchTemplate(w, r, data)
}

func (s *Server) loadGetInTouchTemplate(w http.ResponseWriter, r *http.Request, data GetInTouchTempData) {
	log := logging.FromContext(r.Context()).WithField("method", "loadGetInTouchTemplate")
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
