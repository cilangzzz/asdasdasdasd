package model

import "time"

type User struct {
	Id               int64     `json:"id" xorm:"autoincr pk" `
	Username         string    `json:"username"binding:"required" xorm:"varchar(50) notnull unique" schema:"username"`
	IdNumber         string    `json:"id_number" xorm:"varchar(50) notnull" schema:"id_number"`
	CompanyName      string    `json:"company_name" xorm:"varchar(50) notnull" schema:"company_name"`
	SocialId         string    `json:"social_id" xorm:"varchar(50) notnull" schema:"social_id"`
	Sex              string    `json:"sex" xorm:"varchar(10) notnull" schema:"sex"`
	UnemploymentTime string    `json:"unemployment_time" xorm:"varchar(50) notnull" schema:"unemployment_time"`
	UnemploymentMon  string    `json:"unemployment_mon" xorm:"varchar(50) notnull" schema:"unemployment_mon"`
	UnemploymentSMon string    `json:"unemployment_s_mon" xorm:"varchar(50) notnull" schema:"unemployment_s_mon"`
	UnemploymentDMon string    `json:"unemployment_d_mon" xorm:"varchar(50) notnull" schema:"unemployment_d_mon"`
	WorkInjuryTime   string    `json:"work_injury_time" xorm:"varchar(50) notnull" schema:"work_injury_time"`
	WorkInjuryMon    string    `json:"work_injury_mon" xorm:"varchar(50) notnull" schema:"work_injury_mon"`
	WorkInjurySMon   string    `json:"work_injury_s_mon" xorm:"varchar(50) notnull" schema:"work_injury_s_mon"`
	WorkInjuryDMon   string    `json:"work_injury_d_mon" xorm:"varchar(50) notnull" schema:"work_injury_d_mon"`
	CreateTime       time.Time `json:"create_time" xorm:"created" schema:"create_time"`
	UpdateTime       time.Time `json:"update_time" xorm:"updated" schema:"update_time"`
}
