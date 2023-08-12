package gin_controller_test

// controllerをpresenterのテストを両方行っている

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	gin_controller "github.com/FlowingSPDG/get5-web-go/backend/controller/gin"
)

func TestGetVersion(t *testing.T) {
	// mock controllerの作成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	controller := gin_controller.NewGetVersionController()
	router := setupGET("/api/v1/version", controller.Handle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/version", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"version": "0.2.0"}`, w.Body.String())
}
