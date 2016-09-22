// generate_cpp.go
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GetInitialValueCpp(varType string) string {
	switch varType {
	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64":
		return "0"
	case "string":
		return "\"\""
	}
	return "\"\""
}

func ConvertVarTypeCpp(varType string) string {
	switch varType {
	case "int8":
		return "int8_t"
	case "uint8":
		return "uint8_t"
	case "int16":
		return "int16_t"
	case "uint16":
		return "uint16_t"
	case "int32":
		return "int32_t"
	case "uint32":
		return "uint32_t"
	case "int64":
		return "int64_t"
	case "uint64":
		return "uint64_t"
	case "string":
		return "std::string"
	}
	return varType
}

func WriteMessageSizeCpp(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "size += sizeof(uint16_t);\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "size += %v.length();\n", varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "size += sizeof(%v);\n", ConvertVarTypeCpp(varType))
}

func WriteMessageEncodeCpp(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "size_t %v_size = %v.length();\n", varName, varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "std::memcpy(*buf, &%v_size, sizeof(uint16_t));\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "(*buf) += sizeof(uint16_t);\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "std::memcpy(*buf, %v.c_str(), %v_size);\n", varName, varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "(*buf) += %v_size;\n", varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "std::memcpy(*buf, &%v, sizeof(%v));\n", varName, ConvertVarTypeCpp(varType))
	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "(*buf) += sizeof(%v);\n", ConvertVarTypeCpp(varType))
}

func WriteMessageDecodeCpp(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "if (sizeof(uint16_t) > size) { return false; }\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "uint16_t %v_len = 0;\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "std::memcpy(&%v_len, *buf, sizeof(uint16_t));\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "(*buf) += sizeof(uint16_t); size -= sizeof(uint16_t);\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "if (size < %v_len) { return false; }\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "%v.assign((char*)*buf, %v_len);\n", varName, varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "(*buf) += %v_len; size -= %v_len;\n", varName, varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "if (sizeof(%v) > size) { return false; }\n", ConvertVarTypeCpp(varType))
	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "std::memcpy(&%v, *buf, sizeof(%v));\n", varName, ConvertVarTypeCpp(varType))
	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "(*buf) += sizeof(%v); size -= sizeof(%v);\n", ConvertVarTypeCpp(varType), ConvertVarTypeCpp(varType))
}

func GenerateCppCode(el *TokenElement) {
	fmt.Fprintf(os.Stdout, "struct %v {\n", el.Name)

	if el.IsMsg {
		fmt.Fprintf(os.Stdout, "	enum { MSG_ID = %v };\n", el.Id)
	}

	// Declare Variable
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			fmt.Fprintf(os.Stdout, "	%v %v;\n", ConvertVarTypeCpp(varType), varName)
			continue
		}

		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)
			fmt.Fprintf(os.Stdout, "	std::list<%v> %v;\n", ConvertVarTypeCpp(arrVarType), varName)
			continue
		}

		fmt.Fprintf(os.Stdout, "	%v %v;\n", ConvertVarTypeCpp(varType), varName)
	}

	// Constructor
	fmt.Fprintf(os.Stdout, "	%v() {\n", el.Name)
	fmt.Fprintf(os.Stdout, "		Clear();\n")
	fmt.Fprintf(os.Stdout, "	}\n\n")

	// Clear
	fmt.Fprintf(os.Stdout, "	void Clear() {\n")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			fmt.Fprintf(os.Stdout, "		%v = %v;\n", varName, GetInitialValueCpp(varType))
			continue
		}

		if isArrayType(varType) {
			fmt.Fprintf(os.Stdout, "		%v.clear();\n", varName)
			continue
		}

		fmt.Fprintf(os.Stdout, "		%v.Clear();\n", varName)
	}
	fmt.Fprintf(os.Stdout, "	} // end Clear()\n\n")

	// Size
	fmt.Fprintf(os.Stdout, "	int32_t Size() const {\n")
	fmt.Fprintf(os.Stdout, "		int32_t size = 0;\n")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			WriteMessageSizeCpp(varType, varName, TAB2)
			continue
		}

		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprintf(os.Stdout, "		size += sizeof(uint16_t);\n")
			if isPrimitiveType(arrVarType) {
				fmt.Fprintf(os.Stdout, "		for(std::list<%v>::const_iterator iter = %v.begin(); iter != %v.end(); ++iter) {\n", ConvertVarTypeCpp(arrVarType), varName, varName)
				//fmt.Fprintf(os.Stdout, "			const %v & val = *iter;\n", ConvertVarTypeCpp(arrVarType))
				WriteMessageSizeCpp(arrVarType, "(*iter)", TAB3)
				fmt.Fprintf(os.Stdout, "		}\n")
			} else {
				fmt.Fprintf(os.Stdout, "		for(std::list<%v>::const_iterator iter = %v.begin(); iter != %v.end(); ++iter) {\n", arrVarType, varName, varName)
				//fmt.Fprintf(os.Stdout, "			const %v & o = *iter;\n", arrVarType)
				fmt.Fprintf(os.Stdout, "			size += %v_Serializer::Size(*iter);\n", arrVarType)
				fmt.Fprintf(os.Stdout, "		}\n")
			}
			continue
		}

		// User Define Struct
		fmt.Fprintf(os.Stdout, "		size += %v_Serializer::Size(%v);\n", varType, varName)
	}
	fmt.Fprintf(os.Stdout, "		return size;\n")
	fmt.Fprintf(os.Stdout, "	} // end Size()\n\n")

	// Encode
	fmt.Fprintf(os.Stdout, "	bool Encode(std::vector<char>& buf) const {\n")
	fmt.Fprintln(os.Stdout, "		size_t size = Size();")
	fmt.Fprintln(os.Stdout, "		if (0 == size) { return true; }")
	fmt.Fprintln(os.Stdout, "		if (size > buf.size()) { buf.resize(size); }")
	fmt.Fprintln(os.Stdout, "		char * pBuf = &(buf[0]);")
	fmt.Fprintln(os.Stdout, "		if (false == Encode(&pBuf)) { return false; }")
	fmt.Fprintln(os.Stdout, "		return true;")
	fmt.Fprintf(os.Stdout, "	} // end Encode()\n")

	fmt.Fprintf(os.Stdout, "	bool Encode(char** buf) const {\n")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			WriteMessageEncodeCpp(varType, varName, TAB2)
			continue
		}
		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprintf(os.Stdout, "		size_t %v_size = %v.size();\n", varName, varName)
			fmt.Fprintf(os.Stdout, "		std::memcpy(*buf, &%v_size, sizeof(uint16_t));\n", varName)
			fmt.Fprintf(os.Stdout, "		(*buf) += sizeof(uint16_t);\n")
			if isPrimitiveType(arrVarType) {
				fmt.Fprintf(os.Stdout, "		for (std::list<%v>::const_iterator iter = %v.begin(); iter != %v.end(); ++iter) {\n", ConvertVarTypeCpp(arrVarType), varName, varName)
				fmt.Fprintf(os.Stdout, "			const %v & val = *iter;\n", ConvertVarTypeCpp(arrVarType))
				WriteMessageEncodeCpp(arrVarType, "val", TAB3)
				fmt.Fprintf(os.Stdout, "		}\n")
			} else {
				fmt.Fprintf(os.Stdout, "		for(std::list<%v>::const_iterator iter = %v.begin(); iter != %v.end(); ++iter) {\n", arrVarType, varName, varName)
				//fmt.Fprintf(os.Stdout, "			const %v & o = *iter;\n", arrVarType)
				fmt.Fprintf(os.Stdout, "			if (false == %v_Serializer::Encode(buf, *iter)) {\n", arrVarType)
				fmt.Fprintf(os.Stdout, "				return false;\n")
				fmt.Fprintf(os.Stdout, "			}\n")
				fmt.Fprintf(os.Stdout, "		}\n")
			}
			continue
		}

		// User Define Struct
		fmt.Fprintf(os.Stdout, "		if (false == %v_Serializer::Encode(buf, %v)) { return false; }\n", varType, varName)
	}
	fmt.Fprintln(os.Stdout, "		return true;")
	fmt.Fprintf(os.Stdout, "	} // end Encode()\n")

	// Decode
	fmt.Fprintf(os.Stdout, "	bool Decode(const std::vector<char>& buf) {\n")
	fmt.Fprintln(os.Stdout, "		size_t size = buf.size();")
	fmt.Fprintln(os.Stdout, "		if (0 == size) { return true; }")
	fmt.Fprintln(os.Stdout, "		const char * pBuf = &(buf[0]);")
	fmt.Fprintln(os.Stdout, "		if (false == Decode(&pBuf, size)) { return false; }")
	fmt.Fprintln(os.Stdout, "		return true;")
	fmt.Fprintf(os.Stdout, "	} // end Decode()\n")

	fmt.Fprintf(os.Stdout, "	bool Decode(const char** buf, size_t & size) {\n")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			WriteMessageDecodeCpp(varType, varName, TAB2)
			continue
		}
		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprintf(os.Stdout, "		if (sizeof(uint16_t) > size) { return false; }\n")
			fmt.Fprintf(os.Stdout, "		uint16_t %v_len = 0;\n", varName)
			fmt.Fprintf(os.Stdout, "		std::memcpy(&%v_len, *buf, sizeof(uint16_t));\n", varName)
			fmt.Fprintf(os.Stdout, "		(*buf) += sizeof(uint16_t); size -= sizeof(uint16_t);\n")
			if isPrimitiveType(arrVarType) {
				fmt.Fprintf(os.Stdout, "		for (uint16_t i = 0; i < %v_len; ++i) {\n", varName)
				fmt.Fprintf(os.Stdout, "			%v val;\n", ConvertVarTypeCpp(arrVarType))
				WriteMessageDecodeCpp(arrVarType, "val", TAB3)
				fmt.Fprintf(os.Stdout, "			%v.push_back(val);\n", varName)
				fmt.Fprintf(os.Stdout, "		}\n")
			} else {
				fmt.Fprintf(os.Stdout, "		for (uint16_t i = 0; i < %v_len; ++i) {\n", varName)
				fmt.Fprintf(os.Stdout, "			%v val;\n", arrVarType)
				fmt.Fprintf(os.Stdout, "			if (false == %v_Serializer::Decode(val, buf, size)) {\n", arrVarType)
				fmt.Fprintf(os.Stdout, "				return false;\n")
				fmt.Fprintf(os.Stdout, "			}\n")
				fmt.Fprintf(os.Stdout, "			%v.push_back(val);\n", varName)
				fmt.Fprintf(os.Stdout, "		}\n")
			}
			continue
		}

		// User Define Struct
		fmt.Fprintf(os.Stdout, "		if (false == %v_Serializer::Decode(%v, buf, size)) { return false; }\n", varType, varName)
	}
	fmt.Fprintln(os.Stdout, "		return true;")
	fmt.Fprintf(os.Stdout, "	} // end Encode()\n")

	fmt.Fprintf(os.Stdout, "}; // end struct\n")

	if !el.IsMsg {
		fmt.Fprintf(os.Stdout, "struct %v_Serializer {\n", el.Name)
		fmt.Fprintf(os.Stdout, "	static bool Encode(char** buf, const %v & o) { return o.Encode(buf); }\n", el.Name)
		fmt.Fprintf(os.Stdout, "	static bool Decode(%v & o, const char** buf, size_t & size) { return o.Decode(buf, size); }\n", el.Name)
		fmt.Fprintf(os.Stdout, "	static int32_t Size(const %v & o) { return o.Size(); }\n", el.Name)
		fmt.Fprintln(os.Stdout, "};\n")
	}
}

func CompileCppCode(stmt *TokenStmt, fname string) {
	fmt.Fprintf(os.Stdout, "#ifndef __%v_HPP__\n", strings.ToUpper(fname))
	fmt.Fprintf(os.Stdout, "#define __%v_HPP__\n\n", strings.ToUpper(fname))

	fmt.Fprintln(os.Stdout, "#include <string>")
	fmt.Fprintln(os.Stdout, "#include <vector>")
	fmt.Fprintln(os.Stdout, "#include <list>")
	fmt.Fprintln(os.Stdout, "#include <stdint.h>")
	fmt.Fprintln(os.Stdout, "#include <cstring>")

	for _, el := range stmt.Elements {
		GenerateCppCode(el)
	}

	fmt.Fprintf(os.Stdout, "#endif // __%v_HPP__\n", strings.ToUpper(fname))
}
