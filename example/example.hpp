//////////////////////////////////////////////////////////////////
// Automatically-generated file. Do not edit!
//////////////////////////////////////////////////////////////////

#ifndef __EXAMPLE_HPP__
#define __EXAMPLE_HPP__

#include <string>
#include <vector>
#include <list>
#include <stdint.h>
#include <cstring>
struct Item {
	uint64_t item_seq;
	std::string item_name;
	Item() {
		Clear();
	}

	void Clear() {
		item_seq = 0;
		item_name = "";
	} // end Clear()

	int32_t Size() const {
		int32_t size = 0;
		size += sizeof(uint64_t);
		size += sizeof(uint16_t);
		size += item_name.length();
		return size;
	} // end Size()

	bool Encode(std::vector<char>& buf) const {
		size_t size = Size();
		if (0 == size) { return true; }
		if (size > buf.size()) { buf.resize(size); }
		char * pBuf = &(buf[0]);
		if (false == Encode(&pBuf)) { return false; }
		return true;
	} // end Encode()
	bool Encode(char** buf) const {
		std::memcpy(*buf, &item_seq, sizeof(uint64_t));
		(*buf) += sizeof(uint64_t);
		size_t item_name_size = item_name.length();
		std::memcpy(*buf, &item_name_size, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t);
		std::memcpy(*buf, item_name.c_str(), item_name_size);
		(*buf) += item_name_size;
		return true;
	} // end Encode()
	bool Decode(const std::vector<char>& buf) {
		size_t size = buf.size();
		if (0 == size) { return true; }
		const char * pBuf = &(buf[0]);
		if (false == Decode(&pBuf, size)) { return false; }
		return true;
	} // end Decode()
	bool Decode(const char** buf, size_t & size) {
		if (sizeof(uint64_t) > size) { return false; }
		std::memcpy(&item_seq, *buf, sizeof(uint64_t));
		(*buf) += sizeof(uint64_t); size -= sizeof(uint64_t);
		if (sizeof(uint16_t) > size) { return false; }
		uint16_t item_name_len = 0;
		std::memcpy(&item_name_len, *buf, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t); size -= sizeof(uint16_t);
		if (size < item_name_len) { return false; }
		item_name.assign((char*)*buf, item_name_len);
		(*buf) += item_name_len; size -= item_name_len;
		return true;
	} // end Encode()
}; // end struct
struct Item_Serializer {
	static bool Encode(char** buf, const Item & o) { return o.Encode(buf); }
	static bool Decode(Item & o, const char** buf, size_t & size) { return o.Decode(buf, size); }
	static int32_t Size(const Item & o) { return o.Size(); }
};

struct TestMsg {
	int8_t m_s8;
	uint8_t m_u8;
	int16_t m_s16;
	uint16_t m_u16;
	int32_t m_s32;
	uint32_t m_u32;
	int64_t m_s64;
	uint64_t m_u64;
	std::string m_str;
	std::list<Item> m_items;
	TestMsg() {
		Clear();
	}

	void Clear() {
		m_s8 = 0;
		m_u8 = 0;
		m_s16 = 0;
		m_u16 = 0;
		m_s32 = 0;
		m_u32 = 0;
		m_s64 = 0;
		m_u64 = 0;
		m_str = "";
		m_items.clear();
	} // end Clear()

	int32_t Size() const {
		int32_t size = 0;
		size += sizeof(int8_t);
		size += sizeof(uint8_t);
		size += sizeof(int16_t);
		size += sizeof(uint16_t);
		size += sizeof(int32_t);
		size += sizeof(uint32_t);
		size += sizeof(int64_t);
		size += sizeof(uint64_t);
		size += sizeof(uint16_t);
		size += m_str.length();
		size += sizeof(uint16_t);
		for(std::list<Item>::const_iterator iter = m_items.begin(); iter != m_items.end(); ++iter) {
			size += Item_Serializer::Size(*iter);
		}
		return size;
	} // end Size()

	bool Encode(std::vector<char>& buf) const {
		size_t size = Size();
		if (0 == size) { return true; }
		if (size > buf.size()) { buf.resize(size); }
		char * pBuf = &(buf[0]);
		if (false == Encode(&pBuf)) { return false; }
		return true;
	} // end Encode()
	bool Encode(char** buf) const {
		std::memcpy(*buf, &m_s8, sizeof(int8_t));
		(*buf) += sizeof(int8_t);
		std::memcpy(*buf, &m_u8, sizeof(uint8_t));
		(*buf) += sizeof(uint8_t);
		std::memcpy(*buf, &m_s16, sizeof(int16_t));
		(*buf) += sizeof(int16_t);
		std::memcpy(*buf, &m_u16, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t);
		std::memcpy(*buf, &m_s32, sizeof(int32_t));
		(*buf) += sizeof(int32_t);
		std::memcpy(*buf, &m_u32, sizeof(uint32_t));
		(*buf) += sizeof(uint32_t);
		std::memcpy(*buf, &m_s64, sizeof(int64_t));
		(*buf) += sizeof(int64_t);
		std::memcpy(*buf, &m_u64, sizeof(uint64_t));
		(*buf) += sizeof(uint64_t);
		size_t m_str_size = m_str.length();
		std::memcpy(*buf, &m_str_size, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t);
		std::memcpy(*buf, m_str.c_str(), m_str_size);
		(*buf) += m_str_size;
		size_t m_items_size = m_items.size();
		std::memcpy(*buf, &m_items_size, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t);
		for(std::list<Item>::const_iterator iter = m_items.begin(); iter != m_items.end(); ++iter) {
			if (false == Item_Serializer::Encode(buf, *iter)) {
				return false;
			}
		}
		return true;
	} // end Encode()
	bool Decode(const std::vector<char>& buf) {
		size_t size = buf.size();
		if (0 == size) { return true; }
		const char * pBuf = &(buf[0]);
		if (false == Decode(&pBuf, size)) { return false; }
		return true;
	} // end Decode()
	bool Decode(const char** buf, size_t & size) {
		if (sizeof(int8_t) > size) { return false; }
		std::memcpy(&m_s8, *buf, sizeof(int8_t));
		(*buf) += sizeof(int8_t); size -= sizeof(int8_t);
		if (sizeof(uint8_t) > size) { return false; }
		std::memcpy(&m_u8, *buf, sizeof(uint8_t));
		(*buf) += sizeof(uint8_t); size -= sizeof(uint8_t);
		if (sizeof(int16_t) > size) { return false; }
		std::memcpy(&m_s16, *buf, sizeof(int16_t));
		(*buf) += sizeof(int16_t); size -= sizeof(int16_t);
		if (sizeof(uint16_t) > size) { return false; }
		std::memcpy(&m_u16, *buf, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t); size -= sizeof(uint16_t);
		if (sizeof(int32_t) > size) { return false; }
		std::memcpy(&m_s32, *buf, sizeof(int32_t));
		(*buf) += sizeof(int32_t); size -= sizeof(int32_t);
		if (sizeof(uint32_t) > size) { return false; }
		std::memcpy(&m_u32, *buf, sizeof(uint32_t));
		(*buf) += sizeof(uint32_t); size -= sizeof(uint32_t);
		if (sizeof(int64_t) > size) { return false; }
		std::memcpy(&m_s64, *buf, sizeof(int64_t));
		(*buf) += sizeof(int64_t); size -= sizeof(int64_t);
		if (sizeof(uint64_t) > size) { return false; }
		std::memcpy(&m_u64, *buf, sizeof(uint64_t));
		(*buf) += sizeof(uint64_t); size -= sizeof(uint64_t);
		if (sizeof(uint16_t) > size) { return false; }
		uint16_t m_str_len = 0;
		std::memcpy(&m_str_len, *buf, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t); size -= sizeof(uint16_t);
		if (size < m_str_len) { return false; }
		m_str.assign((char*)*buf, m_str_len);
		(*buf) += m_str_len; size -= m_str_len;
		if (sizeof(uint16_t) > size) { return false; }
		uint16_t m_items_len = 0;
		std::memcpy(&m_items_len, *buf, sizeof(uint16_t));
		(*buf) += sizeof(uint16_t); size -= sizeof(uint16_t);
		for (uint16_t i = 0; i < m_items_len; ++i) {
			Item val;
			if (false == Item_Serializer::Decode(val, buf, size)) {
				return false;
			}
			m_items.push_back(val);
		}
		return true;
	} // end Encode()
}; // end struct
#endif // __EXAMPLE_HPP__
