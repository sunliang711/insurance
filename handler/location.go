package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sunliang711/insurance/model"
	"github.com/sunliang711/insurance/types"
	"github.com/sunliang711/insurance/utils"
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
		Code: types.OK,
		Msg:  "OK",
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
		{Label: "交通事故", Value: "1"},
		{Label: "住院", Value: "2"},
		{Label: "航班延误", Value: "3"},
	}
	c.JSON(200, types.Response{
		Code: types.OK, Msg: "OK", Data: options,
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
		{Label: "意外", Value: "1"},
		{Label: "疾病", Value: "2"},
	}
	c.JSON(200, types.Response{
		Code: types.OK, Msg: "OK", Data: options,
	})
}

// Claim TODO
// @Summary 理赔申请
// @Tags 理赔申请
// @Description 理赔申请
// @Accept  json
// @Produce  json
// @Param claim body types.Claim true "COMMENT"
// @Success 200 {object} types.Response "COMMENT"
// @Failure 400 {object} types.Response "COMMENT"
// @Router /claim [post]
// 2019/08/14 11:36:02
// Claim 理赔申请
// 2019/08/14 11:35:40
func Claim(c *gin.Context) {
	log.Debugf("Claim handler")
	var claim types.Claim
	if err := c.BindJSON(&claim); err != nil {
		c.JSON(400, types.Response{
			Code: types.BadRequest, Msg: "request format error", Data: nil,
		})
		log.Error("request format error")
		return
	}
	log.Debugf("claim: %+v", claim)
	if err := model.AddClaim(&claim); err != nil {
		c.JSON(500, types.Response{
			Code: types.ServerInternalError, Msg: "server internal error", Data: nil,
		})
		log.Error("server internal error")
		return
	}
	c.JSON(200, types.Response{
		Code: types.OK, Msg: "OK", Data: nil,
	})
}

// Decrypt 接受decrypt数据并揭秘
// data":[{"capsule":"QmRuNX95x4ZXqpCcTtJfoukW3jyFo6yLox62wBiZCPYc6t","ciphertext":"QmUid62eMk1Y6oDEYfXo7oiQjTtzv17J889qjhDHnBmaMG","filename":"1565875831_0_sun.png","decrypt_mode":"1","deckey":""},{"capsule":"QmRj3vPidHbXwevBHRecEFuHQYRQozL3D9bHdsX4qc5Q6X","ciphertext":"QmbPmtrwiM487zQyRaVPUwCoAKjJRmax1F5Tyk5d6eqGv1","filename":"1565882396_0_sun.png","decrypt_mode":"1","deckey":""},{"capsule":"QmaEwAfaVgn5Dxq8b1gdpTHeXtG4wjDHhS9mpSqhxjBSbZ","ciphertext":"QmXN6wrTYbjESafbgLqLgARaUoyFE6kqd7yau25CJAmqkj","filename":"1565882400_0_sun.png","decrypt_mode":"1","deckey":""},{"capsule":"QmdMzJT1U8ReKDyZHkVJKA8S87dWjnnc5pxqVnJR3i6D7C","ciphertext":"QmbpqXvphcEsHG979J1g6gtR9jEJdTdiULfLxv1P5CWJ6Q","filename":"1565882403_0_sun.png","decrypt_mode":"1","deckey":""},{"capsule":"QmNmL3iuQLjkkbVyyn3CpShJCmBFxsNkJ9oWpuYn6YAfZA","ciphertext":"QmYSCEdPgkovRbEXjo5X7iQmPLDsJU4iCGRj2WYVLsbe4V","filename":"1565882407_0_sun.png","decrypt_mode":"1","deckey":""}]}
//{"capsule":"","ciphertext":"","filename":"","decrypt_mode":"1","deckey":""}
// 2019/08/12 10:18:51
func Decrypt(c *gin.Context) {
	var dinput []types.DecryptInput
	if err := c.BindJSON(&dinput); err != nil {
		msg := fmt.Sprintf("Bad request data")
		c.JSON(400, types.Response{
			Code: 1,
			Msg:  msg,
		})
		log.Error(msg)
		return
	}
	log.Debugf("decryptInput: %+v", dinput)

	c.JSON(200, types.Response{
		Code: 0,
		Msg:  "OK",
	})
	for _, item := range dinput {
		item.DecKey = types.SK

		//call decrypt api
		//TODO
		utils.PostJson("http://localhost:5000/decrypt", map[string]string{
			"token": "burenshizheshigesha",
		}, item)

	}

}
