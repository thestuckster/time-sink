package internal

//
//import (
//	"database/sql"
//	"sync"
//)
//
//var lock = sync.Mutex{}
//
//type DbSingleton struct {
//	Db *sql.DB
//}
//
//var instance *DbSingleton
//
//func GetDbSingleton() *DbSingleton {
//	if instance == nil {
//		lock.Lock()
//		defer lock.Unlock()
//
//		if instance == nil {
//			db, err := OpenDb()
//			if err != nil {
//				panic(err)
//			}
//
//			instance = &DbSingleton{
//				Db: db,
//			}
//		}
//	}
//
//	return instance
//}
//
//func CloseDbConnection() {
//	db := *GetDbSingleton().Db
//	db.Close()
//}
