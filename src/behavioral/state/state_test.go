package state

import (
	"testing"
)

func TestIterator(t *testing.T) {
	ctx := Context{}
	ctx.SetState(Registration{})
	ctx.Do()
	ctx.SetState(Authentication{})
	ctx.Do()
	ctx.SetState(Authorization{})
	ctx.Do()
	ctx.SetState(Logout{})
	ctx.Do()
}
