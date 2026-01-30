package bark

import (
	"context"
	"flare/internal/model"
	"flare/internal/service"
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	service.RegisterBark(New())
}

func New() *sBark {
	return &sBark{}
}

type sBark struct{}

func (s *sBark) Push(ctx context.Context, in *model.PushInput) error {
	val, err := g.Cfg().Get(ctx, "bark.server")
	if err != nil {
		return err
	}
	dk, err := g.Cfg().Get(ctx, "bark.device_key")
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s", val.String(), dk.String())

	res, err := g.Client().ContentJson().Post(ctx, url, in)
	if err != nil {
		return err
	}
	defer res.Close()

	g.Log().Debugf(ctx, "dump push: %s", res.Raw())

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("push err")
	}
	return nil
}
