package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projego/config"
	"projego/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TugtugController struct{}

func (controller *TugtugController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Welcome\n")
	// ini bisa diganti pakai function
	// db, err := gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
	db, err := config.DBConnection()

	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	var tugtug []models.Tugtug
	db.Find(&tugtug)

	// datas := map[string]interface{}{}
	datas := map[string]interface{}{
		"Tugtug": tugtug,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

}

func (controller *TugtugController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Welcome\n")
	db, err := gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		// fmt.Println(r.FormValue("task"))
		if r.FormValue("assignee") == "" || r.FormValue("task") == "" || r.FormValue("deadline") == "" {
			log.Println("Ada isian yang kosong")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			tugas := models.Tugtug{
				Assignee: r.FormValue("assignee"),
				Task:     r.FormValue("task"),
				Deadline: r.FormValue("deadline"),
				IsFinish: false,
			}

			result := db.Create(&tugas)
			if result.Error != nil {
				log.Println(result.Error)
			}

			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else if r.Method == "GET" {
		files := []string{
			"./views/base.html",
			"./views/create.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Print(err.Error())
			return
		}

		datas := map[string]interface{}{}

		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Print(err.Error())
			return
		}
	}

}

func (controller *TugtugController) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// http.Error(w, params.ByName("id"), http.StatusOK)
	// fmt.Fprint(w, "Welcome\n")
	db, err := gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
	} else if r.Method == "GET" {
		files := []string{
			"./views/base.html",
			"./views/edit.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Print(err.Error())
			return
		}

		var tugas models.Tugtug
		db.Where("ID = ?", params.ByName("id")).Find(&tugas)
		data := map[string]interface{}{
			"Tugtug":   tugas,
			"IDUpdate": params.ByName("id"),
		}

		fmt.Print(data)
		err = htmlTemplate.ExecuteTemplate(w, "base", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Print(err.Error())
			return
		}
	}
}

func (controller *TugtugController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// http.Error(w, params.ByName("id"), http.StatusOK)
	// fmt.Fprint(w, "Welcome\n")
	if r.FormValue("assignee") == "" || r.FormValue("task") == "" || r.FormValue("deadline") == "" {
		log.Println("Ada isian yang kosong")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		db, err := gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		tugtugID := params.ByName("id_update")
		var tugtug models.Tugtug
		db.Where("ID = ?", tugtugID).First(&tugtug)

		tugtug.Assignee = r.FormValue("assignee")
		tugtug.Task = r.FormValue("task")
		tugtug.Deadline = r.FormValue("deadline")

		db.Save(&tugtug)
		fmt.Println(tugtug)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (controller *TugtugController) Finished(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	tugtugID := params.ByName("id_update")
	var tugtug models.Tugtug
	db.Where("ID = ?", tugtugID).First(&tugtug)
	// db.Find(&tugtug, params.ByName("id_update"))

	tugtug.IsFinish = !tugtug.IsFinish

	db.Save(&tugtug)
	fmt.Println(tugtug)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (controller *TugtugController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.DBConnection()
	if err != nil {
		panic(err.Error())
	}

	tugtugID := params.ByName("id_update")
	var tugtug models.Tugtug
	db.Delete(&tugtug, tugtugID)

	http.Redirect(w, r, "/", http.StatusFound)
}
