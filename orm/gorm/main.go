package main

import (
	"fmt"
	"log"
	"os"
	"time"

	v1 "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	v2 "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	StudentID   uint64
	StudentName string
	Age         int
	Score       int
	Grade       Grade `gorm:"foreignKey:StudentID"`
}

func (User) Table() string {
	return "t_user"
}

type Grade struct {
	ID        uint64
	StudentID uint64
	SubjectID int
	Score     int
}

type UserGrade struct {
	User
	Grade
}

func (Grade) TableName() string {
	return "t_grade"
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=UTC"

	dbv1, err := v1.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbv1.LogMode(true)

	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        //
		},
	)

	dbv2, err := v2.Open(mysql.Open(dsn), &v2.Config{
		Logger: log,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	query := "SELECT * FROM t_grade tg INNER JOIN t_user tu ON tg.student_id = tu.student_id WHERE subject_id=?;"

	var u1 []UserGrade
	var u2 User
	if err := dbv1.Raw(query, 1).Scan(&u1).Error; err != nil {
		fmt.Println(err)
	}

	if err := dbv2.Where("student_id=?", 1).Find(&u2).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", u1)
	fmt.Printf("%+v\n", u2)
}
