package oidc

import (
	"context"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// func newValidator() (*validator.Validator, error) {
// 	u, err := url.Parse(viper.GetString("auth0-url"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	provider := jwks.NewCachingProvider(u, 5*time.Minute)

// 	jwtValidator, err := validator.New(
// 		provider.KeyFunc,
// 		validator.RS256,
// 		u.String(),
// 		[]string{viper.GetString("client-id")},
// 		validator.WithAllowedClockSkew(time.Minute),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return jwtValidator, nil
// }

// func ValidateToken(token string, init *InitData, w http.ResponseWriter, r *http.Request) (*validator.ValidatedClaims, bool) {

// 	v, err := newValidator()
// 	if ErrorHandler(w, err, "app.NewValidator Failure") {
// 		return nil, false
// 	}
// 	tokOut, err := v.ValidateToken(init.Context, token)
// 	if ErrorHandler(w, err, "JWT Validation Failure") {
// 		return nil, false
// 	}

// 	return tokOut.(*validator.ValidatedClaims), true
// }

type OidcClient struct {
	Url          string
	ClientId     string
	ClientSecret string
}

func ValidateToken(ctx context.Context, token string, oc *OidcClient) (*validator.ValidatedClaims, error) {

	u, err := url.Parse(oc.Url)
	if err != nil {
		return nil, err
	}
	provider := jwks.NewCachingProvider(u, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		oc.Url,
		[]string{oc.ClientId},
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	validateToken, err := jwtValidator.ValidateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return validateToken.(*validator.ValidatedClaims), nil
}
