package repositorydatabase

import "store/infrastucture/data/infradatabasesqlx"

// RepositoryDispMobile repository for dispositives mobiles
type RepositoryDispMobile struct {
	infradatabasesqlx.IInfraRepositorySqlx
	//interfaces.IRepositoryBase
}

func (rd RepositoryDispMobile) CarregarGoogleId(codigoVeiculo int) string {

	// sql := `
	// 	select oav.google_id
	// 	from veiculos_dispositivos vd
	// 	inner join dispositivos d on vd.codigodispositivo = d.codigo
	// 	inner join tipos_dispositivos td on d.codigotipodispositivo = td.codigo
	// 	inner join op_assist_vinculacao oav on oav.mobile_id = d.numero_serie
	// 	where vd.codigoveiculo = :codigoveiculo
	// 	and vd.removido = 'F'
	// 	and d.apagado = 'F'
	// 	and td.identificador = 11
	// 	and oav.datahora_desvinculacao is null`

	sql := `
	select c.name
      from clients c
     where c.id = :id`

	s := struct {
		ID int
	}{
		codigoVeiculo,
	}

	var ret string
	rd.PrepareNamedGetOne(&ret, sql, s)

	return ret
}
