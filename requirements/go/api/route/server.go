package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func BankServer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		return
	}
	res, err := json.Marshal(ExpectedTransations[id].Transation)
	if err != nil {
		fmt.Fprintf(w, `{"error": "wrong params"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}