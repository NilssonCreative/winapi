package dx11

import "structs"

type DXGI_SAMPLE_DESC struct {
	_ structs.HostLayout

	Count   uint32
	Quality uint32
}

type DXGI_MAPPED_RECT struct {
	_ structs.HostLayout

	Pitch int32
	PBits uintptr
}
