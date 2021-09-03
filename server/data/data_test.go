package data

import (
	"testing"
)

func TestUser_Create(t *testing.T) {
	u1 := User{
		Name:     "root",
		MaxRole:  Owner,
		Password: "12345678",
	}
	if err := u1.Create(); err != nil {
		t.Fatal(err)
	}

	t.Log(u1)
	u2 := User{
		Name:     "admin1",
		MaxRole:  Admin,
		Password: "6469659165901",
	}
	if err := u2.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u2)

	u3 := User{
		Name:     "admin2",
		MaxRole:  Admin,
		Password: "47165916596",
	}
	if err := u3.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u3)

	u4 := User{
		Name:    "guest1",
		MaxRole: Guest,
	}
	if err := u4.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u4)

	u5 := User{
		Name:    "guest2",
		MaxRole: Guest,
	}
	if err := u5.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u5)

	u6 := User{
		Name:    "guest3",
		MaxRole: Guest,
	}
	if err := u6.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u6)

	u7 := User{
		Name:    "guest4",
		MaxRole: Guest,
	}
	if err := u7.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u7)

	us, err := AllUsers()
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range us {
		t.Log(u)
	}
}

func TestAllUsers(t *testing.T) {
	us, err := AllUsers()
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range us {
		t.Log(u)
	}
}

func TestUserByUUId(t *testing.T) {
	us, err := AllUsers()
	if err != nil {
		t.Fatal(err)
	}
	u, err := UserByUUId(us[1].UUId) // admin1
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUser_CreateSession(t *testing.T) {
	u := User{Id: 2} // admin1

	s1, err := u.CreateSession()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s1)

	s2, err := u.Session()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s2)
}

func TestAllSessions(t *testing.T) {
	ss, err := AllSessions()
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range ss {
		t.Log(s)
	}
}

func TestSessionByUUid(t *testing.T) {
	u := User{Id: 2} // admin1

	s1, err := u.Session() // session of admin1
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s1)

	s2, err := SessionByUUid(s1.UUId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s2)
}

func TestUser_CreateRoom(t *testing.T) {
	u := User{Id: 2} // admin1

	r1, err := u.CreateRoom("Chat_Room_1", "210993")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r1)

	r2, err := u.CreateRoom("Chat_Room_2", "5671295")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r2)

	rs, err := u.CreatedRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}

	rs, err = AllRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}
}

func TestAllRooms(t *testing.T) {
	rs, err := AllRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}
}

func TestRoomByUUid(t *testing.T) {
	u := User{Id: 2} // admin1

	rs, err := u.CreatedRooms()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rs[0])

	r, err := RoomByUUid(rs[0].UUId) // Chat_room_1
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestUser_JoinRoom(t *testing.T) {
	r1 := Room{Id: 1} // Chat_room_1

	r2 := Room{Id: 2} // Chat_Room_2

	u := User{Id: 7} // guest4

	if err := u.JoinRoom(r1); err != nil {
		t.Fatal(err)
	}
	if err := u.JoinRoom(r2); err != nil {
		t.Fatal(err)
	}

	rs, err := u.JoinedRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}

	rs, err = AllRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}
}

func TestUser_LeaveRoom(t *testing.T) {
	r1 := Room{Id: 1} // Chat_Room_1

	u := User{Id: 7} // guest4

	if err := u.LeaveRoom(r1); err != nil {
		t.Fatal(err)
	}

	rs, err := u.JoinedRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}

	rs, err = AllRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}
}

func TestUser_Delete(t *testing.T) {
	u := User{Id: 7} // guest4
	if err := u.Delete(); err != nil {
		t.Fatal(err)
	}

	rs, err := AllRooms()
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range rs {
		t.Log(r)
	}
}
