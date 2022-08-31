package proxy

import "testing"

func TestUserProxy(t *testing.T) {
	proxy := NewUserProxy(&User{})

	err := proxy.Login("xxx", "666")
	if err != nil {
		t.Fatal("n1 does not exist")
	}
}
