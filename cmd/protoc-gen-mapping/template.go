package main

import (
	"bytes"
	"html/template"
	"strings"
)

var httpTemplate = `
{{$svrName := .ServiceName}}
{{range .MethodSets}}
{{if .HasBody -}}
{method: "{{.Method}}", body: true, path: "{{.Path}}", operation: "/{{$svrName}}/{{.OriginalName}}", reqMessageFunc: func() proto.Message {return new({{.Request}})}, replyMessageFunc: func() proto.Message {return new({{.Reply}})}},
{{else -}}
{method: "{{.Method}}", body: false, path: "{{.Path}}", operation: "/{{$svrName}}/{{.OriginalName}}", reqMessageFunc: func() proto.Message {return new({{.Request}})}, replyMessageFunc: func() proto.Message {return new({{.Reply}})}},
{{end -}}
{{end}}
`

type serviceDesc struct {
	ServiceType string
	ServiceName string
	Metadata    string
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	Name         string
	OriginalName string
	Num          int
	Request      string
	Reply        string
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.ReplaceAll(strings.Trim(buf.String(), "\r\n\n"), ",\n\n", ",\n")
}
