// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"flare/internal/model"
)

type (
	IBark interface {
		Push(ctx context.Context, in *model.PushInput) error
	}
)

var (
	localBark IBark
)

func Bark() IBark {
	if localBark == nil {
		panic("implement not found for interface IBark, forgot register?")
	}
	return localBark
}

func RegisterBark(i IBark) {
	localBark = i
}
