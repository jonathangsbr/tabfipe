package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	veiculo "github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculo"
	veiculoHelper "github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculoHelper"
)

var CodTabelaReferenciaRecente uint16

type tabelaReferencia struct {
	Codigo uint16
	Mes    string
}

func init() {
	codTabelaReferenciaRecentearr, err := getTabelaReferenciaI()
	if err != nil {
		log.Fatalf("erro ao inicializar gatewayHelper {listarTabelaReferencia()}, causa: \n\t%s", err)
	}
	CodTabelaReferenciaRecente = codTabelaReferenciaRecentearr[0].Codigo
}

func getTabelaReferenciaI() ([]tabelaReferencia, error) {
	urlEndPoint = "api/veiculos//ConsultarTabelaDeReferencia"
	tabelaReferencia := []tabelaReferencia{}

	resp, err := http.Post(urlEntryPoint+urlEndPoint, postHeaderContentType, nil)
	if err != nil || resp.StatusCode != 200 {
		err := fmt.Errorf("erro ao fazer a requisição para: \n\t%s, statusCode: %d", resp.Request.URL, resp.StatusCode)
		return tabelaReferencia, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("erro ao ler dados do request")
		return tabelaReferencia, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal([]byte(body), &tabelaReferencia)
	if err != nil {
		err := fmt.Errorf("erro ao converter JSON para array do tipo tabelaReferencia")
		return tabelaReferencia, err
	}

	return tabelaReferencia, nil
}

type marca struct {
	CodigoTabelaReferencia uint16
	CodigoTipoVeiculo      uint8
}

func GetMarcasI(tipoVeiculo uint8) ([]veiculoHelper.VeiculoMarca, error) {
	urlEndPoint = "api/veiculos//ConsultarMarcas"
	marcas := []veiculoHelper.VeiculoMarca{}

	pl, err := json.Marshal(marca{CodigoTabelaReferencia: CodTabelaReferenciaRecente, CodigoTipoVeiculo: tipoVeiculo})
	if err != nil {
		err := fmt.Errorf("erro ao converter o tipo marcas para JSON")
		return marcas, err
	}
	payload := bytes.NewBuffer(pl)

	resp, err := http.Post(urlEntryPoint+urlEndPoint, postHeaderContentType, payload)
	if err != nil || resp.StatusCode != 200 {
		err := fmt.Errorf("erro ao fazer a requisição para: \n\t%s, statusCode: %d", resp.Request.URL, resp.StatusCode)
		return marcas, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("erro ao ler dados do request")
		return marcas, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal([]byte(body), &marcas)
	if err != nil {
		err := fmt.Errorf("erro ao converter JSON para array do tipo marca")
		return marcas, err
	}

	return marcas, nil
}

type modeloAno struct {
	Modelos []veiculoHelper.Modelo `json:"modelos"`
}

func GetModelosI(v veiculo.VeiculoModel) ([]veiculoHelper.Modelo, error) {
	urlEndPoint = "api/veiculos//ConsultarModelos"
	modeloAno := modeloAno{}

	pl, err := json.Marshal(v)
	if err != nil {
		err := fmt.Errorf("erro ao converter o tipo veiculoModelo para JSON")
		return modeloAno.Modelos, err
	}
	payload := bytes.NewBuffer(pl)

	resp, err := http.Post(urlEntryPoint+urlEndPoint, postHeaderContentType, payload)
	if err != nil || resp.StatusCode != 200 {
		err := fmt.Errorf("erro ao fazer a requisição para: \n\t%s, statusCode: %d", resp.Request.URL, resp.StatusCode)
		return modeloAno.Modelos, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("erro ao ler dados do request")
		return modeloAno.Modelos, err
	}

	err = json.Unmarshal([]byte(body), &modeloAno)
	if err != nil {
		err := fmt.Errorf("erro ao converter JSON para o tipo modelo")
		return modeloAno.Modelos, err
	}
	defer resp.Body.Close()

	return modeloAno.Modelos, nil
}

func GetAnosI(v veiculo.VeiculoModel) ([]veiculoHelper.Ano, error) {
	urlEndPoint = "api/veiculos//ConsultarAnoModelo"
	ano := []veiculoHelper.Ano{}

	pl, err := json.Marshal(v)
	if err != nil {
		err := fmt.Errorf("erro ao converter o tipo veiculo para JSON")
		return ano, err
	}
	payload := bytes.NewBuffer(pl)

	resp, err := http.Post(urlEntryPoint+urlEndPoint, postHeaderContentType, payload)
	if err != nil || resp.StatusCode != 200 {
		err := fmt.Errorf("erro ao fazer a requisição para: \n\t%s, statusCode: %d", resp.Request.URL, resp.StatusCode)
		return ano, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("erro ao ler dados do request")
		return ano, err
	}
	defer resp.Body.Close()

	err = json.Unmarshal([]byte(body), &ano)
	if err != nil {
		err := fmt.Errorf("erro ao converter JSON par ao tipo ano")
		return ano, err
	}

	return ano, nil
}
