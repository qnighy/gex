package manifest_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"

	"github.com/izumin5210/gex/pkg/manifest"
)

func TestParser_Parse(t *testing.T) {
	fs := afero.NewMemMapFs()
	parser := manifest.NewParser(fs)

	var (
		toolsGo = `// Code generated by github.comm/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

import (
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/volatiletech/sqlboiler"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql"
)
`
	)
	path := "/home/src/awesomeapp/tools"

	err := afero.WriteFile(fs, path, []byte(toolsGo), 0644)
	if err != nil {
		t.Fatalf("faield to write %s: %v", path, err)
	}

	out, err := parser.Parse(path)

	if err != nil {
		t.Fatalf("Parse() returned an error: %v", err)
	}

	want := []manifest.Tool{
		manifest.Tool("github.com/gogo/protobuf/protoc-gen-gogofast"),
		manifest.Tool("github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"),
		manifest.Tool("github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"),
		manifest.Tool("github.com/volatiletech/sqlboiler"),
		manifest.Tool("github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql"),
	}

	if diff := cmp.Diff(out.Tools(), want); diff != "" {
		t.Errorf("manifest differs: (-want +got)\n%s", diff)
	}
}
