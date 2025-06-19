package dx11

import (
	"structs"
)

type D3D11_TEXTURE2D_DESC struct {
	_ structs.HostLayout

	Width          uint32
	Height         uint32
	MipLevels      uint32
	ArraySize      uint32
	Format         uint32
	SampleDesc     DXGI_SAMPLE_DESC
	Usage          uint32
	BindFlags      uint32
	CPUAccessFlags uint32
	MiscFlags      uint32
}
