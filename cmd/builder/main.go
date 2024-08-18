package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type EnumValue struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

type EnumType struct {
	Name  string      `json:"name"`
	Value []EnumValue `json:"value"`
}

type Field struct {
	Name       string `json:"name"`
	Number     int    `json:"number"`
	Label      string `json:"label"`
	Type       string `json:"type"`
	TypeName   string `json:"typeName"`
	OneOfIndex *int   `json:"oneofIndex"`
}

type OneOfDecl struct {
	Name string `json:"name"`
}

type MessageType struct {
	Name       string        `json:"name"`
	Field      []Field       `json:"field"`
	NestedType []MessageType `json:"nestedType"`
	EnumType   []EnumType    `json:"enumType"`
	OneOfDecl  []OneOfDecl   `json:"oneofDecl"`
}

type Descriptor struct {
	Name        string        `json:"name"`
	Package     string        `json:"package"`
	Dependency  []string      `json:"dependency"`
	MessageType []MessageType `json:"messageType"`
	EnumType    []EnumType    `json:"enumType"`
	Syntax      string        `json:"syntax"`
}

func main() {
	var inputDir string
	var outputDir string

	flag.StringVar(&inputDir, "input_dir", "", "Path to the directory containing the JSON files")
	flag.StringVar(&outputDir, "output_dir", "", "Path to the directory where the proto files will be written")
	flag.Parse()

	files, err := os.ReadDir(inputDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fmt.Println(file.Name())
			splits := strings.Split(file.Name(), ".")
			packageName := strings.ToLower(splits[4])
			ProcessDescriptorFile(packageName, filepath.Join(inputDir, file.Name()), outputDir)
		}
	}
}

type Fn func(MessageType, Fn)

func ProcessDescriptorFile(packageName string, filepath string, outputDir string) {
	protoDescriptorJsonFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	protoDescriptor := Descriptor{}
	err = json.NewDecoder(protoDescriptorJsonFile).Decode(&protoDescriptor)
	if err != nil {
		panic(err)
	}

	sanitizedFilename := strings.ReplaceAll(protoDescriptor.Name, "includes/", "")

	//fn := func(messageType MessageType, fn Fn) {
	//	for _, nestedType := range messageType.NestedType {
	//		fn(nestedType, fn)
	//	}
	//
	//	if !strings.HasSuffix(messageType.Name, "Request") &&
	//		!strings.HasSuffix(messageType.Name, "Response") &&
	//		!strings.HasSuffix(messageType.Name, "Event") {
	//		return
	//	}
	//
	//	pkg := strings.Replace(sanitizedFilename, ".proto", "", -1)
	//	fmt.Printf(`r.types["type.ankama.com/%s.%s"] = (&%s.%s{}).ProtoReflect().Type()`, protoDescriptor.Package, messageType.Name, pkg, messageType.Name)
	//	fmt.Println()
	//}
	//
	//for _, messageType := range protoDescriptor.MessageType {
	//	fn(messageType, fn)
	//}
	//
	//return

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(`syntax = "%s";`, protoDescriptor.Syntax))
	sb.WriteString("\n\n")

	sb.WriteString(fmt.Sprintf(`option go_package = "go-xp-dofus-unity-proto-builder/src/protocol/%s/%s";`, packageName, strings.Replace(sanitizedFilename, ".proto", "", 1)))
	sb.WriteString("\n\n")

	if len(protoDescriptor.Dependency) > 0 {
		for _, dependency := range protoDescriptor.Dependency {
			sanitizeDependency := strings.ReplaceAll(dependency, "includes/", "")
			sb.WriteString(fmt.Sprintf(`import "%s";`, sanitizeDependency))
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("package %s;\n\n", protoDescriptor.Package))

	for _, enumType := range protoDescriptor.EnumType {
		WriteEnumType(&sb, enumType)
	}

	for _, messageType := range protoDescriptor.MessageType {
		WriteMessageType(&sb, messageType, 0)
	}

	err = os.MkdirAll(path.Join(outputDir, packageName), os.ModePerm)
	protoFile, err := os.Create(path.Join(outputDir, packageName, sanitizedFilename))
	if err != nil {
		panic(err)
	}

	_, err = protoFile.WriteString(sb.String())
	if err != nil {
		panic(err)
	}
}

func WriteEnumType(sb *strings.Builder, enumType EnumType) {
	sb.WriteString(fmt.Sprintf("enum %s {\n", enumType.Name))
	for _, value := range enumType.Value {
		sb.WriteString(fmt.Sprintf("\t%s = %d;\n", value.Name, value.Number))
	}
	sb.WriteString("}\n\n")
}

func writeIndent(sb *strings.Builder, indent int) {
	for i := 0; i < indent; i++ {
		sb.WriteString("\t")
	}
}

func WriteMessageType(sb *strings.Builder, messageType MessageType, indent int) {
	writeIndent(sb, indent)
	sb.WriteString(fmt.Sprintf("message %s {\n", messageType.Name))

	oneOfGroups := make(map[int][]Field)
	for _, field := range messageType.Field {
		if field.OneOfIndex != nil {
			oneOfGroups[*field.OneOfIndex] = append(oneOfGroups[*field.OneOfIndex], field)
		} else {
			writeField(sb, field, indent+1, false)
		}
	}

	for index, fields := range oneOfGroups {
		oneOfName := messageType.OneOfDecl[index].Name

		if len(fields) == 1 {
			//sb.WriteString(fmt.Sprintf("optional %s %s = %d;\n", fields[0].TypeName, fields[0].Name, fields[0].Number))
			fields[0].Label = "LABEL_OPTIONAL_ONEOF"
			writeField(sb, fields[0], indent+1, false)
			continue
		} else {
			writeIndent(sb, indent+1)
			sb.WriteString(fmt.Sprintf("oneof %s {\n", oneOfName))
			for _, field := range fields {
				writeField(sb, field, indent+2, true)
			}
			writeIndent(sb, indent+1)
			sb.WriteString("}\n")
		}
	}

	for _, nestedType := range messageType.NestedType {
		WriteMessageType(sb, nestedType, indent+1)
	}

	for _, enumType := range messageType.EnumType {
		writeIndent(sb, indent+1)
		sb.WriteString(fmt.Sprintf("enum %s {\n", enumType.Name))
		for _, value := range enumType.Value {
			writeIndent(sb, indent+2)
			sb.WriteString(fmt.Sprintf("%s = %d;\n", value.Name, value.Number))
		}
		writeIndent(sb, indent+1)
		sb.WriteString("}\n\n")
	}

	writeIndent(sb, indent)
	sb.WriteString("}\n\n")
}

func writeField(sb *strings.Builder, field Field, indent int, isOneOf bool) {
	label := ""

	if !isOneOf {
		switch field.Label {
		case "LABEL_OPTIONAL":
			//label = "optional "
		case "LABEL_REQUIRED":
			label = "required "
		case "LABEL_REPEATED":
			label = "repeated "
		case "LABEL_OPTIONAL_ONEOF":
			label = "optional "
		default:
			panic("Unknown label " + field.Label)
		}
	}

	protoType := ""
	switch field.Type {
	case "TYPE_INT32":
		protoType = "int32"
	case "TYPE_INT64":
		protoType = "int64"
	case "TYPE_UINT32":
		protoType = "uint32"
	case "TYPE_UINT64":
		protoType = "uint64"
	case "TYPE_BOOL":
		protoType = "bool"
	case "TYPE_STRING":
		protoType = "string"
	case "TYPE_FLOAT":
		protoType = "float"
	case "TYPE_ENUM":
		protoType = field.TypeName
	case "TYPE_MESSAGE":
		protoType = field.TypeName
	default:
		panic("Unknown type " + field.Type)
	}

	writeIndent(sb, indent)
	sb.WriteString(fmt.Sprintf("%s%s %s = %d;\n", label, protoType, field.Name, field.Number))
}
