package session

import (
	"container/list"
	"sync"
	"time"
)

var pder = &Default{
	list: list.New(),
}

// Store - map of clients and their sessions
type Store struct {
	sid   string
	atime time.Time
	value map[interface{}]interface{}
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("default", pder)
}

// Set - register kv into manager
func (st *Store) Set(key, value interface{}) error {
	st.value[key] = value
	pder.Update(st.sid)
	return nil
}

// Get - lookup value by key
func (st *Store) Get(key interface{}) interface{} {
	pder.Update(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete - remove value pointed to by key
func (st *Store) Delete(key interface{}) error {
	delete(st.value, key)
	pder.Update(st.sid)
	return nil
}

// ID - return session ID
func (st *Store) ID() string {
	return st.sid
}

// Default - The main session used for clients on the site
type Default struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

// Init - create session
func (pder *Default) Init(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &Store{sid: sid, atime: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

// Read - Request session by id
func (pder *Default) Read(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*Store), nil
	}
	return pder.Init(sid)
}

// Destroy - Remove session by id
func (pder *Default) Destroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
	}
	return nil
}

// GC - Clean up all expired sessions
func (pder *Default) GC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil {
			return
		}
		if (element.Value.(*Store).atime.Unix() + maxlifetime) >= time.Now().Unix() {
			return
		}
		pder.list.Remove(element)
		delete(pder.sessions, element.Value.(*Store).sid)
	}
}

// Update - move session labelled with ID to top of list
func (pder *Default) Update(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*Store).atime = time.Now()
		pder.list.MoveToFront(element)
	}
	return nil
}
