package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransation(t *testing.T) {
	t.Run("Return user id 1 transations", func(t *testing.T) {
		user := ExpectedTransations[1]
		request := newRequestTransations(user.ID)
		response := httptest.NewRecorder()
		
		BankServer(response, request)
		receive := handleResponseBody(response)

		checkBodyRequest(t, receive, user.Transation)
	})

	t.Run("Return user id 2 transations", func(t *testing.T) {
		user := ExpectedTransations[2]
		request := newRequestTransations(user.ID)
		response := httptest.NewRecorder()
		
		BankServer(response, request)
		receive := handleResponseBody(response)

		checkBodyRequest(t, receive, user.Transation)
	})
}

func newRequestTransations(id int) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/clientes/%d/extrato", id), nil)
	return request
}

func handleResponseBody(response *httptest.ResponseRecorder) TransationsResponse {
	var receive TransationsResponse
	err := json.Unmarshal(response.Body.Bytes(), &receive)
	if err != nil {
		panic(err)
	}
	return receive
}

func checkBodyRequest(t *testing.T, receive, expected TransationsResponse) {
	t.Helper()
	if receive.Balance != expected.Balance ||
	receive.LatestTransactions[0] != expected.LatestTransactions[0] ||
	receive.LatestTransactions[1] != expected.LatestTransactions[1] {
		t.Errorf("expected %v, but has %v", expected, receive)
	}
}