package db

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transation struct {
	UserId int64
	Value int64 `json:"valor"`
	Type byte `json:"tipo"`
	Description string `json:"descricao"`
	CreateAt time.Time
}

type RequestTransation struct {
	Value int `json:"valor"`
	Type byte `json:"tipo"`
	Description string `json:"descricao"`
}

type Balance struct {
	Total int64 `json:"total"`
	CreateAt time.Time `json:"data_extrato"`
	Limit int64 `json:"limite"`
}

type TransationsResponse struct {
	Balance Balance `json:"saldo"`
	LatestTransactions []Transation `json:"ultimas_transacoes"`
}

type DBTransations struct {
	ID int
	Transation TransationsResponse
}

type IDBQueries interface {
	GetExtract(id int) ([]byte, error)
	PostTransation(dbT Transation) ([]byte, error)
}

type DBQueries struct {

}

func (db *DBQueries) PostTransation(dbT Transation) ([]byte, error) {
	return []byte{}, nil
}

func (db *DBQueries) GetExtract(id int) ([]byte, error) {
	res, err := json.Marshal(TransationsResponse{
		Balance: Balance{Total: -9098, CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 217753, time.UTC), Limit: 1000000},
		LatestTransactions: []Transation{
			{Value: 10, Type: 'c', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 543030, time.UTC)},
			{Value: 90000, Type: 'd', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 41, 543030, time.UTC)},
		},
	})
	if err != nil {
		errMsg := fmt.Sprint(`{"Error": "Unable to perform action."}`)
		return []byte(errMsg), err
	}
	return res, nil
}

// type User struct {
// 	Id int64 `json:"id"`
// 	Name string `json:"name"`
// 	DebitBalance int64 `json:"debit_balance"`
// 	CreditLimit int64 `json:"credit_limit"`
// 	CreateAt time.Time `json:"criado_em"`
// }

// type Transation struct {
// 	UserId int64 `json:user_id`
// 	Value int64 `json:"valor"`
// 	Type byte `json:"tipo"`
// 	Description string `json:"descricao"`
// 	CreateAt time.Time `json:"realizada_em"`
// }

// type Balance struct {
// 	Total int64 `json:"total"`
// 	CreateAt time.Time `json:"data_extrato"`
// 	Limit int64 `json:"limite"`
// }

// type TransationsResponse struct {
// 	Balance Balance `json:"saldo"`
// 	LatestTransactions []Transation `json:"ultimas_transacoes"`
// }

// type IDBQueries interface {
// 	GetExtract(id int) ([]byte, error)
// 	PostTransation(dbT Transation) error
// 	New()
// }

// type DBQueries struct {
// 	User []User
// 	Balance []Balance
// 	Transation []Transation
// }

// func (db *DBQueries) New() {
// 	db.User = []User{
// 		{Id: 1, Name: "o barato sai caro", DebitBalance: 0, CreditLimit: 1000 * 100, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC)},
// 		{Id: 2, Name: "zan corp ltda", DebitBalance: 0, CreditLimit: 800 * 100, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC)},
// 		{Id: 3, Name: "les cruders", DebitBalance: 0, CreditLimit: 10000 * 100, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC)},
// 		{Id: 4, Name: "padaria joia de cocaia", DebitBalance: 0, CreditLimit: 100000 * 100, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC)},
// 		{Id: 5, Name: "kid mais", DebitBalance: 0, CreditLimit: 5000 * 100, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC)},
// 	}
// 	db.Balance = []Balance{
// 		{Total: -9098, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC), Limit: 1000000},
// 		{Total: 999900, CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 217753, time.UTC), Limit: 49990},
// 	}
// 	db.Transation = []Transation{
// 		{UserId: 1, Value: 10, Type: 'c', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 543030, time.UTC)},
// 		{UserId: 1, Value: 90000, Type: 'd', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 543030, time.UTC)},
// 		{UserId: 2, Value: 10, Type: 'c', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 543030, time.UTC)},
// 		{UserId: 2, Value: 100, Type: 'd', Description: "descricao", CreateAt: time.Date(2024, 1, 17, 2, 34, 38, 543030, time.UTC)},
// 	}
// }

// func (db *DBQueries) GetExtract(id int) ([]byte, error) {
// 	res, ok := db.db[id]
// 	if ok != true {
// 		errMsg := `{"error": "Error"}`
// 		return []byte(errMsg), errors.New("Error")
// 	}
// 	return json.Marshal(res.Transation)
// }