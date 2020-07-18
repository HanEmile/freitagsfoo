package main

import (
	"net"
	"net/http"
	"text/template"
	"time"

	"git.darknebu.la/chaosdorf/freitagsfoo/src/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	// "git.darknebu.la/chaosdorf/freitagsfoo/src/db"
)

func initHTTPServer() {

	// define the mux router and the routes
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/upcoming", upcomingHandler).Methods("GET")
	r.HandleFunc("/propose", proposeHandler).Methods("GET")
	r.HandleFunc("/talk/{uuid}", talkHandler).Methods("GET")

	// host static files from the ./hosted/static folder
	fs := http.FileServer(http.Dir("./hosted/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// host the uploaded slides
	uploaded := http.FileServer(http.Dir(viper.GetString("uploadpath")))
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", uploaded))

	// api
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/propose", apiProposeHandler).Methods("POST")
	api.HandleFunc("/fetch", apiFetchHandler).Methods("GET")
	api.HandleFunc("/fetch/{key}", apiFetchSpecificHandler).Methods("GET")

	// define the http server
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	logrus.Infof("HTTP server listening on %s:%s", host, port)
	httpServer := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: r,
	}

	// start the http server
	logrus.Fatal(httpServer.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pgdb := db.Connect()
	defer db.Disconnect(pgdb)

	// fetch the next three talks
	firstThreeTalks, err := db.UpcomingTalksLimited(pgdb)
	if err != nil {
		logrus.Warn(err)
		return
	}

	upcomingCount, err := db.CountUpcomingTalks(pgdb)
	if err != nil {
		logrus.Warn(err)
		return
	}

	content := map[string]interface{}{
		"upcomingTalks": firstThreeTalks,
		"upcomingCount": upcomingCount,
		"All":           false,
	}

	// define a template
	t := template.New("")
	t, err = t.ParseGlob("./hosted/tmpl/*.html")
	if err != nil {
		logrus.Warn(err)
		return
	}

	// execute the template
	err = t.ExecuteTemplate(w, "index", content)
	if err != nil {
		logrus.Warn(err)
	}
}

func upcomingHandler(w http.ResponseWriter, r *http.Request) {
	pgdb := db.Connect()
	defer db.Disconnect(pgdb)

	// fetch the next talks
	firstThreeTalks, err := db.UpcomingTalks(pgdb)
	if err != nil {
		logrus.Warn(err)
		return
	}

	upcomingCount, err := db.CountUpcomingTalks(pgdb)
	if err != nil {
		logrus.Warn(err)
		return
	}

	content := map[string]interface{}{
		"upcomingTalks": firstThreeTalks,
		"upcomingCount": upcomingCount,
		"All":           true,
	}

	// define a template
	t := template.New("")
	t, err = t.ParseGlob("./hosted/tmpl/*.html")
	if err != nil {
		logrus.Warn(err)
		return
	}

	// execute the template
	err = t.ExecuteTemplate(w, "upcoming", content)
	if err != nil {
		logrus.Warn(err)
	}
}

func proposeHandler(w http.ResponseWriter, r *http.Request) {

	// find the date of the next friday
	weekday := time.Now().Weekday()
	date := time.Now()
	for weekday != time.Friday {
		date = date.Add(time.Duration(24) * time.Hour)
		weekday = date.Weekday()
	}

	// format it for inserting into the date min paramter of <input>
	layout := "2006-01-02"
	nextFriday := date.Format(layout)

	content := map[string]interface{}{
		"nextFriday": nextFriday,
	}

	t := template.New("")
	t, err := t.ParseGlob("./hosted/tmpl/*.html")
	if err != nil {
		logrus.Warn(err)
		return
	}

	err = t.ExecuteTemplate(w, "propose", content)
	if err != nil {
		logrus.Warn(err)
	}
}

func talkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pgdb := db.Connect()
	defer db.Disconnect(pgdb)

	talk, err := db.TalkByUUID(pgdb, vars["uuid"])
	if err != nil {
		logrus.Warn(err)
		return
	}

	content := map[string]interface{}{
		"talk": talk,
	}

	t := template.New("")
	t, err = t.ParseGlob("./hosted/tmpl/*.html")
	if err != nil {
		logrus.Warn(err)
		return
	}

	err = t.ExecuteTemplate(w, "singleTalk", content)
	if err != nil {
		logrus.Warn(err)
	}
}
