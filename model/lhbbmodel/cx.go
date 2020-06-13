package lhbbmodel

import "apt/init/initsql"

type LHBB_NCP_CX struct {
	CX_ID int `gorm:"column:CX_ID" validate:"" json:"CX_ID"`
	CX_MC string `gorm:"column:CX_MC" validate:"" json:"CX_MC"`
	CX_ZT int `gorm:"column:CX_ZT" validate:"" json:"CX_ZT"`
}

func (LHBB_NCP_CX) TableName() string {
	return "LHBB_NCP_CX"
}

func GetCxList() *[]LHBB_NCP_CX  {
	var cx []LHBB_NCP_CX
	initsql.DEFAULTDB.Find(&cx)
	return &cx
}