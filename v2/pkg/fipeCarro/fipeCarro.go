package fipeCarro

import (
	"github.com/jonathangsbr/tabfipe-api-gateway/v2/entity/veiculo"
	"github.com/jonathangsbr/tabfipe-api-gateway/v2/entity/veiculoHelper"
	"github.com/jonathangsbr/tabfipe-api-gateway/v2/pkg/fipeVeiculo"
)

var codTipoVeiculo uint8 = 1

type carro struct{}

func NewCarro() *carro {
	c := &carro{}
	c.init()
	return c
}

func (c *carro) init() {
	fipeVeiculo.SetCodTipoVeiculo(codTipoVeiculo)
}

func (c *carro) GetCodTipoVeiculo() uint8 {
	return fipeVeiculo.GetCodTipoVeiculo()
}

func (c *carro) GetCodTabelaReferenciaRecente() uint16 {
	return fipeVeiculo.GetCodTabelaReferenciaRecente()
}

func (c *carro) GetVeiculoModeloBase() veiculo.VeiculoModel {
	return fipeVeiculo.GetVeiculoModeloBase()
}

func (c *carro) GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error) {
	return fipeVeiculo.GetMarcasObj()
}

func (c *carro) GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error) {
	return fipeVeiculo.GetModelos(m)
}

func (c *carro) GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error) {
	return fipeVeiculo.GetAnos(v, m)

}

func (c *carro) GetVeiculo(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error) {
	return fipeVeiculo.GetVeiculo(v, a)
}
