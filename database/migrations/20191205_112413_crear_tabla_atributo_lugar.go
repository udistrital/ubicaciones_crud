package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaAtributoLugar_20191205_112413 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaAtributoLugar_20191205_112413{}
	m.Created = "20191205_112413"

	migration.Register("CrearTablaAtributoLugar_20191205_112413", m)
}

// Run the migrations
func (m *CrearTablaAtributoLugar_20191205_112413) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS ubicaciones.atributo_lugar ( id serial NOT NULL, nombre character varying(50) NOT NULL, descripcion character varying(250), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion timestamp NOT NULL, fecha_modificacion timestamp NOT NULL, CONSTRAINT pk_atributo_lugar PRIMARY KEY (id), CONSTRAINT uq_nombre_atributo_lugar UNIQUE (nombre) );")
	m.SQL("ALTER TABLE ubicaciones.atributo_lugar OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE ubicaciones.atributo_lugar IS 'Tabla que almacena los atributos que puede tener un lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.id IS 'Identificador unico de la tabla';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.nombre IS 'Indica el nombre del atributo para un lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.descripcion IS 'Descripción opcional del atributo del lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere. ';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.activo IS 'Estado del atributo del lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD.';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.fecha_creacion IS 'Fecha de creación del registro';")
	m.SQL("COMMENT ON COLUMN ubicaciones.atributo_lugar.fecha_modificacion IS 'Fecha de la útlima modificación del registro';")
	m.SQL("COMMENT ON CONSTRAINT pk_atributo_lugar ON ubicaciones.atributo_lugar  IS 'Llave primaria de la tabla';")
	m.SQL("COMMENT ON CONSTRAINT uq_nombre_atributo_lugar ON ubicaciones.atributo_lugar  IS 'Constraint unique para que el nombre de atributo_lugar no se repita';")
}

// Reverse the migrations
func (m *CrearTablaAtributoLugar_20191205_112413) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS ubicaciones.atributo_lugar")
}
