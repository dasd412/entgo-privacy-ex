package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"privacy-ex/pkg/ent"
	"privacy-ex/pkg/graph/httperror"
	"strings"
)

var requestData struct {
	OperationName string `json:"operationName"`
}

// 인증이 필요 없는 API 목록
var publicOperations = map[string]bool{
	"signup":             true,
	"login":              true,
	"IntrospectionQuery": true, //graphql playground 용
	"post":               true, // 기본적으로 모든 사용자가 조회 가능
}

func getOperationName(r *http.Request) (string, error) {
	// graphql 요청 본문 (json) 읽기
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	//json 파싱
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return "", err
	}

	//요청 본문을 다시 복원
	r.Body = io.NopCloser(strings.NewReader(string(body)))

	return requestData.OperationName, nil
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			operationName, err := getOperationName(r)

			if err != nil {
				httperror.SetErrorResponse(w, r.Context(), &httperror.HTTPError{
					StatusCode: http.StatusBadRequest,
					Message:    "Invalid operation: " + err.Error(),
				})
				return
			}

			// 인증이 필요없는 요청이면 JWT 검증을 건너뜀
			if publicOperations[operationName] {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				httperror.SetErrorResponse(w, r.Context(), &httperror.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Authorization header missing",
				})
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			isRefresh := false

			if operationName == "refreshToken" {
				isRefresh = true
			}

			token, err := ValidateJwt(tokenString, isRefresh)

			if err != nil {
				httperror.SetErrorResponse(w, r.Context(), &httperror.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token",
				})
				return
			}

			//사용자 id 추출
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				httperror.SetErrorResponse(w, r.Context(), &httperror.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token claims",
				})
				return
			}

			userId, ok := claims["sub"].(string)
			if !ok {
				httperror.SetErrorResponse(w, r.Context(), &httperror.HTTPError{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token subject",
				})
				return
			}

			// 사용자 id를 context에 저장
			ctx := WithUserId(r.Context(), userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Context에서 userId 가져오기
			userId, err := GetUserId(r.Context())

			if err != nil || userId == -1 {
				next.ServeHTTP(w, r)
				return
			}

			// DB에서 user 정보 조회하여 Role 가져오기
			client := ent.FromContext(r.Context())

			user, err := client.User.Get(r.Context(), userId)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Authority 정보 Context에 저장
			ctx := WithUserAuthority(r.Context(), NewAuthority(user.Role))
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
