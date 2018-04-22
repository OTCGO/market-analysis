package mongo

import (
	"fmt"
	"market-analysis/config"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

func New(configration *config.Config) (db *mgo.Database, err error) {
	session, err := mgo.Dial(configration.MongoConfiguration.Url)
	//SetPoolLimit
	session.SetPoolLimit(configration.MongoConfiguration.PoolLimit)
	if err != nil {
		errors.Wrap(err, "mongo connent error")
	}

	db = session.Clone().DB(configration.MongoConfiguration.Db)

	// session.Close()
	defer session.Close()

	return db, err
}

var sessionMongo *mgo.Session
var DataBase string

func InitMongo(configration *config.Config) {
	session, err := mgo.Dial(configration.MongoConfiguration.Url)
	//SetPoolLimit
	session.SetPoolLimit(configration.MongoConfiguration.PoolLimit)
	DataBase = configration.MongoConfiguration.Db
	if err != nil {
		fmt.Printf("连接mongodb失败:%s\n", err)
	}
	sessionMongo = session
}

/**
 * GetSession
 */
func GetSession() *mgo.Session {
	if sessionMongo == nil {
		return nil
	}
	return sessionMongo.Clone()
}

/**
 * GetSesExecMongosion
 */
func ExecMongo(collection string, s func(*mgo.Collection) error) error {
	session := GetSession()
	//defer close  session
	defer session.Close()
	c := session.DB(DataBase).C(collection)
	return s(c)
}
