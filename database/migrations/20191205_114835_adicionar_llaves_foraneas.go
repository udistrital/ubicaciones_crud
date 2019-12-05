package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AdicionarLlavesForaneas_20191205_114835 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AdicionarLlavesForaneas_20191205_114835{}
	m.Created = "20191205_114835"

	migration.Register("AdicionarLlavesForaneas_20191205_114835", m)
}

// Run the migrations
func (m *AdicionarLlavesForaneas_20191205_114835) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE ubicaciones.lugar ADD CONSTRAINT fk_lugar_tipo_lugar FOREIGN KEY (tipo_lugar_id) REFERENCES ubicaciones.tipo_lugar (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.valor_atributo_lugar ADD CONSTRAINT fk_valor_atributo_lugar_atributo_lugar FOREIGN KEY (atributo_lugar_id) REFERENCES ubicaciones.atributo_lugar (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.valor_atributo_lugar ADD CONSTRAINT fk_valor_atributo_lugar_lugar FOREIGN KEY (lugar_id) REFERENCES ubicaciones.lugar (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.relacion_lugares ADD CONSTRAINT fk_lugar_hijo FOREIGN KEY (lugar_hijo_id) REFERENCES ubicaciones.lugar (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE ubicaciones.relacion_lugares ADD CONSTRAINT fk_lugar_padre FOREIGN KEY (lugar_padre_id) REFERENCES ubicaciones.lugar (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
}

// Reverse the migrations
func (m *AdicionarLlavesForaneas_20191205_114835) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE ubicaciones.lugar DROP CONSTRAINT IF EXISTS fk_lugar_tipo_lugar CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.valor_atributo_lugar DROP CONSTRAINT IF EXISTS fk_valor_atributo_lugar_atributo_lugar CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.valor_atributo_lugar DROP CONSTRAINT IF EXISTS fk_valor_atributo_lugar_lugar CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.relacion_lugares DROP CONSTRAINT IF EXISTS fk_lugar_hijo CASCADE;")
	m.SQL("ALTER TABLE ubicaciones.relacion_lugares DROP CONSTRAINT IF EXISTS fk_lugar_padre CASCADE;")
}
