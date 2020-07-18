package db

import (
	"git.darknebu.la/chaosdorf/freitagsfoo/src/structs"
	pg "github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/sirupsen/logrus"
)

// CreateSchema creates the defined schemas in the database
func CreateSchema(db *pg.DB) error {
	// define the models
	models := []interface{}{
		(*structs.Talk)(nil),
	}

	// create a table for each model defined above
	for i, model := range models {

		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
		logrus.Tracef("\t Created the %T table", models[i])
	}

	return nil
}

// DeleteAllTables deletes all tables in the given database
func DeleteAllTables(db *pg.DB) error {
	// define the models
	models := []interface{}{
		(*structs.Talk)(nil),
	}

	// delete all models
	for i, model := range models {
		err := db.DropTable(model, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}
		logrus.Tracef("\t Deleted the %T table", models[i])
	}

	return nil
}
