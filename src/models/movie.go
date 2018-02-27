package models

import (
	"github.com/eddiefisher/mgotoster/src/dao"
	"github.com/globalsign/mgo/bson"
)

const (
	COLLECTION = "movies"
)

//Movie Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

type Movies []Movie

func (m *Movies) FindAll() error {
	return dao.DB.C(COLLECTION).Find(nil).All(m)
}

func (m *Movie) FindById(id string) error {
	return dao.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&m)
}

func (m *Movie) Insert(movie Movie) error {
	return dao.DB.C(COLLECTION).Insert(&movie)
}

func (m *Movie) Delete(movie Movie) error {
	return dao.DB.C(COLLECTION).Remove(&movie)
}

func (m *Movie) Update(movie Movie) error {
	return dao.DB.C(COLLECTION).UpdateId(movie.ID, &movie)
}
