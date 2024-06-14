package fipeMoto

import (
	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculo"
	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculoHelper"
	"github.com/jonathangsbr/tabfipe-api-gateway/pkg/fipeVeiculo"
)

var codTipoVeiculo uint8 = 2

type moto struct{}

func NewMoto() *moto {
	c := &moto{}
	c.init()
	return c
}

func (c *moto) init() {
	fipeVeiculo.SetCodTipoVeiculo(codTipoVeiculo)
}

func (c *moto) GetCodTipoVeiculo() uint8 {
	return fipeVeiculo.GetCodTipoVeiculo()
}

func (c *moto) GetCodTabelaReferenciaRecente() uint16 {
	return fipeVeiculo.GetCodTabelaReferenciaRecente()
}

func (c *moto) GetVeiculoModeloBase() veiculo.VeiculoModel {
	return fipeVeiculo.GetVeiculoModeloBase()
}

func (c *moto) GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error) {
	return fipeVeiculo.GetMarcasObj()
}

func (c *moto) GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error) {
	return fipeVeiculo.GetModelos(m)
}

func (c *moto) GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error) {
	return fipeVeiculo.GetAnos(v, m)

}

func (c *moto) GetVeiculo(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error) {
	return fipeVeiculo.GetVeiculo(v, a)
}
