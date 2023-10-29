package main

import (
	"github.com/Grimmore/Sentient3/modelos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Función principal que ejecuta el resto de funciones del programa
func main() {
	//carga la configuración de env.json
	getEnv()
	//crea la base da datos y sus tablas
	modelos.Migrate()
	//Añade al array directorios que sera recorrido por el watcher los directorios de la base de datos
	LeerDirs()
	//crea en un hilo el watcher y le añade los elementosque se encuentren en la configuración del json
	watcher()

	//crea un server en el puerto especificado en el json y en el comparte via una api el contenido de la base de datos usando el formato json, a su vez mantiene en ejecución el resto del programa
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	Init(router)
	err := router.Run(":5447")
	if err != nil {
		panic(err)
	}
}
