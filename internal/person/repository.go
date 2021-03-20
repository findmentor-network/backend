package person

import (
	"context"
	mongohelper "github.com/findmentor-network/backend/pkg/mongoextentions"
	"github.com/findmentor-network/backend/pkg/pagination"
)

type (
	Repository interface {
		Get(context.Context, *mongohelper.Query, *pagination.Pages) ([]*Person, error)
	}
)
