package db

import (
	"encoding/json"
	"sync"
)

type Uid string
type Rid string
type Token string

type Room struct {
	Rid   `json:"room_id"`
	Uid   `json:"admin_id"`
	Token `json:"token"`
}

type RoomAttr struct {
	Uid   Uid
	Token Token
}

type RoomTable struct {
	rwMutex sync.RWMutex
	data    []Room
	roomSet map[Rid]bool
	roomMap map[Rid]RoomAttr
}

func NewRoomTable() *RoomTable {
	return &RoomTable{
		roomSet: map[Rid]bool{},
		roomMap: map[Rid]RoomAttr{},
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
	if r.roomSet[room.Rid] {
		return false
	}
	r.roomSet[room.Rid] = true
	r.roomMap[room.Rid] = RoomAttr{room.Uid, room.Token}
	r.data = append(r.data, room)
	return true
}
