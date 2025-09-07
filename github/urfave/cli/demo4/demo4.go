package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

var (
	Revision = "fssffsfs"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("version=%s revision=%s\n", cmd.Root().Version, Revision)
	}

	cmd := &cli.Command{
		Name:    "partay",
		Version: "v19.99.0",
		Flags: []cli.Flag{
			&cli.TimestampFlag{
				Name: "meeting",
				Config: cli.TimestampConfig{
					Timezone: time.Local,
					Layouts:  []string{"2006-01-02T15:04:05"},
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("%s", cmd.Timestamp("meeting").String())
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
