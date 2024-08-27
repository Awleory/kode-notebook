package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type CtxValue int

const (
	ctxUserId CtxValue = iota
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"uri":    r.RequestURI,
		}).Info()
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login, pass, ok := r.BasicAuth()
		var err error
		if ok {
			if userId, err := h.usersService.VerifyUser(r.Context(), login, pass); err == nil {
				newCtx := context.WithValue(r.Context(), ctxUserId, userId)
				r = r.WithContext(newCtx)
				next.ServeHTTP(w, r)
				return
			}
		} else {
			err = errors.New("Unauthorized")
		}

		logError("authMiddleware", err)
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func getCtxUserId(ctx context.Context) (int, error) {
	if userId, ok := ctx.Value(ctxUserId).(int); ok {
		return userId, nil
	}

	return -1, errors.New("no user id")
}
