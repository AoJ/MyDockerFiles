// Code generated by protoc-gen-gogo.
// source: append_entries_request.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io1 "io"
import code_google_com_p_gogoprotobuf_proto2 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"

import fmt3 "fmt"
import strings2 "strings"
import reflect2 "reflect"

import fmt4 "fmt"
import strings3 "strings"
import code_google_com_p_gogoprotobuf_proto3 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import sort1 "sort"
import strconv1 "strconv"
import reflect3 "reflect"

import fmt5 "fmt"
import bytes1 "bytes"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type AppendEntriesRequest struct {
	Term             *uint64     `protobuf:"varint,1,req" json:"Term,omitempty"`
	PrevLogIndex     *uint64     `protobuf:"varint,2,req" json:"PrevLogIndex,omitempty"`
	PrevLogTerm      *uint64     `protobuf:"varint,3,req" json:"PrevLogTerm,omitempty"`
	CommitIndex      *uint64     `protobuf:"varint,4,req" json:"CommitIndex,omitempty"`
	LeaderName       *string     `protobuf:"bytes,5,req" json:"LeaderName,omitempty"`
	Entries          []*LogEntry `protobuf:"bytes,6,rep" json:"Entries,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *AppendEntriesRequest) Reset()      { *m = AppendEntriesRequest{} }
func (*AppendEntriesRequest) ProtoMessage() {}

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil && m.PrevLogIndex != nil {
		return *m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil && m.PrevLogTerm != nil {
		return *m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
}
func (m *AppendEntriesRequest) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io1.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return proto.ErrWrongType
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Term = &v
		case 2:
			if wireType != 0 {
				return proto.ErrWrongType
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevLogIndex = &v
		case 3:
			if wireType != 0 {
				return proto.ErrWrongType
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PrevLogTerm = &v
		case 4:
			if wireType != 0 {
				return proto.ErrWrongType
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CommitIndex = &v
		case 5:
			if wireType != 2 {
				return proto.ErrWrongType
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.LeaderName = &s
			index = postIndex
		case 6:
			if wireType != 2 {
				return proto.ErrWrongType
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io1.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io1.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &LogEntry{})
			m.Entries[len(m.Entries)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto2.Skip(data[index:])
			if err != nil {
				return err
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *AppendEntriesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings2.Join([]string{`&AppendEntriesRequest{`,
		`Term:` + valueToStringAppendEntriesRequest(this.Term) + `,`,
		`PrevLogIndex:` + valueToStringAppendEntriesRequest(this.PrevLogIndex) + `,`,
		`PrevLogTerm:` + valueToStringAppendEntriesRequest(this.PrevLogTerm) + `,`,
		`CommitIndex:` + valueToStringAppendEntriesRequest(this.CommitIndex) + `,`,
		`LeaderName:` + valueToStringAppendEntriesRequest(this.LeaderName) + `,`,
		`Entries:` + strings2.Replace(fmt3.Sprintf("%v", this.Entries), "LogEntry", "LogEntry", 1) + `,`,
		`XXX_unrecognized:` + fmt3.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAppendEntriesRequest(v interface{}) string {
	rv := reflect2.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect2.Indirect(rv).Interface()
	return fmt3.Sprintf("*%v", pv)
}
func (m *AppendEntriesRequest) Size() (n int) {
	var l int
	_ = l
	if m.Term != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.Term))
	}
	if m.PrevLogIndex != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.PrevLogIndex))
	}
	if m.PrevLogTerm != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.PrevLogTerm))
	}
	if m.CommitIndex != nil {
		n += 1 + sovAppendEntriesRequest(uint64(*m.CommitIndex))
	}
	if m.LeaderName != nil {
		l = len(*m.LeaderName)
		n += 1 + l + sovAppendEntriesRequest(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovAppendEntriesRequest(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAppendEntriesRequest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAppendEntriesRequest(x uint64) (n int) {
	return sovAppendEntriesRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
	return sovAppendEntriesRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedAppendEntriesRequest(r randyAppendEntriesRequest, easy bool) *AppendEntriesRequest {
	this := &AppendEntriesRequest{}
	v1 := uint64(r.Uint32())
	this.Term = &v1
	v2 := uint64(r.Uint32())
	this.PrevLogIndex = &v2
	v3 := uint64(r.Uint32())
	this.PrevLogTerm = &v3
	v4 := uint64(r.Uint32())
	this.CommitIndex = &v4
	v5 := randStringAppendEntriesRequest(r)
	this.LeaderName = &v5
	if r.Intn(10) != 0 {
		v6 := r.Intn(10)
		this.Entries = make([]*LogEntry, v6)
		for i := 0; i < v6; i++ {
			this.Entries[i] = NewPopulatedLogEntry(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedAppendEntriesRequest(r, 7)
	}
	return this
}

type randyAppendEntriesRequest interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAppendEntriesRequest(r randyAppendEntriesRequest) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringAppendEntriesRequest(r randyAppendEntriesRequest) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneAppendEntriesRequest(r)
	}
	return string(tmps)
}
func randUnrecognizedAppendEntriesRequest(r randyAppendEntriesRequest, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldAppendEntriesRequest(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldAppendEntriesRequest(data []byte, r randyAppendEntriesRequest, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateAppendEntriesRequest(data, uint64(key))
		data = encodeVarintPopulateAppendEntriesRequest(data, uint64(r.Int63()))
	case 1:
		data = encodeVarintPopulateAppendEntriesRequest(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateAppendEntriesRequest(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateAppendEntriesRequest(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateAppendEntriesRequest(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateAppendEntriesRequest(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *AppendEntriesRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *AppendEntriesRequest) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != nil {
		data[i] = 0x8
		i++
		i = encodeVarintAppendEntriesRequest(data, i, uint64(*m.Term))
	}
	if m.PrevLogIndex != nil {
		data[i] = 0x10
		i++
		i = encodeVarintAppendEntriesRequest(data, i, uint64(*m.PrevLogIndex))
	}
	if m.PrevLogTerm != nil {
		data[i] = 0x18
		i++
		i = encodeVarintAppendEntriesRequest(data, i, uint64(*m.PrevLogTerm))
	}
	if m.CommitIndex != nil {
		data[i] = 0x20
		i++
		i = encodeVarintAppendEntriesRequest(data, i, uint64(*m.CommitIndex))
	}
	if m.LeaderName != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintAppendEntriesRequest(data, i, uint64(len(*m.LeaderName)))
		i += copy(data[i:], *m.LeaderName)
	}
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			data[i] = 0x32
			i++
			i = encodeVarintAppendEntriesRequest(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64AppendEntriesRequest(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32AppendEntriesRequest(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintAppendEntriesRequest(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *AppendEntriesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings3.Join([]string{`&protobuf.AppendEntriesRequest{` + `Term:` + valueToGoStringAppendEntriesRequest(this.Term, "uint64"), `PrevLogIndex:` + valueToGoStringAppendEntriesRequest(this.PrevLogIndex, "uint64"), `PrevLogTerm:` + valueToGoStringAppendEntriesRequest(this.PrevLogTerm, "uint64"), `CommitIndex:` + valueToGoStringAppendEntriesRequest(this.CommitIndex, "uint64"), `LeaderName:` + valueToGoStringAppendEntriesRequest(this.LeaderName, "string"), `Entries:` + fmt4.Sprintf("%#v", this.Entries), `XXX_unrecognized:` + fmt4.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringAppendEntriesRequest(v interface{}, typ string) string {
	rv := reflect3.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect3.Indirect(rv).Interface()
	return fmt4.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringAppendEntriesRequest(e map[int32]code_google_com_p_gogoprotobuf_proto3.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort1.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv1.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings3.Join(ss, ",") + "}"
	return s
}
func (this *AppendEntriesRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt5.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AppendEntriesRequest)
	if !ok {
		return fmt5.Errorf("that is not of type *AppendEntriesRequest")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt5.Errorf("that is type *AppendEntriesRequest but is nil && this != nil")
	} else if this == nil {
		return fmt5.Errorf("that is type *AppendEntriesRequestbut is not nil && this == nil")
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return fmt5.Errorf("Term this(%v) Not Equal that(%v)", *this.Term, *that1.Term)
		}
	} else if this.Term != nil {
		return fmt5.Errorf("this.Term == nil && that.Term != nil")
	} else if that1.Term != nil {
		return fmt5.Errorf("Term this(%v) Not Equal that(%v)", this.Term, that1.Term)
	}
	if this.PrevLogIndex != nil && that1.PrevLogIndex != nil {
		if *this.PrevLogIndex != *that1.PrevLogIndex {
			return fmt5.Errorf("PrevLogIndex this(%v) Not Equal that(%v)", *this.PrevLogIndex, *that1.PrevLogIndex)
		}
	} else if this.PrevLogIndex != nil {
		return fmt5.Errorf("this.PrevLogIndex == nil && that.PrevLogIndex != nil")
	} else if that1.PrevLogIndex != nil {
		return fmt5.Errorf("PrevLogIndex this(%v) Not Equal that(%v)", this.PrevLogIndex, that1.PrevLogIndex)
	}
	if this.PrevLogTerm != nil && that1.PrevLogTerm != nil {
		if *this.PrevLogTerm != *that1.PrevLogTerm {
			return fmt5.Errorf("PrevLogTerm this(%v) Not Equal that(%v)", *this.PrevLogTerm, *that1.PrevLogTerm)
		}
	} else if this.PrevLogTerm != nil {
		return fmt5.Errorf("this.PrevLogTerm == nil && that.PrevLogTerm != nil")
	} else if that1.PrevLogTerm != nil {
		return fmt5.Errorf("PrevLogTerm this(%v) Not Equal that(%v)", this.PrevLogTerm, that1.PrevLogTerm)
	}
	if this.CommitIndex != nil && that1.CommitIndex != nil {
		if *this.CommitIndex != *that1.CommitIndex {
			return fmt5.Errorf("CommitIndex this(%v) Not Equal that(%v)", *this.CommitIndex, *that1.CommitIndex)
		}
	} else if this.CommitIndex != nil {
		return fmt5.Errorf("this.CommitIndex == nil && that.CommitIndex != nil")
	} else if that1.CommitIndex != nil {
		return fmt5.Errorf("CommitIndex this(%v) Not Equal that(%v)", this.CommitIndex, that1.CommitIndex)
	}
	if this.LeaderName != nil && that1.LeaderName != nil {
		if *this.LeaderName != *that1.LeaderName {
			return fmt5.Errorf("LeaderName this(%v) Not Equal that(%v)", *this.LeaderName, *that1.LeaderName)
		}
	} else if this.LeaderName != nil {
		return fmt5.Errorf("this.LeaderName == nil && that.LeaderName != nil")
	} else if that1.LeaderName != nil {
		return fmt5.Errorf("LeaderName this(%v) Not Equal that(%v)", this.LeaderName, that1.LeaderName)
	}
	if len(this.Entries) != len(that1.Entries) {
		return fmt5.Errorf("Entries this(%v) Not Equal that(%v)", len(this.Entries), len(that1.Entries))
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(that1.Entries[i]) {
			return fmt5.Errorf("Entries this[%v](%v) Not Equal that[%v](%v)", i, this.Entries[i], i, that1.Entries[i])
		}
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt5.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *AppendEntriesRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AppendEntriesRequest)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return false
		}
	} else if this.Term != nil {
		return false
	} else if that1.Term != nil {
		return false
	}
	if this.PrevLogIndex != nil && that1.PrevLogIndex != nil {
		if *this.PrevLogIndex != *that1.PrevLogIndex {
			return false
		}
	} else if this.PrevLogIndex != nil {
		return false
	} else if that1.PrevLogIndex != nil {
		return false
	}
	if this.PrevLogTerm != nil && that1.PrevLogTerm != nil {
		if *this.PrevLogTerm != *that1.PrevLogTerm {
			return false
		}
	} else if this.PrevLogTerm != nil {
		return false
	} else if that1.PrevLogTerm != nil {
		return false
	}
	if this.CommitIndex != nil && that1.CommitIndex != nil {
		if *this.CommitIndex != *that1.CommitIndex {
			return false
		}
	} else if this.CommitIndex != nil {
		return false
	} else if that1.CommitIndex != nil {
		return false
	}
	if this.LeaderName != nil && that1.LeaderName != nil {
		if *this.LeaderName != *that1.LeaderName {
			return false
		}
	} else if this.LeaderName != nil {
		return false
	} else if that1.LeaderName != nil {
		return false
	}
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(that1.Entries[i]) {
			return false
		}
	}
	if !bytes1.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
