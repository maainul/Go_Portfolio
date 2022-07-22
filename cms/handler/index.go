package handler

import (
	"html/template"
	"net/http"
	"time"

	"portfolio/svcUtils/logging"

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

type IndexTempData struct {
	CSRFField   template.HTML
	FormAction  string
	FormErrors  map[string]string
	FormMessage map[string]string
	FormName    string
	Form        GetInTouch
}

func (s *Server) indexHander(w http.ResponseWriter, r *http.Request) {
	logging.FromContext(r.Context()).WithField("method", "indexHander")
	data := IndexTempData{
		CSRFField:  csrf.TemplateField(r),
		Form:       GetInTouch{},
		FormAction: index,
		FormName:   "index.html",
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
	data := IndexTempData{
		CSRFField:   csrf.TemplateField(r),
		FormAction:  index,
		FormMessage: map[string]string{"SuccessMessage": "Thanks for your valueable Message."},
		FormName:    "index.html",
	}
	s.loadIndexTemplate(w, r, data)
}
