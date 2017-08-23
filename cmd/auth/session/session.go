package session

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/albert-widi/transaction_example/cmd/auth/repo"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/redis"
	"github.com/satori/go.uuid"
)

// User session struct
type User struct {
	Admin  bool  `json:"admin"`
	UserID int64 `json:"user_id"`
}

const prefix = "hfdevtest:sess_id:"

func SetSessionAndCookie(w *http.ResponseWriter, userObj User) error {
	log.Debug("SetSessionAndCookie")
	user, err := json.Marshal(userObj)
	if err != nil {
		return err
	}
	sessKey := generateSessionKey()
	log.Debugf("Save to redis, user: %s", string(user))
	redisStore, err := redis.Get(repo.SessionRedis)
	if err != nil {
		return err
	}
	if err = redisStore.Set(prefix+sessKey, string(user), 15000); err != nil {
		return err
	}
	setCookie(w, sessKey)
	return nil
}

func generateSessionKey() string {
	hasher := md5.New()
	rndString := fmt.Sprintf("%s%d", generateUUID(), time.Now().Unix())
	hasher.Write([]byte(rndString))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateUUID() string {
	out := uuid.NewV4().String()
	return strings.TrimSpace(strings.Replace(string(out), "\n", "", -1))
}

func GetUser(sessionKey string) (*User, error) {
	redisStore, err := redis.Get(repo.SessionRedis)
	if err != nil {
		return nil, err
	}
	resp, err := redisStore.Get(prefix + sessionKey)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal([]byte(resp), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
