package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	veiculo "github.com/jonathangsbr/tabfipe-api-gateway/entity"
)

var urlEntryPoint string = "https://veiculos.fipe.org.br/api/veiculos//ConsultarValorComTodosParametros"

func EntryPointRequest(v veiculo.Veiculo) (veiculo.VeiculoResponse, error) {
	veiculo := veiculo.VeiculoResponse{}

	body, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Erro ao converter o tipo veiculo para JSON")
		return veiculo, err
	}
	payload := bytes.NewBuffer(body)

	resp, err := http.Post(urlEntryPoint, "application/json", payload)
	if err != nil {
		fmt.Println("Erro ao fazer requisição")
		return veiculo, err
	}

	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler dados do request")
		return veiculo, err
	}

	err = json.Unmarshal([]byte(bod), &veiculo)
	if err != nil {
		fmt.Println("Erro ao converter JSON para o tipo veiculo")
		return veiculo, err
	}

	defer resp.Body.Close()
	return veiculo, nil
}
