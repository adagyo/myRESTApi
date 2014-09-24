package main

import (
	"encoding/json"
	"github.com/adagyo/myRESTApi/entities"
	"github.com/adagyo/myRESTApi/fixtures"
	"github.com/adagyo/myRESTApi/utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"strconv"
)

var (
	conf    utils.Config
	db      *mgo.Database
	session *mgo.Session
)

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	requestParams := utils.ParseRequestParameters(request)
	users := entities.GetUsers(db, requestParams)
	j, _ := json.Marshal(users)
	writer.Write([]byte(j))
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	requestParams := mux.Vars(request)

	uid, err := strconv.Atoi(requestParams["id"])
	if err != nil {
		writer.Write([]byte("id is not an integer"))
	} else {
		u2find := entities.GetUserByUid(db, uid)
		if u2find.ErrMsg != "" {
			writer.WriteHeader(404)
			writer.Write([]byte(u2find.ErrMsg))
		} else {
			j, _ := json.Marshal(u2find)
			writer.Write([]byte(j))
		}
	}
}

func main() {
	// Load configuration
	utils.LoadConfig(&conf)

	// Connect to Mongo and select the database
	var ErrNo int
	session, db, ErrNo = utils.Connect(&conf)
	switch ErrNo {
	case 1:
		log.Fatal("[FATAL] Could not connect to mongo URL '" + conf.MgoURL + "'.")
	case 2:
		log.Fatal("[FATAL] Database '" + conf.MgoDB + "' does not exist.")
	}
	defer session.Close()

	if conf.LoadFixtures == true {
		fixtures.LoadUsers(db)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users", ListUsers).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]*}", GetUser).Methods("GET")

	log.Println("Server is listening on port 8000")
	http.ListenAndServe(":8000", router)
}
