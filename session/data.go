package session

import (
	"container/list"
	"sync"
	"time"
)

var pder = &Default{
	list: list.New(),
}

type SessionStore struct {
	sid   string
	atime time.Time
	value map[interface{}]interface{}
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("default", pder)
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

type Default struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

func (pder *Default) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, atime: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

func (pder *Default) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	}
	return pder.SessionInit(sid)
}

func (pder *Default) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
	}
	return nil
}

func (pder *Default) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).atime.Unix() + maxlifetime) >= time.Now().Unix() {
			break
		}
		pder.list.Remove(element)
		delete(pder.sessions, element.Value.(*SessionStore).sid)
	}
}

func (pder *Default) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).atime = time.Now()
		pder.list.MoveToFront(element)
	}
	return nil
}
