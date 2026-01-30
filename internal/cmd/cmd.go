package cmd

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "flare",
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
		Run:                        rootRun,
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

	// sub cmd
	cmd.AddCommand(newVersion())

	return cmd
}

func rootRun(cmd *cobra.Command, args []string) {}

func newVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "version",
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
		Run:                        versionRun,
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

func versionRun(cmd *cobra.Command, args []string) {
	buildInfo := gbuild.Info()
	fmt.Printf("Version   : %s\n", buildInfo.Version)
	fmt.Printf("GoFrame   : %s\n", buildInfo.GoFrame)
	fmt.Printf("Go        : %s\n", buildInfo.Golang)
	fmt.Printf("Build time: %s\n", buildInfo.Time)
	fmt.Printf("Commit    : %s\n", buildInfo.Git)
}
