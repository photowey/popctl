package java

import (
	"github.com/photowey/popctl/internal/cmd/java"
	"github.com/photowey/popctl/internal/cmd/java/framework"
	"github.com/photowey/popctl/internal/cmd/java/reverse"
)

func OnFramework(args *java.Args) {
	framework.OnEvent(args)
}

func OnReverseEngineering(args *java.Args) {
	reverse.OnEvent(args)
}
