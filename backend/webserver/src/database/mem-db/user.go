package memdb

import (
	"go-sandbox/src/models"
	"log"
)

type MemoryDatabaseUserDaoImpl struct {
	users []models.Login
}

var userDaoInstance *MemoryDatabaseUserDaoImpl

func (m MemoryDatabaseUserDaoImpl) GetInstance() *MemoryDatabaseUserDaoImpl {
	log.Println("Getting memdb Instance")
	if userDaoInstance != nil {
		return userDaoInstance
	}

	userDaoInstance = &MemoryDatabaseUserDaoImpl{users: make([]models.Login, 0)}
	return userDaoInstance
}

func (m MemoryDatabaseUserDaoImpl) AddUser(user models.Login) {
	m.users = append(m.users, user)
}
