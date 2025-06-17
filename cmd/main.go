package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-example",
	Short: "Go Example 是一个示例命令行应用",
	Long: `Go Example 是一个使用 Cobra 库构建的示例命令行应用。
它展示了如何创建和使用命令行工具。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("欢迎使用 Go Example!")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示应用版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Example v1.0.0")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "向指定的人打招呼",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("你好, %s!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(greetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
