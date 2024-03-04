package route

import (
	"net/http"

	"github.com/ThreeDP/rinha-de-backend/db"
)

type BankServer struct {
	Store db.IDBQueries
}

func (s *BankServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postTransation(w, r)
	case http.MethodGet:
		s.getExtract(w, r)
	}
}
