package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type JwtTokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// JWT 비밀키 설정 (보안상 환경변수로 저장하는 것이 좋음)
var secretKey = []byte("your_secret_key")
var refreshSecretKey = []byte("your_refresh_secret_key")

func GenerateTokenPair(userId int) (*JwtTokenPair, error) {
	accessClaims := jwt.MapClaims{
		"sub": strconv.Itoa(userId),                 // 사용자 Id(subject)
		"exp": time.Now().Add(time.Hour * 1).Unix(), // 만료시간 (1시간)
		"iat": time.Now().Unix(),                    // 발급 시간 (IssuedAt)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessTokenString, err := accessToken.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"sub": strconv.Itoa(userId),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)

	if err != nil {
		return nil, err
	}

	return &JwtTokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
