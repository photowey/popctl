package reverse

import (
	"fmt"
	"os"

	"github.com/photowey/popctl/internal/cmd/java"
)

func OnEvent(args *java.Args) {
	pwd, _ := os.Getwd()
	args.Pwd = pwd
	fmt.Println("$ pwd")
	fmt.Println(pwd)

	validateInput(args)
	gen(args)
}
