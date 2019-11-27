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

// Manager - Represents a particular cookie to a client
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

// NewManager - returns a Cookie manager, ready to use or an error
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

// Provider - interface our cookie manager calls to manage lifetimes
type Provider interface {
	Init(sid string) (Session, error)
	Read(sid string) (Session, error)
	Destroy(sid string) error
	GC(maxLifeTime int64)
}

// Session - a client connection
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	ID() string
}

// Register - Add a provider to our stack (e.g., Default)
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

// Destroy - Call provider Destroy
func (manager *Manager) Destroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.Destroy(cookie.Value)
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

// Start - Call provider Start
func (manager *Manager) Start(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ = manager.provider.Init(sid)
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
	session, _ = manager.provider.Read(sid)
	return
}

// GC - call all providers GC
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.GC(manager.maxlifetime)
	time.AfterFunc(10*time.Second, func() {
		manager.GC()
	})
}

func (manager *Manager) sessionID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Unable to generate UUID %q", err)
	}
	return u.String()
}
