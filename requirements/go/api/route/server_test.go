package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
)

type TestBankStore struct {
	ExpectedTransations map[int]DBTransations
}

func (t *TestBankStore) GetTransations(id int) ([]byte, error) {
	et, ok := t.ExpectedTransations[id]
	if ok == false {
		err := errors.New("Error")
		errMsg := fmt.Sprintf(`{"error": %s}`, err)
		return []byte(errMsg), err
	}
	res, err := json.Marshal(et.Transation)
	if err != nil {
		errMsg := fmt.Sprintf(`{"error": %s}`, err)
		return []byte(errMsg), err
	}
	return res, nil
}

func TestGetTransation(t *testing.T) {
	store := TestBankStore{}
	store.New()
	server := &BankServer{&store}

	t.Run("Return user id 1 transations", func(t *testing.T) {
		user := store.ExpectedTransations[1]
		request := newRequestTransations(user.ID)
		response := httptest.NewRecorder()
		
		server.ServeHTTP(response, request)
		received := handleResponseBody(response)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkBodyRequest(t, received, user.Transation)
	})

	t.Run("Return user id 2 transations", func(t *testing.T) {
		user := store.ExpectedTransations[2]
		request := newRequestTransations(user.ID)
		response := httptest.NewRecorder()
		
		server.ServeHTTP(response, request)
		received := handleResponseBody(response)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkBodyRequest(t, received, user.Transation)
	})

	t.Run("return 404 when don't find id", func(t *testing.T) {
		userId := 6
		request := newRequestTransations(userId)
		response := httptest.NewRecorder()
		
		server.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func newRequestTransations(id int) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/clientes/%d/extrato", id), nil)
	return request
}

func handleResponseBody(response *httptest.ResponseRecorder) TransationsResponse {
	var received TransationsResponse
	err := json.Unmarshal(response.Body.Bytes(), &received)
	if err != nil {
		panic(err)
	}
	return received
}

func checkStatusCode(t *testing.T, received, expected int) {
	t.Helper()
	if received != expected {
		t.Errorf("expected status %d, but has status %d", expected, received)
	}
}

func checkBodyRequest(t *testing.T, received, expected TransationsResponse) {
	t.Helper()
	if received.Balance != expected.Balance ||
	received.LatestTransactions[0] != expected.LatestTransactions[0] ||
	received.LatestTransactions[1] != expected.LatestTransactions[1] {
		t.Errorf("expected %v, but has %v", expected, received)
	}
}

func (t *TestBankStore) New() {
	t.ExpectedTransations = map[int]DBTransations{
		1: {
			ID: 1,
			Transation: TransationsResponse{
				Balance: Balance{Total: -9098, CreateAt: "2024-01-17T02:34:41.217753Z", Limit: 1000000},
				LatestTransactions: []Transation{
					{Value: 10, Type: 'c', Description: "descricao", CreateAt: "2024-01-17T02:34:38.543030Z"},
					{Value: 90000, Type: 'd', Description: "descricao", CreateAt: "2024-01-17T02:34:38.543030Z"},
				},
			},
		},
		2: {
			ID: 2,
			Transation: TransationsResponse{
				Balance: Balance{Total: 999900, CreateAt: "2024-01-17T02:34:41.217753Z", Limit: 49990},
				LatestTransactions: []Transation{
					{Value: 10, Type: 'c', Description: "descricao", CreateAt: "2024-01-17T02:34:38.543030Z"},
					{Value: 100, Type: 'd', Description: "descricao", CreateAt: "2024-01-17T02:34:38.543030Z"},
				},
			},
		},
	}
}