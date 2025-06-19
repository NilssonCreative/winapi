package dx11

import (
	"structs"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

var (
	d3d11DLL = windows.NewLazySystemDLL("d3d11.dll")

	IID_ID3D11Texture2D, _ = windows.GUIDFromString("{6f15aaf2-d208-4e89-9ab4-489535d34f9c}")
)

const D3D11_SDK_VERSION = 7

type D3D11_CREATE_DEVICE_FLAG uint32

const (
	D3D11_CREATE_DEVICE_SINGLETHREADED                                D3D11_CREATE_DEVICE_FLAG = 0x1
	D3D11_CREATE_DEVICE_DEBUG                                         D3D11_CREATE_DEVICE_FLAG = 0x2
	D3D11_CREATE_DEVICE_SWITCH_TO_REF                                 D3D11_CREATE_DEVICE_FLAG = 0x4
	D3D11_CREATE_DEVICE_PREVENT_INTERNAL_THREADING_OPTIMIZATIONS      D3D11_CREATE_DEVICE_FLAG = 0x8
	D3D11_CREATE_DEVICE_BGRA_SUPPORT                                  D3D11_CREATE_DEVICE_FLAG = 0x20
	D3D11_CREATE_DEVICE_DEBUGGABLE                                    D3D11_CREATE_DEVICE_FLAG = 0x40
	D3D11_CREATE_DEVICE_PREVENT_ALTERING_LAYER_SETTINGS_FROM_REGISTRY D3D11_CREATE_DEVICE_FLAG = 0x80
	D3D11_CREATE_DEVICE_DISABLE_GPU_TIMEOUT                           D3D11_CREATE_DEVICE_FLAG = 0x100
	D3D11_CREATE_DEVICE_VIDEO_SUPPORT                                 D3D11_CREATE_DEVICE_FLAG = 0x800
)

var ID3D11DeviceID = ole.NewGUID("{db6f6ddb-ac77-4e88-8253-819df9bbf140}")

type ID3D11Device struct {
	ole.IUnknown
	vtbl *ID3D11DeviceVtbl
}

func (obj *ID3D11Device) CreateTexture2D(desc *D3D11_TEXTURE2D_DESC, ppTexture2D **ID3D11Texture2D) error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CreateTexture2D,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(0),
		uintptr(unsafe.Pointer(ppTexture2D)),
	)
	return ole.NewError(ret)
}

func (v *ID3D11Device) VTable() *ID3D11DeviceVtbl {
	return (*ID3D11DeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ID3D11Device) GetImmediateContext() (pImmediateContext *ID3D11DeviceContext) {
	syscall.SyscallN(v.VTable().GetImmediateContext, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&pImmediateContext)))
	return pImmediateContext
}

var ID3D11DeviceContextID = ole.NewGUID("{c0bfa96c-e089-44fb-8eaf-26f8796190da}")

type ID3D11DeviceContext struct {
	_    structs.HostLayout
	vtbl *ID3D11DeviceContextVtbl
}

func (obj *ID3D11DeviceContext) CopyResourceDXGI(dst, src *IDXGIResource) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CopyResource,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	return int32(ret)
}
func (obj *ID3D11DeviceContext) CopyResource2D(dst, src *ID3D11Texture2D) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CopyResource,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
	)
	return int32(ret)
}
func (obj *ID3D11DeviceContext) CopySubresourceRegion2D(dst *ID3D11Texture2D, dstSubResource, dstX, dstY, dstZ uint32, src *ID3D11Texture2D, srcSubResource uint32, pSrcBox *D3D11_BOX) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CopySubresourceRegion,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(dstSubResource),
		uintptr(dstX),
		uintptr(dstY),
		uintptr(dstZ),
		uintptr(unsafe.Pointer(src)),
		uintptr(srcSubResource),
		uintptr(unsafe.Pointer(pSrcBox)),
	)
	return int32(ret)
}

func (obj *ID3D11DeviceContext) CopySubresourceRegion(dst *ID3D11Resource, dstSubResource, dstX, dstY, dstZ uint32, src *ID3D11Resource, srcSubResource uint32, pSrcBox *D3D11_BOX) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.CopySubresourceRegion,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(dstSubResource),
		uintptr(dstX),
		uintptr(dstY),
		uintptr(dstZ),
		uintptr(unsafe.Pointer(src)),
		uintptr(srcSubResource),
		uintptr(unsafe.Pointer(pSrcBox)),
	)
	return int32(ret)
}
func (obj *ID3D11DeviceContext) Release() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return int32(ret)
}

var pD3DCreateDevice = d3d11DLL.NewProc("D3D11CreateDevice")

// CreateDevice
// https://learn.microsoft.com/en-us/windows/win32/api/d3d11/nf-d3d11-d3d11createdevice
func D3D11CreateDevice(
	pAdapter *IDXGIAdapter,
	DriverType D3D_DRIVER_TYPE,
	Software win.HMODULE,
	Flags D3D11_CREATE_DEVICE_FLAG,
	pFeatureLevels *D3D_FEATURE_LEVEL,
	FeatureLevels int,
	SDKVersion uint32,
	ppDevice **ID3D11Device,
	pFeatureLevel *D3D_FEATURE_LEVEL,
	ppImmediateContext **ID3D11DeviceContext,
) error {
	r1, _, _ := pD3DCreateDevice.Call(
		uintptr(unsafe.Pointer(pAdapter)),
		uintptr(DriverType),
		uintptr(Software),
		uintptr(Flags),
		uintptr(unsafe.Pointer(pFeatureLevels)),
		uintptr(FeatureLevels),
		uintptr(SDKVersion),
		uintptr(unsafe.Pointer(ppDevice)),
		uintptr(unsafe.Pointer(pFeatureLevel)),
		uintptr(unsafe.Pointer(ppImmediateContext)),
	)
	if r1 != win.S_OK {
		return ole.NewError(r1)
	}
	return nil
}

var pCreateDirect3D11DeviceFromDXGIDevice = d3d11DLL.NewProc("CreateDirect3D11DeviceFromDXGIDevice")

func CreateDirect3D11DeviceFromDXGIDevice(dxgiDevice *IDXGIDevice, graphicsDevice **ole.IInspectable) error {
	r1, _, err := pCreateDirect3D11DeviceFromDXGIDevice.Call(uintptr(unsafe.Pointer(dxgiDevice)), uintptr(unsafe.Pointer(graphicsDevice)))
	if r1 != win.S_OK {
		return err
	}
	return nil
}

type ID3D11Resource struct {
	_    structs.HostLayout
	vtbl *ID3D11ResourceVtbl
}

func (obj *ID3D11Resource) Release() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return int32(ret)
}
