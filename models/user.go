package models

import (
	"github.com/rn-consider/grpcservice/dao"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string // 根据需要修改为非指针类型，如果数据库字段允许为空
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateUser 创建User字段
func CreateUser(user *User) error {
	if err := dao.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 根据主键删除User字段
func DeleteUser(ID int) error {
	if err := dao.DB.Where("ID = ?", ID).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdatedAtUser(user *User) error {
	if err := dao.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUser() ([]*User, error) {
	var userList []*User
	if err := dao.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func GetAUser(id int) (*User, error) {
	user := &User{}
	if err := dao.DB.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
