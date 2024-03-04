package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExtract(t *testing.T) {
	store := TestBankStore{}
	store.New()
	server := &BankServer{&store}

	t.Run("Return user id 1 transations", func(t *testing.T) {
		user := store.ExpectedTransations[1]
		expected, _ := json.Marshal(user.Transation)
		response := httptest.NewRecorder()
		request := createRequest(
			http.MethodGet,
			fmt.Sprintf("/clientes/%d/extrato", user.ID),
			[]byte{},
		)
		
		server.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkBodyResponse(t, response.Body.Bytes(), expected)
	})

	t.Run("Return user id 2 transations", func(t *testing.T) {
		user := store.ExpectedTransations[2]
		expected, _ := json.Marshal(user.Transation)
		response := httptest.NewRecorder()
		request := createRequest(
			http.MethodGet,
			fmt.Sprintf("/clientes/%d/extrato", user.ID),
			[]byte{},
		)
		
		server.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkBodyResponse(t, response.Body.Bytes(), expected)
	})

	t.Run("return 404 when don't find id", func(t *testing.T) {
		userId := 6
		response := httptest.NewRecorder()
		request := createRequest(
			http.MethodGet,
			fmt.Sprintf("/clientes/%d/extrato", userId),
			[]byte{},
		)
		
		server.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusNotFound)
		checkErrorResponse(t, response.Body.Bytes(), []byte(`{"Error": "Unable to perform action."}`))
	})
}