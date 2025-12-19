package middleware

import (
	"go-ec-sample/consts"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		isAdmin := session.Get(consts.SessionKeyIsAdmin)

		if isAdmin != true {
			ctx.String(http.StatusForbidden, "You are not admin.")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
