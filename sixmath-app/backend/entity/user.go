package entity

import "time"

type User struct {
	UserId          int       `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id,omitempty"`
	Username        string    `gorm:"column:username;unique;type:varchar(50)" json:"username,omitempty"`
	Email           string    `gorm:"column:email;type:varchar(100)" json:"email,omitempty"`
	Handphone       string    `gorm:"column:handphone;type:varchar(20)" json:"handphone,omitempty"`
	Password        string    `gorm:"column:password;type:varchar(255)" json:"password,omitempty"`
	Role            string    `gorm:"column:role;type:varchar(10)" json:"role,omitempty"`
	Gender          string    `gorm:"column:gender;type:varchar(10)" json:"gender,omitempty"`
	Certificate     string    `gorm:"column:certificate;type:mediumtext" json:"certificate,omitempty"`
	Avatar          string    `gorm:"column:avatar;type:mediumtext" json:"avatar,omitempty"`
	StudentUsername string    `gorm:"column:student_username;type:varchar(50)" json:"student_username,omitempty"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}