package route

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/ThreeDP/rinha-de-backend/db"
)

func (s *BankServer) postTransation(w http.ResponseWriter, r *http.Request) {
	uri, thisPath := s.parseURLPath("/clientes/:id/transacoes", r.URL.Path)
	if thisPath {
		id, err := strconv.Atoi(uri["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"error": "Invalid request"}`))
			return
		}

		body := db.Transation{}
		if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"error": "Invalid request"}`))
			return
		}
		body.UserId = int64(id)

		res, err := s.Store.PostTransation(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"error": "Unable to perform action."}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	w.WriteHeader(http.StatusOK)
}