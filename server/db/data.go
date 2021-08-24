package db

import (
	"encoding/json"
	"sync"
)

type Room struct {
	Rid   string `json:"room_id"`
	Uid   string `json:"admin_id"`
	Token string `json:"token"`
}

type RoomAttr struct {
	Uid   string
	Token string
}

type RoomTable struct {
	rwMutex  sync.RWMutex
	data     []Room
	roomsSet map[string]bool
	roomsMap map[string]RoomAttr
}

func NewRoomTable() *RoomTable {
	return &RoomTable{
		roomsSet: map[string]bool{},
		roomsMap: map[string]RoomAttr{},
	}
}

// String returns JSON format string.
func (r *RoomTable) String() string {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()
	if js, err := json.MarshalIndent(r.data, "", "    "); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Serialize returns JSON format without indent.
func (r *RoomTable) Serialize() string {
	if js, err := json.Marshal(r.data); err != nil {
		return err.Error()
	} else {
		return string(js)
	}
}

// Parse parses type Conf from type []byte in JSON format.
func (r *RoomTable) Parse(js []byte) error {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()
	if len(js) == 0 {
		return nil
	}
	return json.Unmarshal(js, &(r.data))
}

func (r *RoomTable) Insert(room Room) bool {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()
	if r.roomsSet[room.Rid] {
		return false
	}
	r.roomsSet[room.Rid] = true
	r.roomsMap[room.Rid] = RoomAttr{room.Uid, room.Token}
	r.data = append(r.data, room)
	return true
}
