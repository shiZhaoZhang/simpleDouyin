package main

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//评论信息
type Student struct {
	Id        int64  `gorm:"primaryKey;autoIncrement"` //评论唯一标识符
	UserId    int64  `gorm:"not null"`                 //发起评论的用户ID
	VideoId   int64  `gorm:"index:, not null"`         //评论所在的视频ID
	ToUserId  int64  `gorm:"not null"`                 //视频发布者的ID
	Content   string `gorm:"not null"`                 //评论内容
	CreatedAt int64  `gorm:"autoCreateTime:milli"`     //评论创建时间
}

var DSN string = "douyin:123456@tcp(:3306)/douyindata?charset=utf8&parseTime=True&loc=Local"

func BenchmarkVideoSearch(b *testing.B) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Panic("Database connect error : ", err)
	}

	for i := 0; i < b.N; i++ {

		var commentlist []Student
		//db.Where("to_user_id = ? ", 1).Find(&commentlist)
		db.Where("video_id = ? ", 1).Find(&commentlist)
	}
}
func BenchmarkTouseridSearch(b *testing.B) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Panic("Database connect error : ", err)
	}

	for i := 0; i < b.N; i++ {

		var commentlist []Student
		db.Where("to_user_id = ? ", 1).Find(&commentlist)
	}
}

/*
func BenchmarkInsert(b *testing.B) {
	//db, err := gorm.Open(mysql.Open("douyin:123456@tcp(:3306)/douyindata?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Panic("Database connect error : ", err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(Student{})
	rand.Seed(time.Now().UnixNano())

	for i := 0; i != b.N; i++ {
		randomNum := rand.Int63n(100) + 1
		com := Student{
			UserId:   1,
			VideoId:  randomNum,
			ToUserId: randomNum,
			Content:  "test",
		}
		if err := db.Create(&com).Error; err != nil {
			b.FailNow()
		}

	}
}
*/
