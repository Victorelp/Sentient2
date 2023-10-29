package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/Grimmore/Sentient3/modelos"
)

//se declaran de forma global las varibles que seran rellenadas con la información de env.json

//en directorios se planea guardar todas las rutas a añadir al watcher
var directorios []string

//guarda el puerto donde se montará la api
var puerto string

//guarda las rutas que serán excluidas del watcher aun siendo una subruta de algún elemento de directorios
var excluciones []string

func LeerDirs() {
	a := modelos.GetAllElements()
	for _, i := range a {
		_, b := os.ReadDir(i.Path)
		if b == nil {
			if i.Excluded {
				excluciones = append(excluciones, i.Path)
			} else {
				directorios = append(directorios, i.Path)
			}
		} else {
			continue
		}

	}
}

func getEnv() {
	//se crea la variable que almacenara el contenido del json que se asume existe en la ruta
	file, err := os.Open("./env.json")
	//de no existir o no poder abrirse se crea un nuevo json con configuraciones por defecto y termina el metodo
	if err != nil {
		CrearJson()
		//guarda en file el contenido del archivo pasado por parámetro, err guarda el error que puede retornar el método al fallar en abrir el archivo
		file, err := os.Open("./env.json")
		if err != nil {
			return
		}
		//se cierra el file una vez se carga la configuración puesto que ya cumplió su función y de otra forma no se puede modificar o borrar, etc...
		defer file.Close()
		//se lee el contenido del json almacenado en file
		env, err := io.ReadAll(file)
		if err != nil {
			return
		}
		//lee la información contenida en el json que se le pasa por parámetro
		leerJson(env)
		return
	}
	//esta comprobación es necesaria en caso de que haya errores al leer la información.
	//guarda en env la información contenida en la variable pasada por parámetro, guarda en err el error en caso de no poder leer dicha info
	env, err := io.ReadAll(file)
	//de no ser posible leer el json que ya se cargo creara un nuevo json
	if err != nil {
		defer file.Close()
		CrearJson()
		file, err := os.Open("./env.json")
		if err != nil {
			return
		}
		defer file.Close()
		env, err := io.ReadAll(file)
		if err != nil {
			return
		}
		leerJson(env)
		return
	}
	//si el json se abrió correctamente y su contenido pudo ser leido
	defer file.Close()
	leerJson(env)
}

func leerJson(env []byte) {
	//crea un mapa el cual hace corresponder una clave string a un objeto de tipo any (cualquier tipo)
	m := map[string]any{}
	//parse el contenido del json y lo almacena el el puntero m
	err := json.Unmarshal(env, &m)
	if err != nil {
		log.Fatal(err)

	}
	//guarda en las variables el contenido del mapa con las claves "directorios","exclusiones","puerto" respectivamente los cuales corresponden con los campos de mismo nombre de env.json

	port := m["puerto"]

	//comprueba que los campos se hayan podido cargar de forma correcta o no esten vacios

	if port == nil {
		log.Fatal("Problema en el campo puerto de env.json")
	}

	var ok bool
	//es necesario usar una funcion auxiliar para castear el contenido de directorio ya que este devuelve []interface{}, ok guardara un booleano que comprobará la realización de la operación

	puerto, ok = port.(string)
	if !ok {
		log.Fatal("Configuracion del puerto no es correcta")
	}
}

//necesaria para convertir el arreglo interfaces en un arreglo de strings
func casteo(a any) (b []string, c bool) {
	aux := []string{}
	x, ok := a.([]interface{})
	if !ok {
		return
	}
	//para cada elemento y en rango del arreglo x donde _ es una variable anónima que simbolizaria la posición, en caso de ser necesaria debe ser redeclarada a por ejemplo: i
	for _, y := range x {
		//iguala el arreglo de strings aux a la concatenacion del mismo con el elemento y que representa la instancia actual de x
		aux = append(aux, y.(string))
	}
	return aux, true
}
