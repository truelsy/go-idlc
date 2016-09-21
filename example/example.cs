using System;
using System.Collections.Generic;
using System.Text;
using System.IO;
#pragma warning disable 108, 219

namespace msg {
public class Item {
	public ulong item_seq = 0;
	public string item_name = "";
	public int Size() {
		int size = 0;
		try {
			size += sizeof(ulong);
			size += sizeof(short);
			if (null != item_name) { size += Encoding.UTF8.GetByteCount(item_name); }
		} catch (System.Exception) {
			return -1;
		}
		return size;
	} // end Size()

	public bool Encode(MemoryStream buf) {
		try {
			buf.Write(BitConverter.GetBytes(item_seq), 0, sizeof(ulong));
			if (null != item_name) {
				int len = Encoding.UTF8.GetByteCount(item_name);
				buf.Write(BitConverter.GetBytes(len), 0, sizeof(short));
				buf.Write(Encoding.UTF8.GetBytes(item_name), 0, len);
			} else {
				buf.Write(BitConverter.GetBytes(0), 0, sizeof(short));
			}
		} catch (System.Exception) {
			return false;
		}
		return true;
	} // end Encode()

	public bool Decode(MemoryStream buf) {
		try {
			if (sizeof(ulong) > buf.Length - buf.Position) { return false; }
			item_seq = BitConverter.ToUInt64(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(ulong);

			if (sizeof(short) > buf.Length - buf.Position) { return false; }
			int item_name_len = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(short);
			if (item_name_len > buf.Length - buf.Position) { return false; }
			byte[] item_name_buf = new byte[item_name_len];
			Array.Copy(buf.GetBuffer(), (int)buf.Position, item_name_buf, 0, item_name_len);
			item_name = System.Text.Encoding.UTF8.GetString(item_name_buf);
			buf.Position += item_name_len;

		} catch (System.Exception) {
			return false;
		}
		return true;
	} // end Decode()

} // end class

public struct Item_Serializer {
	public static bool Encode(MemoryStream buf, Item o) { return o.Encode(buf); }
	public static bool Decode(ref Item o, MemoryStream buf) { return o.Decode(buf); }
	public static int Size(Item o) { return o.Size(); }
}

public class TestMsg {
	public const int MSG_ID = 7000;
	public sbyte m_s8 = 0;
	public byte m_u8 = 0;
	public short m_s16 = 0;
	public ushort m_u16 = 0;
	public int m_s32 = 0;
	public uint m_u32 = 0;
	public long m_s64 = 0;
	public ulong m_u64 = 0;
	public string m_str = "";
	public List<Item> m_items = new List<Item>();
	public int Size() {
		int size = 0;
		try {
			size += sizeof(sbyte);
			size += sizeof(byte);
			size += sizeof(short);
			size += sizeof(ushort);
			size += sizeof(int);
			size += sizeof(uint);
			size += sizeof(long);
			size += sizeof(ulong);
			size += sizeof(short);
			if (null != m_str) { size += Encoding.UTF8.GetByteCount(m_str); }
			size += sizeof(short);
			foreach(var iter in m_items) {
				Item o = iter;
				size += Item_Serializer.Size(o);
			}
		} catch (System.Exception) {
			return -1;
		}
		return size;
	} // end Size()

	public bool Encode(MemoryStream buf) {
		try {
			buf.Write(BitConverter.GetBytes(m_s8), 0, sizeof(sbyte));
			buf.Write(BitConverter.GetBytes(m_u8), 0, sizeof(byte));
			buf.Write(BitConverter.GetBytes(m_s16), 0, sizeof(short));
			buf.Write(BitConverter.GetBytes(m_u16), 0, sizeof(ushort));
			buf.Write(BitConverter.GetBytes(m_s32), 0, sizeof(int));
			buf.Write(BitConverter.GetBytes(m_u32), 0, sizeof(uint));
			buf.Write(BitConverter.GetBytes(m_s64), 0, sizeof(long));
			buf.Write(BitConverter.GetBytes(m_u64), 0, sizeof(ulong));
			if (null != m_str) {
				int len = Encoding.UTF8.GetByteCount(m_str);
				buf.Write(BitConverter.GetBytes(len), 0, sizeof(short));
				buf.Write(Encoding.UTF8.GetBytes(m_str), 0, len);
			} else {
				buf.Write(BitConverter.GetBytes(0), 0, sizeof(short));
			}
			buf.Write(BitConverter.GetBytes(m_items.Count), 0, sizeof(short));
			foreach(var iter in m_items) {
				Item o = iter;
				if (false == Item_Serializer.Encode(buf, o)) { return false; }
			}
		} catch (System.Exception) {
			return false;
		}
		return true;
	} // end Encode()

	public bool Decode(MemoryStream buf) {
		try {
			if (sizeof(sbyte) > buf.Length - buf.Position) { return false; }
			m_s8 = (sbyte)buf.GetBuffer()[buf.Position];
			buf.Position += sizeof(sbyte);

			if (sizeof(byte) > buf.Length - buf.Position) { return false; }
			m_u8 = (byte)buf.GetBuffer()[buf.Position];
			buf.Position += sizeof(byte);

			if (sizeof(short) > buf.Length - buf.Position) { return false; }
			m_s16 = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(short);

			if (sizeof(ushort) > buf.Length - buf.Position) { return false; }
			m_u16 = BitConverter.ToUInt16(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(ushort);

			if (sizeof(int) > buf.Length - buf.Position) { return false; }
			m_s32 = BitConverter.ToInt32(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(int);

			if (sizeof(uint) > buf.Length - buf.Position) { return false; }
			m_u32 = BitConverter.ToUInt32(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(uint);

			if (sizeof(long) > buf.Length - buf.Position) { return false; }
			m_s64 = BitConverter.ToInt64(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(long);

			if (sizeof(ulong) > buf.Length - buf.Position) { return false; }
			m_u64 = BitConverter.ToUInt64(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(ulong);

			if (sizeof(short) > buf.Length - buf.Position) { return false; }
			int m_str_len = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(short);
			if (m_str_len > buf.Length - buf.Position) { return false; }
			byte[] m_str_buf = new byte[m_str_len];
			Array.Copy(buf.GetBuffer(), (int)buf.Position, m_str_buf, 0, m_str_len);
			m_str = System.Text.Encoding.UTF8.GetString(m_str_buf);
			buf.Position += m_str_len;

			if (sizeof(short) > buf.Length - buf.Position) { return false; }
			int m_items_len = BitConverter.ToInt16(buf.GetBuffer(), (int)buf.Position);
			buf.Position += sizeof(short);
			for (int i = 0; i < m_items_len; i++) {
				Item o = new Item();
				if (false == Item_Serializer.Decode(ref o, buf)) { return false; }
				m_items.Add(o);
			}
		} catch (System.Exception) {
			return false;
		}
		return true;
	} // end Decode()

} // end class

} // end namespace
#pragma warning restore 108, 219
