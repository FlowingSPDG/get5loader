package gin_controller_test

import "github.com/gin-gonic/gin"

func setupGET(path string, f func(c *gin.Context)) *gin.Engine {
	r := gin.New()
	r.GET(path, f)
	return r
}

func setupPOST(path string, f func(c *gin.Context)) *gin.Engine {
	r := gin.New()
	r.POST(path, f)
	return r
}

func setupPUT(path string, f func(c *gin.Context)) *gin.Engine {
	r := gin.New()
	r.PUT(path, f)
	return r
}

func setupDELETE(path string, f func(c *gin.Context)) *gin.Engine {
	r := gin.New()
	r.DELETE(path, f)
	return r
}
