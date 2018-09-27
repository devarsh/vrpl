package auth

import (
	"context"
	"fmt"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/jwt"
	jwtModel "github.com/devarsh/vrpl/jwt/model"
	"github.com/devarsh/vrpl/resp"
	"github.com/devarsh/vrpl/user"
	userModel "github.com/devarsh/vrpl/user/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"time"
)

type Key int

var TokenKey Key = 1

type AuthManager struct {
	tm *jwt.TokenManager
	um *user.UserManager
}

func NewAuthManager(signer string, key string, duration time.Duration, issuer string, db *gorm.DB) *AuthManager {
	tm := jwt.NewTokenManager(signer, key, duration, issuer, db)
	um := user.NewUserManager(db)
	return &AuthManager{
		tm: tm,
		um: um,
	}
}

func (am *AuthManager) Routers() chi.Router {
	r := chi.NewRouter()
	r.Post("/addUser", am.AddUser)
	r.Post("/login", am.Login)
	r.Post("/logout", am.Logout)
	return r
}

func (am *AuthManager) AddUser(w http.ResponseWriter, r *http.Request) {
	user := userModel.User{}
	if err := render.Bind(r, &user); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(user); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	userID, err := am.um.AddUser(&user)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: userID})
}

func (am *AuthManager) Login(w http.ResponseWriter, r *http.Request) {
	credentials := userModel.UserLogin{}
	if err := render.Bind(r, &credentials); err != nil {
		render.Render(w, r, &resp.Resp{Result: customErr.ErrInvalidRequestError(err)})
		return
	}
	if err := customErr.ValidateStruct(credentials); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	userInst, err := am.um.VerifyPassword(&credentials)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	token, err := am.tm.AddToken(userInst)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
		return
	}
	render.Render(w, r, &resp.Resp{Result: token})
}

func (am *AuthManager) Logout(w http.ResponseWriter, r *http.Request) {
	tokenInst, err := GetTokenFromRequest(r)
	if err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
	}
	if err = am.tm.InvalidateToken(tokenInst); err != nil {
		render.Render(w, r, &resp.Resp{Result: err})
	}
	render.Render(w, r, &resp.Resp{Status: true})
}

func (am *AuthManager) Authenticate(handler http.HandlerFunc) http.Handler {
	x := func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := am.ExtractJWTFromAuthHeader(r)
		if err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
		}
		token, err := am.tm.VerifyToken(tokenStr)
		if err != nil {
			render.Render(w, r, &resp.Resp{Result: err})
		}
		SetTokenOnRequest(r, token)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(x)
}

func (am *AuthManager) ExtractJWTFromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", customErr.ErrInvalidRequestError(fmt.Errorf("Invalid Request Autheorization Header missing"))
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", customErr.ErrInvalidRequestError(fmt.Errorf("Authorization Header format must be Bearer {token}"))
	}
	return authHeaderParts[1], nil
}

func SetTokenOnRequest(r *http.Request, token *jwtModel.TokenPayload) {
	*r = *r.WithContext(context.WithValue(r.Context(), TokenKey, token))
}

func GetTokenFromRequest(r *http.Request) (*jwtModel.TokenPayload, error) {
	val := r.Context().Value(TokenKey)
	if val != nil {
		token, ok := val.(*jwtModel.TokenPayload)
		if ok {
			return token, nil
		}
	}
	return nil, customErr.ErrInvalidRequestError(fmt.Errorf("Invalid Request Autheorization Header missing"))
}
