package services

import (
	"gramya/constants"
	"gramya/dao"
	"gramya/model"
)

func AddPerson(wc model.WelcomeRequest) bool {
	person := dao.Person{Id: wc.Id, FName: wc.FName, MName: wc.MName, LName: wc.LName, Dob: wc.Dob, Authors: wc.Authors, Tags: wc.Tags}
	ap := dao.InsertPerson(&person)
	var ar bool
	if wc.WelcomeType == constants.Child {
		relation := dao.Relation{Id: wc.Id, FName: wc.FName, Father: wc.Pid, Authors: wc.Authors}
		ar = dao.InsertRelation(&relation)
	} else {
		relation := dao.Relation{Id: wc.Pid, Wife: wc.Id}
		ar = dao.UpdateRelation(&relation)
	}
	return ap && ar
}
