package database

import (
	"fmt"
	"strconv"

	"github.com/TeaMeow/KitSvc/config"
	"github.com/jinzhu/gorm"
)

func Connect(c config.Database) *gorm.DB {

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s?charset=%s&parseTime=%s&loc=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.Charset,
		strconv.FormatBool(c.ParseTime),
		c.Loc,
	))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}