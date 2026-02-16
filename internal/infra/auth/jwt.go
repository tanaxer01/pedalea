package auth

import "github.com/golang-jwt/jwt/v5"

type Auth struct {
	SecretKey string
}

func NewAuth(secretKey string) *Auth {
	return &Auth{SecretKey: secretKey}
}

func (a *Auth) GenerateJwtToken(params map[string]any) (string, error) {
	claims := jwt.MapClaims(params)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *Auth) GetTokenClaim(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(a.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return map[string]any(claims), nil
}
