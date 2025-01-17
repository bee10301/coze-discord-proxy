package controller

import (
	"coze-discord-proxy/discord"
	"coze-discord-proxy/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ChannelCreate 创建频道
// @Summary 创建频道
// @Description 创建频道
// @Tags channel
// @Accept json
// @Produce json
// @Param channelModel body model.ChannelReq true "channelModel"
// @Success 200 {object} model.ChannelResp "Successful response"
// @Router /channel/create [post]
func ChannelCreate(c *gin.Context) {

	var channelModel model.ChannelReq
	err := json.NewDecoder(c.Request.Body).Decode(&channelModel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}

	var channelId string

	if channelModel.ParentId == "" {
		channelId, err = discord.ChannelCreate(discord.GuildId, channelModel.Name)
	} else {
		channelId, err = discord.ChannelCreateComplex(discord.GuildId, channelModel.ParentId, channelModel.Name)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "discord创建频道异常",
		})
	} else {
		var channel model.ChannelResp
		channel.Id = channelId
		channel.Name = channelModel.Name
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    channel,
		})
	}
	return
}

// ChannelDel 删除频道
// @Summary 删除频道
// @Description 删除频道
// @Tags channel
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} string "Successful response"
// @Router /channel/del/{id} [get]
func ChannelDel(c *gin.Context) {
	channelId := c.Param("id")

	if channelId == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}

	channelId, err := discord.ChannelDel(channelId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "discord删除频道异常",
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "discord删除频道成功",
		})
	}
	return
}
