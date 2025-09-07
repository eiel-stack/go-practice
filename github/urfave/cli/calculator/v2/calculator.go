package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "calculator",
		Usage: "一个简单的命令行计算器，支持+、-、*、/、%运算",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"+"},
				Usage:   "计算两个数的和",
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:     "a",
						Usage:    "第一个操作数",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "b",
						Usage:    "第二个操作数",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					result := c.Float64("a") + c.Float64("b")
					fmt.Printf("%.2f + %.2f = %.2f\n", c.Float64("a"), c.Float64("b"), result)
					return nil
				},
			},
			{
				Name:    "subtract",
				Aliases: []string{"-"},
				Usage:   "计算两个数的差",
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:     "a",
						Usage:    "被减数",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "b",
						Usage:    "减数",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					result := c.Float64("a") - c.Float64("b")
					fmt.Printf("%.2f - %.2f = %.2f\n", c.Float64("a"), c.Float64("b"), result)
					return nil
				},
			},
			{
				Name:    "multiply",
				Aliases: []string{"*"},
				Usage:   "计算两个数的积",
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:     "a",
						Usage:    "第一个乘数",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "b",
						Usage:    "第二个乘数",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					result := c.Float64("a") * c.Float64("b")
					fmt.Printf("%.2f * %.2f = %.2f\n", c.Float64("a"), c.Float64("b"), result)
					return nil
				},
			},
			{
				Name:    "divide",
				Aliases: []string{"/"},
				Usage:   "计算两个数的商",
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:     "a",
						Usage:    "被除数",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "b",
						Usage:    "除数",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					b := c.Float64("b")
					if b == 0 {
						return fmt.Errorf("错误：除数不能为0")
					}
					result := c.Float64("a") / b
					fmt.Printf("%.2f / %.2f = %.2f\n", c.Float64("a"), b, result)
					return nil
				},
			},
			{
				Name:    "mod",
				Aliases: []string{"%"},
				Usage:   "计算两个整数的余数",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "a",
						Usage:    "被除数（整数）",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "b",
						Usage:    "除数（整数）",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					b := c.Int("b")
					if b == 0 {
						return fmt.Errorf("错误：除数不能为0")
					}
					result := c.Int("a") % b
					fmt.Printf("%d %% %d = %d\n", c.Int("a"), b, result)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}
