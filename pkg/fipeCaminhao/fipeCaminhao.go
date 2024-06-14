package fipeCaminhao

import (
	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculo"
	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculoHelper"
	"github.com/jonathangsbr/tabfipe-api-gateway/pkg/fipeVeiculo"
)

var codTipoVeiculo uint8 = 3

type caminhao struct{}

func NewCaminhao() *caminhao {
	c := &caminhao{}
	c.init()
	return c
}

func (c *caminhao) init() {
	fipeVeiculo.SetCodTipoVeiculo(codTipoVeiculo)
}

func (c *caminhao) GetCodTipoVeiculo() uint8 {
	return fipeVeiculo.GetCodTipoVeiculo()
}

func (c *caminhao) GetCodTabelaReferenciaRecente() uint16 {
	return fipeVeiculo.GetCodTabelaReferenciaRecente()
}

func (c *caminhao) GetVeiculoModeloBase() veiculo.VeiculoModel {
	return fipeVeiculo.GetVeiculoModeloBase()
}

func (c *caminhao) GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error) {
	return fipeVeiculo.GetMarcasObj()
}

func (c *caminhao) GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error) {
	return fipeVeiculo.GetModelos(m)
}

func (c *caminhao) GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error) {
	return fipeVeiculo.GetAnos(v, m)

}

func (c *caminhao) GetVeiculo(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error) {
	return fipeVeiculo.GetVeiculo(v, a)
}
