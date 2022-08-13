package main

import (
	"fmt"
	"log"
	"net/http"
	"projego/controllers"
	"projego/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(models.Tugtug{})
	if err != nil {
		panic("Koneksi ke DB gagal")
	}

	tugtugController := &controllers.TugtugController{}

	router := httprouter.New()
	router.GET("/", tugtugController.Index)
	router.GET("/create", tugtugController.Create)
	router.POST("/create", tugtugController.Create)
	router.GET("/edit/:id", tugtugController.Edit)
	router.POST("/update/:id_update", tugtugController.Update)
	router.POST("/selesai/:id_update", tugtugController.Finished)
	router.GET("/hapus/:id_update", tugtugController.Delete)

	port := ":8081"
	fmt.Printf("Aplikasi jalan di http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))

	fmt.Println("Disini")
}
