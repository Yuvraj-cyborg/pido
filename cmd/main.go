package main

import (
	"context"
	"log"
	"os"

	"github.com/Yuvraj-cyborg/pido/internal/types"
	"github.com/Yuvraj-cyborg/pido/internal/utils"
	"github.com/urfave/cli/v3"
)

func main() {
	var compressCmd = &cli.Command{
		Name:  "compress",
		Usage: "Compress images, PDFs, and text files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir",
				Usage: "Directory to scan for files (optional)",
			},
			&cli.IntFlag{
				Name:     "quality",
				Usage:    "Compression quality percentage (0-100)",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "out",
				Usage: "Optional output folder for compressed files",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			opts := types.CompressionOptions{
				Dir:     cmd.String("dir"),
				Quality: cmd.Int("quality"),
				OutDir:  cmd.String("out"),
				Files:   cmd.Args().Slice(),
			}
			return utils.DispatchCompression(opts)
		},
	}

	cmd := &cli.Command{
		Name:  "pido",
		Usage: "A CLI for compressing files (images, PDFs, text)",
		Commands: []*cli.Command{
			compressCmd,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
