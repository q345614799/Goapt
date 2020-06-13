package lhbbmodel

import (
	"time"
	"apt/init/initsql"
	"github.com/go-playground/validator/v10"
)

type  LHBB_NCP_YSPCB struct {
	YSPCB_ID string `gorm:"column:YSPCB_ID" validate:"len=10" json:"CX_ID"`
	YSPCB_SIID string `gorm:"column:YSPCB_SIID" json:"YSPCB_SIID"`
	YSPCB_CPH string `gorm:"column:YSPCB_CPH" json:"YSPCB_CPH"`
	YSPCB_CD string `gorm:"column:YSPCB_CD" json:"YSPCB_CD"`
	YSPCB_ZL float64 `gorm:"column:YSPCB_ZL" json:"YSPCB_ZL"`
	YSPCB_ZM string `gorm:"column:YSPCB_ZM" json:"YSPCB_ZM"`
	YSPCB_CJSJ time.Time `gorm:"column:YSPCB_CJSJ" json:"YSPCB_CJSJ"`
	YSPCB_CJZID string `gorm:"column:YSPCB_CJZID" json:"YSPCB_CJZID"`
	YSPCB_GXSJ time.Time `gorm:"column:YSPCB_GXSJ" json:"YSPCB_GXSJ"`
	YSPCB_GXZID string `gorm:"column:YSPCB_GXZID" json:"YSPCB_GXZID"`
	YSPCB_ZT byte `gorm:"column:YSPCB_ZT" json:"YSPCB_ZT"`
	YSPCB_HZMC string `gorm:"column:YSPCB_HZMC" json:"YSPCB_HZMC"`
	YSPCB_HZDH string `gorm:"column:YSPCB_HZDH" json:"YSPCB_HZDH"`
	YSPCB_YJDDSJ time.Time `gorm:"column:YSPCB_YJDDSJ" json:"YSPCB_YJDDSJ"`
	CX int `gorm:"column:CX" json:"CX"`
	YSPCB_SJMC string `gorm:"column:YSPCB_SJMC" json:"YSPCB_SJMC"`
	YSPCB_SJDH string `gorm:"column:YSPCB_SJDH" json:"YSPCB_SJDH"`
	YSPCB_FLAG_LDSH byte `gorm:"column:YSPCB_FLAG_LDSH" json:"YSPCB_FLAG_LDSH"`
	YSPCB_CDSF string `gorm:"column:YSPCB_CDSF" json:"YSPCB_CDSF"`
	TJRK int `gorm:"column:TJRK" json:"TJRK"`
	LHBB_NCP_PC_GXB []LHBB_NCP_PC_GXB `gorm:"ForeignKey:NGXB_PCID;ASSOCIATION_FOREIGNKEY:YSPCB_ID"`
}

type LHBB_NCP_PC_GXB struct {
	NGXB_ID string `gorm:"column:NGXB_ID" json:"NGXB_ID"`
	NGXB_PCID string `gorm:"column:NGXB_PCID" json:"NGXB_PCID"`
	NGXB_NCPID int `gorm:"column:NGXB_NCPID" json:"NGXB_NCPID"`
	NGXB_BBZL float64 `gorm:"column:NGXB_BBZL" json:"NGXB_BBZL"`
	NGXB_SCJD string `gorm:"column:NGXB_SCJD" json:"NGXB_SCJD"`
	NGXB_SCDQ string `gorm:"column:NGXB_SCDQ" json:"NGXB_SCDQ"`
	NGXB_DWH string `gorm:"column:NGXB_DWH" json:"NGXB_DWH"`
	NGXB_SCZL float64 `gorm:"column:NGXB_SCZL" json:"NGXB_SCZL"`
	NGXB_ZT byte `gorm:"column:NGXB_ZT" json:"NGXB_ZT"`
	FLAG_HG byte `gorm:"column:FLAG_HG" json:"FLAG_HG"`
	FLAG_SCZT byte `gorm:"column:FLAG_SCZT" json:"FLAG_SCZT"`
	NGXB_SCDQ_SDJ string `gorm:"column:NGXB_SCDQ_SDJ" json:"NGXB_SCDQ_SDJ"`
	NGXB_SCDQ_QXJ string `gorm:"column:NGXB_SCDQ_QXJ" json:"NGXB_SCDQ_QXJ"`
	NGXB_SCDQ_SFJ string `gorm:"column:NGXB_SCDQ_SFJ" json:"NGXB_SCDQ_SFJ"`
}





func (LHBB_NCP_YSPCB) TableName() string {
	return "LHBB_NCP_YSPCB"
}

func (LHBB_NCP_PC_GXB) TableName() string {
	return "LHBB_NCP_PC_GXB"
}


func (u *LHBB_NCP_YSPCB)Add() (err error, uenter *LHBB_NCP_YSPCB) {
	err = initsql.DEFAULTDB.Debug().Create(u).Error
	return
}

func (u *LHBB_NCP_YSPCB)addsjjy()(err error)  {
	validate := validator.New()
	err = validate.Struct(u)
	return
}

func GetLhbbBySID(YSPCB_SIID string)(u *LHBB_NCP_YSPCB) {

	u = &LHBB_NCP_YSPCB{}

	initsql.DEFAULTDB.Debug().Preload("LHBB_NCP_PC_GXB").Where("YSPCB_SIID = ?",YSPCB_SIID).Find(u)
	return
}