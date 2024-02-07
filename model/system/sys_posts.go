package system

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
UserID int `json:"userId"` // 
PostTitle string `json:"title"` // 
ChannelId []int `json:"channelId"` //
PostContent string `json:"content"` // 
PostPics []string `json:"pics"` // 
LikeCount int `json:"likeCount"` // 
ViewCount int `json:"viewCount"` // 
}

type OssPolicy struct {
	Accessid string `json:"accessid"` //,
    Policy string `json:"policy"` //,
    Signature string `json:"signature"` //,
    Dir string `json:"dir"` //,
    Host string `json:"host"` //,
    Expire string `json:"expire"` //,
}