package dao

import (
	"gramya/db"
	"gramya/utils"
	"time"
)

type Relation struct {
	Id      int
	FName   string
	Father  int
	Wife    int `default:"-1"`
	Cd      time.Time
	Ud      time.Time
	Authors string
}

func InsertRelation(r *Relation) bool {
	st, err := db.DB.Prepare("INSERT INTO relations (id, fName, father, wife, ud, authors) VALUES (?, ?, ?, ?, ?, ?)")
	utils.Panic(err)
	_, err = st.Exec(r.Id, r.FName, r.Father, r.Wife, time.Now(), r.Authors)
	utils.Panic(err)
	return true
}

func UpdateRelation(r *Relation) bool {
	st, err := db.DB.Prepare("UPDATE relations set wife=? where id = ?")
	utils.Panic(err)
	_, err = st.Exec(r.Wife, r.Id)
	utils.Panic(err)
	return true
}
