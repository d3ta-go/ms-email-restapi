package router

import (
	"github.com/d3ta-go/ms-email-restapi/interface/http-apps/restapi/echo/features/email"
	internalMiddleware "github.com/d3ta-go/system/interface/http-apps/restapi/echo/middleware"
	"github.com/labstack/echo/v4"
)

// SetEmail set Email Router
func SetEmail(eg *echo.Group, f *email.FEmail) {

	gc := eg.Group("/email")
	gc.Use(internalMiddleware.JWTVerifier(f.GetHandler()))

	gc.POST("/send", f.SendEmail)

	gc.GET("/templates/list-all", f.ListAllEmailTemplate)
	gc.GET("/template/:code", f.FindEmailTemplateByCode)
	gc.POST("/template", f.CreateEmailTemplate)
	gc.PUT("/template/update/:code", f.UpdateEmailTemplate)
	gc.PUT("/template/set-active/:code", f.SetActiveEmailTemplate)
	gc.DELETE("/template/:code", f.DeleteEmailTemplate)
}
