package main_test

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
	"github.com/stretchr/testify/require"
)

func TestHandleFlags(t *testing.T) {
	tests := []struct {
		args   []string
		result bool
	}{
		{[]string{"app", "-help"}, true},
		{[]string{"app", "-version"}, true},
		{[]string{"app", "-wjat"}, true},
	}

	for _, test := range tests {
		f := ParseFlags(new(bytes.Buffer), test.args)
		require.Equal(t, test.result, HandleFlags(f))
	}
}

func TestName(t *testing.T) {
	name := "group.proto"
	t.Log(filepath.ToSlash(name))
	t.Log(strings.TrimRight(name, filepath.Ext(name)))
	t.Log(filepath.Ext(name))
	t.Log(strings.TrimSuffix(name, filepath.Ext(name)))

}
