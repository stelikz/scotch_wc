package database

import(
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/stelikz/scotch_wc/messages"
    "log"
    )

func Show(s string, a messages.Messages, db string, co string) messages.Messages{
	session, _ := mgo.Dial(s)


	c := session.DB(db).C(co)
	err2 := c.Find(bson.M{}).All(&a)
	if err2 != nil {
		log.Println(err2)
	}
	return a
}

func Store(s string, msg messages.Message, db string, co string) {
	session, _ := mgo.Dial(s)

	c := session.DB(db).C(co)
	c.Insert(msg)

}