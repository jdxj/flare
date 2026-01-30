package bark

import (
	"flare/internal/consts"
	"flare/internal/model"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx = gctx.GetInitCtx()
	s   = New()
)

func TestPush(t *testing.T) {
	err := s.Push(ctx, &model.PushInput{
		Title:    "test flare",
		Subtitle: "",
		Body:     "fff",
		Level:    consts.BarkLevelActive,
	})
	if err != nil {
		t.Fatal(err)
	}
}
