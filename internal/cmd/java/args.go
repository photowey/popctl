package java

type Args struct {
	App             string // the spring application name of the microservice webapp
	Path            string // the path of code generated
	Pwd             string // current working directory
	Args            []string
	Author          string // the name of author
	Email           string // the email of author
	ProjectCode     string // the code of the project
	ProjectName     string // the name of the project
	AppTemplate     string // the template app name of the project
	EngineTemplate  string // the template engine name of the project
	EnableTemplate  string // the template enable service name of the project
	ProjectTemplate string // the template name of the project
	CompanyEmail    string // the email suffix of company
	Version         string // the version of the microservice
	Date            string // now
	LocalIp         string // local Ip
}
