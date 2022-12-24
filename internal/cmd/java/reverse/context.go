package reverse

type Context struct {
	App       string
	PascalApp string
	Project   string
	Package   string
	Author    string
	Version   string
	Date      string

	*Table
	*Controller
	*Service
	*ServiceImpl
	*Repository
	*Entity
	*Dto
	*Payload
	*Query
	*Engine
	*EngineImpl
	*Assembler
}

type Table struct {
	Name     string
	FullName string
	Comment  string
}

type Controller struct {
	Name    string
	Comment string
	Mapping string
}

type (
	Service struct {
		Name    string
		Comment string
	}
	ServiceImpl struct {
		Name    string
		Comment string
	}
)

type (
	Repository struct {
		Name    string
		Comment string
	}
	Entity struct {
		Name      string
		CamelName string
		Comment   string
		Imports   []string
		Fields    []*Field
	}
	Dto struct {
		Name    string
		Comment string
		Imports []string
		Fields  []*Field
	}
	Payload struct {
		Name          string
		Comment       string
		Imports       []string
		AddImports    []string
		DeleteImports []string
		UpdateImports []string
		Fields        []*Field
		AddFields     []*Field
		UpdateFields  []*Field
		DeleteFields  []*Field
	}
	Query struct {
		Name    string
		Comment string
		Imports []string
		Fields  []*Field
	}
	Field struct {
		Name         string
		Comment      string
		CommentPlain string
		PropertyType string
		PropertyName string
		String       string
		Long         string
		BigDecimal   string
		Length       int
		MinBit       int
		MaxBit       int
		NotNull      bool
		PrimaryKey   bool
	}
)

type (
	Engine struct {
		Name         string
		Entities     []EngineEntity
		Repositories []string
		Services     []string
		Assemblers   []string
	}
	EngineImpl struct {
		Name         string
		Entities     []EngineEntity
		Repositories []string
		Services     []string
		Assemblers   []string
	}
	EngineEntity struct {
		ServiceType    string
		ServiceName    string
		RepositoryType string
		RepositoryName string
		AssemblerType  string
		AssemblerName  string
	}
)

type Assembler struct {
	Name string
}
