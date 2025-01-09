package db

import (
	"kafka-to-db/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	GormDB *gorm.DB
}

func Connect(dsn string) (*DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations
	err = gormDB.AutoMigrate(&models.Message{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database and ran migrations.")
	return &DB{GormDB: gormDB}, nil
}

func (db *DB) SaveMessage(message models.Message) error {
	var existingMessage models.Message

	// First, try to find the record based on the unique key (message.Key)
	if err := db.GormDB.Where("key = ?", message.Key).First(&existingMessage).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// If not found, create a new message
			if err := db.GormDB.Create(&message).Error; err != nil {
				log.Println("Error creating message:", err)
				return err
			}
		} else {
			// Other errors (like database errors)
			log.Println("Error finding message:", err)
			return err
		}
	} else {
		// If found, update the existing record
		if err := db.GormDB.Model(&existingMessage).Updates(message).Error; err != nil {
			log.Println("Error updating message:", err)
			return err
		}
	}

	return nil
}
