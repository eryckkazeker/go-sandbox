package factory

import (
	"go-sandbox/src/database"
	memdb "go-sandbox/src/database/mem-db"
	"go-sandbox/src/database/mysql"
	"log"
)

func UserDao(engine string) database.UserDao {
	var dao database.UserDao
	switch engine {
	case "memdb":
		dao = &memdb.MemoryDatabaseUserDaoImpl{}
	default:
		log.Fatalf("%s engine not implemented yet", engine)
		return nil
	}

	return dao
}

func MovieDao(engine string) database.MovieDao {
	var dao database.MovieDao
	switch engine {
	case "memdb":
		var mem = &memdb.MemoryDatabaseMovieDaoImpl{}
		dao = mem.GetInstance()
		break
	case "mysql":
		dao = mysql.MySqlMovieDaoImpl{}
	default:
		log.Fatalf("%s engine not implemented yet", engine)
		return nil
	}

	return dao
}
