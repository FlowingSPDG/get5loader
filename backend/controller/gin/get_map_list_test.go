package gin_controller_test

// controllerをpresenterのテストを両方行っている

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	gin_controller "github.com/FlowingSPDG/get5loader/backend/controller/gin"
)

func TestGetMapList(t *testing.T) {
	tt := []struct {
		name  string
		input struct {
			activeMapList  []string
			reserveMapList []string
		}
		expected struct {
			code int
			body string
		}
	}{
		{
			name: "success",
			input: struct {
				activeMapList  []string
				reserveMapList []string
			}{
				activeMapList: []string{
					"de_inferno",
					"de_mirage",
					"de_nuke",
					"de_overpass",
					"de_vertigo",
					"de_ancient",
					"de_anubis",
				},
				reserveMapList: []string{},
			},
			expected: struct {
				code int
				body string
			}{
				code: http.StatusOK,
				body: `{"active":["de_inferno","de_mirage","de_nuke","de_overpass","de_vertigo","de_ancient","de_anubis"],"reserve":[]}`,
			},
		},
	}

	for _, tc := range tt {
		// mock controllerの作成
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		controller := gin_controller.NewGetMaplistController(tc.input.activeMapList, tc.input.reserveMapList)
		t.Run(tc.name, func(t *testing.T) {
			router := setupGET("/api/v1/maps", controller.Handle)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v1/maps", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expected.code, w.Code)
			assert.JSONEq(t, tc.expected.body, w.Body.String())
		})
	}
}
