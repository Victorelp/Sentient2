package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/Grimmore/Sentient3/modelos"
	"github.com/fsnotify/fsnotify"
)

var watch *fsnotify.Watcher

//Funcion que inicializa el watcher que supervizará los directorios que se encuentran en env.json
func watcher() {
	if len(directorios) == 0 {
		return
	}
	var err error
	if watch != nil {
		watch.Close()
	}

	watch, err = fsnotify.NewWatcher()
	//Trata el error de no ser posible inicilizar el warcher
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}

	//crea un canal para pazar el watcher
	done := make(chan bool)
	go func() {
		//ciera el canal al final del hilo
		defer close(done)

		for {
			select {
			//Events envia el filesystem de los eventos de cambio.
			case event, ok := <-watch.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
				//LLama a la funcion auxiliar que tratara con la base de datos y creará una nueva instancia del objeto
				modelos.NewEvento(event.Name, event.Op.String())
				//Si el evento de cambio la ruta es de tipo Create comprobara si la ruta es un directorio y de serlo lo añadira junto a sus subdirectorios al listado de rutas a supervizar
				if event.Op.Has(fsnotify.Create) {
					info, err := os.Stat(event.Name)
					if err == nil && info.IsDir() {
						leerIterativo(event.Name)
					}
				}
				//recibe y trata los errores del watcher en la ruta
			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				log.Println("error: ", err)
			}
		}

	}()
	// recorre el listado de directorios de env.json y los añadira junto a sus subdirectorios al listado de rutas a supervizar
	for _, d := range directorios {
		err = leerIterativo(d)

		if err != nil {
			log.Println("error: ", err)
		}
	}

}

//Comprueba si el directorio se encuentra en la lista de exclusiones
func esEscluido(v string) bool {
	for _, e := range excluciones {
		if e == v {
			return true
		}
	}
	return false
}

//Funcion que añade de forma iterativa los directorios y subdirectorios del mismo al watcher, donde b es la ruta a supervizar
func leerIterativo(b string) error {
	cola := []string{b}
	for len(cola) > 0 {
		p := cola[0]
		cola = cola[1:]
		//de ser un directorio excluido salta a proxima iteracion el bucle
		if esEscluido(p) {
			continue
		}
		//se añade la ruta al watcher para ser supervizada
		err := watch.Add(p)
		if err != nil {
			fmt.Println(p)
			return err
		}
		//medainte la libreria os se obtienen los subdirectorios de la ruta b
		dire, err := os.ReadDir(p)
		if err != nil {
			return err
		}
		for _, d := range dire {
			//se comprueba q d sea un directorio
			if d.IsDir() {
				//mediante path.join une la ruta padre b con el subdirectorio hijo y luego lo añade al arreglo cola con el metodo append
				cola = append(cola, path.Join(p, d.Name()))
			}
		}
	}

	return nil
}
