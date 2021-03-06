# go-idlc
IDL(Interface Define Language) Compiler Written by Go Lang.

## 설치
```
$ go get github.com/truelsy/go-idlc
```

## 특징
## 코드 생성
 * Written Raw(IDL) File (example.idl)

```go
struct Item {
	item_seq	uint64
	item_name	string
}

message TestMsg : 7000 {
	m_s8	int8
	m_u8	uint8
	m_s16	int16
	m_u16	uint16	
	m_s32	int32
	m_u32	uint32
	m_s64	int64
	m_u64	uint64	
	m_str	string
	m_items	[]Item
}
```

 * Generate Go Code (example.go)
```
$ go-idlc -l=go example.idl
```

 * Generate C# Code (example.cs)
```
$ go-idlc -l=cs example.idl
```

 * Generate C++ Code (example.hpp)
```
$ go-idlc -l=cpp example.idl
```

 * Generate Python Code (example.py)
```
$ go-idlc -l=py example.idl
```

## 예제
 * Go
```go
func main() {
	buffer := make([]byte, 1024)

	// Message Encode
	{
		writer := &msg.TestMsg{}
		writer.M_s8 = 1
		writer.M_u8 = 2
		writer.M_s16 = 3
		writer.M_u16 = 4
		writer.M_s32 = 5
		writer.M_u32 = 6
		writer.M_s64 = 7
		writer.M_u64 = 8
		writer.M_str = "truelsy"
		for i := 0; i < 3; i++ {
			item := Item{}
			item.Item_seq = uint64(i)
			item.Item_name = "Item_Name_" + strconv.Itoa(i)
			writer.M_items = append(writer.M_items, item)
		}

		offset, err := writer.Encode(buffer)
		if err != nil {
      			return
		}

		fmt.Println("writer offset :", offset)
	}

	// Message Decode
	{
		reader := &msg.TestMsg{}
		err := reader.Decode(buffer)
		if err != nil {
      			return
		}

		fmt.Println("m_s8 :", reader.M_s8)
		fmt.Println("m_u8 :", reader.M_u8)
		fmt.Println("m_s16 :", reader.M_s16)
		fmt.Println("m_u16 :", reader.M_u16)
		fmt.Println("m_s32 :", reader.M_s32)
		fmt.Println("m_u32 :", reader.M_u32)
		fmt.Println("m_s64 :", reader.M_s64)
		fmt.Println("m_u64 :", reader.M_u64)
		fmt.Println("m_str :", reader.M_str)

		for _, item := range reader.M_items {
			fmt.Println("item_seq :", item.Item_seq)
			fmt.Println("item_name :", item.Item_name)
		}
	}
}
```

 * C#
```csharp
static void Main(string[] args)
{
    MemoryStream stream = new MemoryStream();

    // Message Encode
    {
        msg.TestMsg writer = new msg.TestMsg();
        writer.m_s8 = 1;
        writer.m_u8 = 2;
        writer.m_s16 = 3;
        writer.m_u16 = 4;
        writer.m_s32 = 5;
        writer.m_u32 = 6;
        writer.m_s64 = 7;
        writer.m_u64 = 8;
        writer.m_str = "truelsy";
        for (int i = 0; i < 3; i++)
        {
            msg.Item item = new msg.Item();
            item.item_seq = (ulong)i;
            item.item_name = "Item_Name_" + i;
            writer.m_items.Add(item);
        }
        if (!writer.Encode(stream))
        {
            return;
        }
    }

    // Message Decode
    {
        stream.Seek(0, SeekOrigin.Begin);

        msg.TestMsg reader = new msg.TestMsg();
        if (!reader.Decode(stream))
        {
            return;
        }

        System.Console.WriteLine("m_s8 : {0}", reader.m_s8);
        System.Console.WriteLine("m_u8 : {0}", reader.m_u8);
        System.Console.WriteLine("m_s16 : {0}", reader.m_s16);
        System.Console.WriteLine("m_u16 : {0}", reader.m_u16);
        System.Console.WriteLine("m_s32 : {0}", reader.m_s32);
        System.Console.WriteLine("m_u32 : {0}", reader.m_u32);
        System.Console.WriteLine("m_s64 : {0}", reader.m_s64);
        System.Console.WriteLine("m_u64 : {0}", reader.m_u64);
        System.Console.WriteLine("m_str : {0}", reader.m_str);

        foreach (msg.Item item in reader.m_items)
        {
            System.Console.WriteLine("item_seq : {0}", item.item_seq);
            System.Console.WriteLine("item_name : {0}", item.item_name);
        }
    }
}
```
 * C++
```cpp
int main() {
	std::vector<char> buf;

	// Message Encode
	{
		TestMsg writer;
		writer.m_s8 = 1;
		writer.m_u8 = 2;
		writer.m_s16 = 3;
		writer.m_u16 = 4;
		writer.m_s32 = 5;
		writer.m_u32 = 6;
		writer.m_s64 = 7;
		writer.m_u64 = 8;
		writer.m_str = "truelsy";
		for (int i = 0; i < 3; i++)
		{
			Item o;
			o.item_seq = static_cast<uint64_t>(i);
			o.item_name = "Item_Name_" + std::to_string(i);
			writer.m_items.push_back(o);
		}

		if (false == writer.Encode(buf))
		{
			std::cout << "Encode Error" << std::endl;
			return -1;
		}
	}

	// Message Decode
	{
		TestMsg reader;
		if (false == reader.Decode(buf))
		{
			std::cout << "Decode Error" << std::endl;
			return -1;
		}

		printf("s8 : %d\n", reader.m_s8);
		printf("u8 : %d\n", reader.m_u8);
		printf("s16 : %d\n", reader.m_s16);
		printf("u16 : %d\n", reader.m_u16);
		printf("s32 : %d\n", reader.m_s32);
		printf("u32 : %d\n", reader.m_u32);
		printf("s64 : %ld\n", reader.m_s64);
		printf("u64 : %ld\n", reader.m_u64);
		printf("str : %s\n", reader.m_str.c_str());

		for (auto iter = reader.m_items.begin(); iter != reader.m_items.end(); ++iter)
		{
			printf("item_seq : %ld\n", (*iter).item_seq);
			printf("item_name : %s\n", (*iter).item_name.c_str());
		}
	}
}
```
 * Python
```python
#!/usr/bin/env python

from example import *

# Encode
writer = TestMsg()
writer.m_s8 = 1
writer.m_u8 = 2
writer.m_s16 = 3
writer.m_u16 = 4
writer.m_s32 = 5
writer.m_u32 = 6
writer.m_s64 = 7
writer.m_u64 = 8
writer.m_str = "truelsy"

for i in range(3) :
    item = Item()
    item.item_seq = i
    item.item_name = "Item_Name_%d" % i
    writer.m_items.append(item)

data = writer.Encode()

# Decode
reader = TestMsg()
reader.Decode(data)

print "m_s8 : %d" % reader.m_s8
print "m_u8 : %d" % reader.m_u8
print "m_s16 : %d" % reader.m_s16
print "m_u16 : %d" % reader.m_u16
print "m_s32 : %d" % reader.m_s32
print "m_u32 : %d" % reader.m_u32
print "m_s64 : %d" % reader.m_s64
print "m_u64 : %d" % reader.m_u64
print "m_str : %s" % reader.m_str

for item in reader.m_items :
    print "item_seq : %d" % item.item_seq
    print "item_name : %s" % item.item_name
```
