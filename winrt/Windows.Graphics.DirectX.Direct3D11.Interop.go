package winrt

import (
	"syscall"
	"unsafe"

	"github.com/TKMAX777/winapi/dx11"
	"github.com/go-ole/go-ole"
)

// IDirect3DDxgiInterfaceAccess
// https://learn.microsoft.com/en-us/windows/win32/api/windows.graphics.directx.direct3d11.interop/ns-windows-graphics-directx-direct3d11-interop-idirect3ddxgiinterfaceaccess

//var IDirect3DDxgiInterfaceAccessID = ole.NewGUID("{A0B2F7C1-8D5F-4650-9D3E-9EAE3D9BC670}")

var IDirect3DDxgiInterfaceAccessClass = "Windows.Graphics.DirectX.Direct3D11.Interop.IDirect3DDxgiInterfaceAccess"

type IDirect3DDxgiInterfaceAccess struct {
	ole.IInspectable
}

type IDirect3DDxgiInterfaceAccessVtbl struct {
	ole.IInspectableVtbl
}

func (v *IDirect3DDxgiInterfaceAccess) VTable() *IDirect3DDxgiInterfaceAccessVtbl {
	return (*IDirect3DDxgiInterfaceAccessVtbl)(unsafe.Pointer(v.RawVTable))
}

// IID for the raw D3D11 texture
var ID3D11Texture2DID = ole.NewGUID("{C22D3D7A-3EF5-4CD9-A84E-C1A1B46F165F}")

type ID3D11Texture2D struct {
	ole.IUnknown
}
type ID3D11Texture2DVtbl struct {
	ole.IUnknownVtbl
	GetDesc uintptr // func(this *ID3D11Texture2D, desc *D3D11_TEXTURE2D_DESC) HRESULT
}

func (v *ID3D11Texture2D) VTable() *ID3D11Texture2DVtbl {
	return (*ID3D11Texture2DVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ID3D11Texture2D) GetDesc() (*dx11.D3D11_TEXTURE2D_DESC, error) {
	if v == nil {
		return &dx11.D3D11_TEXTURE2D_DESC{}, nil
	}

	var desc *dx11.D3D11_TEXTURE2D_DESC
	r1, _, _ := syscall.SyscallN(v.VTable().GetDesc, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&desc)))
	if r1 != ole.S_OK {
		return nil, ole.NewError(r1)
	}
	return desc, nil
}
