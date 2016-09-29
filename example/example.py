##################################################################
## Automatically-generated file. Do not edit!
##################################################################

import struct

class Effect :
	def __init__(self) :
		self.seq = 0
	def Encode(self) :
		data = ''
		data += struct.pack('q', self.seq)
		return data
	def Decode(self, buf) :
		self.seq = struct.unpack('q', buf[0:8])[0]; buf = buf[8:]
		return buf
def EffectEncode(o) :
	return o.Encode()
def EffectDecode(buf) :
	o = Effect()
	buf = o.Decode(buf)
	return [o, buf]

class Item :
	def __init__(self) :
		self.item_seq = 0
		self.item_name = ''
		self.effect = Effect()
	def Encode(self) :
		data = ''
		data += struct.pack('Q', self.item_seq)
		data += struct.pack('H', len(self.item_name))
		fmt = str(len(self.item_name)) + 's'
		data += struct.pack(fmt, self.item_name)
		data += EffectEncode(self.effect)
		return data
	def Decode(self, buf) :
		self.item_seq = struct.unpack('Q', buf[0:8])[0]; buf = buf[8:]
		length = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
		self.item_name = buf[0:length]; buf = buf[length:]
		[self.effect, buf] = EffectDecode(buf)
		return buf
def ItemEncode(o) :
	return o.Encode()
def ItemDecode(buf) :
	o = Item()
	buf = o.Decode(buf)
	return [o, buf]

class TestMsg :
	MSG_ID = 7000
	def __init__(self) :
		self.m_s8 = 0
		self.m_u8 = 0
		self.m_s16 = 0
		self.m_u16 = 0
		self.m_s32 = 0
		self.m_u32 = 0
		self.m_s64 = 0
		self.m_u64 = 0
		self.m_str = ''
		self.m_items = []
		self.m_strarr = []
		self.m_u32arr = []
	def Encode(self) :
		data = ''
		data += struct.pack('b', self.m_s8)
		data += struct.pack('B', self.m_u8)
		data += struct.pack('h', self.m_s16)
		data += struct.pack('H', self.m_u16)
		data += struct.pack('i', self.m_s32)
		data += struct.pack('I', self.m_u32)
		data += struct.pack('q', self.m_s64)
		data += struct.pack('Q', self.m_u64)
		data += struct.pack('H', len(self.m_str))
		fmt = str(len(self.m_str)) + 's'
		data += struct.pack(fmt, self.m_str)
		data += struct.pack('H', len(self.m_items))
		for o in self.m_items :
			data += ItemEncode(o)
		data += struct.pack('H', len(self.m_strarr))
		for o in self.m_strarr :
			data += struct.pack('H', len(o))
			fmt = str(len(o)) + 's'
			data += struct.pack(fmt, o)
		data += struct.pack('H', len(self.m_u32arr))
		for o in self.m_u32arr :
			data += struct.pack('I', o)
		return data
	def Decode(self, buf) :
		self.m_s8 = struct.unpack('b', buf[0:1])[0]; buf = buf[1:]
		self.m_u8 = struct.unpack('B', buf[0:1])[0]; buf = buf[1:]
		self.m_s16 = struct.unpack('h', buf[0:2])[0]; buf = buf[2:]
		self.m_u16 = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
		self.m_s32 = struct.unpack('i', buf[0:4])[0]; buf = buf[4:]
		self.m_u32 = struct.unpack('I', buf[0:4])[0]; buf = buf[4:]
		self.m_s64 = struct.unpack('q', buf[0:8])[0]; buf = buf[8:]
		self.m_u64 = struct.unpack('Q', buf[0:8])[0]; buf = buf[8:]
		length = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
		self.m_str = buf[0:length]; buf = buf[length:]
		list_count = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
		for i in range(list_count) :
			[o, buf] = ItemDecode(buf)
			self.m_items.append(o)
		list_count = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
		for i in range(list_count) :
			length = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
			o = buf[0:length]; buf = buf[length:]
			self.m_strarr.append(o)
		list_count = struct.unpack('H', buf[0:2])[0]; buf = buf[2:]
		for i in range(list_count) :
			o = struct.unpack('I', buf[0:4])[0]; buf = buf[4:]
			self.m_u32arr.append(o)
		return buf

