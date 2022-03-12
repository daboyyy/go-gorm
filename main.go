package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

type SqlLoger struct {
	logger.Interface
}

func (l SqlLoger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}

func main() {
	dsn := "root:P@ssw0rd@tcp(13.76.16.73:3306)/bond?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLoger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(Test{})
}

// custom field
type Test struct {
	ID uint
	Code uint   `gorm:"comment:This is Code"`
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

// custom table name
func (t Test) TableName() string {
	return "MyTest"
}
