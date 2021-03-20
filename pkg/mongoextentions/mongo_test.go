package mongohelper

import (
	"github.com/pkg/errors"
	"testing"
)

func TestNewDatabase(usecase *testing.T) {
	usecase.Run("When creating a new mongo client", func(t *testing.T) {

		tests := []struct {
			name     string
			args     []string
			wantErr  bool
			whichErr error
		}{
			//{name: "with valid parameters_should not return error", args: []string{"mongodb://root:example@localhost:27017", "admin"}, wantErr: false},
			//{name: "with invalid parameters_should return error", args: []string{"mongodb://an_invalid_address:27017?connectTimeoutMS=300", "admin"}, wantErr: true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := NewDatabase(tt.args[0], tt.args[1])

				if err != nil && tt.wantErr == true && errors.Cause(err) != ConnectionError {
					t.Errorf("NewDatabase() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
}
