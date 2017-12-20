package router

import (
	"github.com/gin-gonic/gin"
)

type Hello struct {
}

func (h *Hello) HTTP(c *gin.Context) {
	panic("not implemented")
}

func (h *Hello) GRPC(c *gin.Context) {
	panic("not implemented")
}
