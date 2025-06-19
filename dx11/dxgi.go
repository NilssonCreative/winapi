package dx11

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

var IDXGIObjectID = ole.NewGUID("{aec22fb8-76f3-4639-9be0-28eb43a67a2e}")

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
