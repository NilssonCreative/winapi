package dx11

import (
	"structs"
	"syscall"
	"unsafe"

	"github.com/TKMAX777/winapi/com"
	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

type ID3D11Texture2D struct {
	_    structs.HostLayout
	vtbl *ID3D11Texture2DVtbl
}

func (obj *ID3D11Texture2D) GetDesc(desc *D3D11_TEXTURE2D_DESC) error {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetDesc,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(desc)),
	)
	return ole.NewError(ret)
}
func (obj *ID3D11Texture2D) Release() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		uintptr(unsafe.Pointer(obj)),
	)
	return int32(ret)
}
func (obj *ID3D11Texture2D) QueryInterface(iid windows.GUID, pp interface{}) int32 {
	return com.ReflectQueryInterface(obj, obj.vtbl.QueryInterface, &iid, pp)
}
