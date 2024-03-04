package route

import (
	"testing"
)

func TestParseURLPath(t *testing.T) {
	store := TestBankStore{}
	store.New()
	server := &BankServer{&store}

	t.Run("Test pass a normal Path", func(t *testing.T) {
		uri := "/client/info"
		path := "/client/info"

		received, thisPath := server.parseURLPath(uri, path)
		expected := map[string]string{}

		compareMapStrings(t, received, expected)
		if thisPath == false {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("Test pass a normal Path with :id", func(t *testing.T) {
		uri := "/client/:id/bill"
		path := "/client/123/bill"

		received, thisPath := server.parseURLPath(uri, path)
		expected := map[string]string{"id": "123"}

		compareMapStrings(t, received, expected)
		if thisPath == false {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("Test pass two equal :id in same path return only first :id value", func(t *testing.T) {
		uri := "/client/:id/:id"
		path := "/client/123/123"

		received, thisPath := server.parseURLPath(uri, path)
		expected := map[string]string{"id": "123"}

		compareMapStrings(t, received, expected)
		if thisPath == false {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("Test pass two equal :id in different path return only first :id value", func(t *testing.T) {
		uri := "/client/:id/:id"
		path := "/client/123/321"

		received, thisPath := server.parseURLPath(uri, path)

		if received != nil {
			t.Errorf("Expected nil, but has %v", received)
		}
		if thisPath == true {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("Test pass :id and :idBill returns map with two values", func(t *testing.T) {
		uri := "/client/:id/:idBill"
		path := "/client/123/321"

		received, thisPath := server.parseURLPath(uri, path)
		expected := map[string]string{"id": "123", "idBill": "321"}

		compareMapStrings(t, received, expected)
		if thisPath == false {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("validade path amount directorys", func(t *testing.T){
		uri := "/client/:id/:idBill"
		path := "/client/123/321/nopath"

		received, thisPath := server.parseURLPath(uri, path)

		if received != nil {
			t.Errorf("Expected nil, but has %v", received)
		}
		if thisPath == true {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("validade path amount directorys with extra /", func(t *testing.T){
		uri := "/client/:id/:idBill"
		path := "/client/123/321/"

		received, thisPath := server.parseURLPath(uri, path)
		expected := map[string]string{"id": "123", "idBill": "321"}

		compareMapStrings(t, received, expected)
		if thisPath == false {
			t.Error("the path is not the same as the uri")
		}
	})

	t.Run("validade path directory values", func(t *testing.T){
		uri := "/client/"
		path := "/login/"

		received, thisPath := server.parseURLPath(uri, path)

		if received != nil {
			t.Errorf("Expected nil, but has %v", received)
		}
		if thisPath == true {
			t.Error("the path is not the same as the uri")
		}
	})
}