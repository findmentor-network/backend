package controller

import (
	"fmt"
	"github.com/findmentor-network/backend/internal/person"
	"github.com/findmentor-network/backend/pkg/log"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_resource_GetAll(usecase *testing.T) {

	e := echo.New()
	e.Logger = log.SetupLogger()
	//e.Use(echoextention.Recover())

	expectedData := []*person.Person{
		{RegisteredAt: "", Name: "", TwitterHandle: " ", Github: "", Linkedin: "", Interests: "", Goals: "", Mentor: ""},
	}
	ctrl := gomock.NewController(usecase)
	mockRepository := person.NewMockRepository(ctrl)
	mockRepository.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedData, nil).AnyTimes()

	controller := NewController(mockRepository)
	NewHandlers(e, controller)

	usecase.Run("API Getter", func(t *testing.T) {

		tests := []struct {
			name           string
			url            string
			wantStatusCode int
		}{
			{name: "With valid url and params should return success", url: "/api/v1/persons?page=1&size=3", wantStatusCode: http.StatusOK},
			{name: "With valid filter url and params should return success", url: "/api/v1/persons/filter?page=5&size=1", wantStatusCode: http.StatusOK},
			{name: "With valid filter url and params should return success", url: "/api/v1/persons/filter?isHireable=true", wantStatusCode: http.StatusOK},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(tt.url), nil)
				t.Logf("Url:%s", tt.url)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				ctx := e.NewContext(req, rec)
				res := rec.Result()
				defer res.Body.Close()
				if err := controller.Get(ctx); err == nil {
					assert.NotNil(t, rec.Body.String())
				}
				t.Logf("Result:%s, %d", rec.Body.String(),rec.Result().StatusCode)
				assert.Equal(t, tt.wantStatusCode, rec.Code)
			})
		}

	})
}
