// generate_py.go
package main

import (
	"fmt"
	"os"
	"regexp"
)

func GetInitialValuePy(varType string) string {
	switch varType {
	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64":
		return "0"
	case "string":
		return "''"
	}
	return "''"
}

func WriteMessageEncodePy(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "data += struct.pack('H', len(%v))\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "fmt = str(len(%v)) + 's'\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "data += struct.pack(fmt, %v)\n", varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	switch varType {
	case "int8":
		fmt.Fprintf(os.Stdout, "data += struct.pack('b', %v)\n", varName)
	case "uint8":
		fmt.Fprintf(os.Stdout, "data += struct.pack('B', %v)\n", varName)
	case "int16":
		fmt.Fprintf(os.Stdout, "data += struct.pack('h', %v)\n", varName)
	case "uint16":
		fmt.Fprintf(os.Stdout, "data += struct.pack('H', %v)\n", varName)
	case "int32":
		fmt.Fprintf(os.Stdout, "data += struct.pack('i', %v)\n", varName)
	case "uint32":
		fmt.Fprintf(os.Stdout, "data += struct.pack('I', %v)\n", varName)
	case "int64":
		fmt.Fprintf(os.Stdout, "data += struct.pack('q', %v)\n", varName)
	case "uint64":
		fmt.Fprintf(os.Stdout, "data += struct.pack('Q', %v)\n", varName)
	}
}

func WriteMessageDecodePy(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "length = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "%v = buf[0:length]; buf = buf[length:]\n", varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	switch varType {
	case "int8":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('b', buf[0:1])[0]; buf = buf[1:]\n", varName)
	case "uint8":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('B', buf[0:1])[0]; buf = buf[1:]\n", varName)
	case "int16":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('h', buf[0:2])[0]; buf = buf[2:]\n", varName)
	case "uint16":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]\n", varName)
	case "int32":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('i', buf[0:4])[0]; buf = buf[4:]\n", varName)
	case "uint32":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('I', buf[0:4])[0]; buf = buf[4:]\n", varName)
	case "int64":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('q', buf[0:8])[0]; buf = buf[8:]\n", varName)
	case "uint64":
		fmt.Fprintf(os.Stdout, "%v = struct.unpack('Q', buf[0:8])[0]; buf = buf[8:]\n", varName)
	}
}

func GeneratePyCode(el *TokenElement) {
	fmt.Fprintf(os.Stdout, "class %v :\n", el.Name)
	if el.IsMsg {
		fmt.Fprint(os.Stdout, TAB1)
		fmt.Fprintf(os.Stdout, "MSG_ID = %v\n", el.Id)
	}
	fmt.Fprint(os.Stdout, TAB1)
	fmt.Fprintf(os.Stdout, "def __init__(self) :\n")

	// Declare Variable
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			fmt.Fprint(os.Stdout, TAB2)
			fmt.Fprintf(os.Stdout, "self.%v = %v\n", varName, GetInitialValuePy(varType))
			continue
		}

		if isArrayType(varType) {
			//re := regexp.MustCompile("[a-zA-Z0-9_]+")
			//arrVarType := re.FindString(varType)
			fmt.Fprintf(os.Stdout, TAB2)
			fmt.Fprintf(os.Stdout, "self.%v = []\n", varName)
			continue
		}

		fmt.Fprint(os.Stdout, TAB2)
		fmt.Fprintf(os.Stdout, "self.%v = %v()\n", varName, varType)
	}

	// Encode
	fmt.Fprint(os.Stdout, TAB1)
	fmt.Fprintf(os.Stdout, "def Encode(self) :\n")
	fmt.Fprint(os.Stdout, TAB2)
	fmt.Fprintf(os.Stdout, "data = ''\n")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			WriteMessageEncodePy(varType, "self."+varName, TAB2)
			continue
		}
		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprint(os.Stdout, TAB2)
			fmt.Fprintf(os.Stdout, "data += struct.pack('H', len(self.%v))\n", varName)
			fmt.Fprint(os.Stdout, TAB2)
			fmt.Fprintf(os.Stdout, "for o in self.%v :\n", varName)
			if isPrimitiveType(arrVarType) {
				WriteMessageEncodePy(arrVarType, "o", TAB3)
			} else {
				fmt.Fprintf(os.Stdout, TAB3)
				fmt.Fprintf(os.Stdout, "data += %vEncode(o)\n", arrVarType)
			}
			continue
		}

		// User Define Struct
		fmt.Fprint(os.Stdout, TAB2)
		fmt.Fprintf(os.Stdout, "data += %vEncode(self.%v)\n", varType, varName)
	}
	fmt.Fprint(os.Stdout, TAB2)
	fmt.Fprintf(os.Stdout, "return data\n")

	// Decode
	fmt.Fprint(os.Stdout, TAB1)
	fmt.Fprintf(os.Stdout, "def Decode(self, buf) :\n")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			WriteMessageDecodePy(varType, "self."+varName, TAB2)
			continue
		}
		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprint(os.Stdout, TAB2)
			fmt.Fprintf(os.Stdout, "list_count = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]\n")
			fmt.Fprint(os.Stdout, TAB2)
			fmt.Fprintf(os.Stdout, "for i in range(list_count) :\n")
			if isPrimitiveType(arrVarType) {
				WriteMessageDecodePy(arrVarType, "o", TAB3)
			} else {
				fmt.Fprint(os.Stdout, TAB3)
				fmt.Fprintf(os.Stdout, "[o, buf] = %vDecode(buf)\n", arrVarType)
			}
			fmt.Fprint(os.Stdout, TAB3)
			fmt.Fprintf(os.Stdout, "self.%v.append(o)\n", varName)
			continue
		}

		// User Define Struct
		fmt.Fprint(os.Stdout, TAB2)
		fmt.Fprintf(os.Stdout, "[self.%v, buf] = %vDecode(buf)\n", varName, varType)
	}
	fmt.Fprint(os.Stdout, TAB2)
	fmt.Fprintf(os.Stdout, "return buf\n")

	if !el.IsMsg {
		fmt.Fprintf(os.Stdout, "def %vEncode(o) :\n", el.Name)
		fmt.Fprint(os.Stdout, TAB1)
		fmt.Fprintf(os.Stdout, "return o.Encode()\n")

		fmt.Fprintf(os.Stdout, "def %vDecode(buf) :\n", el.Name)
		fmt.Fprint(os.Stdout, TAB1)
		fmt.Fprintf(os.Stdout, "o = %v()\n", el.Name)
		fmt.Fprint(os.Stdout, TAB1)
		fmt.Fprintf(os.Stdout, "buf = o.Decode(buf)\n")
		fmt.Fprint(os.Stdout, TAB1)
		fmt.Fprintf(os.Stdout, "return [o, buf]\n")
	}
}

func CompilePyCode(stmt *TokenStmt) {
	fmt.Fprintln(os.Stdout, "import struct\n")

	for _, el := range stmt.Elements {
		GeneratePyCode(el)
		fmt.Fprintln(os.Stdout)
	}
}
