package model

import (
	"gramya/constants"
	"gramya/utils"
)

type WelcomeRequest struct {
	Id          int
	Pid         int
	FName       string
	MName       string
	LName       string
	Dob         string
	Authors     string
	Tags        string
	WelcomeType constants.WelcomeType `iota`
}

type ResponseOk struct {
	Success bool
	Msg     string
}

func (wr *WelcomeRequest) Validate() bool {
	return wr.Id != wr.Pid &&
		(wr.WelcomeType == constants.Child || wr.WelcomeType == constants.Bride) &&
		utils.NotNullAndLetter(wr.FName) &&
		utils.NullOrLetter(wr.MName) &&
		utils.NotNullAndLetter(wr.LName) &&
		utils.NotNullAndDateFormat(wr.Dob) &&
		utils.NullOrCommaSeparatedLetter(wr.Authors) &&
		utils.NullOrCommaSeparatedLetter(wr.Tags)
}
