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
		msgData := make(map[string]*Message)
		for _, file := range template.Files {
			for _, message := range file.Messages {
				msgData[message.FullName] = message
			}
		}
		for _, file := range template.Files {
			name := strings.TrimSuffix(file.Name, filepath.Ext(file.Name))
			for _, service := range file.Services {
				for i, method := range service.Methods {
					_, desc := splitDescription(string(method.Description))
					buf.Reset()
					writeHeader(&buf, 2+i, method.Name, true)
					buf.WriteString("# ")
					buf.WriteString(method.Name)
					buf.WriteString("\n\n")
					buf.WriteString("### Feature Introduction\n\n")
					for _, d := range desc {
						buf.WriteString(fmt.Sprintf("- %s\n", d))
					}
					buf.WriteString("\n\n")
					writeMethod(&buf, false, method)
					buf.WriteString(" \n\n")
					buf.WriteString(fmt.Sprintf("#### %s\n", method.RequestType))
					writeMessage(&buf, msgData[method.RequestFullType])
					buf.WriteString(" \n\n")
					buf.WriteString(fmt.Sprintf("#### %s\n", method.ResponseType))
					writeMessage(&buf, msgData[method.ResponseFullType])
					buf.WriteString("\n\n")
					resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
						Name:    proto.String(filepath.Join(dir, name, filename(method.Name))),
						Content: proto.String(buf.String()),
					})
				}
			}
			buf.Reset()
			writeHeader(&buf, 1, name, false)
			if file.HasServices {
				buf.WriteString(fmt.Sprintf("# Api\n\n"))
				for _, service := range file.Services {
					writeMethod(&buf, true, service.Methods...)
				}
			}
			buf.WriteString("\n\n")
			if file.HasEnums {
				buf.WriteString("## Enums\n\n")
				for _, enum := range file.Enums {
					writeEnum(&buf, enum)
					buf.WriteString("\n\n")
				}
			}

			if file.HasMessages {
				buf.WriteString("## Messages\n\n")
				for _, message := range file.Messages {
					buf.WriteString(fmt.Sprintf(`<a name="%s"></a>`, message.Name))
					buf.WriteString(fmt.Sprintf("\n### %s\n\n", message.Name))
					writeMessage(&buf, message)
					buf.WriteString("\n\n")
				}
				resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
					Name:    proto.String(filepath.Join(dir, name, filename("index"))),
					Content: proto.String(buf.String()),
				})
			}

		}
	}
	resp.SupportedFeatures = proto.Uint64(SupportedFeatures)
	return resp, nil
}

func filename(name string) string {
	return name + ".mdx"
}

func linkpath(mode string, name string, anchor string) string {
	str := "docs/" + mode + "/" + name + ".mdx"
	if anchor != "" {
		str += "#" + anchor
	}
	return str
}

func mdlink(name string, link string) string {
	return fmt.Sprintf("[%s](%s)", name, link)
}

func typeLink(name string) string {
	arr := strings.Split(name, ".")
	if len(arr) > 1 && arr[0] == "openim" {
		lastName := arr[len(arr)-1]
		return fmt.Sprintf("[%s](%s)", lastName, linkpath(arr[len(arr)-2], "index", lastName))
	}
	if len(arr) > 0 && arr[0] == "google" {
		return fmt.Sprintf("[%s](%s)", arr[len(arr)-1], linkpath("google", strings.ToLower(arr[len(arr)-1]), ""))
	}
	return fmt.Sprintf("[%s](%s)", arr[len(arr)-1], linkpath("scalar", "scalar", ""))
}

func methodLink(mode, name string) string {
	return fmt.Sprintf("[%s](%s)", name, linkpath(mode, name, ""))
}

func splitDescription(desc string) (string, []string) {
	res := strings.Split(desc, "\n")
	if len(res) == 0 {
		return "", nil
	}
	return res[0], res[1:]
}

func writeHeader(buf *bytes.Buffer, position int, title string, hide bool) {
	buf.WriteString("---\n")
	buf.WriteString(fmt.Sprintf("sidebar_position: %d\n", position))
	buf.WriteString(fmt.Sprintf("title: %s\n", title))
	buf.WriteString(fmt.Sprintf("hide_title: %+v\n", hide))
	buf.WriteString("---\n\n")
}

func writeEnum(buf *bytes.Buffer, enum *Enum) {
	buf.WriteString("| Name | Number | Description |\n")
	buf.WriteString("| ---- | ------ | ----------- |\n")
	for _, value := range enum.Values {
		buf.WriteString(fmt.Sprintf("| %s | %s | %s |\n", value.Name, value.Number, strings.ReplaceAll(string(value.Description), "\n", "<br>")))
	}
}

func writeMessage(buf *bytes.Buffer, msg *Message) {
	buf.WriteString("| Field | Type | Label | Description |\n")
	buf.WriteString("| ----- | ---- | ----- | ----------- |\n")
	for _, field := range msg.Fields {
		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", field.Name, typeLink(field.FullType), field.Label, strings.ReplaceAll(string(field.Description), "\n", "<br>")))
	}
}

func writeMethod(buf *bytes.Buffer, link bool, methods ...*ServiceMethod) {
	buf.WriteString("| Method Name | Request Type | Response Type | Description |\n")
	buf.WriteString("| ----------- | ------------ | ------------- | ------------- |\n")
	for _, method := range methods {
		var mode string
		if arr := strings.Split(method.RequestFullType, "."); len(arr) > 2 {
			mode = arr[len(arr)-2]
		}
		simple, _ := splitDescription(string(method.Description))
		name := method.Name
		if link {
			name = methodLink(mode, method.Name)
		}
		req := typeLink(method.RequestFullType)
		resp := typeLink(method.ResponseFullType)
		buf.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", name, req, resp, simple))
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
