package session

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
)

var provides = make(map[string]Provider)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	if provider, ok := provides[provideName]; ok {
		m := &Manager{
			provider:    provider,
			cookieName:  cookieName,
			maxlifetime: maxlifetime,
		}
		return m, nil
	}
	return nil, fmt.Errorf("session: unknown provide %q", provideName)
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

func Register(name string, provider Provider) {
	if provider == nil {
		log.Fatal("session: Register provider is nil")
	}
	if _, ok := provides[name]; !ok {
		provides[name] = provider
		return
	}
	log.Fatal("session: Register called twice for provider " + name)
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionDestroy(cookie.Value)
	expiration := time.Now()
	c := http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiration,
		MaxAge:   -1,
	}
	http.SetCookie(w, &c)
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxlifetime),
		}
		http.SetCookie(w, &cookie)
		return
	}
	sid, _ := url.QueryUnescape(cookie.Value)
	session, _ = manager.provider.SessionRead(sid)
	return
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime), func() { manager.GC() })
}

func (manager *Manager) sessionId() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Unable to generate UUID %q", err)
	}
	return u.String()
}
