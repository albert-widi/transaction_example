package session

import (
	"net/http"
	"time"

	"github.com/albert-widi/transaction_example/log"
)

const cookie = "_SID_TXNAPP_"

func GetCookie(r *http.Request) *http.Cookie {
	cookie, err := r.Cookie(cookie)
	if err != nil {
		log.Debugf("[authenticate] failed to fetch request cookie : %s", err.Error())
		return nil
	}
	log.Debugf("Cookie string : %s", cookie.String())
	return cookie
}

func setCookie(w *http.ResponseWriter, sessionKey string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: cookie, Value: sessionKey, Expires: expiration, Path: "/"}
	http.SetCookie(*w, &cookie)
}

func unsetCookie(w *http.ResponseWriter, sessionKey string) {
	expiration := time.Now()
	cookie := http.Cookie{Name: cookie, Value: sessionKey, Expires: expiration}
	http.SetCookie(*w, &cookie)
}
