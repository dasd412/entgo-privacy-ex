package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

var requestData struct {
	OperationName string `json:"operationName"`
}

// 인증이 필요 없는 API 목록
var publicOperations = map[string]bool{
	"signup":             true,
	"login":              true,
	"IntrospectionQuery": true,
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
				return
			}

			// 인증이 필요없는 요청이면 JWT 검증을 건너뜀
			if publicOperations[operationName] {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(
					w,
					"Authorization header missing",
					http.StatusUnauthorized,
				)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			isRefresh := false
			if operationName == "refreshToken" {
				isRefresh = true
			}

			token, err := ValidateJwt(tokenString, isRefresh)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			//사용자 id 추출
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			userId, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "Invalid token subject", http.StatusUnauthorized)
				return
			}

			// 사용자 id를 context에 저장
			ctx := context.WithValue(r.Context(), UserIDKey, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func GetUserId(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return "", errors.New("user ID not found in context")
	}
	return userId, nil
}
