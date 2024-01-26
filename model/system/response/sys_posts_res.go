package response

import (
	"time"
)

type Post struct {
PostID int `json:"id"`  // 
UserName string `json:"userName"` //
UserAvatar string `json:"userAvatar"` // 
PostTitle string `json:"title"` // 
PostContent string `json:"content"` // 
PostPics []string `json:"pics"` // 
UpdateTime time.Time `json:"createTime"` // 
LikeCount uint `json:"likeCount"` // 
ViewCount int `json:"viewCount"` // 
}