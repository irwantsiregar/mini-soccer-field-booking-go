package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	// Create default users
	user := models.User{
		UUID:    uuid.New(),
		Name:	 "Administrator",
		Username: "admin",
		Password: string(password),
		Email:    "admin@example.com",
		PhoneNumber: "085762168572",
		RoleID:   constants.Admin,
	}

	err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error

	if err != nil {
		logrus.Errorf("Failed to seed user %s: %v", user.Username, err)
		panic(err)
	}

	logrus.Infof("Seeded user %s successfully", user.Username)
}
