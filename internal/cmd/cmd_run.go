package cmd

import (
	"context"
	"flare/internal/consts"
	"flare/internal/model"
	"flare/internal/service"
	"fmt"
	"strings"
	"time"

	"github.com/coreos/go-systemd/v22/sdjournal"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/cobra"
)

func newRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "run",
		Aliases:                    []string{},
		SuggestFor:                 []string{},
		Short:                      "",
		GroupID:                    "",
		Long:                       "",
		Example:                    "",
		ValidArgs:                  []cobra.Completion{},
		ValidArgsFunction:          nil,
		Args:                       nil,
		ArgAliases:                 []string{},
		BashCompletionFunction:     "",
		Deprecated:                 "",
		Annotations:                map[string]string{},
		Version:                    "",
		Run:                        run,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	checkSSHLogin(cmd.Context())
}

func checkSSHLogin(ctx context.Context) error {
	r, err := sdjournal.NewJournal()
	if err != nil {
		return err
	}
	defer r.Close()

	err = r.AddMatch("SYSLOG_IDENTIFIER=sshd-session")
	if err != nil {
		return err
	}

	_ = r.SeekTail()
	_, _ = r.Previous()

	fmt.Println("start")

	for {
		n, err := r.Next()
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}
		if n == 0 {
			_ = r.Wait(time.Second)
			continue
		}

		entry, err := r.GetEntry()
		if err != nil {
			g.Log().Error(ctx, err)
			continue
		}

		msg := entry.Fields["MESSAGE"]
		host := entry.Fields["_HOSTNAME"]
		if !strings.Contains(msg, "Accepted publickey") {
			continue
		}

		err = service.Bark().Push(ctx, &model.PushInput{
			Title: fmt.Sprintf("%s 上有SSH登录", host),
			Body:  msg,
			Level: consts.BarkLevelActive,
		})
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}
