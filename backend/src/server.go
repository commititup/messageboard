package main

import (
	"log"
	"net/http"
	"./config"
	"./db"
	"./messages"
	"github.com/gorilla/mux"
)


//APP ENTRY POINT
func main() {

	// READ STARTUP CONFIGURATIONS
	ConfigPath := "/etc/app/config.json"
	con, _ := config.Readconfigfile(ConfigPath, false)

	//INITIALIZE  DB CONNECTION
	_, err := db.InitDB("postgres://" + con.Database.Username + ":" + con.Database.Password + "@" + con.Database.Hostname + "/" + con.Database.Dbname +"?sslmode=disable")

	if err != nil {
		log.Panic(err)
	} else {
		log.Print("Inititated DB connection To use")
	}

	//DEFINING ROUTES
	router := mux.NewRouter()
	router.HandleFunc("/messages", messages.AllMessages).Methods("GET")
	router.HandleFunc("/message", messages.AddMessage).Methods("POST")
	router.HandleFunc("/message", messages.DeleteMessage).Methods("DELETE")
	router.HandleFunc("/message/{id}", messages.GetDetailMessage).Methods("GET")
	http.Handle("/", router)

	//SERVE STATIC FILES/FRONTEND
	fs := http.FileServer(http.Dir(con.AppRoot))
	router.PathPrefix("/").Handler(http.StripPrefix("/",fs))

	//START THE SERVER 
	log.Printf("server started listing on %s", con.Port)
	err = http.ListenAndServe(":"+con.Port, router)
	if err != nil {
		log.Fatal(err)
	}
}
