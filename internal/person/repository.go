package person

import (
	"context"
	"github.com/findmentor-network/backend/pkg/pagination"
)

type (
	Repository interface {
		Get(context.Context, *pagination.Pages) ([]*Person, error)
	}
)
