package main

import (
	"fmt"
	"math/rand"
	"time"

	"git.darknebu.la/chaosdorf/freitagsfoo/src/db"
	"git.darknebu.la/chaosdorf/freitagsfoo/src/structs"
	pg "github.com/go-pg/pg/v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// initialize the DB by first deleting all the tables and then creating the
// tables again
func initDB(pgdb *pg.DB) {
	deleteAllTables(pgdb)
	createAllTables(pgdb)

	createSomeTalks(pgdb)
}

// delete all tables in the database
func deleteAllTables(pgdb *pg.DB) {
	err := db.DeleteAllTables(pgdb)
	if err != nil {
		logrus.Fatalf("Could not delete the database schema: %s", err)

	}
}

// create the schema
func createAllTables(pgdb *pg.DB) {
	err := db.CreateSchema(pgdb)
	if err != nil {
		logrus.Fatalf("Could not create the database schema: %s", err)
	}
}

////////////////////////////////////////////////////////////////////////////////
// This section is just there to insert stuff into the database for testing
////////////////////////////////////////////////////////////////////////////////

func createSomeTalks(pgdb *pg.DB) {

	rand.Seed(time.Now().UnixNano())
	nicknames := []string{
		"Thorin",
		"Fili",
		"Kili",
		"Balin",
		"Dwalin",
		"Oin",
		"Gloin",
		"Dori",
		"Nori",
		"Ori",
		"Bifur",
		"Bofur",
		"Bombur",
	}
	for i := 0; i < 10; i++ {

		date := time.Now().Add(-3 * 7 * 24 * time.Hour).Add(time.Duration(i) * 7 * 24 * time.Hour)

		layout := "2006-01-02"
		formattedDate := date.Format(layout)

		talk := &structs.Talk{
			UUID:          uuid.New(),
			Title:         fmt.Sprintf("Wie baut man eigentlich Raketen ohne Brennstoff? nr. %d", i),
			Description:   fmt.Sprintf("Sunt rerum illo corrupti. Similique qui rem debitis. Accusamus et rerum sint et amet eos nemo. Et enim omnis et. Tempora et corrupti aut ea et vel. \n Dolor est quae sed molestiae nisi esse aliquid atque. Voluptas vero et ducimus voluptatem in eaque. Quo illum et delectus vel sed molestias quidem. Consequuntur unde dolores quis sunt exercitationem eos et provident. Animi eaque temporibus alias. %d", i),
			Slides:        "./uploads/black.png",
			Nickname:      nicknames[rand.Intn(len(nicknames))],
			Date:          date,
			FormattedDate: formattedDate,
			Upcoming:      true,
		}
		err := db.InsertTalk(pgdb, talk)
		if err != nil {
			logrus.Fatalf("Could not insert talk into the database: %s", err)
		}
	}
}
