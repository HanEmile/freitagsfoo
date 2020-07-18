package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"git.darknebu.la/chaosdorf/freitagsfoo/src/db"
	"git.darknebu.la/chaosdorf/freitagsfoo/src/structs"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func apiProposeHandler(w http.ResponseWriter, r *http.Request) {

	// save the uploaded slides from the form to disk
	slidesPath, err := saveUploadedSlides(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	// parse the form
	r.ParseForm()
	title := r.Form["title"][0]
	description := r.Form["description"][0]
	nickname := r.Form["nickname"][0]
	date := r.Form["date"][0]

	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
		return
	}

	// yes, we've parsed the date and formatted it again, but this makes sure
	// that the user input is really valid and not some bad XSS attempt
	formattedDate := parsedDate.Format(layout)

	// fill the talk struct with the information regarding the talk
	talk := &structs.Talk{
		UUID:          uuid.New(),
		Title:         title,
		Description:   description,
		Slides:        slidesPath,
		Nickname:      nickname,
		Date:          parsedDate,
		FormattedDate: formattedDate,
		Upcoming:      true,
	}

	// insert the talk into the database
	pgdb := db.Connect()
	defer db.Disconnect(pgdb)
	err = db.InsertTalk(pgdb, talk)

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

// saveUploadedSlides saves the uploaded slides to disk returning the path
func saveUploadedSlides(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)

	// get the file
	file, handler, err := r.FormFile("slides")
	if err != nil {
		return "", err
	}
	defer file.Close()

	strings.ReplaceAll("oink oink oink", handler.Filename, "moo")

	filename := strings.ReplaceAll(handler.Filename, "..", "")
	uploadPath := viper.GetString("uploadpath")
	filePath := fmt.Sprintf("%s%s", uploadPath, filename)

	// open a file on disk for storing the uploaded file
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
		return "", err
	}
	defer f.Close()

	// copy the uploaded file to the file created
	io.Copy(f, file)

	return filePath, nil
}

func apiFetchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "fetch")
	return
}

func apiFetchSpecificHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "fetch specific")
	return
}
