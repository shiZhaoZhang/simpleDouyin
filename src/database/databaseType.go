package database

import (
	"gorm.io/gorm"
)

//用户信息
type User struct {
	Id             int64  `gorm:"primaryKey;autoIncrement"`                       //用户唯一标志符号
	Name           string `gorm:"uniqueIndex:, type:varchar(128);not null;index"` //用户名
	Password       string `gorm:"type:varchar(128);not null"`                     //用户密码
	Follow_count   int64  `gorm:"not null;default:0"`                             //关注数
	Follower_count int64  `gorm:"not null;default:0"`                             //粉丝数
	Favorite_count int64  `gorm:"not null;default:0"`                             //喜欢数
	Total_favorite int64  `gorm:"not null;default:0"`                             //被赞数
	Avatar         string //用户头像链接Url
	Signature      string //用户个性签名
	Encryption     string //使用的加密手段
	Iter           int    //加密算法迭代次数
}

//视频信息
type Video struct {
	Id             int64          `gorm:"primaryKey;autoIncrement"` //视频唯一标志符
	UserId         int64          `gorm:"index:, not null"`         //视频发布者ID
	PlayUrl        string         `gorm:"not null"`                 //视频URL
	CoverUrl       string         //视频封面URL
	Title          string         //视频标题
	Favorite_count int64          `gorm:"not null;default:0"`           //视频点赞数
	Comment_count  int64          `gorm:"not null;default:0"`           //视频评论数
	CreatedAt      int64          `gorm:"index:, autoCreateTime:milli"` //视频创建时间。选择milli是因为发现本机的mysql操作是以millisecond的，即使使用nano，后边的值也只会补上0,例如使用nano：1653548764819000000
	UpdatedAt      int64          `gorm:"autoCreateTime:milli"`         //更新时间。
	DeletedAt      gorm.DeletedAt `gorm:"index"`                        //删除时间
}

//点赞信息
type Favorite struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`               //点赞数据的唯一标识符
	UserId   int64 `gorm:"index:idx_member, priority:1, not null"` //发起点赞操作的用户ID
	ToUserId int64 `gorm:"not null"`                               //受到点赞的用户ID
	VideoId  int64 `gorm:"index:idx_member, priority:2, not null"` //受到点赞的视频ID
}

//评论信息
type Comment struct {
	Id        int64  `gorm:"primaryKey;autoIncrement"` //评论唯一标识符
	UserId    int64  `gorm:"not null"`                 //发起评论的用户ID
	VideoId   int64  `gorm:"index:, not null"`         //评论所在的视频ID
	ToUserId  int64  `gorm:"not null"`                 //视频发布者的ID
	Content   string `gorm:"not null"`                 //评论内容
	CreatedAt int64  `gorm:"autoCreateTime:milli"`     //评论创建时间
}

//关注信息
type Relation struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`               //关注唯一标识符
	UserId   int64 `gorm:"index:idx_member, priority:1, not null"` //发起关注者ID
	ToUserId int64 `gorm:"index:idx_member, priority:2, not null"` //被关注者的ID
}
