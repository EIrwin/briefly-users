package data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func CreateUser(user *User) error {
	endpoint, db := getConnectionInfo()

	session, err := mgo.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	//Collection Users
	u := session.DB(db).C("users")

	// Index
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := u.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}

	user.Id = bson.NewObjectId()
	err = u.Insert(user)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetUser(id string) (User, error) {
	endpoint, db := getConnectionInfo()

	session, err := mgo.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB(db).C("users")
	var user User
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user, nil
}

func getConnectionInfo() (string, string) {
	endpoint := "mongodb://dev:briefly123!@ds015584.mlab.com:15584/briefly"
	db := "briefly"
	return endpoint, db
}

type User struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	HourlyRate  float32       `json:"hourlyRate"`
	Salary      float32       `json:"salary"`
	IntervalAmt float32       `json:"intervalAmt"`
}
