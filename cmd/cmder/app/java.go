package app

import (
	"fmt"

	"github.com/photowey/popctl/cmd/cmder/java"

	javax "github.com/photowey/popctl/internal/cmd/java"
	"github.com/spf13/cobra"
)

const (
	ProjectNameTemplate           = "hicoomore-intepay-platform-template"
	ApplicationNameTemplate       = "Template"
	EngineNameTemplate            = "TemplateEngine"
	EnableTemplateServiceTemplate = "EnableTemplateService"
	CompanyEmailTemplate          = "@uphicoo.com"
)

var (
	projectCode  string
	projectName  string
	appName      string
	versionInput string
	author       string

	javaCmd = &cobra.Command{
		Use:   "java",
		Short: "Generate Java related applications",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello popctl java~")
		},
	}

	frameworkCmd = &cobra.Command{
		Use:     "framework",
		Aliases: []string{"f", "fw"},
		Short:   "Generate java Spring microservice webapp framework",
		Run: func(cmd *cobra.Command, args []string) {
			java.OnFramework(&javax.Args{
				App:             appName,
				Args:            args,
				ProjectCode:     projectCode,
				ProjectName:     projectName,
				Author:          author,
				Version:         versionInput,
				AppTemplate:     ApplicationNameTemplate,
				EngineTemplate:  EngineNameTemplate,
				EnableTemplate:  EnableTemplateServiceTemplate,
				ProjectTemplate: ProjectNameTemplate,
				CompanyEmail:    CompanyEmailTemplate,
			})
		},
	}
	reverseCmd = &cobra.Command{
		Use:     "reverse",
		Aliases: []string{"r", "re"},
		Short:   "Database reverse engineering",
		Run: func(cmd *cobra.Command, args []string) {
			java.OnReverseEngineering(&javax.Args{
				App:             appName,
				Args:            args,
				ProjectCode:     projectCode,
				ProjectName:     projectName,
				Author:          author,
				Version:         versionInput,
				AppTemplate:     ApplicationNameTemplate,
				EngineTemplate:  EngineNameTemplate,
				EnableTemplate:  EnableTemplateServiceTemplate,
				ProjectTemplate: ProjectNameTemplate,
				CompanyEmail:    CompanyEmailTemplate,
			})
		},
	}
)

func init() {
	frameworkCmd.PersistentFlags().StringVarP(&projectCode, "code", "c", "", "Project code")
	frameworkCmd.PersistentFlags().StringVarP(&projectName, "projectName", "p", "", "Project name")
	frameworkCmd.PersistentFlags().StringVarP(&appName, "app", "a", "", "Service name")
	frameworkCmd.PersistentFlags().StringVarP(&versionInput, "version", "v", "", "Service version")
	frameworkCmd.PersistentFlags().StringVarP(&author, "author", "u", "", "Author")
	javaCmd.AddCommand(frameworkCmd)
}

func init() {
	reverseCmd.PersistentFlags().StringVarP(&projectCode, "code", "c", "", "Project code")
	reverseCmd.PersistentFlags().StringVarP(&projectName, "projectName", "p", "", "Project name")
	reverseCmd.PersistentFlags().StringVarP(&appName, "app", "a", "", "Service name")
	reverseCmd.PersistentFlags().StringVarP(&versionInput, "version", "v", "", "Service version")
	reverseCmd.PersistentFlags().StringVarP(&author, "author", "u", "", "Author")
	javaCmd.AddCommand(reverseCmd)
}
