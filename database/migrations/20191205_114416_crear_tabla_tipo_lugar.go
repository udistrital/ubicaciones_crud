package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoLugar_20191205_114416 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoLugar_20191205_114416{}
	m.Created = "20191205_114416"

	migration.Register("CrearTablaTipoLugar_20191205_114416", m)
}

// Run the migrations
func (m *CrearTablaTipoLugar_20191205_114416) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS ubicaciones.tipo_lugar ( id serial NOT NULL, nombre character varying(50) NOT NULL, descripcion character varying(250), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion timestamp NOT NULL, fecha_modificacion timestamp NOT NULL, CONSTRAINT pk_tipo_lugar PRIMARY KEY (id), CONSTRAINT uq_nombre_tipo_lugar UNIQUE (nombre) );")
	m.SQL("ALTER TABLE ubicaciones.tipo_lugar OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE ubicaciones.tipo_lugar IS 'Almacena los parametros de tipo_lugar';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.id IS 'Identificador unico de la tabla';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parametro.';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.descripcion IS 'Descripción opcional del parámetro.';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.activo IS 'Campo de tipo boolean que inidica si el parametro esta activo';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.numero_orden IS ' En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.fecha_creacion IS 'Fecha de creación del registro';")
	m.SQL("COMMENT ON COLUMN ubicaciones.tipo_lugar.fecha_modificacion IS 'Fecha de la última modificación del registro';")
	m.SQL("COMMENT ON CONSTRAINT pk_tipo_lugar ON ubicaciones.tipo_lugar  IS 'Llave primaria de la tabla';")
	m.SQL("COMMENT ON CONSTRAINT uq_nombre_tipo_lugar ON ubicaciones.tipo_lugar  IS 'Constraint unique para que el nombre de tipo_lugar no se repita';")
}

// Reverse the migrations
func (m *CrearTablaTipoLugar_20191205_114416) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS ubicaciones.tipo_lugar")
}
