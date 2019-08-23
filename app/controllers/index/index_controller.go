package indexController

import (
    "encoding/json"
    model "gin-test/app/models"
    "gin-test/utils/code"
    "net/http"
    "strings"
    
    "gin-test/common"
    "gin-test/utils"
    "github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Ping
// @Description Test Ping
// @Accept  json
// @Produce  json
// @Tags Test
// @Success 200 {object} common.Response
// @Router /test/ping [get]
func Ping(c *gin.Context) {
    appG := common.Gin{C: c}
    appG.Response(http.StatusOK, code.SUCCESS, "pong", nil)
}

// Test godoc
// @Summary Base64 Decode
// @Accept  json
// @Produce  json
// @Tags Test
// @Param base64 query string true "base64 string"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /test/font [get]
func Test(c *gin.Context) {
    appG := common.Gin{C: c}

    base64 := c.DefaultQuery("base64", "")

    // 替换字符串
    base64String := strings.Replace(base64, " ", "+", -1)

    // base64 解码
    arrByte, error := utils.Base64Decode(base64String)
    if error != nil {
        appG.Response(http.StatusOK, code.INVALID_PARAMS, "base64解码失败" + error.Error(), nil)
        return
    }

    // 结构体
    var imgTextArray []common.ImgText
    err := json.Unmarshal(arrByte, &imgTextArray)

    if utils.HandleError(c, http.StatusInternalServerError, err) {
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "文字解析成功", imgTextArray)
}

func GetTestUsers(c *gin.Context) {
    var (
        testUsers [] *model.Test
    )

    appG := common.Gin{C: c}
    
    //data := make(map[string]interface{})
    
    testUsers, error := model.GetTestUsers(0, 15)
    if error != nil {
        appG.Response(http.StatusOK, code.INVALID_PARAMS, error.Error(), nil)
        return
    }

    appG.Response(http.StatusOK, code.SUCCESS, "ok", testUsers)
}
