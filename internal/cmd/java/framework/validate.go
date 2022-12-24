package framework

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/photowey/popctl/configs"
	"github.com/photowey/popctl/internal/cmd/java"
	"github.com/photowey/popctl/pkg/ipaddr"
	"github.com/photowey/popctl/pkg/stringz"
)

const (
	VersionTemplate = "1.0.0"
	FixedDummyIp    = "192.169.10.24"
)

var promptTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . | green }} ",
	Valid:   "{{ . | cyan }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

func validateInput(args *java.Args) {
	validateProjectCode(args)
	validateProjectName(args)
	validateApp(args)
	validatePath(args)
	validateAuthor(args)
	validateEmail(args)
	validateVersion(args)
	acquireIp(args)
}

func validateProjectCode(args *java.Args) {
	if stringz.IsBlankString(args.ProjectCode) {
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty project code")
			}
			return nil
		}
		prompt := promptui.Prompt{
			Label:     "Please enter the current project code: ",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   "popup-helloapp",
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct project code: %v\n", err)

			// Loop
			validateProjectCode(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.ProjectCode = result
	}
}

func validateProjectName(args *java.Args) {
	if stringz.IsBlankString(args.ProjectName) {
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty project name")
			}
			return nil
		}
		prompt := promptui.Prompt{
			Label:     "Please enter the current project name: ",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   "萌新引导",
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct project name: %v\n", err)

			// Loop
			validateProjectName(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.ProjectName = result
	}
}

func validateApp(args *java.Args) {
	if stringz.IsBlankString(args.App) {
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty app.name")
			}
			return nil
		}

		defaultAppName := stringz.Tail(args.ProjectCode, "-")

		prompt := promptui.Prompt{
			Label:     "Please enter the current project app.name: ",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   defaultAppName,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct project app.name: %v\n", err)

			// Loop
			validateApp(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.App = result
	}
}

func validatePath(args *java.Args) {
	if stringz.IsBlankString(args.Path) {
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty project dir")
			}
			return nil
		}

		defaultPath := args.Pwd

		prompt := promptui.Prompt{
			Label:     "Please enter the project dir: ",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   defaultPath,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct project dir: %v\n", err)

			// Loop
			validatePath(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.Path = result
	}
}

func validateAuthor(args *java.Args) {
	if stringz.IsBlankString(args.Author) {
		if stringz.IsNotBlankString(configs.ProjectFunc().Author) {
			args.Author = configs.ProjectFunc().Author
			return
		}

		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty author's name")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:     "Please enter the author of project: ",
			Validate:  validate,
			Templates: promptTemplates,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the author of project: %v\n", err)

			// Loop
			validateAuthor(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.Author = result
	}
}

func validateEmail(args *java.Args) {
	if stringz.IsBlankString(args.Email) {
		if stringz.IsNotBlankString(configs.ProjectFunc().Email) {
			args.Email = configs.ProjectFunc().Email
			return
		}

		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty author's email")
			}
			return nil
		}

		defaultEmail := args.Author + args.CompanyEmail

		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("Please enter the email of author(%s): ", args.Author),
			Validate:  validate,
			Templates: promptTemplates,
			Default:   defaultEmail,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct email of author(%s): %v\n", args.Author, err)

			// Loop
			validateEmail(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.Email = result
	}
}

func validateVersion(args *java.Args) {
	if stringz.IsBlankString(args.Version) {
		if stringz.IsNotBlankString(configs.ProjectFunc().Version) {
			args.Version = configs.ProjectFunc().Version
			return
		}
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty version of project")
			}
			return nil
		}

		defaultVersion := VersionTemplate

		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("Please enter the version of project(%s): ", args.ProjectCode),
			Validate:  validate,
			Templates: promptTemplates,
			Default:   defaultVersion,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct version of project(%s): %v\n", args.Author, err)

			// Loop
			validateVersion(args)
		}

		fmt.Printf("You answered: %s\n", result)
		args.Version = result
	}
}

func acquireIp(args *java.Args) {
	args.LocalIp = ipaddr.LocalIp
	if args.LocalIp == "" {
		args.LocalIp = FixedDummyIp
	}
}
