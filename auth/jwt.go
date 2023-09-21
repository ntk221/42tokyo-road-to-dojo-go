package auth

import (
	"context"
	_ "embed"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
}

type Store interface {
	Save(ctx context.Context, key string, userID int) error
	Load(ctx context.Context, key string) (int, error)
}

/*
func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in New JWTer: private key: %w", err)
	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in New JWTer: public key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	// j.Clocker = c
	return j, nil
}
*/
