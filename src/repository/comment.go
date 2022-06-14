package repository

import (
	"douyin/src/database"

	"gorm.io/gorm"
)

//添加评论信息
func CommentCreate(com *database.Comment) error {
	return database.MySqlDb.Transaction(func(db *gorm.DB) error {
		if err := database.MySqlDb.Create(com).Error; err != nil {
			return err
		}
		if err := CommentUpdataNumbers(db, com.VideoId, true); err != nil {
			return err
		}
		return nil
	})

}

//使用videoId查询评论列表
func CommentQuerybyVideoID(videoid int64) (commentlist []database.Comment, nums int64) {
	database.MySqlDb.Where("video_id = ? ", videoid).Find(&commentlist)
	nums = int64(len(commentlist))
	return
}

//查询评论是否存在
func CommentQueryByCommentId(comId int64) bool {
	var nums int64
	if database.MySqlDb.Table("comments").Select("count(*)").Where("id = ?", comId).Scan(&nums); nums > 0 {
		return true
	}

	return false

}

//删除评论
func CommentDelete(comment_id, videoId int64) error {
	return database.MySqlDb.Transaction(func(db *gorm.DB) error {
		if err := database.MySqlDb.Where("id = ? ", comment_id).Delete(&database.Comment{}).Error; err != nil {
			return err
		}
		if err := CommentUpdataNumbers(db, videoId, false); err != nil {
			return err
		}
		return nil
	})

}

//添加/取消评论操作
func CommentUpdataNumbers(db *gorm.DB, videoId int64, add bool) error {
	var n int64
	if add {
		n = 1
	} else {
		n = -1
	}
	//更新视频评论数
	if err := db.Model(&database.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", n)).Error; err != nil {
		return err
	}
	return nil
}
