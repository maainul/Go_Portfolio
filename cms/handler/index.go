package handler

import (
	"fmt"
	"net/http"

	"portfolio/svcUtils/logging"
)

func (s *Server) indexHander(w http.ResponseWriter, r *http.Request) {
	logging.FromContext(r.Context()).WithField("method", "index Hander")
	fmt.Println("####################")
	fmt.Println("Index Handler Called")
	fmt.Println("####################")

}
