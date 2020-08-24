package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/carloserocha/sdk-gix-pyxis-go/pyxis"
)

const GetStocks = "ServPubGetSaldoEstoqueEcommerce/V1"

type DefaultResponse struct {
	Status   string `json:"status"`
	Response string `json:"retorno"`
}

type DefaultResponseStock struct {
	Stocks []StockStruct `json:"saldos"`
}

type StockStruct struct {
	Sku int     `json:"cod"`
	Qty float32 `json:"saldo"`
}

// Corpo de Requisição para busca de todos estoques
type AllStock struct {
	GetTotal string `json:"getporcentagemEstoqueTotal"`
}

// Corpo de Requisição para busca de estoque um produto
type StockBySku struct {
	Sku      string `json:"cod"`
	GetTotal string `json:"getporcentagemEstoqueTotal"`
}

// Corpo de Requisição para busca de estoques alterados por data
type StockByDateModified struct {
	Date     string `json:"data"`
	Hour     string `json:"hora"`
	GetTotal string `json:"getporcentagemEstoqueTotal"`
}

func GetStockBySku(sku string, getTotal string) interface{} {
	body := &StockBySku{Sku: sku, GetTotal: getTotal}

	raw, _ := json.Marshal(body)

	resp := pyxis.NewRequest(http.MethodPost, GetStocks, raw)

	p := handleStock(string(resp))

	return p
}

func GetStockByDateModified(t time.Time, getTotal string) interface{} {

	epoch := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day()) // yyyymmmdd format
	period := fmt.Sprintf("%02d%02d00", t.Hour(), t.Minute())        // only hour and minute.

	body := &StockByDateModified{Date: epoch, Hour: period, GetTotal: getTotal}

	raw, _ := json.Marshal(body)

	resp := pyxis.NewRequest(http.MethodPost, GetStocks, raw)

	p := handleStock(string(resp))

	fmt.Println(p)

	return p
}

func GetAllStocks(getTotal string) interface{} {
	body := &AllStock{GetTotal: getTotal}

	raw, _ := json.Marshal(body)

	resp := pyxis.NewRequest(http.MethodPost, GetStocks, raw)

	p := handleStock(string(resp))

	return p
}

func handleStock(body string) DefaultResponseStock {
	d := DefaultResponse{}
	err := json.Unmarshal([]byte(body), &d)
	if err != nil {
		panic(err)
	}

	var p DefaultResponseStock
	err = json.Unmarshal([]byte(d.Response), &p)
	if err != nil {
		panic(err)
	}

	return p
}
