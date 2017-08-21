package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/albert-widi/transaction_example/cmd/auth/session"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/route"
	"github.com/pressly/chi"
)

// API struct
type APIV1 struct{}

// New api
func New() *APIV1 { return new(APIV1) }

// Register new api
func (api *APIV1) Register(r chi.Router) {
	log.Debug("Registering api v1 endpoints...")
	w := route.NewWrapper(r, route.Options{
		Timeout: route.Timeout{
			Timeout:  1,
			Response: map[string]string{"halo": "hola"},
		},
	})
	w.Get("/ping", w.Handle(ping))
	w.Get("/simplelogin", simpleLogin())
	w.Get("/auth", w.Handle(authenticate))
}

func ping(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	resp.Data = map[string]string{
		"data": "This is data",
	}
	return resp, errors.DatabaseTypeNotExists.Err()
}

// SimpleAuth for saving session to redis
func simpleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAdmin, err := strconv.ParseBool(r.FormValue("admin"))
		if err != nil {
			log.Debug("[ADMIN ERROR] ", err.Error())
			isAdmin = false
		}
		log.Debug("Admin: ", isAdmin)

		unixTime := time.Now().Unix()
		user := session.User{
			Admin:  isAdmin,
			UserID: unixTime,
		}
		err = session.SetSessionAndCookie(&w, user)
		if err != nil {
			log.Error("Failed to set session and cookie. Error: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to set user session"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success set cookie"))
	}
}

func authenticate(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	cookie := session.GetCookie(r)
	if cookie == nil {
		log.Debug("No cookie, user have no access")
		return nil, errors.New("No Cookie, no access", http.StatusForbidden)
	}
	sessionKey := (*cookie).Value

	user, err := session.GetUser(sessionKey)
	if err != nil {
		return nil, err
	}

	if user.UserID == 0 {
		log.Debug("UserID is nil")
		return nil, errors.New("UserID nil", http.StatusForbidden)
	}
	resp.Data = user
	return resp, nil
}
