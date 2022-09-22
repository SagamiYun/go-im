package service

import (
	"github.com/gin-gonic/gin"
	"im/helper"
	"im/models"
	"net/http"
	"strconv"
)

func ChatList(c *gin.Context) {
	roomIdentity := c.Query("room_identity")
	if roomIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "房间号不能为空",
		})
		return
	}
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	_, err := models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, roomIdentity)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "非法访问",
		})
		return
	}

	parseIndex, _ := strconv.ParseInt("page_index", 10, 32)
	parseSize, _ := strconv.ParseInt("page_size", 10, 32)

	skip := (parseIndex - 1) * parseSize
	data, err := models.GetMessageListByRoomIdentity(roomIdentity, &parseSize, &skip)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": gin.H{
			"list": data,
		},
	})
}
