package modelos

import "github.com/Grimmore/Sentient3/database"

//crea las tablas de la base de datos
func Migrate() {
	db := database.GetDB()
	//donde & hace referencia a la dirección de memeoria en la que se guarda la estructura
	db.AutoMigrate(&Evento{}, &Element{})
}
