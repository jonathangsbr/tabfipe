package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	veiculo "github.com/jonathangsbr/tabfipe-api-gateway/v2/entity/veiculo"
)

var urlEntryPoint string = "https://veiculos.fipe.org.br/"
var urlEndPoint = ""
var postHeaderContentType = "application/json"

func GetVeiculoI(v veiculo.VeiculoModel) (veiculo.VeiculoResponse, error) {
	urlEndPoint = "api/veiculos//ConsultarValorComTodosParametros"
	veiculo := veiculo.VeiculoResponse{}

	pl, err := json.Marshal(v)
	if err != nil {
		err := fmt.Errorf("erro ao converter o tipo veiculo para JSON")
		return veiculo, err
	}
	payload := bytes.NewBuffer(pl)

	resp, err := http.Post(urlEntryPoint+urlEndPoint, postHeaderContentType, payload)
	if err != nil || resp.StatusCode != 200 {
		err := fmt.Errorf("erro ao fazer a requisição para: \n\t%s, statusCode: %d", resp.Request.URL, resp.StatusCode)
		return veiculo, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("erro ao ler dados do request")
		return veiculo, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal([]byte(body), &veiculo)
	if err != nil {
		err := fmt.Errorf("erro ao converter JSON para o tipo veiculo")
		return veiculo, err
	}

	return veiculo, nil
}
