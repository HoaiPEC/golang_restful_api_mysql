package Models

import (
	"gin-restful-api-mysql/Config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     string
	Address   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *User) TableName() string {
	return "users"
}

func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}

func Create(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)

	return nil
}

func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)

	return nil
}
