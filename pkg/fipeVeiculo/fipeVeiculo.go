package fipeVeiculo

import (
	"strings"

	"github.com/jonathangsbr/tabfipe-api-gateway/v2/entity/veiculo"
	"github.com/jonathangsbr/tabfipe-api-gateway/v2/entity/veiculoHelper"
	"github.com/jonathangsbr/tabfipe-api-gateway/v2/internal/gateway"
)

var codTipoVeiculo uint8 = 0

type FipeVeiculo interface {
	GetCodTipoVeiculo() uint8
	GetCodTabelaReferenciaRecente() uint16
	GetVeiculoModeloBase() veiculo.VeiculoModel
	GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error)
	GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error)
	GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error)
	GetVeiculo(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error)
}

func GetCodTipoVeiculo() uint8 {
	return codTipoVeiculo
}

func SetCodTipoVeiculo(cod uint8) {
	codTipoVeiculo = cod
}

func GetCodTabelaReferenciaRecente() uint16 {
	return gateway.CodTabelaReferenciaRecente
}

func GetVeiculoModeloBase() veiculo.VeiculoModel {
	return veiculo.VeiculoModel{CodigoTipoVeiculo: codTipoVeiculo, TipoConsulta: "tradicional", CodigoTabelaReferencia: GetCodTabelaReferenciaRecente()}
}

func GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error) {
	return gateway.GetMarcasI(codTipoVeiculo)
}

func GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error) {
	veiculo := GetVeiculoModeloBase()
	veiculo.SetMarca(m)
	modelos, err := gateway.GetModelosI(veiculo)
	return veiculo, modelos, err
}

func GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error) {
	v.SetCodigoModelo(m.CodigoModelo)
	anos, err := gateway.GetAnosI(*v)
	for i, v := range anos {
		if strings.Contains(v.AnoComb, "32000") {
			str := strings.Replace(v.AnoComb, "32000", "Zero KM", -1)
			// str2 := strings.Replace(v.AnoComb, "32000", "Zero KM", -1)
			anos[i].AnoComb = str
			// anos[i].AnoValue = str2
		}
	}
	return anos, err

}

func GetVeiculo(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error) {
	v.SetAnoComb(a)
	return gateway.GetVeiculoI(*v)
}
