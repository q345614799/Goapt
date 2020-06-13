package lhbb

import (
	"apt/model/lhbbmodel"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context)  {

	//gxb :=lhbbmodel.LHBB_NCP_PC_GXB{
	//	NGXB_ID:       "1111111111111111111",
	//	NGXB_PCID:     "1111111111111111112",
	//	NGXB_NCPID:    1,
	//	NGXB_BBZL:     500.00,
	//	NGXB_SCJD:     "，哦我",
	//	NGXB_SCDQ:     "北京市-北京市-东城区",
	//	NGXB_DWH:      "",
	//	NGXB_SCZL:     0.00,
	//	NGXB_ZT:       0,
	//	FLAG_HG:       0,
	//	FLAG_SCZT:     0,
	//	NGXB_SCDQ_SDJ: "北京市",
	//	NGXB_SCDQ_QXJ: "东城区",
	//	NGXB_SCDQ_SFJ: "北京",
	//}
	//
	//
	//lhbb := lhbbmodel.LHBB_NCP_YSPCB{
	//	YSPCB_ID:        "1111111111111111112",
	//	YSPCB_SIID:      "1210368807403659265",
	//	YSPCB_CPH:       "赣A82758",
	//	YSPCB_CD:        "山东省-济南市-槐荫区",
	//	YSPCB_ZL:        44360.00,
	//	YSPCB_ZM:        "",
	//	YSPCB_CJSJ:      time.Now(),
	//	YSPCB_CJZID:     "1210367760966864898",
	//	YSPCB_GXSJ:      time.Now(),
	//	YSPCB_GXZID:     "10004",
	//	YSPCB_ZT:        1,
	//	YSPCB_HZMC:      "樊保明",
	//	YSPCB_HZDH:      "13576273831",
	//	YSPCB_YJDDSJ:    time.Time{},
	//	CX:              29,
	//	YSPCB_SJMC:      "张灿华",
	//	YSPCB_SJDH:      "13767159542",
	//	YSPCB_FLAG_LDSH: 0,
	//	YSPCB_CDSF:      "黑龙江省",
	//	TJRK:            2,
	//	LHBB_NCP_PC_GXB: []lhbbmodel.LHBB_NCP_PC_GXB{gxb},
	//}
	get:=lhbbmodel.LHBB_NCP_YSPCB{}

	fmt.Printf("%#v",get)

	//err,_ :=lhbb.Add()
	//
	//if err!=nil {
	//	c.JSON(http.StatusBadRequest,gin.H{
	//		"massage":"测试失败",
	//		"error":err.Error(),
	//	})
	//}
	//
	//
	//
	//c.JSON(http.StatusOK,gin.H{
	//"massage":"测试成功",
	//})

}
