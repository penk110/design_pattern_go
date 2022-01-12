package proxy

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"strings"
)

/*
	ast 抽象语法树
	参考：
		https://github.com/chai2010/go-ast-book

*/

type proxyData struct {
	Package         string         // package name
	ProxyStructName string         // class name
	Methods         []*proxyMethod // method name

}

type proxyMethod struct {
	Name        string // name
	Params      string // 参数，含参数类型
	ParamNames  string // 参数名
	Results     string
	ResultNames string
}

func generate(file string) (string, error) {
	fileSet := token.NewFileSet() // creates a new file set.
	f, err := parser.ParseFile(fileSet, file, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// 获取代理需要的数据
	data := proxyData{
		Package: f.Name.Name,
	}

	// 构建注释和 node 关系
	commentMap := ast.NewCommentMap(fileSet, f, f.Comments)

	for node, group := range commentMap {
		// 根据 @proxy 接口名 匹配和切分出接口名称
		name := getProxyInterfaceName(group)
		if name == "" {
			continue
		}

		// 代理类名称
		data.ProxyStructName = node.(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name.Name
		// 从文件中查找接口
		objects := f.Scope.Lookup(name)

		// 类型转换，没有做断言，可能会出现panic
		t := objects.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType)
		for _, field := range t.Methods.List {
			funcName := field.Type.(*ast.FuncType)
			method := &proxyMethod{
				Name: field.Names[0].Name,
			}
			// 获取方法的参数和返回值
			method.Params, method.ParamNames = getParamsOrResults(funcName.Params)
			method.Results, method.ResultNames = getParamsOrResults(funcName.Results)

			data.Methods = append(data.Methods, method)
		}
	}

	// 生成模板
	temp, err := template.New("").Parse(proxyTpl)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	buf := &bytes.Buffer{}
	if err := temp.Execute(buf, data); err != nil {
		return "", err
	}

	// 使用 go fmt 格式化代码
	source, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}

	return string(source), nil

}

func getProxyInterfaceName(groups []*ast.CommentGroup) string {
	for _, commentGroup := range groups {
		for _, comment := range commentGroup.List {
			if strings.Contains(comment.Text, "// @proxy ") {
				interfaceName := strings.TrimLeft(comment.Text, "// @proxy ")
				return strings.TrimSpace(interfaceName)
			}

		}
	}

	return ""
}

func getParamsOrResults(fieldList *ast.FieldList) (string, string) {
	var params []string
	var paramNames []string

	for i, param := range fieldList.List {
		var names []string
		for _, name := range param.Names {
			names = append(names, name.Name)
		}

		// TODO: ??
		if len(names) == 0 {
			names = append(names, fmt.Sprintf("r%d", i))
		}

		paramNames = append(paramNames, names...)
		// 参数名加参数类型组成完整的参数
		param := fmt.Sprintf("%s %s",
			strings.Join(names, ","),
			param.Type.(*ast.Ident).Name,
		)

		params = append(params, strings.TrimSpace(param))

	}

	return strings.Join(params, ","), strings.Join(paramNames, ",")
}

// 生成代理类的文件模板
const proxyTpl = `
package {{.Package}}
type {{ .ProxyStructName }}Proxy struct {
	child *{{ .ProxyStructName }}
}
func New{{ .ProxyStructName }}Proxy(child *{{ .ProxyStructName }}) *{{ .ProxyStructName }}Proxy {
	return &{{ .ProxyStructName }}Proxy{child: child}
}
{{ range .Methods }}
func (p *{{$.ProxyStructName}}Proxy) {{ .Name }} ({{ .Params }}) ({{ .Results }}) {
	start := time.Now()
	{{ .ResultNames }} = p.child.{{ .Name }}({{ .ParamNames }})
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return {{ .ResultNames }}
}
{{ end }}
`
