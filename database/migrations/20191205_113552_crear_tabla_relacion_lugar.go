package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRelacionLugar_20191205_113552 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRelacionLugar_20191205_113552{}
	m.Created = "20191205_113552"

	migration.Register("CrearTablaRelacionLugar_20191205_113552", m)
}

// Run the migrations
func (m *CrearTablaRelacionLugar_20191205_113552) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS ubicaciones.relacion_lugares ( id serial NOT NULL, lugar_padre_id integer NOT NULL, lugar_hijo_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion timestamp NOT NULL, fecha_modificacion timestamp NOT NULL, CONSTRAINT pk_relacion_lugares PRIMARY KEY (id), CONSTRAINT uq_hijo_id_relacion_lugares UNIQUE (lugar_hijo_id), CONSTRAINT uq_hijo_id_padre_id_relacion_lugares UNIQUE (lugar_padre_id,lugar_hijo_id) );")
	m.SQL("ALTER TABLE ubicaciones.relacion_lugares OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE ubicaciones.relacion_lugares IS 'Tabla que almacena la relacion Jerarquica entre los lugares';")
	m.SQL("COMMENT ON COLUMN ubicaciones.relacion_lugares.id IS 'Identificador unico de la tabla';")
	m.SQL("COMMENT ON COLUMN ubicaciones.relacion_lugares.lugar_padre_id IS 'Identificador del lugar padre';")
	m.SQL("COMMENT ON COLUMN ubicaciones.relacion_lugares.lugar_hijo_id IS 'Identificador del lugar hijo';")
	m.SQL("COMMENT ON COLUMN ubicaciones.relacion_lugares.activo IS 'Indica si la relación entre los lugares se encuentra activa';")
	m.SQL("COMMENT ON COLUMN ubicaciones.relacion_lugares.fecha_creacion IS 'Fecha de creación del registro';")
	m.SQL("COMMENT ON COLUMN ubicaciones.relacion_lugares.fecha_modificacion IS 'Fecha de la última modificación del registro';")
	m.SQL("COMMENT ON CONSTRAINT pk_relacion_lugares ON ubicaciones.relacion_lugares  IS 'Llave primaria de la tabla';")
	m.SQL("COMMENT ON CONSTRAINT uq_hijo_id_relacion_lugares ON ubicaciones.relacion_lugares  IS 'Restriccion para garantizar la jerarquía entre lugares.';")
	m.SQL("COMMENT ON CONSTRAINT uq_hijo_id_padre_id_relacion_lugares ON ubicaciones.relacion_lugares  IS 'Restriccion para que solo pueda existir una unica relacion entre un padre y un hijo';")
}

// Reverse the migrations
func (m *CrearTablaRelacionLugar_20191205_113552) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS ubicaciones.relacion_lugares")
}
