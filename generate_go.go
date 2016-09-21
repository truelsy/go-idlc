// generate_go.go
package main

import (
	"fmt"
	"os"
	"regexp"
	"unicode"
)

func WriteMessageDecodeGo(varType, varName string) {
	if varType == "string" {
		fmt.Fprintln(os.Stdout, "{")
		fmt.Fprintln(os.Stdout, "	strlen := binary.LittleEndian.Uint16(buf[offset:])")
		fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(strlen))")
		fmt.Fprintln(os.Stdout, "	if (size + offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return BOError{size, offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintln(os.Stdout, "	offset += size\n")
		fmt.Fprintln(os.Stdout, "	size = int(strlen)")
		fmt.Fprintln(os.Stdout, "	if (size + offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return BOError{size, offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintf(os.Stdout, "	%v = string(buf[offset : offset + int(strlen)])\n", varName)
		fmt.Fprintln(os.Stdout, "	offset += size")
		fmt.Fprintln(os.Stdout, "}")
		return
	}

	fmt.Fprintln(os.Stdout, "{")
	fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(", varName, "))")
	fmt.Fprintln(os.Stdout, "		if (size + offset) > bufcap {")
	fmt.Fprintln(os.Stdout, "			return BOError{size, offset, bufcap}")
	fmt.Fprintln(os.Stdout, "		}")

	switch varType {
	case "int8":
		fmt.Fprintf(os.Stdout, "	%v = int8(buf[offset])\n", varName)
	case "uint8":
		fmt.Fprintf(os.Stdout, "	%v = uint8(buf[offset])\n", varName)
	case "int16":
		fmt.Fprintf(os.Stdout, "	%v = int16(binary.LittleEndian.Uint16(buf[offset:]))\n", varName)
	case "uint16":
		fmt.Fprintf(os.Stdout, "	%v = binary.LittleEndian.Uint16(buf[offset:])\n", varName)
	case "int32":
		fmt.Fprintf(os.Stdout, "	%v = int32(binary.LittleEndian.Uint32(buf[offset:]))\n", varName)
	case "uint32":
		fmt.Fprintf(os.Stdout, "	%v = binary.LittleEndian.Uint32(buf[offset:])\n", varName)
	case "int64":
		fmt.Fprintf(os.Stdout, "	%v = int64(binary.LittleEndian.Uint64(buf[offset:]))\n", varName)
	case "uint64":
		fmt.Fprintf(os.Stdout, "	%v = binary.LittleEndian.Uint64(buf[offset:])\n", varName)
	}
	fmt.Fprintln(os.Stdout, "	offset += size")
	fmt.Fprintln(os.Stdout, "}")
}

func WriteMessageEncodeGo(varType, varName string) {
	if varType == "string" {
		fmt.Fprintln(os.Stdout, "{")
		fmt.Fprintln(os.Stdout, "	strlen := uint16(len(", varName, "))")
		fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(strlen))")
		fmt.Fprintln(os.Stdout, "	if (size + offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return 0, BOError{size, offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[offset:], strlen)")
		fmt.Fprintln(os.Stdout, "	offset += size\n")
		fmt.Fprintln(os.Stdout, "	size = int(strlen)")
		fmt.Fprintln(os.Stdout, "	if (size + offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return 0, BOError{size, offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintln(os.Stdout, "	copy(buf[offset:], ", varName, ")")
		fmt.Fprintln(os.Stdout, "	offset += size")
		fmt.Fprintln(os.Stdout, "}")
		return
	}

	fmt.Fprintln(os.Stdout, "{")
	fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(", varName, "))")
	fmt.Fprintln(os.Stdout, "		if (size + offset) > bufcap {")
	fmt.Fprintln(os.Stdout, "			return 0, BOError{size, offset, bufcap}")
	fmt.Fprintln(os.Stdout, "		}")

	switch varType {
	case "int8":
		fmt.Fprintln(os.Stdout, "	buf[offset] = byte(", varName, ")")
	case "uint8":
		fmt.Fprintln(os.Stdout, "	buf[offset] = uint8(", varName, ")")
	case "int16":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[offset:], uint16(", varName, "))")
	case "uint16":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[offset:], ", varName, ")")
	case "int32":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint32(buf[offset:], uint32(", varName, "))")
	case "uint32":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint32(buf[offset:], ", varName, ")")
	case "int64":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint64(buf[offset:], uint64(", varName, "))")
	case "uint64":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint64(buf[offset:], ", varName, ")")
	}

	fmt.Fprintln(os.Stdout, "	offset += size")
	fmt.Fprintln(os.Stdout, "}")
}

func WriteStructDecodeGo(varType, varName string) {
	if varType == "string" {
		fmt.Fprintln(os.Stdout, "{")
		fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(uint16(0)))")
		fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintln(os.Stdout, "	strlen := binary.LittleEndian.Uint16(buf[*offset:])")
		fmt.Fprintln(os.Stdout, "	*offset += size\n")
		fmt.Fprintln(os.Stdout, "	size = int(strlen)")
		fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintf(os.Stdout, "	%v = string(buf[*offset : *offset + int(strlen)])\n", varName)
		fmt.Fprintln(os.Stdout, "	*offset += size")
		fmt.Fprintln(os.Stdout, "}")
		return
	}

	fmt.Fprintln(os.Stdout, "{")
	fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(", varName, "))")
	fmt.Fprintln(os.Stdout, "		if (size + *offset) > bufcap {")
	fmt.Fprintln(os.Stdout, "			return BOError{size, *offset, bufcap}")
	fmt.Fprintln(os.Stdout, "		}")

	switch varType {
	case "int8":
		fmt.Fprintf(os.Stdout, "	%v = int8(buf[*offset])\n", varName)
	case "uint8":
		fmt.Fprintf(os.Stdout, "	%v = uint8(buf[*offset])\n", varName)
	case "int16":
		fmt.Fprintf(os.Stdout, "	%v = int16(binary.LittleEndian.Uint16(buf[*offset:]))\n", varName)
	case "uint16":
		fmt.Fprintf(os.Stdout, "	%v = binary.LittleEndian.Uint16(buf[*offset:])\n", varName)
	case "int32":
		fmt.Fprintf(os.Stdout, "	%v = int32(binary.LittleEndian.Uint32(buf[*offset:]))\n", varName)
	case "uint32":
		fmt.Fprintf(os.Stdout, "	%v = binary.LittleEndian.Uint32(buf[*offset:])\n", varName)
	case "int64":
		fmt.Fprintf(os.Stdout, "	%v = int64(binary.LittleEndian.Uint64(buf[*offset:]))\n", varName)
	case "uint64":
		fmt.Fprintf(os.Stdout, "	%v = binary.LittleEndian.Uint64(buf[*offset:])\n", varName)
	}
	fmt.Fprintln(os.Stdout, "	*offset += size")
	fmt.Fprintln(os.Stdout, "}")
}

func WriteStructEncodeGo(varType, varName string) {
	if varType == "string" {
		fmt.Fprintln(os.Stdout, "{")
		fmt.Fprintln(os.Stdout, "	strlen := uint16(len(", varName, "))")
		fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(strlen))")
		fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[*offset:], strlen)")
		fmt.Fprintln(os.Stdout, "	*offset += size\n")
		fmt.Fprintln(os.Stdout, "	size = int(strlen)")
		fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
		fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
		fmt.Fprintln(os.Stdout, "	}")
		fmt.Fprintln(os.Stdout, "	copy(buf[*offset:], ", varName, ")")
		fmt.Fprintln(os.Stdout, "	*offset += size")
		fmt.Fprintln(os.Stdout, "}")
		return
	}

	fmt.Fprintln(os.Stdout, "{")
	fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(", varName, "))")
	fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
	fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
	fmt.Fprintln(os.Stdout, "	}")

	switch varType {
	case "int8":
		fmt.Fprintln(os.Stdout, "	buf[*offset] = byte(", varName, ")")
	case "uint8":
		fmt.Fprintln(os.Stdout, "	buf[*offset] = uint8(", varName, ")")
	case "int16":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[*offset:], uint16(", varName, "))")
	case "uint16":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[*offset:], ", varName, ")")
	case "int32":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint32(buf[*offset:], uint32(", varName, "))")
	case "uint32":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint32(buf[*offset:], ", varName, ")")
	case "int64":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint64(buf[*offset:], uint64(", varName, "))")
	case "uint64":
		fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint64(buf[*offset:], ", varName, ")")
	}
	fmt.Fprintln(os.Stdout, "	*offset += size")
	fmt.Fprintln(os.Stdout, "}")
}

func GenerateGoCode_Message(el *TokenElement) {
	fmt.Fprintln(os.Stdout, "///////////////////////////////////////////////////////////////////////")
	fmt.Fprintln(os.Stdout, "// message", el.Name)

	fmt.Fprintln(os.Stdout, "type", el.Name, "struct {")

	for idx, varName := range el.VarName {
		fmt.Fprintln(os.Stdout, varName, el.VarType[idx])
	}
	fmt.Fprintln(os.Stdout, "}\n")

	fmt.Fprintln(os.Stdout, "func (p", el.Name, ") GetId() uint16 {")
	fmt.Fprintf(os.Stdout, "	return %v_ID\n", el.Name)
	fmt.Fprintln(os.Stdout, "}\n")

	// Message Encode
	fmt.Fprintln(os.Stdout, "func (p *", el.Name, ") Encode(buf []byte) (int, error) {")
	fmt.Fprintln(os.Stdout, "var offset int = 0")
	fmt.Fprintln(os.Stdout, "var bufcap int = cap(buf)")
	fmt.Fprintln(os.Stdout, "var size int = 0\n")

	for idx, varName := range el.VarName {
		varType := el.VarType[idx]

		if isPrimitiveType(varType) {
			WriteMessageEncodeGo(varType, "p."+varName)
			continue
		}

		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)
			fmt.Fprintln(os.Stdout, "{")
			fmt.Fprintln(os.Stdout, "	arrlen := uint16(len(p.", varName, "))")
			fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(arrlen))")
			fmt.Fprintln(os.Stdout, "	if (size + offset) > bufcap {")
			fmt.Fprintln(os.Stdout, "		return 0, BOError{size, offset, bufcap}")
			fmt.Fprintln(os.Stdout, "	}")
			fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[offset:], arrlen)")
			fmt.Fprintln(os.Stdout, "	offset += size")
			fmt.Fprintln(os.Stdout, "	for _, v := range p.", varName, "{")
			if isPrimitiveType(arrVarType) {
				WriteMessageEncodeGo(arrVarType, "v")
			} else {
				fmt.Fprintln(os.Stdout, "		if err := v.Encode(&offset, buf); err != nil {")
				fmt.Fprintln(os.Stdout, "		return 0, err")
				fmt.Fprintln(os.Stdout, "	}")
			}
			fmt.Fprintln(os.Stdout, "	}") // end for
			fmt.Fprintln(os.Stdout, "}\n")
			continue
		}

		// UserDefine Struct
		fmt.Fprintln(os.Stdout, "if err := p.", varName, ".Encode(&offset, buf); err != nil {")
		fmt.Fprintln(os.Stdout, "	return 0, err")
		fmt.Fprintln(os.Stdout, "}")
		fmt.Fprintln(os.Stdout, "\n")
	}
	fmt.Fprintln(os.Stdout, "return offset, nil")
	fmt.Fprintln(os.Stdout, "}\n")

	// Message Decode
	fmt.Fprintln(os.Stdout, "func (p *", el.Name, ") Decode(buf []byte) error {")
	fmt.Fprintln(os.Stdout, "var offset int = 0")
	fmt.Fprintln(os.Stdout, "var bufcap int = cap(buf)")
	fmt.Fprintln(os.Stdout, "var size int = 0\n")

	for idx, varName := range el.VarName {
		varType := el.VarType[idx]

		if isPrimitiveType(varType) {
			WriteMessageDecodeGo(varType, "p."+varName)
			continue
		}

		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)
			fmt.Fprintln(os.Stdout, "{")
			fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(uint16(0)))")
			fmt.Fprintln(os.Stdout, "	if (size + offset) > bufcap {")
			fmt.Fprintln(os.Stdout, "		return BOError{size, offset, bufcap}")
			fmt.Fprintln(os.Stdout, "	}")
			fmt.Fprintln(os.Stdout, "	arrlen := binary.LittleEndian.Uint16(buf[offset:])")
			fmt.Fprintln(os.Stdout, "	offset += size")
			fmt.Fprintln(os.Stdout, "	for i := uint16(0); i < arrlen; i++ {")
			if isPrimitiveType(arrVarType) {
				fmt.Fprintln(os.Stdout, "	var v", arrVarType)
				WriteMessageDecodeGo(arrVarType, "v")
				fmt.Fprintln(os.Stdout, "	p.", varName, "= append(p.", varName, ", v)")
			} else {
				fmt.Fprintln(os.Stdout, "	v :=", arrVarType, "{}")
				fmt.Fprintln(os.Stdout, "	if err := v.Decode(&offset, buf); err != nil {")
				fmt.Fprintln(os.Stdout, "		return err")
				fmt.Fprintln(os.Stdout, "	}")
				fmt.Fprintln(os.Stdout, "	p.", varName, "= append(p.", varName, ", v)")
			}
			fmt.Fprintln(os.Stdout, "	}") // end for
			fmt.Fprintln(os.Stdout, "}\n")
			continue
		}

		// UserDefine Struct
		fmt.Fprintln(os.Stdout, "if err := p.", varName, ".Decode(&offset, buf); err != nil {")
		fmt.Fprintln(os.Stdout, "	return err")
		fmt.Fprintln(os.Stdout, "}")
		fmt.Fprintln(os.Stdout, "\n")
	}
	fmt.Fprintln(os.Stdout, "return nil")
	fmt.Fprintln(os.Stdout, "}\n")
}

func GenerateGoCode_Struct(el *TokenElement) {
	fmt.Fprintln(os.Stdout, "///////////////////////////////////////////////////////////////////////")
	fmt.Fprintln(os.Stdout, "// struct", el.Name)

	fmt.Fprintln(os.Stdout, "type", el.Name, "struct {")

	for idx, varName := range el.VarName {
		fmt.Fprintln(os.Stdout, varName, el.VarType[idx])
	}
	fmt.Fprintln(os.Stdout, "}\n")

	// Struct Encode
	fmt.Fprintln(os.Stdout, "func (p *", el.Name, ") Encode(offset *int, buf []byte) error {")
	fmt.Fprintln(os.Stdout, "var bufcap int = cap(buf)")
	fmt.Fprintln(os.Stdout, "var size int = 0\n")

	for idx, varName := range el.VarName {
		varType := el.VarType[idx]

		if isPrimitiveType(varType) {
			WriteStructEncodeGo(varType, "p."+varName)
			continue
		}

		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)
			fmt.Fprintln(os.Stdout, "{")
			fmt.Fprintln(os.Stdout, "	arrlen := uint16(len(p.", varName, "))")
			fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(arrlen))")
			fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
			fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
			fmt.Fprintln(os.Stdout, "	}")
			fmt.Fprintln(os.Stdout, "	binary.LittleEndian.PutUint16(buf[*offset:], arrlen)")
			fmt.Fprintln(os.Stdout, "	*offset += size")
			fmt.Fprintln(os.Stdout, "	for _, v := range p.", varName, "{")
			if isPrimitiveType(arrVarType) {
				WriteStructEncodeGo(arrVarType, "v")
			} else {
				fmt.Fprintln(os.Stdout, "	if err := v.Encode(offset, buf); err != nil {")
				fmt.Fprintln(os.Stdout, "		return err")
				fmt.Fprintln(os.Stdout, "	}")
			}
			fmt.Fprintln(os.Stdout, "	}") // end for
			fmt.Fprintln(os.Stdout, "}\n")
			continue
		}

		// UserDefine Struct
		fmt.Fprintln(os.Stdout, "if err := p.", varName, ".Encode(offset, buf); err != nil {")
		fmt.Fprintln(os.Stdout, "	return err")
		fmt.Fprintln(os.Stdout, "}")
		fmt.Fprintln(os.Stdout, "\n")
	}
	fmt.Fprintln(os.Stdout, "return nil")
	fmt.Fprintln(os.Stdout, "}\n")

	// Struct Decode
	fmt.Fprintln(os.Stdout, "func (p *", el.Name, ") Decode(offset *int, buf []byte) error {")
	fmt.Fprintln(os.Stdout, "var bufcap int = cap(buf)")
	fmt.Fprintln(os.Stdout, "var size int = 0\n")

	for idx, varName := range el.VarName {
		//		fmt.Fprintln(os.Stdout)
		varType := el.VarType[idx]

		if isPrimitiveType(varType) {
			WriteStructDecodeGo(varType, "p."+varName)
			continue
		}

		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)
			fmt.Fprintln(os.Stdout, "{")
			fmt.Fprintln(os.Stdout, "	size = int(unsafe.Sizeof(uint16(0)))")
			fmt.Fprintln(os.Stdout, "	if (size + *offset) > bufcap {")
			fmt.Fprintln(os.Stdout, "		return BOError{size, *offset, bufcap}")
			fmt.Fprintln(os.Stdout, "	}")
			fmt.Fprintln(os.Stdout, "	arrlen := binary.LittleEndian.Uint16(buf[*offset:])")
			fmt.Fprintln(os.Stdout, "	*offset += size")
			fmt.Fprintln(os.Stdout, "	for i := uint16(0); i < arrlen; i++ {")
			if isPrimitiveType(arrVarType) {
				fmt.Fprintln(os.Stdout, "	var v", arrVarType)
				WriteStructDecodeGo(arrVarType, "v")
				fmt.Fprintln(os.Stdout, "	p.", varName, "= append(p.", varName, ", v)")
			} else {
				fmt.Fprintln(os.Stdout, "	v :=", arrVarType, "{}")
				fmt.Fprintln(os.Stdout, "	if err := v.Decode(offset, buf); err != nil {")
				fmt.Fprintln(os.Stdout, "		return err")
				fmt.Fprintln(os.Stdout, "	}")
				fmt.Fprintln(os.Stdout, "	p.", varName, "= append(p.", varName, ", v)")
			}
			fmt.Fprintln(os.Stdout, "	}") // end for
			fmt.Fprintln(os.Stdout, "}\n")
			continue
		}

		// UserDefine Struct
		fmt.Fprintln(os.Stdout, "if err := p.", varName, ".Decode(offset, buf); err != nil {")
		fmt.Fprintln(os.Stdout, "	return err")
		fmt.Fprintln(os.Stdout, "}")
		fmt.Fprintln(os.Stdout, "\n")
	}
	fmt.Fprintln(os.Stdout, "return nil")
	fmt.Fprintln(os.Stdout, "}\n")
}

func CompileGoCode(stmt *TokenStmt) {
	fmt.Fprintln(os.Stdout, "//////////////////////////////////////////////////////////////////")
	fmt.Fprintln(os.Stdout, "// Automatically-generated file. Do not edit!")
	fmt.Fprintln(os.Stdout, "//////////////////////////////////////////////////////////////////\n")
	fmt.Fprintln(os.Stdout, "package msg")
	fmt.Fprintln(os.Stdout, "import \"encoding/binary\"")
	fmt.Fprintln(os.Stdout, "import \"unsafe\"")
	fmt.Fprintln(os.Stdout, "import \"fmt\"")
	fmt.Fprintln(os.Stdout, "")

	fmt.Fprintln(os.Stdout, "type Messager interface {")
	fmt.Fprintln(os.Stdout, "	GetId() uint16")
	fmt.Fprintln(os.Stdout, "	Encode([]byte) (int, error)")
	fmt.Fprintln(os.Stdout, "	Decode([]byte) error")
	fmt.Fprintln(os.Stdout, "}\n")

	fmt.Fprintln(os.Stdout, "type BOError struct {")
	fmt.Fprintln(os.Stdout, "	size int")
	fmt.Fprintln(os.Stdout, "	offset int")
	fmt.Fprintln(os.Stdout, "	bufcap int")
	fmt.Fprintln(os.Stdout, "}\n")

	fmt.Fprintln(os.Stdout, "func (e BOError) Error() string {")
	fmt.Fprintln(os.Stdout, "	return fmt.Sprintf(\"ERROR - Buffer Overflow! size(%v) offset(%v) bufcap(%v)\", e.size, e.offset, e.bufcap)")
	fmt.Fprintln(os.Stdout, "}\n")

	// Message Id
	fmt.Fprintln(os.Stdout, "///////////////////////////////////////////////////////////////////////")
	fmt.Fprintln(os.Stdout, "// declare message id")
	fmt.Fprintln(os.Stdout, "var (")
	for _, el := range stmt.Elements {
		if el.IsMsg {
			fmt.Fprintf(os.Stdout, "	%v_ID = uint16(%v)\n", el.Name, el.Id)
		}
	}
	fmt.Fprintln(os.Stdout, ")")

	for _, el := range stmt.Elements {

		for idx, varName := range el.VarName {
			a := []rune(varName)
			a[0] = unicode.ToUpper(a[0])
			el.VarName[idx] = string(a)
		}

		if el.IsMsg {
			GenerateGoCode_Message(el)
		} else {
			GenerateGoCode_Struct(el)
		}
	}

}
