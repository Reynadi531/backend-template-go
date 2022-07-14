package user

import (
	"backend-template-go/internal/entities/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) Create(user *model.User) (*model.User, error) {
	return user, u.db.Create(&user).Error
}

func (u userRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u userRepository) Update(user *model.User) (*model.User, error) {
	return user, u.db.Model(&user).Updates(user).Error
}

func (u userRepository) Delete(id string) error {
	return u.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u userRepository) UserExists(email string) (bool, error) {
	var count int64
	err := u.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
