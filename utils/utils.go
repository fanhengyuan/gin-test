package utils

import (
    "encoding/base64"
    "github.com/gin-gonic/gin"
    "log"
    "runtime"
    "time"
    "unicode"
)


func UcFirst(str string) string {
    for _, v := range str {
        u := string(unicode.ToUpper(v))
        return u + str[len(u):]
    }
    return ""
}

func LcFirst(str string) string {
    for _, v := range str {
        u := string(unicode.ToLower(v))
        return u + str[len(u):]
    }
    return ""
}

// get datetime now
func GetDateTimes() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func TimeToDateTimesString(t time.Time) string{
    return t.Format("2006-01-02 15:04:05")
}

// base64 解码
func Base64Decode(raw string) ([]byte, error) {
    encodeBytes, err := base64.StdEncoding.DecodeString(raw)
    if err != nil {
        return nil, err
    }
    return encodeBytes, nil
}

// 打印堆栈信息 方便 debug
func PrintStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    log.Printf("==> %s\n", string(buf[:n]))
}

// 统一错误处理
func HandleError(c *gin.Context, httpCode int, err error) bool {
    if err != nil {
        _, file, line, ok := runtime.Caller(1)

        log.Printf("Error: %s\nFile: %s Line: %d  %t", err, file, line, ok)

        PrintStack()
        c.JSON(httpCode, gin.H{"data": nil, "code": httpCode, "msg": err.Error()})
        return true
    }

    return false
}
