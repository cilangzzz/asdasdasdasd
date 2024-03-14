package dao

import (
	"goDnParse/model"
	"goDnParse/util"
)

// Insert 插入用户
func Insert(user *model.User) error {
	_, err := util.Orm.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUid 获取用户
func GetUserByUid(user *model.User) error {
	orm := util.Orm
	_, err := orm.Get(user)
	return err
}
