package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
	"github.com/uynilo9/logger.go"
)

type (
	add struct {
		Dev bool `arg:"--dev,-d" help:"set dependency(ies) as \"devDependencies\" in package.json"`
		Optional bool `arg:"--optional,-o" help:"set dependency(ies) as \"optionalDependencies\" in package.json"`
		Trusted bool `arg:"--trusted,-t" help:"set dependency(ies) as \"trustedDependencies\" in package.json"`
		Peer bool `arg:"--peer,-p" help:"set dependency(ies) as \"peerDependencies\" in package.json"`
		Packages []string `arg:"positional" help:"the package(s) you want to install" placeholder:"<package>"`
	}
	remove struct {
		Packages []string `arg:"positional" help:"the package(s) you want to uninstall" placeholder:"<package>"`
	}
	args struct { 
		Add *add `arg:"subcommand:add" help:"install dependency(ies) for your project"`
		Remove *remove `arg:"subcommand:remove" help:"uninstall dependency(ies) for your project"`
		Version bool `arg:"--version,-v" help:"display the version and exit"`
		License bool `arg:"--license,-l" help:"display the license and exit"`
	}
)

func (args) Description() string {
	return "🥨 Welcome to Pretzel " + os.Getenv("VERSION") + "\n"
}

func (args) Epilogue() string {
	return "✨ Visit " + os.Getenv("WEBSITE") + " to get more infomation about Pretzel"
}

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	env := filepath.Join(dir, "../.env")
	if godotenv.Load(env) != nil && godotenv.Load() != nil {
		logger.Fatalf("Failed to find or load the dotenv file `%s`\n", env)
	}
	
	raw := strings.Join(append([]string{os.Getenv("A")}, os.Args[1:]...), " ")
	
	var args args
	goarg, err := arg.NewParser(arg.Config{Program: os.Getenv("A")}, &args)
	if err != nil {
		logger.Detailln(err)
		logger.Fatalln("Failed to create the argument parser")
	}
	
	goarg.Parse(os.Args[1:])
	switch {
		case args.Version:
			fmt.Println("🥨 Pretzel " + os.Getenv("VERSION"))
			os.Exit(0)
		case args.License:
			fmt.Println("📜 Apache License 2.0 Copyright " + os.Getenv("YEAR") + " @uynilo9")
			os.Exit(0)
		case args.Add != nil:
			if args.Add.Packages != nil {
				program := filepath.Join(dir, os.Getenv("B"))
				cmd := exec.Command(
					program,
					"--allowed=true",
					"--@=add",
					"--dev=" + strconv.FormatBool(args.Add.Dev),
					"--optional=" + strconv.FormatBool(args.Add.Optional),
					"--trusted=" + strconv.FormatBool(args.Add.Trusted),
					"--peer=" + strconv.FormatBool(args.Add.Peer),
					strings.Join(args.Add.Packages, " "),
				)
				// err := cmd.Run()
				output, err := cmd.CombinedOutput() // NOTE For debugging
				if err != nil {
					logger.Detailln(err)
					logger.Fatalf("Failed to execute the command `%s`\n", raw)
				}
				fmt.Println(string(output)) // NOTE For debugging
			} else {
				logger.Errorln("The argument <package> was required while running the `add` subcommand")
				goarg.WriteHelp(os.Stdout)
			}
		case args.Remove != nil:
			if args.Remove.Packages != nil {
				program := filepath.Join(dir, os.Getenv("B"))
				cmd := exec.Command(
					program,
					"--allowed=true",
					"--@=remove",
					strings.Join(args.Remove.Packages, " "),
				)
				// err := cmd.Run()
				output, err := cmd.CombinedOutput() // NOTE For debugging
				if err != nil {
					logger.Detailln(err)
					logger.Fatalf("Failed to execute the command `%s`\n", raw)
				}
				fmt.Println(string(output)) // NOTE For debugging
			} else {
				logger.Errorln("The argument <package> was required while running the `remove` subcommand")
				goarg.WriteHelp(os.Stdout)		
			}
		default:
			goarg.WriteHelp(os.Stdout)
	}
}