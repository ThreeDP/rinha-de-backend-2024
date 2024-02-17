package route

type Transation struct {
	Value int64 `json:"valor"`
	Type byte `json:"tipo"`
	Description string `json:"descricao"`
	CreateAt string `json:"realizada_em"`
}

type Balance struct {
	Total int64 `json:"total"`
	CreateAt string `json:"data_extrato"`
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
	GetTransations(id int) ([]byte, error) 
}
