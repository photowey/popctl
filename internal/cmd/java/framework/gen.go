package framework

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"

	"github.com/photowey/popctl/internal/cmd/java"
	"github.com/photowey/popctl/pkg/stringz"
)

var (
	path string
	argz *java.Args
	ctx  Context

	green  = color.FgGreen.Render
	blue   = color.FgBlue.Render
	yellow = color.FgYellow.Render
	cyan   = color.FgCyan.Render
)

func gen(args *java.Args) {
	path = args.Path
	argz = args

	populateTemplateContext(args)

	fmt.Println("")
	fmt.Println(green("---------------- $ popctl java framework input report ----------------"))
	fmt.Println(blue("ProjectCode name:"), argz.ProjectCode)
	fmt.Println(blue("Microservice app.name:"), argz.App)
	fmt.Println(blue("ProjectCode path:"), argz.Path)
	fmt.Println(blue("ProjectCode author:"), argz.Author)
	fmt.Println(blue("ProjectCode author's email:"), argz.Email)
	fmt.Println(blue("ProjectCode date:"), ctx.Date)
	fmt.Println(blue("ProjectCode version:"), ctx.Version)
	fmt.Println(green("---------------- $ popctl java framework input report ----------------"))
	fmt.Println("")

	prompt := promptui.Prompt{
		Label:     "ProjectCode info Confirm:",
		IsConfirm: true,
		Default:   "Y",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Confirm failed: %v\n", err)
		return
	}

	fmt.Printf("You choose: %q\n", result)

	if err = write(); err != nil {
		fmt.Printf("Generate framework project failed %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println(green("---------------- $ popctl java framework generate report ----------------"))
	fmt.Println(blue("ProjectCode created successfully"))
	fmt.Println(cyan("Run cmd:"))
	fmt.Println(yellow("$ cd " + argz.Path + string(os.PathSeparator) + argz.ProjectCode))
	fmt.Println(yellow("$ mvn install"))
	fmt.Println(yellow("$ mvn clean -DskipTests source:jar deploy"))
	fmt.Println(green("---------------- $ popctl java framework generate report ----------------"))

	return
}

func populateTemplateContext(args *java.Args) {
	now := time.Now()
	layout := "2006/01/02"

	ctx = Context{
		App:         argz.App,
		PascalApp:   stringz.Pascal(argz.App),
		ProjectCode: argz.ProjectCode,
		ProjectName: argz.ProjectName,
		Package:     strings.ReplaceAll(argz.ProjectCode, "-", "."),
		Author:      args.Author,
		Email:       args.Email,
		Date:        now.Format(layout),
		Version:     args.Version,
		LocalIp:     args.LocalIp,
	}
}
