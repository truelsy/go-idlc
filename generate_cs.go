// generate_cs.go
package main

import (
	"fmt"
	"os"
	"regexp"
)

const (
	TAB3 = "\t\t\t"
	TAB4 = "\t\t\t\t"
)

func GetInitialValue(varType string) string {
	switch varType {
	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64":
		return "0"
	case "string":
		return "\"\""
	}
	return "\"\""
}

func ConvertVarType(varType string) string {
	switch varType {
	case "int8":
		return "sbyte"
	case "uint8":
		return "byte"
	case "int16":
		return "short"
	case "uint16":
		return "ushort"
	case "int32":
		return "int"
	case "uint32":
		return "uint"
	case "int64":
		return "long"
	case "uint64":
		return "ulong"
	}
	return varType
}

func WriteMessageSizeCs(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "size += sizeof(short);\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "if (null != %v) { size += Encoding.UTF8.GetByteCount(%v); }\n", varName, varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "size += sizeof(%v);\n", ConvertVarType(varType))
}

func WriteMessageEncodeCs(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "if (null != %v) {\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "	int len = Encoding.UTF8.GetByteCount(%v);\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "	buf.Write(BitConverter.GetBytes(len), 0, sizeof(short));\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "	buf.Write(Encoding.UTF8.GetBytes(%v), 0, len);\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintln(os.Stdout, "} else {")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "	buf.Write(BitConverter.GetBytes(0), 0, sizeof(short));\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintln(os.Stdout, "}")
		return
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "buf.Write(BitConverter.GetBytes(%v), 0, sizeof(%v));\n", varName, ConvertVarType(varType))
}

func WriteMessageDecodeCs(varType, varName string, tab string) {
	if varType == "string" {
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "if (sizeof(short) > buf.Length - buf.Position) { return false; }\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "int %v_len = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "buf.Position += sizeof(short);\n")
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "if (%v_len > buf.Length - buf.Position) { return false; }\n", varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "byte[] %v_buf = new byte[%v_len];\n", varName, varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "Array.Copy(buf.GetBuffer(), (int)buf.Position, %v_buf, 0, %v_len);\n", varName, varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "%v = System.Text.Encoding.UTF8.GetString(%v_buf);\n", varName, varName)
		fmt.Fprint(os.Stdout, tab)
		fmt.Fprintf(os.Stdout, "buf.Position += %v_len;\n", varName)
		return
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "if (sizeof(%v) > buf.Length - buf.Position) { return false; }\n", ConvertVarType(varType))

	fmt.Fprint(os.Stdout, tab)
	switch varType {
	case "int8":
		fmt.Fprintf(os.Stdout, "%v = (sbyte)buf.GetBuffer()[buf.Position];\n", varName)
	case "uint8":
		fmt.Fprintf(os.Stdout, "%v = (byte)buf.GetBuffer()[buf.Position];\n", varName)
	case "int16":
		fmt.Fprintf(os.Stdout, "%v = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);\n", varName)
	case "uint16":
		fmt.Fprintf(os.Stdout, "%v = BitConverter.ToUInt16(buf.GetBuffer(), (int)buf.Position);\n", varName)
	case "int32":
		fmt.Fprintf(os.Stdout, "%v = BitConverter.ToInt32(buf.GetBuffer(), (int)buf.Position);\n", varName)
	case "uint32":
		fmt.Fprintf(os.Stdout, "%v = BitConverter.ToUInt32(buf.GetBuffer(), (int)buf.Position);\n", varName)
	case "int64":
		fmt.Fprintf(os.Stdout, "%v = BitConverter.ToInt64(buf.GetBuffer(), (int)buf.Position);\n", varName)
	case "uint64":
		fmt.Fprintf(os.Stdout, "%v = BitConverter.ToUInt64(buf.GetBuffer(), (int)buf.Position);\n", varName)
	}

	fmt.Fprint(os.Stdout, tab)
	fmt.Fprintf(os.Stdout, "buf.Position += sizeof(%v);\n", ConvertVarType(varType))
}

func GenerateCsCode(el *TokenElement) {
	fmt.Fprintln(os.Stdout, "public class", el.Name, "{")

	if el.IsMsg {
		fmt.Fprintf(os.Stdout, "	public const int MSG_ID = %v;\n", el.Id)
	}

	// Declare Variable
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		if isPrimitiveType(varType) {
			fmt.Fprintf(os.Stdout, "	public %v %v = %v;\n", ConvertVarType(varType), varName, GetInitialValue(varType))
			continue
		}

		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)
			fmt.Fprintf(os.Stdout, "	public List<%v> %v = new List<%v>();\n", ConvertVarType(arrVarType), varName, ConvertVarType(arrVarType))
			continue
		}

		fmt.Fprintf(os.Stdout, "	public %v %v = new %v();\n", varType, varName, varType)
	}

	// Size
	fmt.Fprintln(os.Stdout, "	public int Size() {")
	fmt.Fprintln(os.Stdout, "		int size = 0;")
	fmt.Fprintln(os.Stdout, "		try {")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		// Primitive
		if isPrimitiveType(varType) {
			WriteMessageSizeCs(varType, varName, TAB3)
			continue
		}

		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprintf(os.Stdout, "			size += sizeof(short);\n")
			if isPrimitiveType(arrVarType) {
				fmt.Fprintf(os.Stdout, "			foreach(var iter in %v) {\n", varName)
				fmt.Fprintf(os.Stdout, "				%v o = iter;\n", ConvertVarType(arrVarType))
				WriteMessageSizeCs(arrVarType, "o", TAB4)
				fmt.Fprintf(os.Stdout, "			}\n")
			} else {
				fmt.Fprintf(os.Stdout, "			foreach(var iter in %v) {\n", varName)
				fmt.Fprintf(os.Stdout, "				%v o = iter;\n", arrVarType)
				fmt.Fprintf(os.Stdout, "				size += %v_Serializer.Size(o);\n", arrVarType)
				fmt.Fprintf(os.Stdout, "			}\n")
			}
			continue
		}

		// User Define Struct
		fmt.Fprintf(os.Stdout, "			size += %v_Serializer.Size(%v);\n", varType, varName)
	}
	fmt.Fprintln(os.Stdout, "		} catch (System.Exception) {")
	fmt.Fprintln(os.Stdout, "			return -1;")
	fmt.Fprintln(os.Stdout, "		}")
	fmt.Fprintln(os.Stdout, "		return size;")
	fmt.Fprintln(os.Stdout, "	} // end Size()\n")

	// Encode
	fmt.Fprintln(os.Stdout, "	public bool Encode(MemoryStream buf) {")
	fmt.Fprintln(os.Stdout, "		try {")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		// Primitive
		if isPrimitiveType(varType) {
			WriteMessageEncodeCs(varType, varName, TAB3)
			continue
		}

		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprintf(os.Stdout, "			buf.Write(BitConverter.GetBytes(%v.Count), 0, sizeof(short));\n", varName)
			if isPrimitiveType(arrVarType) {
				fmt.Fprintf(os.Stdout, "			foreach(var iter in %v) {\n", varName)
				fmt.Fprintf(os.Stdout, "				%v o = iter;\n", ConvertVarType(arrVarType))
				WriteMessageEncodeCs(arrVarType, "o", TAB4)
				fmt.Fprintf(os.Stdout, "			}\n")
			} else {
				fmt.Fprintf(os.Stdout, "			foreach(var iter in %v) {\n", varName)
				fmt.Fprintf(os.Stdout, "				%v o = iter;\n", arrVarType)
				fmt.Fprintf(os.Stdout, "				if (false == %v_Serializer.Encode(buf, o)) { return false; }\n", arrVarType)
				fmt.Fprintf(os.Stdout, "			}\n")
			}
			continue
		}

		// User Define Struct
		fmt.Fprintf(os.Stdout, "			if (false == %v_Serializer.Encode(buf, %v)) { return false; }\n", varType, varName)
	}
	fmt.Fprintln(os.Stdout, "		} catch (System.Exception) {")
	fmt.Fprintln(os.Stdout, "			return false;")
	fmt.Fprintln(os.Stdout, "		}")
	fmt.Fprintln(os.Stdout, "		return true;")
	fmt.Fprintln(os.Stdout, "	} // end Encode()\n")

	// Decode
	fmt.Fprintln(os.Stdout, "	public bool Decode(MemoryStream buf) {")
	fmt.Fprintln(os.Stdout, "		try {")
	for idx, varName := range el.VarName {
		varType := el.VarType[idx]
		// Primitive
		if isPrimitiveType(varType) {
			WriteMessageDecodeCs(varType, varName, TAB3)
			fmt.Fprintln(os.Stdout, "")
			continue
		}

		// Array
		if isArrayType(varType) {
			re := regexp.MustCompile("[a-zA-Z0-9_]+")
			arrVarType := re.FindString(varType)

			fmt.Fprintf(os.Stdout, "			if (sizeof(short) > buf.Length - buf.Position) { return false; }\n")
			fmt.Fprintf(os.Stdout, "			int %v_len = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);\n", varName)
			fmt.Fprintf(os.Stdout, "			buf.Position += sizeof(short);\n")

			if isPrimitiveType(arrVarType) {
				fmt.Fprintf(os.Stdout, "			for (int i = 0; i < %v_len; i++) {\n", varName)
				fmt.Fprintf(os.Stdout, "				%v o = %v;\n", ConvertVarType(arrVarType), GetInitialValue(arrVarType))
				WriteMessageDecodeCs(arrVarType, "o", TAB4)
				fmt.Fprintf(os.Stdout, "				%v.Add(o);\n", varName)
				fmt.Fprintf(os.Stdout, "			}\n")
			} else {
				fmt.Fprintf(os.Stdout, "			for (int i = 0; i < %v_len; i++) {\n", varName)
				fmt.Fprintf(os.Stdout, "				%v o = new %v();\n", arrVarType, arrVarType)
				fmt.Fprintf(os.Stdout, "				if (false == %v_Serializer.Decode(ref o, buf)) { return false; }\n", arrVarType)
				fmt.Fprintf(os.Stdout, "				%v.Add(o);\n", varName)
				fmt.Fprintf(os.Stdout, "			}\n")
			}
			continue
		}

		// User Define Struct
		fmt.Fprintf(os.Stdout, "			if (false == %v_Serializer.Decode(ref %v, buf)) { return false; }\n", varType, varName)
	}
	fmt.Fprintln(os.Stdout, "		} catch (System.Exception) {")
	fmt.Fprintln(os.Stdout, "			return false;")
	fmt.Fprintln(os.Stdout, "		}")
	fmt.Fprintln(os.Stdout, "		return true;")
	fmt.Fprintln(os.Stdout, "	} // end Decode()\n")

	fmt.Fprintln(os.Stdout, "} // end class\n")

	if !el.IsMsg {
		fmt.Fprintf(os.Stdout, "public struct %v_Serializer {\n", el.Name)
		fmt.Fprintf(os.Stdout, "	public static bool Encode(MemoryStream buf, %v o) { return o.Encode(buf); }\n", el.Name)
		fmt.Fprintf(os.Stdout, "	public static bool Decode(ref %v o, MemoryStream buf) { return o.Decode(buf); }\n", el.Name)
		fmt.Fprintf(os.Stdout, "	public static int Size(%v o) { return o.Size(); }\n", el.Name)
		fmt.Fprintln(os.Stdout, "}\n")
	}
}

func CompileCsCode(stmt *TokenStmt) {
	fmt.Fprintln(os.Stdout, "using System;")
	fmt.Fprintln(os.Stdout, "using System.Collections.Generic;")
	fmt.Fprintln(os.Stdout, "using System.Text;")
	fmt.Fprintln(os.Stdout, "using System.IO;")
	fmt.Fprintln(os.Stdout, "#pragma warning disable 108, 219\n")

	fmt.Fprintln(os.Stdout, "namespace msg {")
	for _, el := range stmt.Elements {
		GenerateCsCode(el)
	}
	fmt.Fprintln(os.Stdout, "} // end namespace")
	fmt.Fprintln(os.Stdout, "#pragma warning restore 108, 219")
}
