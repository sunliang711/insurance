package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sunliang711/insurance/types"
)

// Location 全国省市
// @Summary 全国省市
// @Tags 全国省市
// @Description 返回全国省市列表
// @Produce  json
// @Success 200 {object} types.Response "省市列表"
// @Router /location [get]
// 2019/08/09 10:23:40
// Location TODO
// 2019/08/09 10:26:30
func Location(c *gin.Context) {
	// var proCity []types.Province
	// if err := json.Unmarshal([]byte(types.ProvinceWithCities), &proCity); err != nil {
	// 	msg := fmt.Sprintf("Unmarshal json error: %v", err)
	// 	c.JSON(500, types.Response{1, msg, nil})
	// 	return
	// }
	// c.JSON(200, types.Response{0, "OK", proCity})
	c.JSON(200, types.Response{
		Code: 0,
		Msg:  "ok",
		Data: types.ProCity,
	})
}

// Reason 出险原因
// @Summary 出险原因
// @Tags 出险原因
// @Description TODO
// @Accept  json
// @Produce  json
// @Success 200 {object} types.Response "COMMENT"
// @Router /reason [get]
// 2019/08/12 14:40:28
// Reason 出险原因
// 2019/08/12 10:10:32
func Reason(c *gin.Context) {
	options := []types.Option{
		{"交通事故", "1"},
		{"住院", "2"},
		{"航班延误", "3"},
	}
	c.JSON(200, types.Response{
		0, "OK", options,
	})
}

// Typ 出险类型
// @Summary 出险类型
// @Tags 出险类型
// @Description TODO
// @Accept  json
// @Produce  json
// @Success 200 {object} types.Response "COMMENT"
// @Router /type [get]
// 2019/08/12 14:41:37
// Typ 出险类型
// 2019/08/12 10:17:20
func Typ(c *gin.Context) {
	options := []types.Option{
		{"意外", "1"},
		{"疾病", "2"},
	}
	c.JSON(200, types.Response{
		0, "OK", options,
	})
}

// DeviceAuth 核赔设备授权
// 2019/08/12 10:18:51
func DeviceAuth(c *gin.Context) {

}
