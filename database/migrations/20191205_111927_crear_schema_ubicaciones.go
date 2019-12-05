package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaUbicaciones_20191205_111927 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaUbicaciones_20191205_111927{}
	m.Created = "20191205_111927"

	migration.Register("CrearSchemaUbicaciones_20191205_111927", m)
}

// Run the migrations
func (m *CrearSchemaUbicaciones_20191205_111927) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS ubicaciones;")
	m.SQL("ALTER SCHEMA ubicaciones OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,ubicaciones;")
}

// Reverse the migrations
func (m *CrearSchemaUbicaciones_20191205_111927) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS ubicaciones CASCADE")
}
