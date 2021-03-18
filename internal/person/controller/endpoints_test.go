package controller

import (
	"fmt"
	"github.com/findmentor-network/backend/internal/person"
	"github.com/findmentor-network/backend/pkg/echoextention"
	"github.com/findmentor-network/backend/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_resource_GetAll(usecase *testing.T) {

	e := echo.New()
	e.Logger = log.SetupLogger()
	e.Use(echoextention.Recover())
	m := mock.Mock{}

	expectedData := []*person.Person{
		{RegisteredAt: "", Name: "", TwitterHandle: " ", Github: "", Linkedin: "", Interests: "", Goals: "", Mentor: ""},
	}

	m.On("Get", mock.Anything,
		mock.Anything).
		Return(expectedData, nil)

	mockRepository := person.NewMockRepository(m)
	controller := NewController(mockRepository)
	NewHandlers(e, controller)

	usecase.Run("API Getter", func(t *testing.T) {

		tests := []struct {
			name           string
			args           []string
			wantStatusCode int
		}{
			{name: "With valid url and params should return success", args: []string{"/api/v1/persons?page=%s&size=%s", "1", "3"}, wantStatusCode: http.StatusOK},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(tt.args[0], tt.args[1], tt.args[2]), nil)
				t.Logf("Url:%s", tt.args[0])
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				ctx := e.NewContext(req, rec)
				res := rec.Result()
				defer res.Body.Close()
				if err := controller.GetAll(ctx); err == nil {
					assert.NotNil(t, rec.Body.String())
				}
				t.Logf("Result:%s", rec.Body.String())
				assert.Equal(t, tt.wantStatusCode, rec.Code)
			})
		}

	})
}
