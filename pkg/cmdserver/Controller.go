package cmdserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
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


func handlerConfigureHost(ctx *gin.Context) {
	var rule = processor.HostMatchRule{}
	if err := ctx.ShouldBindJSON(&rule); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rule.Init()
	if len(rule.Host) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Empty field: host"})
		return
	}
	if !rule.HasActiveFilter() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No active filter"})
		return
	}
	processor.HostMatcherService.Append(rule)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "rule added",
	})
}
