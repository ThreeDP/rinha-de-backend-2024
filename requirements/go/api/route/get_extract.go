package route

import (
	"net/http"
	"strconv"
)

func (s *BankServer) getExtract(w http.ResponseWriter, r *http.Request) {
	uri, thisPath := s.parseURLPath("/clientes/:id/extrato", r.URL.Path)
	if thisPath {
		id, err := strconv.Atoi(uri["id"])
		if err != nil {
			return
		}
		res, err := s.Store.GetExtract(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
