/*
*

	@author: junwang
	@since: 2023/8/24
	@desc: //TODO

*
*/
package generator

import (
"fmt"
"os"
	"github.com/Bifang-Bird/simba/handle"
	"strings"
)

// GenerateCode generates source code based on the given parameters.
func GenerateCode(packageName, structName, grpcgo string) error {
	_, gprcFile, err := handle.LoadFile(grpcgo)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return err
	}

	structs,interfaces := handle.FindStructs(gprcFile,structName)
	fmt.Println("Structs in the file:", structs)
	fmt.Println("Interface in the file:", interfaces)
	var unimplemented string
	for _,structItem:=range structs{
		if strings.HasPrefix(strings.ToLower(structItem),"unimplemented"){
			unimplemented = structItem
			break
		}
	}
	var parentServer = strings.Replace(unimplemented,"Unimplemented","",-1)

	var funcs []string
	for k,v :=range interfaces{
		if k == parentServer{
			funcs = v
		}
	}

	fmt.Println("parentServer", parentServer)
	fmt.Println("unimplementedServer", unimplemented)

	code := fmt.Sprintf("package %s\n\n", packageName)
	code += fmt.Sprintf("import (\n")
	code += fmt.Sprintf("\t\t\"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/gorm\"\n")
	code += fmt.Sprintf("\t\t\"context\"\n")
	code += fmt.Sprintf("\t\t\"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-proto.git/gen\"\n")
	code += fmt.Sprintf("\t\t\"github.com/google/wire\"\n")
	code += fmt.Sprintf("\t\t\"github.com/samber/lo\"\n")
	code += fmt.Sprintf("\t\t\"golang.org/x/exp/slog\"\n")
	code += fmt.Sprintf("\t\t\"google.golang.org/grpc\"\n")
	code += fmt.Sprintf("\t\t\"google.golang.org/grpc/codes\"\n")
	code += fmt.Sprintf("\t\t\"google.golang.org/grpc/reflection\"\n")
	code += fmt.Sprintf(")\n")
	code += fmt.Sprintf("var _ gen.%s = (*%s)(nil)\n\n",parentServer,structName)
	code += fmt.Sprintf("var %sSet = wire.NewSet(New%s)\n",structName,structName)
	code += fmt.Sprintf("type %s struct {\n", structName)
	code += fmt.Sprintf("\t\tgen.%s\n",unimplemented )
	code += fmt.Sprintf("\t\tuc usecases.UseCase\n" )
	code += fmt.Sprintf("}\n")
	code += fmt.Sprintf("func New%s(\n", structName)
	code += fmt.Sprintf("\t\tgrpcServer *grpc.Server,\n")
	code += fmt.Sprintf("\t\tuc usecases.UseCase,\n")
	code += fmt.Sprintf(")gen.%s{\n",parentServer)
	code += fmt.Sprintf("\tsvc := %s{\n",structName)
	code += fmt.Sprintf(" \t\tuc: uc,\n")
	code += fmt.Sprintf("\t}\n")
	code += fmt.Sprintf("\tgen.Register%s(grpcServer,&svc)\n",parentServer)
	code += fmt.Sprintf("\treflection.Register(grpcServer)\n")
	code += fmt.Sprintf("\treturn &svc\n")
	code += fmt.Sprintf("}\n\n")

	if len(funcs)>0{
		for _,itemFunc:=range funcs{
			code += fmt.Sprintf("%s\n",itemFunc)
		}
	}
	outputPath := fmt.Sprintf("%s.go", structName)


	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(code)
	return err
}
