package types

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	_ "github.com/joho/godotenv/autoload"
)

const (
	issuer string = "GPSTrackingServer"
)

var (
	jwtSecret string           = os.Getenv("JWT_SECRET")
	Audience  jwt.ClaimStrings = jwt.ClaimStrings{
		"api",
	}
)

type JwtData struct {
	Token  *jwt.Token
	Claims *JwtClaims
}

type JwtClaims struct {
	customClaims
	jwt.RegisteredClaims
}

type customClaims struct {
	SessionID uuid.UUID `json:"session_id"`
}

func CreateAccessToken(sessionId uuid.UUID) (string, error) {
	var accessExpiration time.Time = time.Now().UTC().Add(1 * time.Hour)
	return CreateJWT(sessionId, accessExpiration)
}

func CreateRefreshToken(sessionId uuid.UUID) (string, error) {
	var refreshExpiration time.Time = time.Now().UTC().Add(24 * time.Hour)
	return CreateJWT(sessionId, refreshExpiration)
}

func CreateJWT(sessionId uuid.UUID, expiration time.Time) (string, error) {
	tokenId, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	claims := JwtClaims{
		customClaims{
			SessionID: sessionId,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    issuer,
			Subject:   sessionId.String(),
			ID:        tokenId.String(),
			Audience:  Audience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
}

func ExtractJWTFromHeader(c *fiber.Ctx, expired func(string)) (*JwtData, error) {
	header := strings.Join(c.GetReqHeaders()["Authorization"], "")

	log.Debug(header)
	if !strings.HasPrefix(header, "Bearer") {
		return nil, errors.New("permission denied")
	}

	tokenString := strings.Split(header, " ")[1]
	token, err := ValidateJWT(tokenString)
	if err != nil {
		log.Debug(err.Error())
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		log.Debug("invalid claims | invalid token")
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		log.Debug("invalid audience")
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != issuer {
		log.Debug("invalid issuer")
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	return &JwtData{
		Token:  token,
		Claims: claims,
	}, nil
}
