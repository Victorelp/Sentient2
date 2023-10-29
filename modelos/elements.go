package modelos

import (
	"errors"

	"github.com/Grimmore/Sentient3/database"
)

type Element struct {
	IsDir    bool
	Name     string
	Path     string
	Excluded bool
	Watching bool
}

//obtiene todos los objetos Evento de la base de datos
func GetAllElements() []Element {
	var element []Element
	db := database.GetDB()
	db.Find(&element)
	return element
}

//Crea un objeto evento y lo añade como tupla a la base da datos
func NewElement(name string, path string, a bool, b bool, c bool) (*Element, error) {
	if name == "" || path == "" {
		return nil, errors.New("empty fields")
	}
	element := Element{
		IsDir:    a,
		Name:     name,
		Path:     path,
		Excluded: b,
		Watching: c,
	}

	db := database.GetDB()
	err := db.Create(&element).Error
	return &element, err
}

//borra una instancia de objeto evento con la id pasada por parametro
func DeleteElements(path string) error {
	db := database.GetDB()
	return db.Where("Path == ?", path).Delete(&Element{}).Error
}

func GetElementByPath(p string) (*Element, error) {
	db := database.GetDB()
	var e Element
	err := db.Model(&e).Where("Path == ?", p).First(&e).Error
	return &e, err
}

func UpdateElement(name, path string, a, b, c bool) error {
	Element := Element{}
	Element.Path = path

	db := database.GetDB()
	model := db.Model(&Element).Where("Path = ?", path)
	if path != "" {
		model = model.Update("Path", path)
	}
	if name != "" {
		model = model.Update("Name", name)
	}
	model = model.Update("IsDir", a)
	model = model.Update("Excluded", b)
	model = model.Update("Watching", c)

	return model.Error
}
