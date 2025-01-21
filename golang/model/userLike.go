package model

import "time"

type UserLike struct {
	userId     uint      `json:"user_id" gorm:"not null"`
	likeUserId uint      `json:"like_user_id"`
	isMatch    bool      `json:"isMatch"`
	isNotLike  bool      `json:"age"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserLikeResponse struct {
	userId      uint   `json:"id" gorm:"primaryKey"`
	likeUserId  uint `json:"email" gorm:"unique"`
	isMatch     bool `json:"name"`
	isNotLike   bool  `json:"age"`
}
