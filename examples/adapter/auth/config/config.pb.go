// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/auth/config/config.proto

/*
	Package config is a generated protocol buffer package.

	config for auth adapter

	It is generated from these files:
		mixer/adapter/auth/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// config for auth adapter
type Params struct {
	// test token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// URL for the config store. It is used to initiate a new Store instance.
	// Following are some examples of the config store URL:
	// * "k8s://"
	// * "fs:///tmp/testdata/configroot"
	ConfigStoreUrl string `protobuf:"bytes,2,opt,name=config_store_url,json=configStoreUrl,proto3" json:"config_store_url,omitempty"`
	// Redis connection string <hostname>:<port number>
	// ex) localhost:6379
	RedisServerUrl string `protobuf:"bytes,3,opt,name=redis_server_url,json=redisServerUrl,proto3" json:"redis_server_url,omitempty"`
	// Maximum number of idle connections to redis
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	ConnectionPoolSize int64 `protobuf:"varint,4,opt,name=connection_pool_size,json=connectionPoolSize,proto3" json:"connection_pool_size,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Params) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Params) GetConfigStoreUrl() string {
	if m != nil {
		return m.ConfigStoreUrl
	}
	return ""
}

func (m *Params) GetRedisServerUrl() string {
	if m != nil {
		return m.RedisServerUrl
	}
	return ""
}

func (m *Params) GetConnectionPoolSize() int64 {
	if m != nil {
		return m.ConnectionPoolSize
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "adapter.auth.config.Params")
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	if this.ConfigStoreUrl != that1.ConfigStoreUrl {
		return false
	}
	if this.RedisServerUrl != that1.RedisServerUrl {
		return false
	}
	if this.ConnectionPoolSize != that1.ConnectionPoolSize {
		return false
	}
	return true
}
func (this *Params) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&config.Params{")
	s = append(s, "Token: "+fmt.Sprintf("%#v", this.Token)+",\n")
	s = append(s, "ConfigStoreUrl: "+fmt.Sprintf("%#v", this.ConfigStoreUrl)+",\n")
	s = append(s, "RedisServerUrl: "+fmt.Sprintf("%#v", this.RedisServerUrl)+",\n")
	s = append(s, "ConnectionPoolSize: "+fmt.Sprintf("%#v", this.ConnectionPoolSize)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.ConfigStoreUrl) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ConfigStoreUrl)))
		i += copy(dAtA[i:], m.ConfigStoreUrl)
	}
	if len(m.RedisServerUrl) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.RedisServerUrl)))
		i += copy(dAtA[i:], m.RedisServerUrl)
	}
	if m.ConnectionPoolSize != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.ConnectionPoolSize))
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.ConfigStoreUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.RedisServerUrl)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ConnectionPoolSize != 0 {
		n += 1 + sovConfig(uint64(m.ConnectionPoolSize))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`ConfigStoreUrl:` + fmt.Sprintf("%v", this.ConfigStoreUrl) + `,`,
		`RedisServerUrl:` + fmt.Sprintf("%v", this.RedisServerUrl) + `,`,
		`ConnectionPoolSize:` + fmt.Sprintf("%v", this.ConnectionPoolSize) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigStoreUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigStoreUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedisServerUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedisServerUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionPoolSize", wireType)
			}
			m.ConnectionPoolSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnectionPoolSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/auth/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x06, 0xe0, 0x3c, 0x0a, 0x91, 0xc8, 0x80, 0x50, 0xc8, 0x10, 0x31, 0x3c, 0x55, 0x0c, 0x28,
	0x53, 0x82, 0x04, 0x27, 0xe0, 0x04, 0x55, 0x23, 0x16, 0x16, 0x2b, 0xa4, 0x26, 0x58, 0xa4, 0x7e,
	0x91, 0xed, 0x22, 0xd4, 0x89, 0x23, 0x70, 0x08, 0x06, 0x8e, 0xc2, 0xd8, 0x91, 0x91, 0x98, 0x85,
	0xb1, 0x47, 0x40, 0xb6, 0x2b, 0x75, 0x7a, 0xf6, 0xff, 0x7f, 0xb2, 0xfc, 0x92, 0xcb, 0xa5, 0x78,
	0xe5, 0xaa, 0x6a, 0x16, 0xcd, 0x60, 0xdc, 0x5c, 0x99, 0xa7, 0xaa, 0x25, 0xf9, 0x28, 0xba, 0xdd,
	0x28, 0x07, 0x45, 0x86, 0xd2, 0xb3, 0x9d, 0x28, 0x9d, 0x28, 0x43, 0x75, 0x9e, 0x75, 0xd4, 0x91,
	0xef, 0x2b, 0x77, 0x0a, 0xf4, 0xe2, 0x03, 0x92, 0x78, 0xd6, 0xa8, 0x66, 0xa9, 0xd3, 0x2c, 0x39,
	0x32, 0xf4, 0xcc, 0x65, 0x0e, 0x53, 0x28, 0x8e, 0xe7, 0xe1, 0x92, 0x16, 0xc9, 0x69, 0x78, 0x80,
	0x69, 0x43, 0x8a, 0xb3, 0x95, 0xea, 0xf3, 0x03, 0x0f, 0x4e, 0x42, 0x5e, 0xbb, 0xf8, 0x4e, 0xf5,
	0x4e, 0x2a, 0xbe, 0x10, 0x9a, 0x69, 0xae, 0x5e, 0xb8, 0xf2, 0x72, 0x12, 0xa4, 0xcf, 0x6b, 0x1f,
	0x3b, 0x79, 0x95, 0x64, 0x2d, 0x49, 0xc9, 0x5b, 0x23, 0x48, 0xb2, 0x81, 0xa8, 0x67, 0x5a, 0xac,
	0x79, 0x7e, 0x38, 0x85, 0x62, 0x32, 0x4f, 0xf7, 0xdd, 0x8c, 0xa8, 0xaf, 0xc5, 0x9a, 0xdf, 0xde,
	0x6c, 0x46, 0x8c, 0xbe, 0x47, 0x8c, 0xb6, 0x23, 0xc2, 0x9b, 0x45, 0xf8, 0xb4, 0x08, 0x5f, 0x16,
	0x61, 0x63, 0x11, 0x7e, 0x2c, 0xc2, 0x9f, 0xc5, 0x68, 0x6b, 0x11, 0xde, 0x7f, 0x31, 0xba, 0x8f,
	0xc3, 0xcf, 0x1e, 0x62, 0xbf, 0xe3, 0xf5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xe2, 0xcc,
	0x58, 0x38, 0x01, 0x00, 0x00,
}
