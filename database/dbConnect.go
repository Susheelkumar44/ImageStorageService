/**
 * @author Susheelkumar S
 * @version V1
 * @create date 2020-10-11
 * @desc Interface for DB Conenctions
 */

package database

import (
	config "ImageStorageService/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	//pq is a blank import for database driver
	_ "github.com/lib/pq"
)

//DBRepo satisfies the interface by implementing all the methods
type DBRepo struct {
	GormDB *gorm.DB
}

func (dc *DBRepo) DBConnect(config config.DBConfig) error {
	var err error
	// Format DB configs to required format connect DB
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.PORT)
	dc.GormDB, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		log.Fatalf("Unable to connect DB %v", err)
		return err
	}
	log.Printf("Postgres started at %s PORT", config.PORT)
	return err
}
