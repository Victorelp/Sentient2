package modelos

import (
	"errors"
	"log"
	"os/user"

	"github.com/Grimmore/Sentient3/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Estructura de go la cual simula un modelo
type Evento struct {
	gorm.Model
	Path    string
	Type    string
	Usuario string
}

//obtiene todos los objetos Evento de la base de datos
func GetAllEvento() []Evento {
	var evento []Evento
	db := database.GetDB()
	db.Find(&evento)
	return evento
}

//Crea un objeto evento y lo añade como tupla a la base da datos
func NewEvento(ruta string, tye string) (*Evento, error) {
	if ruta == "" || tye == "" {
		return nil, errors.New("empty fields")
	}
	a, err2 := user.Current()
	if err2 != nil {
		log.Println(err2)
		return nil, nil
	}
	Evento := Evento{
		Path:    ruta,
		Type:    tye,
		Usuario: a.Username,
	}

	db := database.GetDB()
	err := db.Create(&Evento).Error
	return &Evento, err
}

//borra una instancia de objeto evento con la id pasada por parametro
func DeleteEvento(id string) error {
	db := database.GetDB()
	return db.Where("id == ?", id).Delete(&Evento{}).Error
}
