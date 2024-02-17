package route

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
)

type BankServer struct {
	Store IDBQueries
}

type DBQueries struct {

}

func (db *DBQueries) GetTransations(id int) ([]byte, error) {
	res, err := json.Marshal(TransationsResponse{
		Balance: Balance{Total: -9098, CreateAt: "2024-01-17T02:34:41.217753Z", Limit: 1000000},
		LatestTransactions: []Transation{
			{Value: 10, Type: 'c', Description: "descricao", CreateAt: "2024-01-17T02:34:38.543030Z"},
			{Value: 90000, Type: 'd', Description: "descricao", CreateAt: "2024-01-17T02:34:38.543030Z"},
		},
	})
	if err != nil {
		errMsg := fmt.Sprintf(`{"error": %s}`, err)
		return []byte(errMsg), err
	}
	return res, nil
}

func (s *BankServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lc := len("/clientes/")
	le := len("/extrato")
	lp := len(r.URL.Path)
	id, err := strconv.Atoi(r.URL.Path[lc : lp - le])
	if err != nil {
		return
	}
	res, err := s.Store.GetTransations(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return 
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
