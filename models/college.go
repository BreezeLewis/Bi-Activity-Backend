package models

type College struct {
	ID                  uint64 `gorm:"primaryKey;autoIncrement" json:"id"`                           // 主键，自动递增
	CollegeAccount      string `gorm:"type:varchar(64);unique;not null" json:"college_account"`      // 学院账号，唯一
	CollegeName         string `gorm:"type:varchar(64);not null" json:"college_name"`                // 学院名称
	Password            string `gorm:"type:varchar(255);not null" json:"password"`                  // 密码
	AdminName           string `gorm:"type:varchar(255);" json:"admin_name"`                        // 管理员名称
	AdminIDNumber       string `gorm:"type:varchar(20);" json:"admin_id_number"`                    // 管理员身份证号
	AdminImageID        uint64 `gorm:"type:bigint" json:"admin_image_id"`                           // 管理员头像 ID
	AdminPhone          string `gorm:"type:varchar(64);unique;" json:"admin_phone"`                 // 管理员电话，唯一
	AdminEmail          string `gorm:"type:varchar(255);unique;" json:"admin_email"`                // 管理员邮箱，唯一
	Campus              int    `gorm:"type:int;not null" json:"campus"`                             // 校园 ID
	CollegeAddress      string `gorm:"type:varchar(255);not null" json:"college_address"`           // 学院地址
	CollegeIntroduction string `gorm:"type:text" json:"college_introduction"`                       // 学院简介
	CollegeAvatarID     uint64 `gorm:"type:bigint" json:"college_avatar_id"`                        // 学院头像 ID
}