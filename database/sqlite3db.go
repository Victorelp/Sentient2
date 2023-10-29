package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	//Constante del nombre de la base de datos
	DataBaseRoute = "database.db"
	DataBaseType  = "sqlite3"
)

var DB *gorm.DB

func conectar() {
	var err error
	//instancia la base de datos usando la libreria gorm donde DataBaseType es el tipo de base de datos(sqlite3, MySQL, PostgreSQL) y DataBaseRoute la ruta, err guarda el error que puede retornar el método
	DB, err = gorm.Open(DataBaseType, DataBaseRoute)

	if err != nil {
		log.Fatal(err)
	}
}

//retorna la base de datos
func GetDB() *gorm.DB {
	if DB == nil {
		conectar()
	}
	return DB
}
