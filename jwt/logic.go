package jwt

import (
	"context"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/jwt/model"
	userModel "github.com/devarsh/vrpl/user/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type TokenManager struct {
	gen      *JwtGen
	db       *JwtDB
	duration time.Duration
	issuer   string
}

//Signer methods could be HS256,HS384,HS512
func NewTokenManager(signer string, key string, duration time.Duration, issuer string, db *gorm.DB) *TokenManager {
	method := jwt.GetSigningMethod(signer)
	if db == nil {
		panic("Nil Db Instance passed")
	}
	if method == nil {
		panic("JWT Generator Algo method not specified")
	}
	if key == "" {
		panic("JWT Algo method Key not loaded")
	}
	if duration == time.Duration(0) {
		duration = time.Duration(time.Minute * 30)
	}
	if issuer == "" {
		issuer = "MY_APP"
	}
	return &TokenManager{
		gen:      &JwtGen{key: []byte(key), signingMethod: method},
		db:       &JwtDB{db: db},
		duration: duration,
		issuer:   issuer,
	}
}

func (tm *TokenManager) AddToken(user *userModel.User) (string, error) {
	uid := uuid.Must(uuid.NewV4())
	expireTime := time.Now().Add(tm.duration)
	payload := model.TokenPayload{
		user.ID,
		"",
		&jwt.StandardClaims{
			Id:        uid.String(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    tm.issuer,
		},
	}
	token, err := tm.gen.GenerateJwtToken(&payload)
	if err != nil {
		return "", err
	}
	persist := model.JwtPersist{
		UniqueID:  payload.Id,
		ExpiresAt: expireTime,
		UserId:    payload.UserId,
		Blacklist: 0,
		Token:     token,
	}
	ctx := context.Background()
	_, err = tm.db.Add(ctx, &persist)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (tm *TokenManager) VerifyToken(token string) (*model.TokenPayload, error) {
	payload, err := tm.gen.VerifyJwtToken(token)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	jwt, err := tm.db.GetByUUID(ctx, payload.Id)
	if err != nil {
		return nil, err
	}
	if jwt.Blacklist == 1 {
		return nil, customErr.JwtTokenBlackList()
	}
	return payload, nil
}

func (tm *TokenManager) InvalidateToken(token *model.TokenPayload) error {
	ctx := context.Background()
	jwt, err := tm.db.GetByUUID(ctx, token.Id)
	if err != nil {
		return err
	}
	if jwt.Blacklist == 0 {
		jwt.Blacklist = 1
		_, err = tm.db.Update(ctx, jwt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tm *TokenManager) InvalidateUser(userId uint) error {
	ctx := context.Background()
	_, err := tm.db.UpdateByUserID(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}
