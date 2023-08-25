/*
*

	@author: junwang
	@since: 2023/8/24
	@desc: //TODO

*
*/
package handle

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Mystruct struct {
	Id string
	Nid string
}

func LoadFile(filePath string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	return fset, file, nil
}

func FindStructs(file *ast.File,structName string)([]string ,map[string][]string) {
	structs := []string{}
	interfaces := map[string][]string{}
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					_, ok1 := typeSpec.Type.(*ast.StructType)
					_, okk := typeSpec.Type.(*ast.InterfaceType)
					if ok1 {
						structs = append(structs, typeSpec.Name.Name)
					}
					if okk{
						funcs := findFuncs(typeSpec,structName)
						interfaces[typeSpec.Name.Name] = funcs
					}
				}
			}
		}
		//fset := token.NewFileSet()
		// 遍历文件的声明，查找类型声明
		//for _, decl := range file.Decls {
		//	if genDecl, ok := decl.(*ast.GenDecl); ok {
		//		for _, spec := range genDecl.Specs {
		//			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
		//				// 检查是否为目标类型
		//				if typeSpec.Name.Name == "SmsServiceServer" {
		//					// 检查是否为结构体类型
		//					if structType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
		//						fmt.Printf("接口 \"%s\" 的方法：\n", "SmsServiceServer")
		//						// 遍历结构体字段
		//						for _, field := range structType.Methods.List {
		//							// 检查是否为方法
		//							if funcType, ok := field.Type.(*ast.FuncType); ok {
		//								fmt.Printf("  方法名：%s\n", field.Names[0].Name)
		//								// 处理方法参数
		//								var params []string
		//								var results []string
		//
		//								if funcType.Params != nil {
		//									fmt.Println("    参数：")
		//									for _, param := range funcType.Params.List {
		//										// 获取参数类型
		//										if fieldType, ok := param.Type.(*ast.StarExpr); ok {
		//											// 参数类型是指针类型
		//											actualType := fieldType.X
		//											fmt.Printf("参数 %s 的类型是指针类型，指向的实际类型：%s\n", param.Names, actualType)
		//											paramType := getTypeName(param.Type)
		//											params = append(params, fmt.Sprintf("%s %s", strings.ToLower(strings.TrimPrefix(paramType,"*")),paramType))
		//										} else {
		//											// 参数类型不是指针类型
		//											paramType := getTypeName(param.Type)
		//
		//											fmt.Printf("参数 %s 的类型：%s\n", "", paramType)
		//											params = append(params, fmt.Sprintf("%s %s",  paramType,""))
		//										}
		//
		//									}
		//								}
		//
		//								// 处理方法返回值
		//								if funcType.Results != nil {
		//									fmt.Println("    返回值：")
		//									for _, result := range funcType.Results.List {
		//										fmt.Println("      ", result.Type)
		//
		//										if fieldType, ok := result.Type.(*ast.StarExpr); ok {
		//											// 参数类型是指针类型
		//											actualType := fieldType.X
		//											resultType := getTypeName(result.Type)
		//											fmt.Printf(" %s 的类型是指针类型，指向的实际类型：%s\n", result.Names, actualType)
		//											results = append(results, fmt.Sprintf("%s%s", "",resultType))
		//										} else {
		//											// 参数类型不是指针类型
		//											resultType := getTypeName(result.Type)
		//											fmt.Printf("参数 %s 的类型：%s\n", "", result.Type)
		//											results = append(results, fmt.Sprintf("%s %s", "",resultType))
		//										}
		//
		//
		//									}
		//								}
		//								newFuncStr := fmt.Sprintf("func (%s) %s(%s) (%s) {\n    // TODO: Implement the new function\n\n}", "g *smsGRPCServer",field.Names[0].Name, strings.Join(params, ", "), strings.Join(results, ", "))
		//								fmt.Println("新的函数代码：")
		//								fmt.Println(newFuncStr)
		//							}
		//						}
		//
		//
		//					}
		//
		//				}
		//			}
		//		}
		//	}
		//}
}

	return structs,interfaces
}

//
// findFuncs
//  @Description: SmsServiceServer
//  @param interfaceName
//  @return []string
//
func findFuncs(typeSpec *ast.TypeSpec,structName string)[]string{
	var funcs =[]string{}
		// 检查是否为结构体类型
	if structType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
		fmt.Printf("接口 \"%s\" 的方法：\n", typeSpec.Name.Name)
		// 遍历结构体字段
		for _, field := range structType.Methods.List {
			// 检查是否为方法
			if funcType, ok := field.Type.(*ast.FuncType); ok {
				fmt.Printf("  方法名：%s\n", field.Names[0].Name)
				// 处理方法参数
				var params []string
				var results []string

				if funcType.Params != nil {
					fmt.Println("    参数：")
					for _, param := range funcType.Params.List {
						// 获取参数类型
						if fieldType, ok := param.Type.(*ast.StarExpr); ok {
							// 参数类型是指针类型
							actualType := fieldType.X
							fmt.Printf("参数 %s 的类型是指针类型，指向的实际类型：%s\n", param.Names, actualType)
							paramType := getTypeName(param.Type)
							paramName := strings.Replace(strings.ToLower(strings.TrimPrefix(paramType,"*")),".","",-1)
							params = append(params, fmt.Sprintf("%s %s", paramName,paramType))
						} else {
							// 参数类型不是指针类型
							paramType := getTypeName(param.Type)

							fmt.Printf("参数 %s 的类型：%s\n", "", paramType)
							params = append(params, fmt.Sprintf("%s %s",  paramType,""))
						}

					}
				}

				// 处理方法返回值
				if funcType.Results != nil {
					fmt.Println("    返回值：")
					for _, result := range funcType.Results.List {
						fmt.Println("      ", result.Type)

						if fieldType, ok := result.Type.(*ast.StarExpr); ok {
							// 参数类型是指针类型
							actualType := fieldType.X
							resultType := getTypeName(result.Type)
							fmt.Printf(" %s 的类型是指针类型，指向的实际类型：%s\n", result.Names, actualType)
							results = append(results, fmt.Sprintf("%s%s", "",resultType))
						} else {
							// 参数类型不是指针类型
							resultType := getTypeName(result.Type)
							fmt.Printf("参数 %s 的类型：%s\n", "", result.Type)
							results = append(results, fmt.Sprintf("%s %s", "",resultType))
						}


					}
				}
				newFuncStr := fmt.Sprintf("func (%s) %s(%s) (%s) {\n    // TODO: Implement the new function\n\n}\n", "g *"+structName ,field.Names[0].Name, strings.Join(params, ", "), strings.Join(results, ", "))
				fmt.Println("新的函数代码：")
				fmt.Println(newFuncStr)
				funcs = append(funcs,newFuncStr)
			}
		}

	}
	return funcs
}



func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		// 基本类型
		return t.Name
	case *ast.StarExpr:
		// 指针类型
		return "*gen." + getTypeName(t.X)
	case *ast.ArrayType:
		// 数组类型
		return "[]" + getTypeName(t.Elt)
	case *ast.MapType:
		// 映射类型
		return fmt.Sprintf("map[%s]%s", getTypeName(t.Key), getTypeName(t.Value))
	case *ast.StructType:
		// 结构体类型
		return "struct"
	default:
		return "unknow"
	}
}
