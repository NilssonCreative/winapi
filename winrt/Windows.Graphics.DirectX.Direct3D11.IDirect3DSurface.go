package winrt

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

// IDirect3DSurface
// https://learn.microsoft.com/en-us/uwp/api/windows.graphics.directx.direct3d11.idirect3dsurface?view=winrt-22621

// var IDirect3DSurfaceID = ole.NewGUID("{A0B2F7C1-8D5F-4650-9D3E-9EAE3D9BC670}")
var IID_IDirect3DSurface = ole.NewGUID("{A37624AB-8D5F-4650-9D3E-9EAE3D9BC670}")
var IDirect3DSurfaceClass = "Windows.Graphics.DirectX.Direct3D11.IDirect3DSurface"

type IDirect3DSurface struct {
	ole.IInspectable
}

type IDirect3DSurfaceVtbl struct {
	ole.IInspectableVtbl
	Surface uintptr
}

func (s *IDirect3DSurface) VTable() *IDirect3DSurfaceVtbl {
	return (*IDirect3DSurfaceVtbl)(unsafe.Pointer(s.RawVTable))
}

func (s *IDirect3DSurface) Surface() (*ID3D11Texture2D, error) {
	if s == nil {
		return nil, ole.NewError(ole.E_POINTER)
	}

	var raw *ID3D11Texture2D
	r1, _, _ := syscall.SyscallN(s.VTable().Surface, uintptr(unsafe.Pointer(s)), uintptr(unsafe.Pointer(&raw)))
	if r1 != ole.S_OK {
		return nil, ole.NewError(r1)
	}
	return raw, nil
}
