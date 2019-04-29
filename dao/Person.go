package dao

import (
	"gramya/db"
	"gramya/utils"
	"time"
)

type Person struct {
	Id      int
	FName   string
	MName   string
	LName   string
	Dob     string
	Cd      time.Time
	Ud      time.Time
	Authors string
	Tags    string
}

func InsertPerson(p *Person) bool {
	st, err := db.DB.Prepare("INSERT INTO person (id, fName, mName, lName, dob, ud, authors, tags) values (?,?,?,?,?,?,?,?)")
	utils.Panic(err)
	_, err = st.Exec(p.Id, p.FName, p.MName, p.LName, p.Dob, time.Now(), p.Authors, p.Tags)
	utils.Panic(err)
	return true
}
