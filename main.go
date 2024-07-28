package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jguer/go-alpm/v2"
	"github.com/helioloureiro/golorama"
)

const (
	ON  bool = true
	OFF bool = false
)

type Options struct {
	install bool
	update  bool
	remove  bool
	list    bool
	search  bool
	help    bool
	info    bool
}

var (
	DEFAULTOPTIONS = [...]string{"--noconfirm", "--color=always"}
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Missing arguments")
	}

	opt := initializeOptions()
	command := os.Args[1]
	switch command {
	case "install":
		opt.install = ON
		break
	case "upgrade":
		opt.update = ON
		break
	case "dist-upgrade":
		opt.update = ON
		break
	case "distupgrade":
		opt.update = ON
		break
	case "update":
		opt.update = ON
		break
	case "remove":
		opt.remove = ON
		break
	case "purge":
		opt.remove = ON
		break
	case "list":
		opt.list = ON
		break
	case "search":
		opt.search = ON
		break
	case "help":
		opt.help = ON
		break
	case "info":
		opt.info = ON
		break
	default:
		log.Fatal("Unknown option: " + command)
	}

	takeActions(opt)
}

func initializeOptions() Options {
	var opt Options
	opt.install = OFF
	opt.remove = OFF
	opt.update = OFF
	opt.list = OFF
	opt.search = OFF
	opt.help = OFF
	opt.info = OFF

	return opt
}

func takeActions(opt Options) {
	if opt.help {
		usage(0)
	} else if opt.list {
		ListPackages()
	} else if opt.search {
		SearchPackages()
	} else if opt.info {
		InfoPackage()
	} else if opt.install {
		InstallPackage()
	} else if opt.update {
		UpdatePackages()
	} else if opt.remove {
		RemovePackage()
	}

}

func usage(exitCode int) {
	fmt.Println("Usage: " + os.Args[0] + "<option> <flags>")
	fmt.Println("Option:")
	fmt.Println("  install: to install new package")
	fmt.Println("  update: to update packages")
	fmt.Println("  upgrade: same as update")
	fmt.Println("  dist-upgrade: same as update")
	fmt.Println("  distupgrade: same as update")
	fmt.Println("  search: to search for a package")
	fmt.Println("  list: to list installed packages")
	fmt.Println("  info: to provide info about a package")
	fmt.Println("  help: this help")

	os.Exit(exitCode)
}

func SearchPackages() {
	h, er := alpm.Initialize("/", "/var/lib/pacman")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer h.Release()
	db, _ := h.RegisterSyncDB("core", 0)
	h.RegisterSyncDB("community", 0)
	h.RegisterSyncDB("extra", 0)

	for _, pkg := range db.PkgCache().Slice() {
		fmt.Printf("%s==%s: %s\n",
			golorama.GetCSI(golorama.GREEN)+pkg.Name(), golorama.GetCSI(golorama.BLUE)+pkg.Version(), golorama.GetCSI(golorama.RESET)+pkg.Description())
	}
}

func GetArgsArray(option string) []string {
	arguments := []string{option}
	for _, value := range DEFAULTOPTIONS {
		arguments = append(arguments, value)
	}

	for i, value := range os.Args {
		if i < 2 {
			continue
		}
		arguments = append(arguments, value)
	}
	return arguments
}

func InfoPackage() {
	arguments := GetArgsArray("-Qi")
	output, err := ShellExec(PACMAN, arguments...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

func InstallPackage() {
	CheckRoot()
	arguments := GetArgsArray("-S")
	output, err := ShellExec(PACMAN, arguments...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

func CheckRoot() {
	uid := os.Geteuid()
	if uid != 0 {
		log.Fatal("You must run as root!")
	}
}

func UpdatePackages() {
	CheckRoot()
	arguments := GetArgsArray("-Syu")
	output, err := ShellExec(PACMAN, arguments...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

func RemovePackage() {
	CheckRoot()
	arguments := GetArgsArray("-R")
	output, err := ShellExec(PACMAN, arguments...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
