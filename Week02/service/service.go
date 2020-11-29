package service

import "Go-000/Week02/dao"

func DoDao() error {
	err := dao.UpdateModel()
	return err
}
