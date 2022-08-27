package profiledto

import "waysbuck/models"

type ProfileResponse struct {
	ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string                      `json:"phone" gorm:"type: varchar(255)"`
	Address string                      `json:"address" gorm:"type: text"`
	UserID  int                         `json:"user_id"`
	User    models.UsersProfileResponse `json:"user"`
}
