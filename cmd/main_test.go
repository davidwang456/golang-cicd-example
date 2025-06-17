package main

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-example",
		Short: "Go Example 是一个示例命令行应用",
		Long:  `Go Example 是一个使用 Cobra 库构建的示例命令行应用。\n它展示了如何创建和使用命令行工具。`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("欢迎使用 Go Example!")
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "显示应用版本信息",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Go Example v1.0.0")
		},
	}

	greetCmd := &cobra.Command{
		Use:   "greet [name]",
		Short: "向指定的人打招呼",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			cmd.Printf("你好, %s!\n", name)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(greetCmd)
	return rootCmd
}

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err := root.Execute()
	return buf.String(), err
}

func TestRootCommand(t *testing.T) {
	output, err := executeCommand(newRootCmd())
	if err != nil {
		t.Errorf("执行根命令时发生错误: %v", err)
	}
	if output != "欢迎使用 Go Example!\n" {
		t.Errorf("根命令输出不符合预期: %s", output)
	}
}

func TestVersionCommand(t *testing.T) {
	output, err := executeCommand(newRootCmd(), "version")
	if err != nil {
		t.Errorf("执行版本命令时发生错误: %v", err)
	}
	if output != "Go Example v1.0.0\n" {
		t.Errorf("版本命令输出不符合预期: %s", output)
	}
}

func TestGreetCommand(t *testing.T) {
	output, err := executeCommand(newRootCmd(), "greet", "张三")
	if err != nil {
		t.Errorf("执行问候命令时发生错误: %v", err)
	}
	if output != "你好, 张三!\n" {
		t.Errorf("问候命令输出不符合预期: %s", output)
	}
}

func TestGreetCommandWithoutArgs(t *testing.T) {
	_, err := executeCommand(newRootCmd(), "greet")
	if err == nil {
		t.Error("期望在没有参数时返回错误，但没有")
	}
}
