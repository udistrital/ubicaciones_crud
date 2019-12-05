package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaLugar_20191205_112903 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaLugar_20191205_112903{}
	m.Created = "20191205_112903"

	migration.Register("CrearTablaLugar_20191205_112903", m)
}

// Run the migrations
func (m *CrearTablaLugar_20191205_112903) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS ubicaciones.lugar ( id serial NOT NULL, nombre character varying(100) NOT NULL, tipo_lugar_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion timestamp NOT NULL, fecha_modificacion timestamp NOT NULL, CONSTRAINT pk_lugar PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE ubicaciones.lugar OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE ubicaciones.lugar IS 'Tabla que almacena los lugares vinculados a una ubicación de forma Jerarquica  Pais, Depto, Municipio, Ciudad, Localidades, Barrios, etc.';")
	m.SQL("COMMENT ON COLUMN ubicaciones.lugar.id IS 'Identificador unico de la tabla';")
	m.SQL("COMMENT ON COLUMN ubicaciones.lugar.nombre IS 'Corresponde al nombre del lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.lugar.activo IS 'Campo que indica si el lugar es activo o no';")
	m.SQL("COMMENT ON COLUMN ubicaciones.lugar.fecha_creacion IS 'Fecha de creación del registro';")
	m.SQL("COMMENT ON COLUMN ubicaciones.lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';")
	m.SQL("COMMENT ON CONSTRAINT pk_lugar ON ubicaciones.lugar IS 'Llave primaria de la tabla';")
}

// Reverse the migrations
func (m *CrearTablaLugar_20191205_112903) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS ubicaciones.lugar")
}
