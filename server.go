package main

import (
	"log"
	"net/http"

	"github.com/Grimmore/Sentient3/database"
	"github.com/Grimmore/Sentient3/modelos"
	"github.com/Grimmore/Sentient3/storage"
	"github.com/gin-gonic/gin"
)

var DirAScan string = ""

func Init(router *gin.Engine) {
	router.Static("/assets", "static")
	router.LoadHTMLGlob("templates/*")

	web := router.Group("/web")

	web.GET("/sentient/api/dir", func(ctx *gin.Context) {
		a, e := storage.Explore(DirAScan)
		if e == nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.POST("/sentient/api/dir", func(ctx *gin.Context) {
		DirAScan = ctx.PostForm("info")
	})

	web.POST("/sentient/api/dir/add", func(ctx *gin.Context) {
		_, e := modelos.GetElementByPath(ctx.PostForm("path"))
		b := modelos.Element{
			IsDir:    ctx.PostForm("isdir") == "true",
			Name:     ctx.PostForm("name"),
			Path:     ctx.PostForm("path"),
			Watching: ctx.PostForm("watch") == "true",
			Excluded: ctx.PostForm("excluded") == "true",
		}

		if e != nil {
			log.Printf("Error: %v", e)
			_, s := modelos.NewElement(b.Name, b.Path, b.IsDir, b.Excluded, b.Watching)
			if s != nil {
				log.Printf("Error al crear: %v", s)
			}
		} else {
			if !b.Excluded && !b.Watching {
				s := modelos.DeleteElements(b.Path)
				if s != nil {
					log.Printf("Error al crear: %v", s)
				}
			}
			s := modelos.UpdateElement(b.Name, b.Path, b.IsDir, b.Excluded, b.Watching)
			if s != nil {
				log.Printf("Error al crear: %v", s)
			}
		}

		watch.Close()
		LeerDirs()
		watcher()

		h := modelos.GetAllElements()

		//especifica que el contenido de la api es un json
		ctx.Request.Header.Add("Content-Type", "application/json")
		//escribe en la ruta de la api el json
		ctx.JSON(http.StatusOK, h)

	})

	web.GET("/sentient/api/log", func(ctx *gin.Context) {
		db := database.GetDB()
		eventos := []modelos.Evento{}
		//busca en la base de datos el modelo Evento
		err := db.Find(&eventos).Error
		if err != nil {
			log.Println(err)
			return
		}
		if err == nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, eventos)
		}
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/alerts", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/directorios", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
