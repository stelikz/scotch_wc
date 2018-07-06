package database

import(
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "messages"
    "log"
    )

func Show(s string, a messages.Messages, db string, co string) {
	session, _ := mgo.Dial(s)

	anotherSession := session.Copy()
	defer anotherSession.Close()

	c := session.DB(db).C(co)
	err2 := c.Find(bson.M{}).All(&a)
	if err2 != nil {
		log.Println(err2)
	}
	log.Println(a)

}

func Store(s string, msg messages.Message, db string, co string) {
	session, _ := mgo.Dial(s)

	anotherSession := session.Copy()
	defer anotherSession.Close()

	c := session.DB(db).C(co)
	c.Insert(msg)

}