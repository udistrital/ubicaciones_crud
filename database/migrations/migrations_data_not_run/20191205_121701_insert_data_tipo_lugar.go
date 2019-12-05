package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertDataTipoLugar_20191205_121701 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertDataTipoLugar_20191205_121701{}
	m.Created = "20191205_121701"

	migration.Register("InsertDataTipoLugar_20191205_121701", m)
}

// Run the migrations
func (m *InsertDataTipoLugar_20191205_121701) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO ubicaciones.tipo_lugar (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) VALUES (1, 'PAIS', 'País', 'P', TRUE, 1.00, LOCALTIMESTAMP, LOCALTIMESTAMP), (2, 'CIUDAD', 'Ciudad', 'C', TRUE, 2.00, LOCALTIMESTAMP, LOCALTIMESTAMP), (3, 'LOCALIDAD', 'Localidad', 'L', TRUE, 3.00, LOCALTIMESTAMP, LOCALTIMESTAMP), (4, 'DEPARTAMENTO', 'Departamento', 'D', TRUE, 4.00, LOCALTIMESTAMP, LOCALTIMESTAMP);")
}

// Reverse the migrations
func (m *InsertDataTipoLugar_20191205_121701) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM ubicaciones.tipo_lugar WHERE id BETWEEN 1 AND 4;")
}
