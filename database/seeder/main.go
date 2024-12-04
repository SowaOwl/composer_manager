package seeder

import "gorm.io/gorm"

func Seed(db *gorm.DB) {
	publicTypesSeed(db)
	testSeed(db)
}
