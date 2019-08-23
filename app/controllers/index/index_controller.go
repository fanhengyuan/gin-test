package indexController

import (
    "encoding/json"
    "gin-test/utils/code"
    "net/http"
    "strings"
    
    "gin-test/common"
    "gin-test/utils"
    "github.com/gin-gonic/gin"
)

// @Summary Base64 Decode
// @Produce  json
// @Param base64 query string true "base64 string"
// @Router /font [get]
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
