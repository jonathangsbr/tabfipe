package fipeCarro

import (
	"encoding/json"
	"strings"

	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculo"
	"github.com/jonathangsbr/tabfipe-api-gateway/entity/veiculoHelper"
	"github.com/jonathangsbr/tabfipe-api-gateway/internal/gateway"
)

var codTipoCarro uint8 = 1

func GetCodTipoCarro() uint8 {
	return codTipoCarro
}

func GetCodTabelaReferenciaRecente() uint16 {
	return gateway.CodTabelaReferenciaRecente
}

func GetVeiculoModeloBase() veiculo.VeiculoModel {
	return veiculo.VeiculoModel{CodigoTipoVeiculo: codTipoCarro, TipoConsulta: "tradicional", CodigoTabelaReferencia: GetCodTabelaReferenciaRecente()}
}

func GetMarcasObj() ([]veiculoHelper.VeiculoMarca, error) {
	return gateway.GetMarcasI(codTipoCarro)
}

func GetMarcasJSON() (string, error) {
	marcas, err := gateway.GetMarcasI(codTipoCarro)
	if err != nil {
		return "{}", err
	}
	m, err := json.Marshal(marcas)
	return string(m), err
}

func GetModelos(m veiculoHelper.VeiculoMarca) (veiculo.VeiculoModel, []veiculoHelper.Modelo, error) {
	// if
	veiculo := GetVeiculoModeloBase()
	veiculo.SetMarca(m)
	modelos, err := gateway.GetModelosI(veiculo)
	return veiculo, modelos, err
}

func GetAnos(v *veiculo.VeiculoModel, m veiculoHelper.Modelo) ([]veiculoHelper.Ano, error) {
	v.SetCodigoModelo(m.CodigoModelo)
	anos, err := gateway.GetAnosI(*v)
	for _, v := range anos {
		if strings.Contains(v.AnoComb, "3200") {
			str := strings.Replace("3200", v.AnoComb, "Zero KM", 1)
			v.AnoComb = str
		}
	}
	return anos, err

}

func GetCarro(v *veiculo.VeiculoModel, a veiculoHelper.Ano) (veiculo.VeiculoResponse, error) {
	v.SetAnoComb(a)
	return gateway.GetVeiculoI(*v)
}
