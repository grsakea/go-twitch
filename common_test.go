package twitch

import (
	"os"
	"testing"
)

func TestNewSession(t *testing.T) {
	oldEnv := os.Getenv("CLIENT_ID")
	os.Setenv("CLIENT_ID", "foobar")
	defer os.Setenv("CLIENT_ID", oldEnv)

	s := NewSession()
	if s.ClientID != "foobar" {
		t.Fatal("Failed creating session")
	}

}

func TestImplementInterface(t *testing.T) {
	var twitch Interface = (*Session)(nil)
	t.Log(twitch)
}
