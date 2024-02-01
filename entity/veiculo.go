package veiculo

import "strconv"

type Veiculo struct {
	CodigoTabelaReferencia uint16 `json:"codigoTabelaReferencia"`
	CodigoMarca            uint16 `json:"codigoMarca"`
	CodigoModelo           uint16 `json:"codigoModelo"`
	CodigoTipoVeiculo      uint8  `json:"codigoTipoVeiculo"`
	AnoModelo              string `json:"anoModelo"`
	CodigoTipoCombustivel  uint8  `json:"codigoTipoCombustivel"`
	TipoConsulta           string `json:"tipoConsulta"`
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

func (v VeiculoResponse) ToString() string {
	return "Valor: " + v.Valor +
		"\nMarca: " + v.Marca +
		"\nModelo: " + v.Modelo +
		"\nAnoModelo: " + strconv.Itoa(int(v.AnoModelo)) +
		"\nCombustivel: " + v.Combustivel +
		"\nCodigoFipe: " + v.CodigoFipe +
		"\nMesReferencia: " + v.MesReferencia +
		"\nAutenticacao: " + v.Autenticacao +
		"\nTipoVeiculo: " + strconv.Itoa(int(v.TipoVeiculo)) +
		"\nSiglaCombustivel: " + v.SiglaCombustivel +
		"\nDataConsulta: " + v.DataConsulta
}
