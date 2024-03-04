package route

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
	"bytes"
	"github.com/ThreeDP/rinha-de-backend/db"
)

/* Help Config Functions */
func createRequest(method, url string, body []byte) *http.Request {
	request, _ := http.NewRequest(method, url, bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	return request
}

/* Help Test Functions */
func checkStatusCode(t *testing.T, received, expected int) {
	t.Helper()
	if received != expected {
		t.Errorf("expected status %d, but has status %d", expected, received)
	}
}

func checkBodyResponse(t *testing.T, received, expected []byte) {
	t.Helper()
	if !reflect.DeepEqual(received, expected) {
		t.Errorf("expected %s, but has %s", string(expected), string(received))
	}
}

func checkErrorResponse(t *testing.T, received, expected []byte) {
	t.Helper()
	if !reflect.DeepEqual(received, expected) {
		t.Errorf("Expected Error mensage %s, but has %s", expected, received)
	}
}

func compareMapStrings(t *testing.T, received, expected map[string]string) {
	t.Helper()
	if reflect.DeepEqual(received, expected) == false {
		t.Errorf("Expected %v, but has %v", expected, received)
	}
}

/* DB Test Functions */
type TestBankStore struct {
	Trasations []db.Transation
	ExpectedTransations map[int]db.DBTransations
}

func (t *TestBankStore) GetExtract(id int) ([]byte, error) {
	et, ok := t.ExpectedTransations[id]
	if ok == false {
		err := errors.New("Unable to perform action.")
		errMsg := fmt.Sprintf(`{"Error": "%s"}`, err)
		return []byte(errMsg), err
	}
	res, err := json.Marshal(et.Transation)
	if err != nil {
		errMsg := fmt.Sprintf(`{"Error": "%s"}`, err)
		return []byte(errMsg), err
	}
	return res, nil
}

func (t *TestBankStore) PostTransation(dbT db.Transation) ([]byte, error) {
	t.Trasations = append(t.Trasations, dbT)

	return []byte(`{"limite" : 100000, "saldo" : -9098}`), nil
}

func (t *TestBankStore) New() {
	t.ExpectedTransations = map[int]db.DBTransations{
		1: {
			ID: 1,
			Transation: db.TransationsResponse{
				Balance: db.Balance{Total: -9098, CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 217753, time.UTC), Limit: 1000000},
				LatestTransactions: []db.Transation{
					{Value: 10, Type: 'c', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 543030, time.UTC)},
					{Value: 90000, Type: 'd', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 543030, time.UTC)},
				},
			},
		},
		2: {
			ID: 2,
			Transation: db.TransationsResponse{
				Balance: db.Balance{Total: 999900, CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 217753, time.UTC), Limit: 49990},
				LatestTransactions: []db.Transation{
					{Value: 10, Type: 'c', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 543030, time.UTC)},
					{Value: 100, Type: 'd', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 543030, time.UTC)},
				},
			},
		},
	}
}