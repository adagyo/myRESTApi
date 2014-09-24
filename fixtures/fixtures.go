package fixtures

import (
	"fmt"
	"github.com/adagyo/myRESTApi/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func LoadUsers(db *mgo.Database) {
	usersCollection := db.C("users")

	nbUsers, _ := usersCollection.Count()
	if nbUsers > 0 {
		fmt.Println("Dropping all data from collection 'users'")
		usersCollection.RemoveAll(bson.M{})
	}

	fmt.Println("Populating collection 'users' with 100 documents")
	for i := 0; i < 100; i++ {
		usersCollection.Insert(&entities.User{
			UserId:   i,
			Login:    "user_" + strconv.Itoa(i),
			Password: "secret",
			Name:     "User #" + strconv.Itoa(i),
		})
	}

}
