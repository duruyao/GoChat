package conf

import "testing"

func TestInit(t *testing.T) {
	t.Log(Addr())
	t.Log(MaxUsers())
	t.Log(MaxRooms())
	t.Log(MaxUsersPreRoom())
	t.Log(MaxRoomsPerAdmin())
}
