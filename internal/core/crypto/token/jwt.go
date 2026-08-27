package crypto_token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}
type jwtClaims struct {
	Role      string    `json:"role"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

func NewJWT(config Config) *JWT {
	return &JWT{
		secret:     []byte(config.Secret),
		accessTTL:  config.AccessTTL,
		refreshTTL: config.RefreshTTL,
	}
}
func (j *JWT) ttl(tokenType TokenType) (time.Duration, error) {
	switch tokenType {
	case TokenTypeAccess:
		return j.accessTTL, nil
	case TokenTypeRefresh:
		return j.refreshTTL, nil
	default:
		return 0, fmt.Errorf("unsupported token type: %q", tokenType)
	}
}

func (j *JWT) Generate(userID uuid.UUID, role string, tokenType TokenType) (string, error) {
	now := time.Now().UTC()

	ttl, err := j.ttl(tokenType)
	if err != nil {
		return "", err
	}

	if userID == uuid.Nil {
		return "", fmt.Errorf("user id is empty")
	}
	if role == "" {
		return "", fmt.Errorf("role is empty")
	}
	claims := jwtClaims{
		Role:      role,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(j.secret)
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signedToken, nil
}
func (j *JWT) Validate(rawToken string) (Claims, error) {
	var claims jwtClaims

	token, err := jwt.ParseWithClaims(
		rawToken,
		&claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf(
					"unexpected signing method: %s",
					token.Method.Alg(),
				)
			}

			return j.secret, nil
		},
	)

	if err != nil {
		return Claims{}, fmt.Errorf("parse token: %w", err)
	}

	if !token.Valid {
		return Claims{}, fmt.Errorf("invalid token")
	}
	if !claims.TokenType.IsValid() {
		return Claims{}, fmt.Errorf("invalid token type: %q", claims.TokenType)
	}
	if claims.Subject == "" {
		return Claims{}, fmt.Errorf("user id is empty")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return Claims{}, fmt.Errorf("parse user id: %w", err)
	}

	if claims.Role == "" {
		return Claims{}, fmt.Errorf("role is empty")
	}

	return Claims{
		UserID:    userID,
		Role:      claims.Role,
		TokenType: claims.TokenType,
	}, nil
}
