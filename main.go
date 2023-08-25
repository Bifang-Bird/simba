/*
*

	@author: junwang
	@since: 2023/8/24
	@desc: //TODO

*
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/Bifang-Bird/simba/generator"
)

func main() {

	methodName := flag.String("m", "", "Method name")
	dstDir := flag.String("d", "", "Dst name")
	packageName := flag.String("p", "", "Package name")
	structName := flag.String("s", "", "Struct name")
	grpcgo := flag.String("g", "", "grpcgo name")
	flag.Parse()
	fmt.Println("loading file:", *grpcgo)

	//filePath := "/Users/junwang/go/pkg/mod/codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-proto.git@v0.0.192/gen/sms_grpc.pb.go"
	//_, gprcFile, erre := handle.LoadFile(filePath)
	//if erre != nil {
	//	fmt.Println("Error loading file:", erre)
	//}
	//structs,interfaces := handle.FindStructs(gprcFile)
	//fmt.Println("Structs in the file:", structs)
	//fmt.Println("Interface in the file:", interfaces)
	switch *methodName {
	case "grpc":
		generateGrpcServiceImpl(packageName,structName,grpcgo)
	case "init":
		load(dstDir)
	default:
		fmt.Println("未知的方法名称")
	}
}
func load(dstDir *string){
	//srcDir := "./simbaproject"
	//err := os.MkdirAll(*dstDir, 0755)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = generator.CopyDirectory(srcDir, *dstDir)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("项目初始化成功")
	repoURL := "https://github.com/Bifang-Bird/lexington.git"
	err:=generator.CloneProject(repoURL,*dstDir)
	if err != nil {
		log.Fatal(err)
	}
}

func generateGrpcServiceImpl(packageName,structName,grpcgo *string){
	if *packageName == "" || *structName == "" {
		fmt.Println("Usage: mytool -package <package-name> -struct <struct-name>")
		os.Exit(1)
	}
	err := generator.GenerateCode(*packageName, *structName, *grpcgo)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated %s.go\n", *structName)
}