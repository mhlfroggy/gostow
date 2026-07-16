// TODO: Ensure flags which are mutually exclusive are set as such
// TODO: Enable support for single-file configurations

package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "Print gostow version",
	}

	cmd := &cli.Command{
		Name:      "gostow",
		Usage:     "An improved dotfile management utility",
		Version:   "v0.0.1-alpha",
		ArgsUsage: "DIR",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dry-run", Value: false, Usage: "Do a dry run, showing what operations would be performed"},
			&cli.StringFlag{Name: "dir", Aliases: []string{"d"}, Usage: "Set the source directory to `DIR` (default is current directory)"},
			&cli.StringFlag{Name: "target", Aliases: []string{"t"}, Usage: "Set the target directory to `DIR` (default is ~/.config)"},
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Value: false, Usage: "Show verbose output"},
			&cli.BoolFlag{Name: "remove", Aliases: []string{"rm"}, Value: false, Usage: "Remove dotfiles from configuration directory instead of linking them"},
			&cli.BoolFlag{Name: "replace", Aliases: []string{"r"}, Value: false, Usage: "Remove all dotfiles from configuration, then relinks them"},
			&cli.StringFlag{Name: "ignore", Usage: "Ignore all files matching `REGEX`"},
			&cli.StringFlag{Name: "override", Usage: "Force linking of files matching `REGEX` even if the config already exists"},
			// &cli.BoolFlag{Name:   "migrate",  Value: false, Usage: "Generate a dotfiles directory and migrate all your dotfiles"},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			var DIR_PATH string
			var CONFIG_PATH string
			var err error
			configDirectory := cmd.Args().Get(0)

			// Logic handling for --dir/-d flag
			// If this flag is present, the source directory for the
			// configuration file is set to the user-supplied location.
			// Default value is current directory
			if cmd.IsSet("dir") {
				DIR_PATH = filepath.Join(cmd.String("dir"), configDirectory)
			} else {
				currentDir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				DIR_PATH = filepath.Join(currentDir, configDirectory)
			}

			// Logic handling for --target/-t flag
			// If this flag is present, the target directory where the new
			// config folder will be created is set to whatever the user
			// provides as an argument following the flag. Default value is
			///home/user/.config
			if cmd.IsSet("target") {
				CONFIG_PATH = cmd.String("target")
			} else {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					log.Fatal(err)
				}
				CONFIG_PATH = filepath.Join(homeDir, ".config")
			}

			// Generate the new directory in the config path
			configDirectoryPath := filepath.Join(CONFIG_PATH, configDirectory)
			os.MkdirAll(configDirectoryPath, 0755)

			// Iterate through all files in the source directory path,
			// symlinking them to the new directory created in the target
			// directory
			confFiles, err := os.ReadDir(DIR_PATH)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range confFiles {
				// Handling for regex if it exists
				// Specifically, this will skip the file if the --ignore/-i flag
				// is present
				if cmd.IsSet("ignore") {
					regexMatch, err := regexp.MatchString(cmd.String("ignore"), file.Name())
					if err != nil {
						log.Fatal(err)
					}
					if regexMatch == true {
						continue
					}
				}

				// Create origin and destination filepaths
				sourceFile := filepath.Join(DIR_PATH, file.Name())
				linkedFile := filepath.Join(CONFIG_PATH, configDirectory, file.Name())

				// Handling of regexp if override is enabled
				// This will wipe the existing file if present
				// If the file already exists, the operation will abort
				_, err := os.Stat(linkedFile)
				fileNotExist := os.IsNotExist(err)
				if cmd.IsSet("replace") && fileNotExist {
					if _, err := os.Lstat(linkedFile); err == nil {
						fmt.Printf("Unlinking %s\n", linkedFile)
						os.Remove(linkedFile)
					}
				} else if !fileNotExist {
					log.Fatal("ERROR: File already exists. Remove before continuing")
				}

				// Unlink file if --remove/-r flag is present
				if cmd.Bool("remove") == true {
					if _, err := os.Lstat(linkedFile); err == nil {
						os.Remove(linkedFile)
					}
					continue
				}

				// Display file actions if --verbose flag is present
				if cmd.Bool("verbose") == true {
					fmt.Printf("Linking %s to %s\n", sourceFile, linkedFile)
				}

				// Perform a dry run if --dry-run flag is present
				if cmd.Bool("dry-run") == true {
					fmt.Printf("Linking %s to %s\n", sourceFile, linkedFile)
					continue
				}

				// Perform symlinking
				os.Symlink(sourceFile, linkedFile)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
