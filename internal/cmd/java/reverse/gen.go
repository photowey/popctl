package reverse

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"

	"github.com/photowey/popctl/configs"
	"github.com/photowey/popctl/internal/cmd/java"
	"github.com/photowey/popctl/internal/cmd/java/reverse/database"
	"github.com/photowey/popctl/pkg/alphabet"
	"github.com/photowey/popctl/pkg/stringz"
)

const (
	JavaTypeLong                = "Long"
	JavaTypeString              = "String"
	JavaTypeBigDecimal          = "BigDecimal"
	JavaTypeBigDecimalImport    = "java.math.BigDecimal"
	JavaTypeLocalDateTime       = "LocalDateTime"
	JavaTypeLocalDateTimeImport = "java.time.LocalDateTime"
	JsonSerialize               = "com.fasterxml.jackson.databind.annotation.JsonSerialize"
	ToStringSerializer          = "com.fasterxml.jackson.databind.ser.std.ToStringSerializer"
	Size                        = "javax.validation.constraints.Size"
	Digits                      = "javax.validation.constraints.Digits"
	NotNull                     = "javax.validation.constraints.NotNull"
	NotBlank                    = "javax.validation.constraints.NotBlank"
	NotEmpty                    = "javax.validation.constraints.NotEmpty"
	SafeHtml                    = "org.hibernate.validator.constraints.SafeHtml" // @Deprecated
)

const (
	dateLayout = "2006/01/02"
)

var (
	path string
	argz *java.Args
	ctx  *Context

	green  = color.FgGreen.Render
	blue   = color.FgBlue.Render
	yellow = color.FgYellow.Render
	cyan   = color.FgCyan.Render
)

func gen(args *java.Args) {
	path = args.Path
	argz = args

	databasePtr, err := database.ReverseEngineering()
	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now()
	layout := dateLayout

	fmt.Println("")
	fmt.Println(green("---------------- $ popctl java reverse input report ----------------"))
	fmt.Println(blue("Project name:"), argz.ProjectCode)
	fmt.Println(blue("Microservice app.name:"), argz.App)
	fmt.Println(blue("Project path:"), argz.Path)
	fmt.Println(blue("Project author:"), argz.Author)
	fmt.Println(blue("Project author's email:"), argz.Email)
	fmt.Println(blue("Project date:"), now.Format(layout))
	fmt.Println(blue("Project version:"), argz.Version)
	fmt.Println(green("---------------- $ popctl java reverse input report ----------------"))
	fmt.Println("")

	prompt := promptui.Prompt{
		Label:     "Project info Confirm:",
		IsConfirm: true,
		Default:   "Y",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Confirm failed: %v\n", err)
		return
	}

	fmt.Printf("You choose: %q\n", result)

	doGenerate(databasePtr, args)

	fmt.Println("")
	fmt.Println(green("---------------- $ popctl java reverse generate report ----------------"))
	fmt.Println(blue("ProjectCode created successfully"))
	fmt.Println(cyan("Run cmd:"))
	fmt.Println(yellow("$ cd " + argz.Path + string(os.PathSeparator) + argz.ProjectCode))
	fmt.Println(yellow("$ mvn install"))
	fmt.Println(yellow("$ mvn clean -DskipTests source:jar deploy"))
	fmt.Println(green("---------------- $ popctl java reverse generate report ----------------"))
}

func doGenerate(databasePtr *database.Database, args *java.Args) {
	initTmplCtx(args)

	engine(databasePtr, args)
	engineImpl(databasePtr, args)

	excludes := configs.DatabaseFunc().Excludes
	includes := configs.DatabaseFunc().Includes
	for _, tablePtr := range databasePtr.Tables {
		if len(excludes) > 0 && stringz.ArrayContains(excludes, tablePtr.Name) {
			continue
		}
		if len(includes) > 0 && !stringz.ArrayContains(includes, tablePtr.Name) {
			continue
		}
		populateTmplCtx(tablePtr, args)

		controller(ctx, tablePtr, args)

		service(ctx, tablePtr, args)
		serviceImpl(ctx, tablePtr, args)
		assembler(ctx, tablePtr, args)
		repository(ctx, tablePtr, args)

		entity(ctx, tablePtr, args)
		dto(ctx, tablePtr, args)

		payload(ctx, tablePtr, args)
		query(ctx, tablePtr, args)

		if err := write(); err != nil {
			fmt.Printf("write tmpl failed, err:%v\n", err)
			return
		}
	}
}

func controller(ctx *Context, tablePtr *database.Table, args *java.Args) {
	mapping := strings.ReplaceAll(tablePtr.Name, "_", "-")
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Controller = &Controller{
		Name:    entityName + "Controller",
		Comment: ctx.Table.Comment,
		Mapping: mapping,
	}
}

func service(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Service = &Service{
		Name:    entityName + "Service",
		Comment: ctx.Table.Comment,
	}
}

func serviceImpl(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.ServiceImpl = &ServiceImpl{
		Name:    entityName + "ServiceImpl",
		Comment: ctx.Table.Comment,
	}
}

func assembler(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Assembler = &Assembler{
		Name: entityName + "Assembler",
	}
}

func repository(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Repository = &Repository{
		Name:    entityName + "Repository",
		Comment: ctx.Table.Comment,
	}
}

func entity(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Entity = &Entity{
		Name:      entityName,
		CamelName: alphabet.CamelCase(entityName),
		Comment:   ctx.Table.Comment,
	}

	imports := make([]string, 0)
	fields := make([]*Field, 0, len(tablePtr.Fields))
	for _, field := range tablePtr.Fields {
		javaType, importz := convertJavaType(field.Type)

		fd := populateTableField(field, javaType)
		d := configs.ProjectFunc().Entity

		needAdd := true
		if stringz.ArrayContains(d.Excludes, fd.Name) {
			needAdd = false
		}

		if stringz.IsNotBlankString(importz) && !stringz.ArrayContains(imports, importz) {
			imports = append(imports, importz)
		}

		if needAdd {
			fields = append(fields, fd)
		}
	}

	ctx.Entity.Imports = imports
	ctx.Entity.Fields = fields
}

func dto(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Dto = &Dto{
		Name:    entityName + "DTO",
		Comment: ctx.Table.Comment,
	}

	imports := make([]string, 0)
	fields := make([]*Field, 0, len(tablePtr.Fields))
	for _, field := range tablePtr.Fields {
		javaType, importz := convertJavaType(field.Type)

		fd := populateTableField(field, javaType)
		d := configs.ProjectFunc().Dto

		needAdd := true
		if stringz.ArrayContains(d.Excludes, fd.Name) {
			needAdd = false
		}

		if needAdd && stringz.IsNotBlankString(importz) && !stringz.ArrayContains(imports, importz) {
			imports = append(imports, importz)
		}

		if needAdd && javaType == JavaTypeLong && !stringz.ArrayContains(imports, JsonSerialize) {
			imports = append(imports, JsonSerialize)
		}
		if needAdd && javaType == JavaTypeLong && !stringz.ArrayContains(imports, ToStringSerializer) {
			imports = append(imports, ToStringSerializer)
		}

		if needAdd && javaType == JavaTypeString && !stringz.ArrayContains(imports, Size) {
			// @SafeHtml Deprecated
			// imports = append(imports, SafeHtml)
			imports = append(imports, Size)
		}

		if needAdd && javaType == JavaTypeBigDecimal && !stringz.ArrayContains(imports, Digits) {
			imports = append(imports, Digits)
		}

		if needAdd {
			fields = append(fields, fd)
		}
	}

	ctx.Dto.Imports = imports
	ctx.Dto.Fields = fields
}

func payload(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Payload = &Payload{
		Name:    entityName + "Payload",
		Comment: ctx.Table.Comment,
	}

	imports := make([]string, 0)
	addImports := make([]string, 0)
	deleteImports := make([]string, 0)
	updateImports := make([]string, 0)

	fields := make([]*Field, 0, len(tablePtr.Fields))
	addFields := make([]*Field, 0, len(tablePtr.Fields))
	deleteFields := make([]*Field, 0, len(tablePtr.Fields))
	updateFields := make([]*Field, 0, len(tablePtr.Fields))

	for _, field := range tablePtr.Fields {
		javaType, importz := convertJavaType(field.Type)

		fd := populateTableField(field, javaType)
		p := configs.ProjectFunc().Payload

		needAdd := true
		needDelete := true
		needUpdate := true

		if stringz.IsNotEmptyStringSlice(p.Add.Includes) && !stringz.ArrayContains(p.Add.Includes, fd.Name) {
			needAdd = false
		}
		if stringz.IsNotEmptyStringSlice(p.Add.Excludes) && stringz.ArrayContains(p.Add.Excludes, fd.Name) {
			needAdd = false
		}

		if stringz.IsNotEmptyStringSlice(p.Delete.Includes) && !stringz.ArrayContains(p.Delete.Includes, fd.Name) {
			needDelete = false
		}
		if stringz.IsNotEmptyStringSlice(p.Delete.Excludes) && stringz.ArrayContains(p.Delete.Excludes, fd.Name) {
			needDelete = false
		}

		if stringz.IsNotEmptyStringSlice(p.Update.Includes) && !stringz.ArrayContains(p.Update.Includes, fd.Name) {
			needUpdate = false
		}
		if stringz.IsNotEmptyStringSlice(p.Update.Excludes) && stringz.ArrayContains(p.Update.Excludes, fd.Name) {
			needUpdate = false
		}

		// ---- imports

		if stringz.IsNotBlankString(importz) && !stringz.ArrayContains(imports, importz) {
			imports = append(imports, importz)
		}

		if javaType == JavaTypeLong && !stringz.ArrayContains(imports, JsonSerialize) {
			imports = append(imports, JsonSerialize)
		}
		if javaType == JavaTypeLong && !stringz.ArrayContains(imports, ToStringSerializer) {
			imports = append(imports, ToStringSerializer)
		}

		if javaType == JavaTypeString && !stringz.ArrayContains(imports, Size) {
			// @SafeHtml Deprecated
			// imports = append(imports, SafeHtml)
			imports = append(imports, Size)
		}

		if javaType == JavaTypeBigDecimal && !stringz.ArrayContains(imports, Digits) {
			imports = append(imports, Digits)
		}

		// ---- addImports

		if needAdd && stringz.IsNotBlankString(importz) && !stringz.ArrayContains(addImports, importz) {
			addImports = append(addImports, importz)
		}

		if needAdd && javaType == JavaTypeLong && !stringz.ArrayContains(addImports, JsonSerialize) {
			addImports = append(addImports, JsonSerialize)
		}
		if needAdd && javaType == JavaTypeLong && !stringz.ArrayContains(addImports, ToStringSerializer) {
			addImports = append(addImports, ToStringSerializer)
		}

		if needAdd && javaType == JavaTypeString && !stringz.ArrayContains(addImports, Size) {
			// @SafeHtml Deprecated
			// addImports = append(addImports, SafeHtml)
			addImports = append(addImports, Size)
		}

		if needAdd && javaType == JavaTypeBigDecimal && !stringz.ArrayContains(addImports, Digits) {
			addImports = append(addImports, Digits)
		}

		// ---- deleteImports

		if needDelete && stringz.IsNotBlankString(importz) && !stringz.ArrayContains(deleteImports, importz) {
			deleteImports = append(deleteImports, importz)
		}

		if needDelete && javaType == JavaTypeLong && !stringz.ArrayContains(deleteImports, JsonSerialize) {
			deleteImports = append(deleteImports, JsonSerialize)
		}
		if needDelete && javaType == JavaTypeLong && !stringz.ArrayContains(deleteImports, ToStringSerializer) {
			deleteImports = append(deleteImports, ToStringSerializer)
		}

		if needDelete && javaType == JavaTypeString && !stringz.ArrayContains(deleteImports, Size) {
			// @SafeHtml Deprecated
			// deleteImports = append(deleteImports, SafeHtml)
			deleteImports = append(deleteImports, Size)
		}

		if needDelete && javaType == JavaTypeBigDecimal && !stringz.ArrayContains(deleteImports, Digits) {
			deleteImports = append(deleteImports, Digits)
		}

		// ---- updateImports

		if needUpdate && stringz.IsNotBlankString(importz) && !stringz.ArrayContains(updateImports, importz) {
			updateImports = append(updateImports, importz)
		}

		if needUpdate && javaType == JavaTypeLong && !stringz.ArrayContains(updateImports, JsonSerialize) {
			updateImports = append(updateImports, JsonSerialize)
		}
		if needUpdate && javaType == JavaTypeLong && !stringz.ArrayContains(updateImports, ToStringSerializer) {
			updateImports = append(updateImports, ToStringSerializer)
		}

		if needUpdate && javaType == JavaTypeString && !stringz.ArrayContains(updateImports, Size) {
			// @SafeHtml Deprecated
			// updateImports = append(updateImports, SafeHtml)
			updateImports = append(updateImports, Size)
		}

		if needUpdate && javaType == JavaTypeBigDecimal && !stringz.ArrayContains(updateImports, Digits) {
			updateImports = append(updateImports, Digits)
		}

		// default
		if !stringz.ArrayContains(imports, NotNull) && !stringz.ArrayContains(imports, NotBlank) {
			imports = append(imports, NotNull)
		}

		if javaType == JavaTypeString {
			fd.NotBlankCheck = true
			if !stringz.ArrayContains(addImports, NotBlank) {
				addImports = append(addImports, NotBlank)
			}
			if !stringz.ArrayContains(updateImports, NotBlank) {
				updateImports = append(updateImports, NotBlank)
			}
			if !stringz.ArrayContains(deleteImports, NotBlank) {
				deleteImports = append(deleteImports, NotBlank)
			}
		}
		if javaType != JavaTypeString {
			fd.NotNullCheck = true
			if !stringz.ArrayContains(addImports, NotNull) {
				addImports = append(addImports, NotNull)
			}
			if !stringz.ArrayContains(updateImports, NotNull) {
				updateImports = append(updateImports, NotNull)
			}
			if !stringz.ArrayContains(deleteImports, NotNull) {
				deleteImports = append(deleteImports, NotNull)
			}
		}

		// @NotEmpty
		// fd.NotEmptyCheck = true

		fields = append(fields, fd)

		if needAdd {
			addFields = append(addFields, fd)
		}
		if needDelete {
			deleteFields = append(deleteFields, fd)
		}
		if needUpdate {
			updateFields = append(updateFields, fd)
		}
	}

	ctx.Payload.Imports = imports
	ctx.Payload.AddImports = addImports
	ctx.Payload.DeleteImports = deleteImports
	ctx.Payload.UpdateImports = updateImports

	ctx.Payload.Fields = fields
	ctx.Payload.AddFields = addFields
	ctx.Payload.DeleteFields = deleteFields
	ctx.Payload.UpdateFields = updateFields
}

func query(ctx *Context, tablePtr *database.Table, args *java.Args) {
	entityName := alphabet.Snake2Pascal(ctx.Table.Name)
	ctx.Query = &Query{
		Name:    entityName + "Query",
		Comment: ctx.Table.Comment,
	}

	imports := make([]string, 0)
	fields := make([]*Field, 0, len(tablePtr.Fields))
	for _, field := range tablePtr.Fields {
		javaType, importz := convertJavaType(field.Type)

		fd := populateTableField(field, javaType)
		// 查询: required = false
		fd.NotNull = false

		q := configs.ProjectFunc().Query

		needAdd := true

		if stringz.ArrayContains(q.Excludes, fd.Name) {
			needAdd = false
		}

		if needAdd && stringz.IsNotBlankString(importz) && !stringz.ArrayContains(imports, importz) {
			imports = append(imports, importz)
		}

		if needAdd && javaType == JavaTypeString && !stringz.ArrayContains(imports, Size) {
			// @SafeHtml Deprecated
			// imports = append(imports, SafeHtml)
			imports = append(imports, Size)
		}

		if needAdd && javaType == JavaTypeBigDecimal && !stringz.ArrayContains(imports, Digits) {
			imports = append(imports, Digits)
		}

		if needAdd {
			fields = append(fields, fd)
		}

	}

	ctx.Query.Imports = imports
	ctx.Query.Fields = fields
}

func engine(databasePtr *database.Database, args *java.Args) {
	ctx.Engine = &Engine{
		Name: alphabet.PascalCase(args.App) + "Engine",
	}

	entities := make([]EngineEntity, 0)
	repositories := make([]string, 0)
	services := make([]string, 0)
	assemblers := make([]string, 0)

	for _, tablePtr := range databasePtr.Tables {
		tableNamex := tablePtr.Name
		for _, prefix := range configs.DatabaseFunc().Prefixes {
			tableNamex = strings.ReplaceAll(tableNamex, prefix, "")
		}

		entityName := alphabet.Snake2Pascal(tableNamex)
		propertyName := alphabet.Snake2Camel(tableNamex)
		ee := EngineEntity{
			ServiceType:    entityName + "Service",
			ServiceName:    propertyName + "Service",
			RepositoryType: entityName + "Repository",
			RepositoryName: propertyName + "Repository",
			AssemblerType:  entityName + "Assembler",
			AssemblerName:  propertyName + "Assembler",
		}

		entities = append(entities, ee)
		repositories = append(repositories, fmt.Sprintf("com.photowey.%s.database.repository.%s", ctx.Package, ee.RepositoryType))
		services = append(services, fmt.Sprintf("com.photowey.%s.service.%s", ctx.Package, ee.ServiceType))
		assemblers = append(assemblers, fmt.Sprintf("com.photowey.%s.service.assembler.%s", ctx.Package, ee.AssemblerType))
	}

	ctx.Engine.Entities = entities
	ctx.Engine.Repositories = repositories
	ctx.Engine.Services = services
	ctx.Engine.Assemblers = assemblers
}

func engineImpl(databasePtr *database.Database, args *java.Args) {
	ctx.EngineImpl = &EngineImpl{
		Name: alphabet.PascalCase(args.App) + "EngineImpl",
	}

	entities := make([]EngineEntity, 0)
	repositories := make([]string, 0)
	services := make([]string, 0)
	assemblers := make([]string, 0)

	for _, tablePtr := range databasePtr.Tables {
		tableNamex := tablePtr.Name
		for _, prefix := range configs.DatabaseFunc().Prefixes {
			tableNamex = strings.ReplaceAll(tableNamex, prefix, "")
		}

		entityName := alphabet.Snake2Pascal(tableNamex)
		propertyName := alphabet.Snake2Camel(tableNamex)
		ee := EngineEntity{
			ServiceType:    entityName + "Service",
			ServiceName:    propertyName + "Service",
			RepositoryType: entityName + "Repository",
			RepositoryName: propertyName + "Repository",
			AssemblerType:  entityName + "Assembler",
			AssemblerName:  propertyName + "Assembler",
		}

		entities = append(entities, ee)
		repositories = append(repositories, fmt.Sprintf("com.photowey.%s.database.repository.%s", ctx.Package, ee.RepositoryType))
		services = append(services, fmt.Sprintf("com.photowey.%s.service.%s", ctx.Package, ee.ServiceType))
		assemblers = append(assemblers, fmt.Sprintf("com.photowey.%s.service.assembler.%s", ctx.Package, ee.AssemblerType))
	}

	ctx.EngineImpl.Entities = entities
	ctx.EngineImpl.Repositories = repositories
	ctx.EngineImpl.Services = services
	ctx.EngineImpl.Assemblers = assemblers
}

func populateTableField(field *database.Field, javaType string) *Field {
	propertyName := alphabet.Snake2Camel(field.Name)
	if field.Name == "gmt_create" {
		propertyName = "createTime"
	}
	if field.Name == "gmt_modified" {
		propertyName = "updateTime"
	}

	CommentPlain := strings.ReplaceAll(field.Comment, "\r\n", "")
	CommentPlain = strings.ReplaceAll(CommentPlain, "\n\n", "")
	CommentPlain = strings.ReplaceAll(CommentPlain, "\n", "")

	fd := &Field{
		Name:         field.Name,
		Comment:      field.Comment,
		CommentPlain: CommentPlain,
		PropertyType: javaType,
		PropertyName: propertyName,
		NotNull:      field.NotNull,
		PrimaryKey:   field.PrimaryKey,
		String:       JavaTypeString,
		Long:         JavaTypeLong,
		BigDecimal:   JavaTypeBigDecimal,
	}
	if fd.PropertyType == "String" {
		if stringz.IsNotBlankString(field.TypLength) {
			fd.Length, _ = strconv.Atoi(field.TypLength[1 : len(field.TypLength)-1])
		}
		if stringz.IsBlankString(field.TypLength) && field.TypLen == -1 {
			fd.Length = 1000 // Default 1000
		}
	}
	if fd.PropertyType == "BigDecimal" {
		formatLen := field.TypLength[1 : len(field.TypLength)-1]
		fd.Length, _ = strconv.Atoi(strings.Split(formatLen, ",")[0])
		fd.MinBit, _ = strconv.Atoi(strings.Split(formatLen, ",")[1])
		fd.MaxBit, _ = strconv.Atoi(strings.Split(formatLen, ",")[0])
	}

	return fd
}

func convertJavaType(jdbcType string) (string, string) {
	switch jdbcType {
	case "int8":
		return JavaTypeLong, ""
	case "int4", "int2":
		return "Integer", ""
	case "varchar", "text":
		return JavaTypeString, ""
	case "timestamp":
		return JavaTypeLocalDateTime, JavaTypeLocalDateTimeImport
	case "numeric":
		return JavaTypeBigDecimal, JavaTypeBigDecimalImport
	}

	return "", ""
}

func initTmplCtx(args *java.Args) {
	now := time.Now()
	layout := dateLayout
	ctx = &Context{
		App:       argz.App,
		PascalApp: alphabet.PascalCase(argz.App),
		Project:   argz.ProjectCode,
		Package:   strings.ReplaceAll(argz.ProjectCode, "-", "."),
		Author:    args.Author,
		Date:      now.Format(layout),
		Version:   args.Version,
	}
}

func populateTmplCtx(tablePtr *database.Table, args *java.Args) {
	tm := tablePtr.Fields[0].TableComment

	tableName := tablePtr.Name
	for _, prefix := range configs.DatabaseFunc().Prefixes {
		tableName = strings.ReplaceAll(tableName, prefix, "")
	}

	ctx.Table = &Table{
		Name:     tableName,
		FullName: tablePtr.Name,
		Comment:  tm,
	}
}
