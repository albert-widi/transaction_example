package middleware

import (
	"context"
	"net/http"

	"github.com/albert-widi/transaction_example/apicalls"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/route"
)

// User session struct
type User struct {
	Admin  bool  `json:"admin"`
	UserID int64 `json:"user_id"`
}

// AuthenticateV1 for V1 authentication
func Authenticate(h route.Handle) route.Handle {
	return func(r *http.Request) (route.HandleObject, error) {
		v1reponse := new(route.V1)

		log.Debugf("Authenticating %s", r.URL.String())
		user := User{}
		cookie, err := r.Cookie("_SID_TXNAPP_")
		if err != nil {
			return v1reponse, err
		}
		err = apicalls.Auth.Authenticate(cookie, &user)
		if err != nil {
			return v1reponse, err
		}

		if user.UserID == 0 {
			return v1reponse, errors.New("User is not authenticated", http.StatusForbidden)
		}
		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)
		return h(r)
	}
}
