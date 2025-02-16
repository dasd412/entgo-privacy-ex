package httperror

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

// HttpErrorPresenter는 발생한 에러를 GraphQL 오류 형식에 맞게 wrapping합니다.
func ErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	// 기본 httperror presenter로 변환
	gqlErr := graphql.DefaultErrorPresenter(ctx, err)

	var httpErr *HTTPError

	if errors.As(err, &httpErr) {
		gqlErr.Message = httpErr.Message

		switch httpErr.StatusCode {
		case http.StatusUnauthorized:
			gqlErr.Extensions = map[string]interface{}{
				"code": "UNAUTHORIZED",
			}
		case http.StatusForbidden:
			gqlErr.Extensions = map[string]interface{}{
				"code": "FORBIDDEN",
			}
		default:
			gqlErr.Extensions = map[string]interface{}{
				"code": "INTERNAL_ERROR",
			}
		}
	}

	return gqlErr
}
