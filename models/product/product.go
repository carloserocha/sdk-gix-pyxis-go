package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/carloserocha/sdk-gix-pyxis-go/pyxis"
)

const GetProducts = "ServPubGetProdutoEcommerce/V1"

type DefaultResponse struct {
	Status   string `json:"status"`
	Response string `json:"retorno"`
}

type DefaultResponseProduct struct {
	Products []ProductStruct `json:"produtos"`
}

// Corpo de Requisição para busca de todos produtos
type AllProduct struct {
	GetDescription string `json:"getficha"`
	GetImage       string `json:"getfoto"`
}

// Corpo de Requisição para busca de um produto
type ProductBySku struct {
	GetDescription string `json:"getficha"`
	GetImage       string `json:"getfoto"`
	Sku            string `json:"cod"`
}

// Corpo de Requisição para busca de produtos alterados por data
type ProductByDateModified struct {
	GetDescription string `json:"getficha"`
	GetImage       string `json:"getfoto"`
	Date           string `json:"data"`
	Hour           string `json:"hora"`
}

type ProductStruct struct {
	Sku               int     `json:"cod"`
	Description       string  `json:"desc"`
	Unity             string  `json:"um"`
	EAN               string  `json:"codBarras"`
	FactorySku        string  `json:"codFabrica"`
	Reference         string  `json:"referencia"`
	Weight            float32 `json:"peso"`
	Status            string  `json:"inat"`
	DropShipping      string  `json:"dropShipping"`
	ShortDescription  string  `json:"ficha"`
	Deposit           string  `json:"priorizaDeposito"`
	TechnicalFeatures struct {
		Liters  float32 `json:"litros"`
		Width   float32 `json:"largura"`
		Height  float32 `json:"altura"`
		Depth   float32 `json:"profundidade"`
		Package float32 `json:"embalagem"`
		Color   string  `json:"cor"`
		Model   string  `json:"modelo"`
		Voltage string  `json:"voltagem"`
	} `json:"caracteristicas"`
	ManagerialDivision struct {
		Factory           string `json:"fabricante"`
		Type              string `json:"tipo"`
		SubType           string `json:"subtipo"`
		Line              string `json:"linha"`
		Family            string `json:"familia"`
		Brand             string `json:"marca"`
		MasterSku         int    `json:"codMaster"`
		MasterDescription string `json:"descMaster"`
	} `json:"divisaogerencial"`
}

const DefaultGetDescription = "S"
const DefaultGetImage = "N"

func GetProductBySku(sku string) interface{} {
	body := &ProductBySku{GetDescription: DefaultGetDescription, GetImage: DefaultGetImage, Sku: sku}

	raw, _ := json.Marshal(body)

	resp := pyxis.NewRequest(http.MethodPost, GetProducts, raw)

	p := handleProduct(string(resp))

	return p
}

func GetProductByDateModified(t time.Time) interface{} {

	epoch := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day()) // yyyymmmdd format
	period := fmt.Sprintf("%02d%02d00", t.Hour(), t.Minute())        // only hour and minute.

	body := &ProductByDateModified{Date: epoch, Hour: period, GetDescription: DefaultGetDescription, GetImage: DefaultGetImage}

	raw, _ := json.Marshal(body)

	resp := pyxis.NewRequest(http.MethodPost, GetProducts, raw)

	p := handleProduct(string(resp))

	return p
}

func GetAllProducts() interface{} {
	body := &AllProduct{GetDescription: DefaultGetDescription, GetImage: DefaultGetImage}

	raw, _ := json.Marshal(body)

	resp := pyxis.NewRequest(http.MethodPost, GetProducts, raw)

	p := handleProduct(string(resp))

	return p
}

// Default Schema
func handleProduct(product string) interface{} {
	// product := strings.ReplaceAll(body, `\"`, `"`)
	// fmt.Println(product, body)
	d := DefaultResponse{}
	err := json.Unmarshal([]byte(product), &d)
	if err != nil {
		panic(err)
	}

	p := DefaultResponseProduct{}
	err = json.Unmarshal([]byte(d.Response), &p)
	if err != nil {
		panic(err)
	}

	return p
}
