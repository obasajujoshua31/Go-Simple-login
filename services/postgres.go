package services

import (
	"Go-Simple-login/providers"
	"github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/postgres"
)


const dialect = "postgres"

type DB struct {
	*gorm.DB
}


func ConnectToDB(connString string) (DB, error){
   db, err := gorm.Open(dialect, connString)

   if err != nil {
   	return DB{}, err
   }

 return DB{db}, nil

}

type User struct {
	Name string `gorm:"size:255"`
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"size:255"`
}

func (d *DB) FindUserByEmail(email string) User{
	var user User
	d.Where("email = ?", email).First(&user)
    return user
}

func (d *DB) CreateNewUser(user User) error {
	pwd, err := providers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pwd
	d.Create(&user)
	return nil
}

func (u *User) IsMatchPassword(pwd string) bool{
	return providers.CheckPasswordHash(pwd, u.Password)
}