package authutil

import (
	"fmt"
	userpb "github.com/Gictorbit/gotodotasks/api/gen/proto/user/v1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type TokenUserClaims struct {
	jwt.RegisteredClaims
	UserID    uint32            `json:"id"`
	UserAgent string            `json:"ua"`
	IP        string            `json:"ip"`
	Email     string            `json:"em"`
	Roles     []userpb.UserRole `json:"rl"`
}

type JWTManager struct {
	secretKey      string
	tokenValidTime time.Duration
	issuer         string
	signingMethod  jwt.SigningMethod
}

func NewJWTManager(secretKey, issuer string, validTime time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:      secretKey,
		tokenValidTime: validTime,
		issuer:         issuer,
		signingMethod:  jwt.SigningMethodHS256,
	}
}

func (jm *JWTManager) NewToken(userInfo *userpb.User) (string, error) {
	userClaims := TokenUserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jm.issuer,
			Subject:   userInfo.Name,
			ID:        uuid.New().String(),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(jm.tokenValidTime)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
		Email:     userInfo.Email,
		UserID:    userInfo.Id,
		IP:        userInfo.Ip,
		UserAgent: userInfo.Agent,
		Roles:     userInfo.Roles,
	}
	return jm.generateToken(userClaims)
}

func (jm *JWTManager) generateToken(claims TokenUserClaims) (string, error) {
	unsignedToken := jwt.NewWithClaims(jm.signingMethod, claims)
	token, err := unsignedToken.SignedString([]byte(jm.secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (jm *JWTManager) VerifyToken(accessToken string) (*TokenUserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return []byte(jm.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*TokenUserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
