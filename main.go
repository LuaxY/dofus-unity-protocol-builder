package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type EnumField struct {
	Name  string
	Type  string
	Index int
}

type Enum struct {
	Namespace string
	Name      string
	Fields    []EnumField
	IsOneOf   bool
}

type MessageSubField struct {
	Name  string
	Type  string
	Index int
}

type MessageField struct {
	Name       string
	Type       string
	Index      int
	IsOneOf    bool
	IsRepeated bool
	SubFields  []MessageSubField
}

type Message struct {
	Namespace string
	Imports   map[string]any
	Name      string
	Fields    []MessageField
}

const defaultNamespace = "core"

func main() {
	export("com.ankama.dofus.server.connection", "output/connection.proto")
	export("com.ankama.dofus.server.game", "output/game.proto")
}

func export(targetNamespace string, outputFile string) {
	os.Mkdir("output", os.ModePerm)
	os.Remove(outputFile)

	protoFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}

	_, _ = protoFile.WriteString(`syntax = "proto3";`)
	_, _ = protoFile.WriteString("\n\n")
	_, _ = protoFile.WriteString(`import "google/protobuf/any.proto";`)
	_, _ = protoFile.WriteString("\n\n")

	file, err := os.Open("dump.cs")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var reEnum = regexp.MustCompile(`(?m)// Namespace: (.*)\npublic enum ([A-Za-z.]+).*\n{((?s).*?)}`)
	var reEnumField = regexp.MustCompile(`(?m)public const ([A-Za-z.]+) (\w+) = (\d+);`)

	listEnums := map[string]Enum{}

	enums := reEnum.FindAllStringSubmatch(string(data), -1)
	for _, enum := range enums {
		namespace := enum[1]
		name := enum[2]
		name = strings.Replace(name, ".", "", -1)
		code := enum[3]
		isOneOf := strings.Contains(name, "OneofCase")

		if namespace == "" {
			namespace = defaultNamespace
		}

		e := Enum{
			Namespace: namespace,
			Name:      name,
			IsOneOf:   isOneOf,
		}

		fields := reEnumField.FindAllStringSubmatch(code, -1)
		for _, field := range fields {
			fieldType := field[1]
			fieldName := field[2]
			fieldIndex, _ := strconv.Atoi(field[3])

			e.Fields = append(e.Fields, EnumField{
				Name:  fieldName,
				Type:  fieldType,
				Index: fieldIndex,
			})
		}

		listEnums[name] = e
	}

	// Remove enums that don't have index 0 filed
	for name, e := range listEnums {
		if e.IsOneOf {
			continue
		}

		hasIndex0 := false
		for _, field := range e.Fields {
			if field.Index == 0 {
				hasIndex0 = true
				break
			}
		}

		if !hasIndex0 {
			delete(listEnums, name)
		}
	}

	var reMessage = regexp.MustCompile(`(?m)// Namespace: (.*)\n.*\npublic sealed class (.*) : IMessage.*\n\{((?s).*?)}\n\n//`)
	var reFields = regexp.MustCompile(`// Fields\n((?s).*?)// (Properties|Methods)`)
	var reProperties = regexp.MustCompile(`// Properties\n((?s).*?)// (Methods|Fields)`)
	var reFieldIndex = regexp.MustCompile(`(?m)public const int (.*)FieldNumber = (\d+);`)
	var reField = regexp.MustCompile(`(?m)private (readonly )?([A-za-z.<>]+) ([A-za-z.]+)_;`)
	var reProperty = regexp.MustCompile(`(?m)public ([A-Za-z.]+) (\w+) { get; set; }`)

	listMessages := map[string]*Message{}

	matches := reMessage.FindAllStringSubmatch(string(data), -1)
	for _, match := range matches {
		namespace := match[1]
		name := match[2]
		name = strings.Replace(name, ".", "", -1)
		code := match[3]

		if namespace == "" {
			namespace = defaultNamespace
		}

		m := &Message{
			Namespace: namespace,
			Name:      name,
			Imports:   map[string]any{},
		}

		mapFieldIndex := make(map[string]string)
		fieldsCode := reFields.FindStringSubmatch(code)
		fieldsIndex := reFieldIndex.FindAllStringSubmatch(fieldsCode[1], -1)
		for _, fieldIndex := range fieldsIndex {
			field := fieldIndex[1]
			field = strings.Replace(field, "_", "", -1)
			index := fieldIndex[2]

			mapFieldIndex[field] = index
		}

		mapPropertyType := make(map[string]string)
		propertiesCode := reProperties.FindStringSubmatch(code)
		properties := reProperty.FindAllStringSubmatch(propertiesCode[1], -1)
		for _, info := range properties {
			typeName := info[1]
			fieldName := info[2]

			mapPropertyType[fieldName] = typeName
		}

		fieldsInfo := reField.FindAllStringSubmatch(fieldsCode[1], -1)
		for _, field := range fieldsInfo {
			typeName := field[2]
			typeName = strings.Replace(typeName, ".", "", -1)
			fieldName := field[3]
			isOneOf := strings.Contains(typeName, "OneofCase")
			isRepeated := strings.Contains(typeName, "RepeatedField")

			if isOneOf {
				fieldName = strings.TrimSuffix(fieldName, "Case")
			}

			if isRepeated {
				typeName = strings.TrimPrefix(typeName, "RepeatedField<")
				typeName = strings.TrimSuffix(typeName, ">")
			}

			mf := MessageField{
				Name:       fieldName,
				Type:       CSharpToProtoType(typeName),
				IsOneOf:    isOneOf,
				IsRepeated: isRepeated,
			}

			if typeName == "object" {
				// TODO: check if it's safe to ignore
				// ignore
				continue
			}

			if isOneOf {
				enum, found := listEnums[typeName]
				if !found {
					fmt.Println("ENUM NOT FOUND", typeName)
					panic("ENUM NOT FOUND")
				} else {
					for _, field := range enum.Fields {
						if field.Index == 0 {
							// ignore None
							continue
						}

						propertyType, found := mapPropertyType[field.Name]
						if !found {
							fmt.Println("PROPERTY NOT FOUND", field.Name)
							panic("PROPERTY NOT FOUND")
						}

						field.Name = strings.ToLower(field.Name[0:1]) + field.Name[1:]

						mf.SubFields = append(mf.SubFields, MessageSubField{
							Name:  field.Name,
							Type:  CSharpToProtoType(propertyType),
							Index: field.Index,
						})
					}
				}
			} else {
				index, found := mapFieldIndex[strings.ToUpper(fieldName[0:1])+fieldName[1:]]
				if !found {
					fmt.Println("INDEX NOT FOUND", fieldName)
					panic("INDEX NOT FOUND")
				}

				mf.Index, _ = strconv.Atoi(index)
			}

			m.Fields = append(m.Fields, mf)
		}

		listMessages[name] = m
	}

	mapMessageNamespace := make(map[string][]*Message)

	for _, m := range listMessages {
		mapMessageNamespace[m.Namespace] = append(mapMessageNamespace[m.Namespace], m)
	}

	messagesToAdd := map[string]*Message{}
	enumsToAdd := map[string]Enum{}

	toAdd := map[string]any{}

	for namespace, listMessages := range mapMessageNamespace {
		namespace = strings.ToLower(namespace)
		if !strings.Contains(namespace, targetNamespace) {
			continue
		}

		//fmt.Println(len(listMessages), namespace)

		for _, m := range listMessages {
			for _, field := range m.Fields {
				if field.IsOneOf {
					for _, subField := range field.SubFields {
						toAdd[subField.Type] = nil
					}
				} else {
					toAdd[field.Type] = nil
				}
			}

			messagesToAdd[m.Name] = m
		}
	}

	added := len(toAdd)
	alreadyImport := map[string]any{}

	for {
		if added == 0 {
			break
		}

		added = 0

		newAdd := map[string]any{}

		for typeName := range toAdd {
			if !isCustomProtoType(typeName) {
				continue
			}

			if _, found := alreadyImport[typeName]; found {
				continue
			}

			alreadyImport[typeName] = nil

			message, found := listMessages[typeName]
			if !found {
				enum, found := listEnums[typeName]
				if !found {
					fmt.Println("TYPE OR ENUM NOT FOUND", typeName)
					panic("TYPE NOT FOUND")
				} else {
					//fmt.Println("ENUM", enum.Name)
					enumsToAdd[enum.Name] = enum
				}
			} else {
				//fmt.Println("MESSAGE", message.Name)
				messagesToAdd[message.Name] = message

				for _, field := range message.Fields {
					if field.IsOneOf {
						for _, subField := range field.SubFields {
							newAdd[subField.Type] = nil
							added += 1
						}
					} else {
						newAdd[field.Type] = nil
						added += 1
					}
				}
			}
		}

		toAdd = newAdd
	}

	for _, m := range messagesToAdd {
		//fmt.Println("MESSAGE", m.Name)

		_, _ = protoFile.WriteString(fmt.Sprintf("message %s {\n", m.Name))

		for _, field := range m.Fields {
			if field.IsOneOf {
				_, _ = protoFile.WriteString(fmt.Sprintf("\toneof %s {\n", field.Name))

				for _, subField := range field.SubFields {
					_, _ = protoFile.WriteString(fmt.Sprintf("\t\t%s %s = %d;\n", subField.Type, subField.Name, subField.Index))
				}

				_, _ = protoFile.WriteString("\t}\n")
			} else if field.IsRepeated {
				_, _ = protoFile.WriteString(fmt.Sprintf("\trepeated %s %s = %d;\n", field.Type, field.Name, field.Index))
			} else {
				_, _ = protoFile.WriteString(fmt.Sprintf("\t%s %s = %d;\n", field.Type, field.Name, field.Index))
			}
		}

		_, _ = protoFile.WriteString("}\n\n")
	}

	for _, e := range enumsToAdd {
		//fmt.Println("ENUM", e.Name)

		if e.IsOneOf {
			continue
		}

		_, _ = protoFile.WriteString(fmt.Sprintf("enum %s {\n", e.Name))

		for _, field := range e.Fields {
			name := PascalToSnake(e.Name + field.Name)
			_, _ = protoFile.WriteString(fmt.Sprintf("\t%s = %d;\n", name, field.Index))

		}

		_, _ = protoFile.WriteString("}\n\n")
	}
}

func CSharpToProtoType(csharpType string) string {
	switch csharpType {
	case "bool":
		return "bool"
	case "byte":
		return "uint32"
	case "sbyte":
		return "int32"
	case "short":
		return "int32"
	case "ushort":
		return "uint32"
	case "int":
		return "int32"
	case "uint":
		return "uint32"
	case "long":
		return "int64"
	case "ulong":
		return "uint64"
	case "float":
		return "float"
	case "double":
		return "double"
	case "string":
		return "string"
	case "object":
		return "object"
	case "Any":
		return "google.protobuf.Any"
	default:
		return strings.Replace(csharpType, ".", "", -1)
	}
}

func isCustomProtoType(protoType string) bool {
	list := []string{
		"bool",
		"uint32",
		"int32",
		"int64",
		"uint64",
		"float",
		"double",
		"string",
		"object",
	}

	for _, t := range list {
		if t == protoType {
			return false
		}
	}

	return true
}

func PascalToSnake(s string) string {
	// Regular expression to find places to insert underscores
	re := regexp.MustCompile("([a-z])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")

	// Convert to upper case
	return strings.ToUpper(snake)
}
