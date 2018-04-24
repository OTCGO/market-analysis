package mongo

import (
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

func InitMongo() {
	configration, _ := config.GetConfig()
	session, err := mgo.Dial(configration.MongoConfiguration.Url)
	//SetPoolLimit
	session.SetPoolLimit(configration.MongoConfiguration.PoolLimit)
	DataBase = configration.MongoConfiguration.Db
	if err != nil {
		errors.Wrap(err, "mongo connet error")
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
