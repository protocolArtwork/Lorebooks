package newdb

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"lbcli"
)

func NewDB(dbpath string) (*gorm.DB, error) {
	// Create a new database connection

	abspath, patherr := filepath.Abs(dbpath)

	if patherr == nil {

		_, staterr := os.Stat(abspath)

		if os.IsNotExist(staterr) {
			db, err := gorm.Open(sqlite.Open(abspath), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}

			lbcli.Action(fmt.Sprintf("Created new lorebook at %s", abspath))

			// Migrate the schema
			err = db.AutoMigrate()
			if err != nil {
				panic("failed to migrate schema")
			}

			lbcli.Log("Migrated schema to lorebook")

			return db, nil
		} else if staterr == nil {
			lbcli.Fatal(fmt.Sprintf("Lorebook already exists at %s", abspath))

			return nil, os.ErrExist
		} else {
			return nil, staterr
		}
	} else {
		return nil, patherr
	}
}
