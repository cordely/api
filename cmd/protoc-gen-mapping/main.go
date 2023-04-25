package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion = flag.Bool("version", false, "print the version and exit")
	omitempty   = flag.Bool("omitempty", true, "omit if google.api is empty")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-mapping %v\n", release)
		return
	}
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		generateFiles(gen, *omitempty)
		return nil
	})
}
