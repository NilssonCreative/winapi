package dx11

import "structs"

type DXGI_RATIONAL struct {
	_ structs.HostLayout

	Numerator   uint32
	Denominator uint32
}

type DXGI_MODE_ROTATION uint32

type DXGI_OUTPUT_DESC struct {
	_ structs.HostLayout

	DeviceName         [32]uint16
	DesktopCoordinates RECT
	AttachedToDesktop  uint32 // BOOL
	Rotation           DXGI_MODE_ROTATION
	Monitor            uintptr
}

type DXGI_MODE_DESC struct {
	_ structs.HostLayout

	Width            uint32
	Height           uint32
	Rational         DXGI_RATIONAL
	Format           uint32 // DXGI_FORMAT
	ScanlineOrdering uint32 // DXGI_MODE_SCANLINE_ORDER
	Scaling          uint32 // DXGI_MODE_SCALING
}

type DXGI_OUTDUPL_DESC struct {
	_ structs.HostLayout

	ModeDesc                   DXGI_MODE_DESC
	Rotation                   uint32 // DXGI_MODE_ROTATION
	DesktopImageInSystemMemory uint32 // BOOL
}

type DXGI_SAMPLE_DESC struct {
	_ structs.HostLayout

	Count   uint32
	Quality uint32
}

type POINT struct {
	_ structs.HostLayout

	X int32
	Y int32
}
type RECT struct {
	_ structs.HostLayout

	Left, Top, Right, Bottom int32
}

type DXGI_OUTDUPL_MOVE_RECT struct {
	_ structs.HostLayout

	Src  POINT
	Dest RECT
}
type DXGI_OUTDUPL_POINTER_POSITION struct {
	_ structs.HostLayout

	Position POINT
	Visible  uint32
}
type DXGI_OUTDUPL_FRAME_INFO struct {
	_ structs.HostLayout

	LastPresentTime           int64
	LastMouseUpdateTime       int64
	AccumulatedFrames         uint32
	RectsCoalesced            uint32
	ProtectedContentMaskedOut uint32
	PointerPosition           DXGI_OUTDUPL_POINTER_POSITION
	TotalMetadataBufferSize   uint32
	PointerShapeBufferSize    uint32
}
type DXGI_MAPPED_RECT struct {
	_ structs.HostLayout

	Pitch int32
	PBits uintptr
}

const (
	DXGI_FORMAT_R8G8B8A8_UNORM DXGI_FORMAT = 28
	DXGI_FORMAT_B8G8R8A8_UNORM DXGI_FORMAT = 87
)

type DXGI_OUTDUPL_POINTER_SHAPE_TYPE uint32

const (
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_MONOCHROME   DXGI_OUTDUPL_POINTER_SHAPE_TYPE = 1
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_COLOR        DXGI_OUTDUPL_POINTER_SHAPE_TYPE = 2
	DXGI_OUTDUPL_POINTER_SHAPE_TYPE_MASKED_COLOR DXGI_OUTDUPL_POINTER_SHAPE_TYPE = 4
)

type DXGI_OUTDUPL_POINTER_SHAPE_INFO struct {
	_ structs.HostLayout

	Type    DXGI_OUTDUPL_POINTER_SHAPE_TYPE
	Width   uint32
	Height  uint32
	Pitch   uint32
	HotSpot POINT
}
