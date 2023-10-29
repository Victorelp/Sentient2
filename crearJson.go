package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

func CrearJson() {
	//pregunta si el sistema operativo es windows
	if runtime.GOOS == "windows" {
		//crea un mapa que asigna la correspondencia de string a un elemento de tipo cualquiera es decir el mapa["directorios"] contiene un arreglo de strings
		m := map[string]any{
			"puerto": "5647",
		}
		//marshal retorna el json codificado en la variable m que se almacenara en a, donde err guarda el error que retorna el método en caso de fallar
		a, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		//Crea un json en la ruta base con el nombre "env.json", contenido guardado en a y con los permisos de lectura y escritura
		os.WriteFile("./env.json", a, 0777)
		return
	}
	//para el resto de sistemas define la ruta base
	m := map[string]any{
		"puerto": "5647",
	}

	a, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./env.json", a, 0777)

}
