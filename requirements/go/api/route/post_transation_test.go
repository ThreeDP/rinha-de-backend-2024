package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ThreeDP/rinha-de-backend/db"
)

func TestPostTransation(t *testing.T) {
	store := TestBankStore{}
	store.New()
	server := &BankServer{&store}

	ExpextedTransations := []db.RequestTransation {
		{Value: 1000, Type: 'c', Description: "descricao"},
	}
	t.Run("return statusOK when create a transation for user id 1", func(t *testing.T) {
		response := httptest.NewRecorder()
		body, _ := json.Marshal(ExpextedTransations[0])
		request := createRequest(
			http.MethodPost,
			fmt.Sprintf("/clientes/%d/transacoes", 1),
			body,
		)

		server.ServeHTTP(response, request)

		checkStatusCode(t, response.Code, http.StatusOK)
		checkBodyResponse(t, response.Body.Bytes(), []byte(`{"limite" : 100000, "saldo" : -9098}`))
	})
}

