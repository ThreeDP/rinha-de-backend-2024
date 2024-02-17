package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	switch r.Method {
		case http.MethodPost:
			s.postTransation(w, r)
		case http.MethodGet:
			s.getTransations(w, r)
	}
}

func (s *BankServer) postTransation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *BankServer) splitFields(str string, step string) []string {
	parts := strings.Split(str, step)
	var res []string
	for _, part := range parts {
		if len(part) > 0 {
			res = append(res, part)
		} 
	}
	return res
}

func (s *BankServer) parseURLPath(uri, path string) (map[string]string, bool) {
	uriKeys := map[string]string{}
	sUri, sPath := s.splitFields(uri, "/"), s.splitFields(path, "/")
	if len(sUri) == len(sPath) {
		for i, u := range sUri {
			key, ok := strings.CutPrefix(u, ":")
			field := u
			if ok {
				_, ok := uriKeys[key]
				if !ok {
					uriKeys[key] = sPath[i]
				}
				field = uriKeys[key]
			}
			if field != sPath[i] {
				return nil, false
			}
		}
		return uriKeys, true
	}
	return nil, false
}

func (s *BankServer) getTransations(w http.ResponseWriter, r *http.Request) {
	uri, thisPath := s.parseURLPath("/clientes/:id/extrato", r.URL.Path)
	if thisPath {
		id, err := strconv.Atoi(uri["id"])
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
}
