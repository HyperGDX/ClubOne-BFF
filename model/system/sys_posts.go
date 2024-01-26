package system

import (
	"time"
)

// type Post struct {
// PostID int `json:"postId"`  //
// UserID int `json:"userId"` //
// UserName string `json:"userName"` //
// UserAvatar base64.Encoding `json:"userAvatar"` //
// PostTitle string `json:"postTitle"` //
// PostContent string `json:"postContent"` //
// PostPics []base64.Encoding `json:"postPics"` //
// UpdateTime time.Time `json:"updateTime"` //
// LikeCount uint `json:"likeCount"` //
// ViewCount int `json:"viewCount"` //
// }

type Post struct {
PostID int `json:"id"`  // 
UserID int `json:"userId"` // 
PostTitle string `json:"title"` // 
PostContent string `json:"content"` // 
PostPics []string `json:"pics"` // 
UpdateTime time.Time `json:"createTime"` // 
LikeCount uint `json:"likeCount"` // 
ViewCount int `json:"viewCount"` // 
}