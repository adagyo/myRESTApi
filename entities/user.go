package entities

import (
	"github.com/adagyo/myRESApi/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	_ "log"
	_ "strconv"
)

type User struct {
	UserId   int
	Login    string
	Password string
	Name     string
	ErrMsg   string
}

func GetUserByUid(db *mgo.Database, uid int) User {
	usersCollection := db.C("users")

	u2find := User{}
	errmsg := usersCollection.Find(bson.M{"userid": uid}).One(&u2find)
	if errmsg != nil {
		u2find.ErrMsg = errmsg.Error()
	}

	return u2find
}

func GetUsers(db *mgo.Database, params utils.QueryRequestParameters) []User {
	usersCollection := db.C("users")
	userList := []User{}
	_ = usersCollection.Find(bson.M{}).Skip(int(params.Offset)).Limit(int(params.Limit)).Sort(params.Sort).All(&userList)

	return userList
}
