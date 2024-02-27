package util

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetMockContext() context.Context {

	w := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(w)

	context.Request = httptest.NewRequest(http.MethodPost, "/", nil)
	
	return context.Request.Context()
}
