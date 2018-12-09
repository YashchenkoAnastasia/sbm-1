package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"gopkg.in/mgo.v2/bson"
	"reflect"
)

const (
	CollectionPost = "posts"
)

// Post model
type Post struct {
	Id    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
	Body  string        `json:"body" bson:"body"`
}

func (p Post) ValidateNewPost() error {
	return validation.ValidateStruct(&p,
		// Title cannot be empty
		validation.Field(&p.Title, validation.Required),
		// Body cannot be empty
		validation.Field(&p.Title, validation.Required),
	)
}

func IsZeroOrNil(x interface{}) bool {
	if x.(reflect.Value).Kind() == reflect.Ptr {
		return x.(reflect.Value).IsNil()
	}
	return x == reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func (old *Post) MergeInPlace(new Post) {
	for ii := 0; ii < reflect.TypeOf(old).Elem().NumField(); ii++ {

		if x := reflect.ValueOf(&new).Elem().Field(ii); !IsZeroOrNil(x) && x.String() != "" {
			reflect.ValueOf(old).Elem().Field(ii).Set(x)
		}
	}
}
