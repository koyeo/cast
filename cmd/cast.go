package cmd

import (
	"github.com/koyeo/cast/cmd/initialize"
	"github.com/koyeo/cast/cmd/list"
	"github.com/koyeo/cast/cmd/run"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cast",
	Short: "Development helper CLI / 开发辅助命令行工具集",
	Long: `Cast - Development helper CLI for local build, server upload, remote exec & pipeline tasks.
Cast - 开发辅助命令行工具集，支持本地构建、服务器上传、远程命令执行、流水线任务等。`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	rootCmd.AddCommand(
		initialize.Cmd,
		run.Cmd,
		list.Cmd,
		//upload.Cmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
