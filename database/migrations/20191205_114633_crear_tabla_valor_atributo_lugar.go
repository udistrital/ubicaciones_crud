package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaValorAtributoLugar_20191205_114633 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaValorAtributoLugar_20191205_114633{}
	m.Created = "20191205_114633"

	migration.Register("CrearTablaValorAtributoLugar_20191205_114633", m)
}

// Run the migrations
func (m *CrearTablaValorAtributoLugar_20191205_114633) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS ubicaciones.valor_atributo_lugar ( id serial NOT NULL, valor character varying(150) NOT NULL, lugar_id integer NOT NULL, atributo_lugar_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion timestamp NOT NULL, fecha_modificacion timestamp NOT NULL, CONSTRAINT pk_valor_atributo_lugar PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE ubicaciones.valor_atributo_lugar OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE ubicaciones.valor_atributo_lugar IS 'Tabla que almacena el valor de los atributos que tiene un lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.id IS 'Identificador unico de la tabla';")
	m.SQL("COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.valor IS 'Valor para el atributo del lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.activo IS 'Indica el estado del registro';")
	m.SQL("COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.fecha_creacion IS 'Fecha de creación del registro';")
	m.SQL("COMMENT ON COLUMN ubicaciones.valor_atributo_lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';")
	m.SQL("COMMENT ON CONSTRAINT pk_valor_atributo_lugar ON ubicaciones.valor_atributo_lugar  IS 'Llave primaria de la tabla';")
}

// Reverse the migrations
func (m *CrearTablaValorAtributoLugar_20191205_114633) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS ubicaciones.valor_atributo_lugar")
}
