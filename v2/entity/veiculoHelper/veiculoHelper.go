package veiculoHelper

type VeiculoMarca struct {
	Marca       string `json:"label"`
	CodigoMarca string `json:"value"`
}

type Modelo struct {
	ModeloNome   string `json:"label"`
	CodigoModelo uint16 `json:"value"`
}

type Ano struct {
	AnoComb  string `json:"label"`
	AnoValue string `json:"value"`
}
