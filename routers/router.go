package routers

import (
    authController "gin-test/app/controllers/v1/auth"
    indexController "gin-test/app/controllers/v1/index"
    reportController "gin-test/app/controllers/v1/report"
    "gin-test/app/middleware"
    "gin-test/docs"
    "gin-test/utils/casbin"
    "gin-test/utils/setting"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "net/http"
    "strings"
)

func InitRouter() *gin.Engine {
    // programatically set swagger info
    docs.SwaggerInfo.Title = "Swagger Example API"
    docs.SwaggerInfo.Description = "This is a sample server Petstore server."
    docs.SwaggerInfo.Version = "1.0"

    configHost := ""
    prefixUrl := setting.AppSetting.PrefixUrl
    if strings.HasPrefix(prefixUrl, "http://") {
        configHost = strings.Replace(prefixUrl, "http://", "", -1)
        docs.SwaggerInfo.Schemes = []string{"http"}
    }else if strings.HasPrefix(prefixUrl, "https://") {
        configHost = strings.Replace(prefixUrl, "https://", "", -1)
        docs.SwaggerInfo.Schemes = []string{"https"}
    }
    docs.SwaggerInfo.Host = configHost
    docs.SwaggerInfo.BasePath = "/v1/api/"

    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    if setting.AppSetting.EnabledCORS {
        r.Use(middleware.CORS())
    }

    // 初始化路由权限 在这初始化的目的： 避免每次访问路由查询数据库
    // 如果更改路由权限 需要重新调用一下这个方法
    casbin.SetupCasbin()

    v1 := r.Group("/v1/api")
    {
        v1.POST("/login", authController.GetAuth)
        report := v1.Group("/report").Use(middleware.TranslationHandler())
        {
            report.POST("", reportController.Report)
        }

        test := v1.Group("/test").Use(
            middleware.TranslationHandler(),
            middleware.JWTHandler(),
            middleware.CasbinHandler())
        {
            test.POST("/ping", indexController.Ping)
            test.GET("/ping", indexController.Ping)
            test.GET("/font", indexController.Test)
            test.GET("/test_users", indexController.GetTestUsers)
            test.POST("/login", indexController.UserLogin)
        }

    }

    // swagger docs
    r.GET("/swagger", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
