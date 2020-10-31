package example

import "testing"

func TestUserJsonExample(t *testing.T) {
	t.Log(UserJsonExample())
}

func TestUserGobExample(t *testing.T) {
	t.Log(UserGobExample())
}
