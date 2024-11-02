package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// VARIABLES

var (
	host     string // Your host (example: localhost)
	user     string // Your user (example: postgres)
	password string // Your password (optional)
	dbname   string // Your DB name (example: postgres)
	sslmode  string // SSL (enable/disable)
)

type Purchase struct {
	ID          uint   `gorm:"type:BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
	Name        string `gorm:"type:VARCHAR(100) NOT NULL"`
	Description string `gorm:"type:VARCHAR(150)"`
	Amount      int    `gorm:"type:INT NOT NULL"`
}

var purchases = []Purchase{
	{Name: "Grocery", Amount: 1500},
	{Name: "PS5", Description: "Gaming Console", Amount: 40000},
	{Name: "Sneakers", Description: "Jordans", Amount: 20000},
}

var db *gorm.DB

// GORM

func ConfigGorm() error {
	var all []Purchase

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, sslmode)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = conn

	err = db.Transaction(func(tx *gorm.DB) error {
		ok := tx.Migrator().HasTable(&Purchase{})
		if !ok {
			err := tx.Migrator().CreateTable(&Purchase{})
			if err != nil {
				return err
			}

			tx.Table("purchases").Create(purchases)
			if tx.Error != nil {
				return err
			}
		}

		tx = tx.Table("purchases").Find(&all)
		if tx.RowsAffected == 0 {
			tx := tx.Table("purchases").Create(purchases)
			if tx.Error != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil
	}
	return nil
}

func SelectAll() []Purchase {
	var all []Purchase

	db.Table("purchases").Find(&all)

	return all
}

func SelectWhere() Purchase {
	var where Purchase

	db.Table("purchases").Where("name=?", "PS5").Find(&where)

	return where
}

func SelectSpecific() []Purchase {
	var specific []Purchase

	db.Table("purchases").Select("name", "amount").Find(&specific)

	return specific
}

func UpdateAll() Purchase {
	var updated Purchase

	db.Table("purchases").Where("name=?", "PS5").Updates(&Purchase{Name: "Flowers", Description: "Present", Amount: 1000})

	db.Table("purchases").Where("name=?", "Flowers").Find(&updated)

	return updated
}

func UpdateName() Purchase {
	var updated Purchase

	db.Table("purchases").Where("name=?", "Sneakers").Update("name", "My Purchase")

	db.Table("purchases").Where("name=?", "My Purchase").Find(&updated)

	return updated
}

func DeleteRow() int64 {
	tx := db.Table("purchases").Where("name=?", "Grocery").Delete(&Purchase{})

	return tx.RowsAffected
}

func InsertRow() Purchase {
	var inserted Purchase

	db.Table("purchases").Create(&Purchase{Name: "Book", Description: "Book to read", Amount: 400})

	db.Table("purchases").Where("name=?", "Book").Find(&inserted)

	return inserted
}
