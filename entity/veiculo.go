package veiculo

type Veiculo struct {
	CodigoTabelaReferencia uint16
	CodigoMarca            uint16
	CodigoModelo           uint16
	AnoModelo              uint16
	CodigoTipoCombustivel  uint16
	TipoVeiculo            string
	ModeloCodigoExterno    string
	TipoConsulta           string
}

type VeiculoResponse struct {
	CodigoTabelaReferencia uint16 `json:"codigoTabelaReferencia"`
	CodigoMarca            uint16 `json:"codigoMarca"`
	CodigoModelo           uint16 `json:"codigoModelo"`
	AnoModelo              uint16 `json:"anoModelo"`
	CodigoTipoCombustivel  uint16 `json:"codigoTipoCombustivel"`
	TipoVeiculo            string `json:"tipoVeiculo"`
	ModeloCodigoExterno    string `json:"modeloCodigoExterno"`
	TipoConsulta           string `json:"tipoConsulta"`
}
