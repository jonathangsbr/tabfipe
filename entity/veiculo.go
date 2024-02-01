package veiculo

type Veiculo struct {
	CodigoTabelaReferencia uint16 `json:"codigoTabelaReferencia"`
	CodigoMarca            uint16 `json:"codigoMarca"`
	CodigoModelo           uint16 `json:"codigoModelo"`
	AnoModelo              string `json:"anoModelo"`
	CodigoTipoCombustivel  uint16 `json:"codigoTipoCombustivel"`
	TipoVeiculo            string `json:"tipoVeiculo"`
	ModeloCodigoExterno    string `json:"modeloCodigoExterno"`
	TipoConsulta           string `json:"tipoConsulta"`
}
