package model

import (
	"gramya/constants"
	"testing"
)

func TestWelcomeRequest_Validate(t *testing.T) {
	type fields struct {
		Id          int
		Pid         int
		FName       string
		MName       string
		LName       string
		Dob         string
		Authors     string
		Tags        string
		WelcomeType constants.WelcomeType
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"IdAndPid", fields{1, 2, "F", "", "L", "2019-01-01", "A", "A", 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wr := &WelcomeRequest{
				Id:          tt.fields.Id,
				Pid:         tt.fields.Pid,
				FName:       tt.fields.FName,
				MName:       tt.fields.MName,
				LName:       tt.fields.LName,
				Dob:         tt.fields.Dob,
				Authors:     tt.fields.Authors,
				Tags:        tt.fields.Tags,
				WelcomeType: tt.fields.WelcomeType,
			}
			if got := wr.Validate(); got != tt.want {
				t.Errorf("WelcomeRequest.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
