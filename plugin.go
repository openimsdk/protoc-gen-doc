package gendoc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

// PluginOptions encapsulates options for the plugin. The type of renderer, template file, and the name of the output
// file are included.
type PluginOptions struct {
	Type            RenderType
	TemplateFile    string
	OutputFile      string
	ExcludePatterns []*regexp.Regexp
	SourceRelative  bool
}

// SupportedFeatures describes a flag setting for supported features.
var SupportedFeatures = uint64(plugin_go.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// Plugin describes a protoc code generate plugin. It's an implementation of Plugin from github.com/pseudomuto/protokit
type Plugin struct{}

// Generate compiles the documentation and generates the CodeGeneratorResponse to send back to protoc. It does this
// by rendering a template based on the options parsed from the CodeGeneratorRequest.
func (p *Plugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	options, err := ParseOptions(r)
	if err != nil {
		return nil, err
	}

	result := excludeUnwantedProtos(protokit.ParseCodeGenRequest(r), options.ExcludePatterns)

	customTemplate := ""

	if options.TemplateFile != "" {
		data, err := ioutil.ReadFile(options.TemplateFile)
		if err != nil {
			return nil, err
		}

		customTemplate = string(data)
	}
	_ = customTemplate
	resp := new(plugin_go.CodeGeneratorResponse)
	fdsGroup := groupProtosByDirectory(result, options.SourceRelative)
	var buf bytes.Buffer
	for dir, fds := range fdsGroup {
		template := NewTemplate(fds)
		// todo 生成文件
		//log.Println("===>", dir, template.Files[0].Name, template.Files[0].Package, template.Files[0].Description)
		//for i, file := range template.Files {
		//	log.Printf("[%d]===> %+v\n", i+1, file.Services[0])
		//	for j, service := range file.Services {
		//		log.Printf("    [%d]===> %+v\n", j+1, service)
		//		for k, method := range service.Methods {
		//			log.Printf("        [%d]===> %+v\n", k+1, method)
		//		}
		//	}
		//}
		msgData := make(map[string]*Message)
		for _, file := range template.Files {
			for _, message := range file.Messages {
				msgData[message.FullName] = message
			}
		}
		for _, file := range template.Files {
			//pkgname := file.Package
			//log.Printf("===> %s %s\n", pkgname, file.Description)
			for _, service := range file.Services {
				for _, method := range service.Methods {
					buf.Reset()
					buf.WriteString("# ")
					buf.WriteString(method.Name)
					buf.WriteString("\n\n")
					buf.WriteString("### Feature Introduction\n\n")
					buf.WriteString("- ")
					buf.WriteString(string(method.Description))
					buf.WriteString("\n\n")
					buf.WriteString("| Method Name | Request Type | Response Type |\n")
					buf.WriteString("| ----------- | ------------ | ------------- |\n")
					buf.WriteString("| ")
					buf.WriteString(method.Name)
					buf.WriteString(" | ")
					writeType(&buf, method.RequestType, method.RequestFullType)
					buf.WriteString(" | ")
					writeType(&buf, method.ResponseType, method.ResponseFullType)
					buf.WriteString(" |\n\n")
					//buf.WriteString("### 请求参数\n")
					buf.WriteString("### ")
					buf.WriteString(method.RequestType)
					buf.WriteString("\n")
					writeMessage(&buf, msgData[method.RequestFullType])
					buf.WriteString("\n\n")
					//buf.WriteString("### 响应参数\n")
					buf.WriteString("### ")
					buf.WriteString(method.ResponseType)
					buf.WriteString("\n")
					writeMessage(&buf, msgData[method.ResponseFullType])
					buf.WriteString("\n\n")
					resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
						Name:    proto.String(filepath.Join(dir, method.Name+".md")),
						Content: proto.String(buf.String()),
					})
				}
			}
		}

		//output, err := RenderTemplate(options.Type, template, customTemplate)
		//if err != nil {
		//	return nil, err
		//}
		//resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		//	Name:    proto.String(filepath.Join(dir, options.OutputFile)),
		//	Content: proto.String(string(output)),
		//})
	}

	resp.SupportedFeatures = proto.Uint64(SupportedFeatures)

	return resp, nil
}

func writeType(buf *bytes.Buffer, name string, full string) {
	buf.WriteString("[")
	buf.WriteString(name)
	buf.WriteString("](#")
	buf.WriteString(full)
	buf.WriteString(")")
}

func writeMessage(buf *bytes.Buffer, msg *Message) {
	//| Field | Type | Label | Description |
	//| ----- | ---- | ----- | ----------- |
	//| groupID | [string](#string) |  | group id you want to transfer |
	//| ownerUserID | [string](#string) |  | new owner user id |
	buf.WriteString("| Field | Type | Label | Description |\n")
	buf.WriteString("| ----- | ---- | ----- | ----------- |\n")
	for _, field := range msg.Fields {
		buf.WriteString("| ")
		buf.WriteString(field.Name)
		buf.WriteString(" | ")
		writeType(buf, field.Type, field.FullType)
		buf.WriteString(" | ")
		buf.WriteString(field.Label)
		buf.WriteString(" | ")
		buf.WriteString(string(field.Description))
		buf.WriteString(" |\n")
	}
}

func groupProtosByDirectory(fds []*protokit.FileDescriptor, sourceRelative bool) map[string][]*protokit.FileDescriptor {
	fdsGroup := make(map[string][]*protokit.FileDescriptor)

	for _, fd := range fds {
		dir := ""
		if sourceRelative {
			dir, _ = filepath.Split(fd.GetName())
		}
		if dir == "" {
			dir = "./"
		}
		fdsGroup[dir] = append(fdsGroup[dir], fd)
	}
	return fdsGroup
}

func excludeUnwantedProtos(fds []*protokit.FileDescriptor, excludePatterns []*regexp.Regexp) []*protokit.FileDescriptor {
	descs := make([]*protokit.FileDescriptor, 0)

OUTER:
	for _, d := range fds {
		for _, p := range excludePatterns {
			if p.MatchString(d.GetName()) {
				continue OUTER
			}
		}

		descs = append(descs, d)
	}

	return descs
}

// ParseOptions parses plugin options from a CodeGeneratorRequest. It does this by splitting the `Parameter` field from
// the request object and parsing out the type of renderer to use and the name of the file to be generated.
//
// The parameter (`--doc_opt`) must be of the format <TYPE|TEMPLATE_FILE>,<OUTPUT_FILE>[,default|source_relative]:<EXCLUDE_PATTERN>,<EXCLUDE_PATTERN>*.
// The file will be written to the directory specified with the `--doc_out` argument to protoc.
func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		Type:           RenderTypeHTML,
		TemplateFile:   "",
		OutputFile:     "index.html",
		SourceRelative: false,
	}

	params := req.GetParameter()
	if strings.Contains(params, ":") {
		// Parse out exclude patterns if any
		parts := strings.Split(params, ":")
		for _, pattern := range strings.Split(parts[1], ",") {
			r, err := regexp.Compile(pattern)
			if err != nil {
				return nil, err
			}
			options.ExcludePatterns = append(options.ExcludePatterns, r)
		}
		// The first part is parsed below
		params = parts[0]
	}
	if params == "" {
		return options, nil
	}

	if !strings.Contains(params, ",") {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	parts := strings.Split(params, ",")
	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	options.TemplateFile = parts[0]
	options.OutputFile = path.Base(parts[1])
	if len(parts) > 2 {
		switch parts[2] {
		case "source_relative":
			options.SourceRelative = true
		case "default":
			options.SourceRelative = false
		default:
			return nil, fmt.Errorf("Invalid parameter: %s", params)
		}
	}
	options.SourceRelative = len(parts) > 2 && parts[2] == "source_relative"

	renderType, err := NewRenderType(options.TemplateFile)
	if err == nil {
		options.Type = renderType
		options.TemplateFile = ""
	}

	return options, nil
}
