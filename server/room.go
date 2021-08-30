package server

import "sync"

type Uid string
type Rid string
type Token string

type Room struct {
	rwMu  sync.RWMutex
	rid   Rid
	admin Uid
	token Token
	users map[Uid]bool
}

type Rooms struct {
	rwMu  sync.RWMutex
	rooms map[Rid]*Room
}
