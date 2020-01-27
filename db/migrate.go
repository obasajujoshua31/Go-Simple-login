package db

import "Go-Simple-login/services"

func Run(conString string) error{

	db, err := services.ConnectToDB(conString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&services.User{})
	return nil
}
