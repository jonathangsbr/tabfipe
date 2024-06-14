package veiculo

import (
	"strconv"
	"strings"

	"github.com/jonathangsbr/tabfipe-api-gateway/v2/entity/veiculoHelper"
)

type VeiculoModel struct {
	CodigoTabelaReferencia uint16 `json:"codigoTabelaReferencia"`
	CodigoMarca            string `json:"codigoMarca"`
	CodigoModelo           uint16 `json:"codigoModelo"`
	CodigoTipoVeiculo      uint8  `json:"codigoTipoVeiculo"`
	AnoModelo              string `json:"anoModelo"`
	CodigoTipoCombustivel  string `json:"codigoTipoCombustivel"`
	TipoConsulta           string `json:"tipoConsulta"`
}

func (v *VeiculoModel) SetTipoVeiculo(t uint8) {
	v.CodigoTipoVeiculo = t
}

func (v *VeiculoModel) SetTabelaReferencia(t uint16) {
	v.CodigoTabelaReferencia = t
}

func (v *VeiculoModel) SetMarca(m veiculoHelper.VeiculoMarca) {
	v.CodigoMarca = m.CodigoMarca
}

func (v *VeiculoModel) SetCodigoModelo(cm uint16) {
	v.CodigoModelo = cm
}

func (v *VeiculoModel) SetAnoComb(a veiculoHelper.Ano) {
	anocomb := strings.Split(a.AnoValue, "-")
	v.AnoModelo = anocomb[0]
	v.CodigoTipoCombustivel = anocomb[1]
}

func (v *VeiculoModel) SetTipoConsulta(s string) {
	v.TipoConsulta = s
}

func (v *VeiculoModel) ToString() string {
	return "CodigoTabelaReferencia: " + strconv.Itoa(int(v.CodigoTabelaReferencia)) +
		"\nCodigoMarca: " + v.CodigoMarca +
		"\nCodigoModelo: " + strconv.Itoa(int(v.CodigoModelo)) +
		"\nCodigoTipoVeiculo: " + strconv.Itoa(int(v.CodigoTipoVeiculo)) +
		"\nAnoModelo: " + v.AnoModelo +
		"\nCodigoTipoCombustivel: " + v.CodigoTipoCombustivel +
		"\nTipoConsulta: " + v.TipoConsulta
}

type VeiculoResponse struct {
	Valor            string `json:"Valor"`
	Marca            string `json:"Marca"`
	Modelo           string `json:"Modelo"`
	AnoModelo        uint16 `json:"AnoModelo"`
	Combustivel      string `json:"Combustivel"`
	CodigoFipe       string `json:"CodigoFipe"`
	MesReferencia    string `json:"MesReferencia"`
	Autenticacao     string `json:"Autenticacao"`
	TipoVeiculo      uint8  `json:"TipoVeiculo"`
	SiglaCombustivel string `json:"SiglaCombustivel"`
	DataConsulta     string `json:"DataConsulta"`
}

func (v *VeiculoResponse) ToString() string {
	return "Valor: " + v.Valor +
		"\nMarca: " + v.Marca +
		"\nModelo: " + v.Modelo +
		"\nAnoModelo: " + anoModeloFormat(v.AnoModelo) +
		"\nCombustivel: " + v.Combustivel +
		"\nCodigoFipe: " + v.CodigoFipe +
		"\nMesReferencia: " + v.MesReferencia +
		"\nAutenticacao: " + v.Autenticacao +
		"\nTipoVeiculo: " + strconv.Itoa(int(v.TipoVeiculo)) +
		"\nSiglaCombustivel: " + v.SiglaCombustivel +
		"\nDataConsulta: " + v.DataConsulta
}

func anoModeloFormat(a uint16) string {
	str := strconv.Itoa(int(a))
	if str == "32000" {
		return "Zero KM"
	} else {
		return str
	}
}
