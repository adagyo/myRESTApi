package utils

import (
	"gopkg.in/mgo.v2"
)

func Connect(conf *Config) (*mgo.Session, *mgo.Database, int) {
	session, err := mgo.Dial(conf.MgoURL); if err != nil {
		return nil, nil, 1
	}
	// TODO: Check if database exists and return err=2 if not
	return session, session.DB(conf.MgoDB), 0
}
