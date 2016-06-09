package main

import (
	"log"
	"os"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "errors"
)

func CreateUser(user User) (User,error) {
    endpoint,db := getConnectionInfo()
    
    session, err := mgo.Dial(endpoint)
    if err != nil {
        msg := "failed to connect to mongo:" + endpoint
        log.Fatal(err)
        return User{},errors.New(msg)
    }
    defer session.Close()
    
    //Todo: Need to setup index

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    u := session.DB(db).C("users")
    err = u.Insert(&user)
    if err != nil {
        msg := err.Error()
        log.Fatal(msg)
        return User{}, errors.New(msg)
    }
    
    return user,nil
}

func GetUser(id string) (User,error) {
    endpoint,db := getConnectionInfo()
    
    session, err := mgo.Dial(endpoint)
    if err != nil {
        msg := "failed to connect to mongo:" + endpoint
        log.Fatal(err)
        return User{},errors.New(msg)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB(db).C("users")
    user := User{}
	err = c.Find(bson.M{"id":id}).One(&user)
    if err != nil {
        msg := err.Error()
        log.Fatal(msg)
        return User{}, errors.New(msg)
    }
    
    return user,nil
}

func getConnectionInfo() (string,string) {
    endpoint := os.Getenv("MONGO_ENDPOINT")
    db := os.Getenv("MONGO_DB_NAME")
    return endpoint,db
}

