package repository

import "gorm.io/gorm"

func InitializeRepository(db *gorm.DB) {
	InitializeUserRepository(db)
	InitializePostRepository(db)
	InitializeCommentRepository(db)
	InitializeReactionRepository(db)
	InitializeImageRepository(db)
}
