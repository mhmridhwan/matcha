// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/imageview.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha "gomatcha.io/matcha/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ImageResizeMode int32

const (
	ImageResizeMode_FIT     ImageResizeMode = 0
	ImageResizeMode_FILL    ImageResizeMode = 1
	ImageResizeMode_STRETCH ImageResizeMode = 2
	ImageResizeMode_CENTER  ImageResizeMode = 3
)

var ImageResizeMode_name = map[int32]string{
	0: "FIT",
	1: "FILL",
	2: "STRETCH",
	3: "CENTER",
}
var ImageResizeMode_value = map[string]int32{
	"FIT":     0,
	"FILL":    1,
	"STRETCH": 2,
	"CENTER":  3,
}

func (x ImageResizeMode) String() string {
	return proto.EnumName(ImageResizeMode_name, int32(x))
}
func (ImageResizeMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type ImageView struct {
	Image      *matcha.ImageOrResource `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	ResizeMode ImageResizeMode         `protobuf:"varint,2,opt,name=resizeMode,enum=matcha.view.ImageResizeMode" json:"resizeMode,omitempty"`
	Tint       *matcha.Color           `protobuf:"bytes,3,opt,name=tint" json:"tint,omitempty"`
	Scale      float64                 `protobuf:"fixed64,5,opt,name=scale" json:"scale,omitempty"`
}

func (m *ImageView) Reset()                    { *m = ImageView{} }
func (m *ImageView) String() string            { return proto.CompactTextString(m) }
func (*ImageView) ProtoMessage()               {}
func (*ImageView) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ImageView) GetImage() *matcha.ImageOrResource {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageView) GetResizeMode() ImageResizeMode {
	if m != nil {
		return m.ResizeMode
	}
	return ImageResizeMode_FIT
}

func (m *ImageView) GetTint() *matcha.Color {
	if m != nil {
		return m.Tint
	}
	return nil
}

func (m *ImageView) GetScale() float64 {
	if m != nil {
		return m.Scale
	}
	return 0
}

func init() {
	proto.RegisterType((*ImageView)(nil), "matcha.view.ImageView")
	proto.RegisterEnum("matcha.view.ImageResizeMode", ImageResizeMode_name, ImageResizeMode_value)
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/imageview.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xdd, 0xfc, 0x69, 0x75, 0xe2, 0x9f, 0xb0, 0x08, 0x86, 0xe2, 0x21, 0xf5, 0x14, 0x0a,
	0x26, 0x50, 0xaf, 0xea, 0x21, 0x21, 0xc5, 0x40, 0xab, 0x65, 0x0d, 0x1e, 0xbc, 0x25, 0x71, 0xa9,
	0x0b, 0xad, 0x1b, 0x36, 0xd1, 0x82, 0x8f, 0xe3, 0x1b, 0xf8, 0x86, 0xb2, 0xbb, 0x31, 0x14, 0xf1,
	0xb4, 0x33, 0xdf, 0xfc, 0xbe, 0x9d, 0xe1, 0x83, 0xc9, 0x8a, 0x6f, 0x8a, 0xb6, 0x7a, 0x2d, 0x42,
	0xc6, 0x23, 0x5d, 0x45, 0x75, 0x19, 0x7d, 0x30, 0xba, 0x8d, 0xd8, 0xa6, 0x58, 0x51, 0x59, 0x85,
	0xb5, 0xe0, 0x2d, 0xc7, 0x4e, 0x47, 0x4a, 0x69, 0x34, 0xfe, 0xdf, 0xa8, 0x3c, 0x9a, 0xbf, 0xf8,
	0x46, 0x70, 0x90, 0xc9, 0xfe, 0x89, 0xd1, 0x2d, 0xbe, 0x04, 0x5b, 0x0d, 0x3d, 0xe4, 0xa3, 0xc0,
	0x99, 0x9e, 0x85, 0x9d, 0x5d, 0x11, 0x0f, 0x82, 0xd0, 0x86, 0xbf, 0x8b, 0x8a, 0x12, 0x4d, 0xe1,
	0x6b, 0x00, 0x41, 0x1b, 0xf6, 0x49, 0x17, 0xfc, 0x85, 0x7a, 0x86, 0x8f, 0x82, 0xe3, 0xe9, 0x79,
	0xb8, 0x73, 0x81, 0x36, 0x92, 0x9e, 0x21, 0x3b, 0x3c, 0x1e, 0x83, 0xd5, 0xb2, 0xb7, 0xd6, 0x33,
	0xd5, 0xae, 0xa3, 0x5f, 0x5f, 0xc2, 0xd7, 0x5c, 0x10, 0x35, 0xc2, 0xa7, 0x60, 0x37, 0x55, 0xb1,
	0xa6, 0x9e, 0xed, 0xa3, 0x00, 0x11, 0xdd, 0x4c, 0x6e, 0xe0, 0xe4, 0xcf, 0xbf, 0x78, 0x08, 0xe6,
	0x2c, 0xcb, 0xdd, 0x3d, 0xbc, 0x0f, 0xd6, 0x2c, 0x9b, 0xcf, 0x5d, 0x84, 0x1d, 0x18, 0x3e, 0xe6,
	0x24, 0xcd, 0x93, 0x3b, 0xd7, 0xc0, 0x00, 0x83, 0x24, 0xbd, 0xcf, 0x53, 0xe2, 0x9a, 0xf1, 0x2d,
	0x8c, 0x18, 0x0f, 0xfb, 0x68, 0xba, 0xa7, 0x2e, 0xd5, 0xc5, 0xb1, 0xb3, 0x2c, 0xfb, 0x3c, 0x9e,
	0x2d, 0x29, 0x7d, 0x19, 0x87, 0x0b, 0xc5, 0x48, 0x69, 0x19, 0x97, 0x03, 0x95, 0xdc, 0xd5, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x69, 0x62, 0x55, 0x97, 0x01, 0x00, 0x00,
}
