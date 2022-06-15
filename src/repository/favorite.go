package repository

import (
	"douyin/src/database"

	"gorm.io/gorm"
)

//添加赞信息
func FavoriteCreate(fav database.Favorite) error {
	return database.MySqlDb.Transaction(func(db *gorm.DB) error {
		if err := db.Create(&fav).Error; err != nil {
			return err
		}
		if err := FavoriteUpdataNumbers(db, fav.VideoId, fav.ToUserId, fav.UserId, true); err != nil {
			return err
		}
		return nil
	})
}

//删除点赞
func FavoriteDelete(userid, videoid, videoUserId int64) error {
	return database.MySqlDb.Transaction(func(db *gorm.DB) error {
		if err := database.MySqlDb.Where("user_id = ? AND video_id = ?", userid, videoid).Delete(&database.Favorite{}).Error; err != nil {
			return err
		}
		if err := FavoriteUpdataNumbers(db, videoid, videoUserId, userid, false); err != nil {
			return err
		}
		return nil
	})
}

//查询点赞是否存在
func FavoriteQueryByUserAndVideo(userid int64, videoid int64) (exist bool) {
	var nums int64
	//select count(*) from favorites where user_id = ? AND video_id = ?
	if database.MySqlDb.Table("favorites").Select("count(*)").Where("user_id = ? AND video_id = ?", userid, videoid).Scan(&nums); nums > 0 {
		exist = true
	} else {
		exist = false
	}
	return
}

//使用userId查看点赞列表视频
func FavoriteQuerybyUserID(userid int64) (videolist []database.Video, nums int64) {
	database.MySqlDb.Model(&database.Video{}).Select("videos.*").Joins("inner join favorites on favorites.video_id = videos.id").Where("favorites.user_id = ? ", userid).Scan(&videolist)
	nums = int64(len(videolist))
	return
}

//添加/取消赞操作，统一更新视频点赞数、用户获赞数、用户喜欢数
func FavoriteUpdataNumbers(db *gorm.DB, videoId, videoUserId, userId int64, add bool) error {
	var n int64
	if add {
		n = 1
	} else {
		n = -1
	}

	//更新视频点赞数
	if err := db.Model(&database.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", n)).Error; err != nil {
		return err
	}
	//更新video用户获赞数
	if err := db.Model(&database.User{}).Where("id = ?", videoUserId).Update("total_favorite", gorm.Expr("total_favorite + ?", n)).Error; err != nil {
		return err
	}
	//更新token用户喜欢数
	if err := db.Model(&database.User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", n)).Error; err != nil {
		return err
	}

	return nil

}
