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

var ExpectedTransations = map[int]DBTransations{
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
