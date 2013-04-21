package handlers

import (
	"html/template"
	"net/http"
	"github.com/jakecoffman/gorunner/models"
	"github.com/jakecoffman/gorunner/utils"
	"github.com/gorilla/mux"
)

func Triggers(w http.ResponseWriter, r *http.Request) {
	triggerList := models.GetTriggerList()

	if r.Method == "GET" {
		t := template.Must(template.New("_base.html").Funcs(utils.FuncMap).ParseFiles(
			"web/templates/_base.html",
			"web/templates/_nav.html",
			"web/templates/triggers.html",
		))

		if err := t.Execute(w, triggerList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		trigger := models.Trigger{Name: name}
		models.Append(triggerList, trigger)
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func Trigger(w http.ResponseWriter, r *http.Request) {
	triggerList := models.GetTriggerList()

	vars := mux.Vars(r)
	trigger, err := models.Get(triggerList, vars["trigger"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		t := template.Must(template.New("_base.html").Funcs(utils.FuncMap).ParseFiles(
			"web/templates/_base.html",
			"web/templates/_nav.html",
			"web/templates/trigger.html",
		))

		if err := t.Execute(w, trigger); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "PUT" {
		t := trigger.(models.Trigger)
		t.Schedule = r.FormValue("cron")
		err = models.Update(triggerList, t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "DELETE" {
		models.Delete(triggerList, vars["trigger"])
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}