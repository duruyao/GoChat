package conf

import "testing"

func TestInit(t *testing.T) {
	t.Log(Addr())
	t.Log(MaxUsers())
	t.Log(MaxGroups())
	t.Log(MaxUsersPreGroup())
	t.Log(MaxGroupsPerAdmin())
}
