package cmdserver

import (
	"github.com/gin-gonic/gin"
	"rzreversescheme/pkg/processor"
)

func handlerStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  "worked",
		"message": "ready to connection",
	})
}


func handlerSchema(ctx *gin.Context) {
	host := ctx.Param("host")
	proc := processor.GetProcessorByHost(host)
	if (proc != nil) {
		ctx.String(200, proc.GetScheme())
	} else {
		ctx.JSON(404, gin.H{
			"status":  "not found",
			"message": "scheme for host " + host + " not found",
		})

	}
}
