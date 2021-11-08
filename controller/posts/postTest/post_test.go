package postTest

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"web_app/controller/posts"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST("/api/v1/post", posts.PostCommunityHandle)
	body := `{
		"title":"这是新帖子",
		"content":"这是内容",
		"communityid":3
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// 判断内容及状态码是不是预期的内容
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "登录")
}
