package fipemoto

import (
	"encoding/json"

	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculo"
	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculoHelper"
	"github.com/jonathangsbr/tabfipe-api-gateway/internal/gateway"
)

var codTipoMoto uint8 = 2

func GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error) {
	return gateway.GetMarcasI(codTipoMoto)
}

func GetMarcasJSON() (string, error) {
	marcas, err := gateway.GetMarcasI(codTipoMoto)
	if err != nil {
		return "{}", err
	}
	m, err := json.Marshal(marcas)
	return string(m), err
}

func GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error) {
	veiculo := veiculo.VeiculoModel{}
	veiculo.SetTipoVeiculo(codTipoMoto)
	veiculo.SetTipoConsulta("tradicional")
	veiculo.SetMarca(m)
	veiculo.SetTabelaReferencia(gateway.CodTabelaReferenciaRecente)
	modelos, err := gateway.GetModelosI(veiculo)
	return veiculo, modelos, err
}

func GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error) {
	v.SetCodigoModelo(m.CodigoModelo)
	anos, err := gateway.GetAnosI(*v)
	return anos, err

}

func GetCarro(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error) {
	v.SetAnoComb(a)
	return gateway.GetVeiculoI(*v)
}
