package entity

type Video struct {
	VideoId     int    `gorm:"column:video_id;primaryKey;autoIncrement;type:int" json:"video_id"`
	Teacher     string `gorm:"teacher;type:varchar(100)" json:"teacher"`
	Title       string `gorm:"column:title;type:varchar(200)" json:"title"`
	URLVideo    string `gorm:"column:url_video;type:varchar(200)" json:"url_video"`
	Description string `gorm:"column:description;type:longtext" json:"description"`
}

func (Video) TableName() string {
	return "videos"
}
