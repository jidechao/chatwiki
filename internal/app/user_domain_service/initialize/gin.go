// Copyright © 2016- 2024 Sesame Network Technology all right reserved

package initialize

import (
	"chatwiki/internal/app/user_domain_service/define"
	"chatwiki/internal/app/user_domain_service/middlewares"
	"chatwiki/internal/app/user_domain_service/route"
	"chatwiki/internal/pkg/lib_web"

	"github.com/gin-gonic/gin"
)

func initGin() {
	if !define.IsDev {
		gin.SetMode(gin.ReleaseMode)
	}
	port := define.Config.WebService[`port`]
	define.WebService = lib_web.InitGin(port, route.Route, middlewares.CasbinAuth())
}
