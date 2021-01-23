package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/cpartogi/izyai/internal/menu/constant"

	"github.com/cpartogi/izyai/pkg/common/config"
	"github.com/cpartogi/izyai/pkg/common/state"
	"github.com/cpartogi/izyai/pkg/common/util"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/urfave/negroni"
)

type key int

type AuthData struct {
	UserID     int64
	UserTypeID int64
	Error      error
}

var (
	errMissingHeader      = "missing Authorization header"
	errInvalidTokenFormat = "invalid token format"
	errInvalidToken       = "invalid token"
)

const UserIDKey key = 1

// Auth add auth. middlewares before main handler
func Auth(handler negroni.HandlerFunc) http.Handler {
	n := negroni.New()
	n.Use(negroni.HandlerFunc(SetCtx))
	n.Use(negroni.HandlerFunc(AuthToken))
	n.Use(negroni.HandlerFunc(handler))
	return n
}

// Public only register handler as middleware without auth. middleware
func Public(handler negroni.HandlerFunc) http.Handler {
	n := negroni.New()
	n.Use(negroni.HandlerFunc(CheckToken))
	n.Use(negroni.HandlerFunc(handler))
	return n
}

// Admin add auth. middlewares before main handler
func Admin(handler negroni.HandlerFunc) http.Handler {
	n := negroni.New()
	n.Use(negroni.HandlerFunc(SetCtx))
	n.Use(negroni.HandlerFunc(AdminToken))
	n.Use(negroni.HandlerFunc(handler))
	return n
}

// SetCtx is middleware for setting context to incoming request
func SetCtx(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	cfg := config.Get()
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.API.Timeout)*time.Millisecond)
	defer cancel()
	r = r.WithContext(ctx)
	next(w, r)
}

// AuthToken is middleware to authenticate incoming request
func AuthToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()

	// do auth
	c := make(chan AuthData, 1)
	go func() {
		tokenString, err := GetToken(r)
		if err != nil {
			c <- AuthData{UserID: 0, UserTypeID: 0, Error: err}
			return
		}

		token, err := ParseToken(tokenString)
		if err != nil {
			c <- AuthData{UserID: 0, UserTypeID: 0, Error: err}
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := int64(claims["id"].(float64))
		userType := int64(claims["type"].(float64))

		c <- AuthData{UserID: userID, UserTypeID: userType, Error: nil}
	}()

	var v AuthData
	select {
	case <-ctx.Done():
		util.ResponseJSON(w, http.StatusGatewayTimeout, state.ErrServerTimeout, nil)
		return
	case v = <-c:
		if v.Error != nil {
			util.ResponseJSON(w, http.StatusUnauthorized, v.Error.Error(), nil)
			return
		}
	}
	ctx = context.WithValue(ctx, "type", v.UserTypeID)
	ctx = context.WithValue(ctx, "id", v.UserID)
	r = r.WithContext(ctx)

	// pass to main handler
	next(w, r)
}

func GetToken(r *http.Request) (string, error) {
	var token string

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return token, errors.New(errMissingHeader)
	}
	authHeaderArr := strings.Split(authHeader, " ")
	if len(authHeaderArr) < 2 || authHeaderArr[0] != "Bearer" {
		return token, errors.New(errInvalidTokenFormat)
	}

	return authHeaderArr[1], nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	cfg := config.Get()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			log.Println("Invalid JWT signing method")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(cfg.Auth.PrivateKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New(errInvalidToken)
	}

	return token, err
}

func CheckToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()

	tokenString, err := GetToken(r)
	if err == nil {
		token, err := ParseToken(tokenString)
		if err == nil {
			claims, _ := token.Claims.(jwt.MapClaims)
			userID := int64(claims["id"].(float64))
			userType := int64(claims["type"].(float64))

			v := AuthData{UserID: userID, UserTypeID: userType, Error: nil}

			ctx = context.WithValue(ctx, "type", v.UserTypeID)
			ctx = context.WithValue(ctx, "id", v.UserID)
			r = r.WithContext(ctx)
		}
	}

	next(w, r)
}

// AdminToken is middleware to authenticate incoming request
func AdminToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()

	// do auth
	c := make(chan AuthData, 1)
	go func() {
		tokenString, err := GetToken(r)
		if err != nil {
			c <- AuthData{UserID: 0, UserTypeID: 0, Error: err}
			return
		}

		token, err := ParseToken(tokenString)
		if err != nil {
			c <- AuthData{UserID: 0, UserTypeID: 0, Error: err}
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := int64(claims["id"].(float64))
		userType := int64(claims["type"].(float64))

		if userType != constant.AdminUserTypeID {
			c <- AuthData{UserID: 0, UserTypeID: 0, Error: errors.New("INVALID_USER_TYPE")}
			return
		}

		c <- AuthData{UserID: userID, UserTypeID: userType, Error: nil}
	}()

	var v AuthData
	select {
	case <-ctx.Done():
		util.ResponseJSON(w, http.StatusGatewayTimeout, state.ErrServerTimeout, nil)
		return
	case v = <-c:
		if v.Error != nil {
			util.ResponseJSON(w, http.StatusUnauthorized, v.Error.Error(), nil)
			return
		}
	}
	ctx = context.WithValue(ctx, "type", v.UserTypeID)
	ctx = context.WithValue(ctx, "id", v.UserID)
	r = r.WithContext(ctx)

	// pass to main handler
	next(w, r)
}
