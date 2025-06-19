package dx11

import (
	"structs"
	"syscall"
	"unsafe"

	"github.com/TKMAX777/winapi/com"
	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

type DXGI_FORMAT uint32

const (
	DXGI_MAP_READ    = 1 << 0
	DXGI_MAP_WRITE   = 1 << 1
	DXGI_MAP_DISCARD = 1 << 2
)

var (
	IDXGIObjectID = ole.NewGUID("{aec22fb8-76f3-4639-9be0-28eb43a67a2e}")

	// iid_IDXGIDevice, _   = windows.GUIDFromString("{54ec77fa-1377-44e6-8c32-88fd5f44c84c}")
	IID_IDXGIDevice1, _ = windows.GUIDFromString("{77db970f-6276-48ba-ba28-070143b4392c}")
	// IID_IDXGIAdapter, _  = windows.GUIDFromString("{2411E7E1-12AC-4CCF-BD14-9798E8534DC0}")
	IID_IDXGIAdapter1, _ = windows.GUIDFromString("{29038f61-3839-4626-91fd-086879011a05}")
	// IID_IDXGIOutput, _   = windows.GUIDFromString("{ae02eedb-c735-4690-8d52-5a8dc20213aa}")
	IID_IDXGIOutput1, _  = windows.GUIDFromString("{00cddea8-939b-4b83-a340-a685226666cc}")
	IID_IDXGIOutput5, _  = windows.GUIDFromString("{80A07424-AB52-42EB-833C-0C42FD282D98}")
	IID_IDXGIFactory1, _ = windows.GUIDFromString("{770aae78-f26f-4dba-a829-253c83d1b387}")
	// IID_IDXGIResource, _ = windows.GUIDFromString("{035f3ab4-482e-4e50-b41f-8a7f8bd8960b}")
	IID_IDXGISurface, _ = windows.GUIDFromString("{cafcb56c-6ac3-4889-bf47-9e23bbd260ec}")
)

type IDXGIAdapter uintptr

type IDXGIObject struct {
	ole.IUnknown
}

type IDXGIObjectVtbl struct {
	ole.IUnknownVtbl
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
}

func (v *IDXGIObject) VTable() *IDXGIObjectVtbl {
	return (*IDXGIObjectVtbl)(unsafe.Pointer(v.RawVTable))
}

var IDXGIDeviceID = ole.NewGUID("{54ec77fa-1377-44e6-8c32-88fd5f44c84c}")

type IDXGIDevice struct {
	IDXGIObject
}

type IDXGIDeviceVtbl struct {
	IDXGIObjectVtbl
	GetAdapter             uintptr
	CreateSurface          uintptr
	QueryResourceResidency uintptr
	SetGPUThreadPriority   uintptr
	GetGPUThreadPriority   uintptr
}

func (v *IDXGIDevice) VTable() *IDXGIDeviceVtbl {
	return (*IDXGIDeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

var (
	IID_IDirect3DDxgiInterfaceAccess = ole.NewGUID("{A9C3260C-7E24-4AD2-A62E-63910832823A}")
)

type IDirect3DDxgiInterfaceAccess struct {
	ole.IInspectable
}

type IDirect3DDxgiInterfaceAccessVtbl struct {
	ole.IInspectableVtbl
	GetInterface uintptr
}

// VTable returns the v-table pointer for this interface
func (v *IDirect3DDxgiInterfaceAccess) VTable() *IDirect3DDxgiInterfaceAccessVtbl {
	return (*IDirect3DDxgiInterfaceAccessVtbl)(unsafe.Pointer(v.RawVTable))
}

// GetInterface calls:
//
//	HRESULT GetInterface(IUnknown* this, REFIID riid, void** ppv)
//
// and returns the requested pointer in out.
func (v *IDirect3DDxgiInterfaceAccess) GetInterface(iid *ole.GUID, out unsafe.Pointer) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetInterface,
		uintptr(unsafe.Pointer(v)),   // this
		uintptr(unsafe.Pointer(iid)), // REFIID
		uintptr(out),                 // void** ppv
	)
	if windows.Handle(hr) != windows.S_OK {
		return ole.NewError(hr)
	}
	return nil
}

type IDXGISurface struct {
	_    structs.HostLayout
	vtbl *IDXGISurfaceVtbl
}

func (obj *IDXGISurface) QueryInterface(iid windows.GUID, pp interface{}) int32 {
	return com.ReflectQueryInterface(obj, obj.vtbl.QueryInterface, &iid, pp)
}
func (obj *IDXGISurface) Map(pLockedRect *DXGI_MAPPED_RECT, mapFlags uint32) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Map,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pLockedRect)),
		uintptr(mapFlags),
	)
	return int32(ret)
}
func (obj *IDXGISurface) Unmap() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Unmap,
		uintptr(unsafe.Pointer(obj)),
	)
	return int32(ret)
}
func (obj *IDXGISurface) Release() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return int32(ret)
}

type IDXGIResource struct {
	_    structs.HostLayout
	vtbl *IDXGIResourceVtbl
}

func (obj *IDXGIResource) QueryInterface(iid windows.GUID, pp interface{}) int32 {
	return com.ReflectQueryInterface(obj, obj.vtbl.QueryInterface, &iid, pp)
}
func (obj *IDXGIResource) Release() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return int32(ret)
}
