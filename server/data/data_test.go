package data

import (
	"math"
	"testing"
)

func TestUser_Create(t *testing.T) {
	u1 := User{
		Name:     "root",
		MaxRole:  RoleTypeOwner,
		Password: "12345678",
	}
	if err := u1.Create(); err != nil {
		t.Fatal(err)
	}

	t.Log(u1)
	u2 := User{
		Name:     "admin1",
		MaxRole:  RoleTypeAdmin,
		Password: "6469659165901",
	}
	if err := u2.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u2)

	u3 := User{
		Name:     "admin2",
		MaxRole:  RoleTypeAdmin,
		Password: "47165916596",
	}
	if err := u3.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u3)

	u4 := User{
		Name:    "guest1",
		MaxRole: RoleTypeGuest,
	}
	if err := u4.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u4)

	u5 := User{
		Name:    "guest2",
		MaxRole: RoleTypeGuest,
	}
	if err := u5.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u5)

	u6 := User{
		Name:    "guest3",
		MaxRole: RoleTypeGuest,
	}
	if err := u6.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u6)

	u7 := User{
		Name:    "guest4",
		MaxRole: RoleTypeGuest,
	}
	if err := u7.Create(); err != nil {
		t.Fatal(err)
	}
	t.Log(u7)

	us, err := Users(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range us {
		t.Log(u)
	}
}

func TestUsers(t *testing.T) {
	us, err := Users(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range us {
		t.Log(u)
	}
}

func TestUserByUniqueKey(t *testing.T) {
	us, err := Users(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	u, err := UserByUniqueKey("UUID", us[1].UUId) // admin1
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

func TestSessions(t *testing.T) {
	ss, err := Sessions(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range ss {
		t.Log(s)
	}
}

func TestSessionByUniqueKey(t *testing.T) {
	u := User{Id: 2} // admin1

	s1, err := u.Session() // session of admin1
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s1)

	s2, err := SessionByUniqueKey("UUID", s1.UUId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s2)
}

func TestUser_CreateGroup(t *testing.T) {
	u := User{Id: 2} // admin1

	g1, err := u.CreateGroup("Chat_Room_1", "210993")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g1)

	g2, err := u.CreateGroup("Chat_Room_2", "5671295")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g2)

	gs, err := u.CreatedGroups()
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}

	gs, err = Groups(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}
}

func TestGroups(t *testing.T) {
	gs, err := Groups(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}
}

func TestGroupByUniqueKey(t *testing.T) {
	u := User{Id: 2} // admin1

	gs, err := u.CreatedGroups()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(gs[0])

	g, err := GroupByUniqueKey("UUID", gs[0].UUId) // Chat_room_1
	if err != nil {
		t.Fatal(err)
	}
	t.Log(g)
}

func TestUser_JoinGroup(t *testing.T) {
	g1 := Group{Id: 1} // Chat_room_1

	g2 := Group{Id: 2} // Chat_Room_2

	u := User{Id: 7} // guest4

	if err := u.JoinGroup(g1); err != nil {
		t.Fatal(err)
	}
	if err := u.JoinGroup(g2); err != nil {
		t.Fatal(err)
	}

	gs, err := u.JoinedGroups()
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}

	gs, err = Groups(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}
}

func TestUser_LeaveGroup(t *testing.T) {
	g1 := Group{Id: 1} // Chat_Room_1

	u := User{Id: 7} // guest4

	if err := u.LeaveGroup(g1); err != nil {
		t.Fatal(err)
	}

	gs, err := u.JoinedGroups()
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}

	gs, err = Groups(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}
}

func TestUser_Delete(t *testing.T) {
	u := User{Id: 7} // guest4
	if err := u.Delete(); err != nil {
		t.Fatal(err)
	}

	gs, err := Groups(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range gs {
		t.Log(g)
	}
}

func TestGroup_HasMember(t *testing.T) {

}
