package database

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/panbhatt/gorm_fiber_sqlserver/config"
	"github.com/panbhatt/gorm_fiber_sqlserver/internals/model"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Err(err)
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(sqlserver.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	log.Info().Msg("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	DB.AutoMigrate(&model.Employee{})

	var employees []model.Employee
	result := DB.Model(model.Employee{}).Find(&employees, []int{1, 2, 3})
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

	employees = nil
	DB.Model(model.Employee{}).Where(" fname like ? ", "%pank%").Not("fname like ?", "%ra%").Find(&employees)

	for _, i := range employees {
		fmt.Println(i.FName)
	}

	employees = nil
	rows, err := DB.Model(model.Employee{}).Rows()
	for rows.Next() {
		var emp model.Employee
		DB.ScanRows(rows, &emp)
		log.Info().Msg(emp.FName + " " + emp.LName)
	}

	var ct int64
	DB.Model(&model.Employee{}).Distinct("FName").Count(&ct)
	log.Info().Msg("Total number of records = " + strconv.FormatInt(ct, 10))
	log.Info().Msg("Database Migrated")
}
