package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "calc",
		Usage: "A simple command line calculator",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"add", "+"},
				Usage:   "Add two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return fmt.Errorf("requires exactly 2 arguments")
					}
					x, y, err := parseArgs(cmd.Args().Get(0), cmd.Args().Get(1))
					if err != nil {
						return err
					}
					fmt.Printf("%.2f + %.2f = %.2f\n", x, y, x+y)
					return nil
				},
			},
			{
				Name:    "subtract",
				Aliases: []string{"sub"},
				Usage:   "Subtract two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return fmt.Errorf("requires exactly 2 arguments")
					}
					x, y, err := parseArgs(cmd.Args().Get(0), cmd.Args().Get(1))
					if err != nil {
						return err
					}
					fmt.Printf("%.2f - %.2f = %.2f\n", x, y, x-y)
					return nil
				},
			},
			{
				Name:    "multiply",
				Aliases: []string{"mul"},
				Usage:   "Multiply two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return fmt.Errorf("requires exactly 2 arguments")
					}
					x, y, err := parseArgs(cmd.Args().Get(0), cmd.Args().Get(1))
					if err != nil {
						return err
					}
					fmt.Printf("%.2f * %.2f = %.2f\n", x, y, x*y)
					return nil
				},
			},
			{
				Name:    "divide",
				Aliases: []string{"div", "/"},
				Usage:   "Divide two numbers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return fmt.Errorf("requires exactly 2 arguments")
					}
					x, y, err := parseArgs(cmd.Args().Get(0), cmd.Args().Get(1))
					if err != nil {
						return err
					}
					if y == 0 {
						return fmt.Errorf("cannot divide by zero")
					}
					fmt.Printf("%.2f / %.2f = %.2f\n", x, y, x/y)
					return nil
				},
			},
			{
				Name:    "modulo",
				Aliases: []string{"mod", "%"},
				Usage:   "Modulo operation on two integers",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return fmt.Errorf("requires exactly 2 arguments")
					}
					x, err := strconv.Atoi(cmd.Args().Get(0))
					if err != nil {
						return fmt.Errorf("first argument must be an integer")
					}
					y, err := strconv.Atoi(cmd.Args().Get(1))
					if err != nil {
						return fmt.Errorf("second argument must be an integer")
					}
					if y == 0 {
						return fmt.Errorf("cannot modulo by zero")
					}
					fmt.Printf("%d %% %d = %d\n", x, y, x%y)
					return nil
				},
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseArgs(a, b string) (float64, float64, error) {
	x, err := strconv.ParseFloat(strings.TrimSpace(a), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("first argument must be a number")
	}
	y, err := strconv.ParseFloat(strings.TrimSpace(b), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("second argument must be a number")
	}
	return x, y, nil
}
