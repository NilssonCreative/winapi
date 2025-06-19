package dx11

import "structs"

type IDXGIDeviceSubObjectVtbl struct {
	_ structs.HostLayout
	IDXGIObjectVtbl

	GetDevice uintptr
}

type IDXGISurfaceVtbl struct {
	_ structs.HostLayout
	IDXGIDeviceSubObjectVtbl

	GetDesc uintptr
	Map     uintptr
	Unmap   uintptr
}

type IDXGIResourceVtbl struct {
	_ structs.HostLayout
	IDXGIDeviceSubObjectVtbl

	GetSharedHandle     uintptr
	GetUsage            uintptr
	SetEvictionPriority uintptr
	GetEvictionPriority uintptr
}
