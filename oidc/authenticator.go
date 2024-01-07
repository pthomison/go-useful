package oidc

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/davecgh/go-spew/spew"
	"github.com/pthomison/utilkit"
	"golang.org/x/oauth2"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

type NewAuthenticatorInput struct {
	BaseUrl      string
	ClientId     string
	CallbackUrl  string
	ClientSecret string
}

// New instantiates the *Authenticator.
func NewAuthenticator(input NewAuthenticatorInput) *Authenticator {
	provider, err := oidc.NewProvider(context.TODO(), input.BaseUrl)
	utilkit.Check(err)

	conf := oauth2.Config{
		ClientID:     input.ClientId,
		ClientSecret: input.ClientSecret,
		RedirectURL:  input.CallbackUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "at_hash"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	spew.Dump("Authenticator: ", token)

	rawIDToken, ok := token.Extra("id_token").(string)
	spew.Dump(rawIDToken)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func CreateAtHash(accessToken string) string {
	// var h hash.Hash
	h := sha256.New()

	h.Write([]byte(accessToken)) // hash documents that Write will never return an error
	sum := h.Sum(nil)[:h.Size()/2]
	actual := base64.RawURLEncoding.EncodeToString(sum)

	return actual
}
