package model

import "time"

type Person struct {
	id      int32
	fName   string
	mName   string
	lName   string
	dob     string
	gender  string
	cd      time.Time
	ud      time.Time
	authors string
	tags    string
}

type Relation struct {
	id      int32
	fName   string
	father  int32
	wife    int32
	cd      string
	ud      string
	authors string
}
