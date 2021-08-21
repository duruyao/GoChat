package db

import (
	"encoding/json"
	"sync"
)

type Data struct {
	Room  string `json:"room"`
	Uid   string `json:"admin_uid"`
	Token string `json:"token"`
}

type roomAttribute struct {
	Uid   string
	Token string
}

type table struct {
	rwMutex  sync.RWMutex
	data     []Data
	roomsSet map[string]bool
	roomsMap map[string]roomAttribute
}

func NewTable() *table {
	return &table{
		roomsSet: map[string]bool{},
		roomsMap: map[string]roomAttribute{},
	}
}

// String returns type string in JSON format.
func (t *table) String() string {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if js, err := json.MarshalIndent(t.data, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// String returns type []byte in JSON format.
func (t *table) Bytes() []byte {
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if js, err := json.MarshalIndent(t.data, "", "    "); err != nil {
		return []byte(err.Error())
	} else {
		return js
	}
}

// Parse parses type Conf from type []byte in JSON format.
func (t *table) Parse(js []byte) error {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, &(t.data))
}

func (t *table) Insert(data Data) bool {
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	if t.roomsSet[data.Room] {
		return false
	}
	t.roomsSet[data.Room] = true
	t.roomsMap[data.Room] = roomAttribute{data.Uid, data.Token}
	t.data = append(t.data, data)
	return true
}
