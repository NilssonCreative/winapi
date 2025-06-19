package winrt

import (
	"syscall"
	"unsafe"

	"github.com/TKMAX777/winapi/dx11"
	"github.com/go-ole/go-ole"
)

// IDirect3DDevice
// https://learn.microsoft.com/en-us/uwp/api/windows.graphics.directx.direct3d11.idirect3ddevice?view=winrt-22621

var IDirect3DDeviceID = ole.NewGUID("{A37624AB-8D5F-4650-9D3E-9EAE3D9BC670}")
var IDirect3DDeviceClass = "Windows.Graphics.DirectX.Direct3D11.IDirect3DDevice"

type IDirect3DDevice struct {
	ole.IInspectable
}

type IDirect3DDeviceVtbl struct {
	ole.IInspectableVtbl
	CheckCounter                         uintptr // func(this *IDirect3DDevice, counterName *uint16, counterType *D3D_COUNTER_TYPE, count *uint32) HRESULT
	CheckCounterInfo                     uintptr // func(this *IDirect3DDevice, counterName *uint16, counterInfo *D3D_COUNTER_INFO) HRESULT
	CheckFeatureSupport                  uintptr // func(this *IDirect3DDevice, feature *D3D_FEATURE, featureSupportData *byte, featureSupportDataSize uint32) HRESULT
	CheckFormatSupport                   uintptr // func(this *IDirect3DDevice, format uint32, formatSupport *uint32) HRESULT
	CheckMultisampleQualityLevels        uintptr // func(this *IDirect3DDevice, format uint32, sampleCount uint32, qualityLevels *uint32) HRESULT
	CreateBlendState                     uintptr // func(this *IDirect3DDevice, blendDesc *D3D11_BLEND_DESC, blendState **ID3D11BlendState) HRESULT
	CreateBuffer                         uintptr // func(this *IDirect3DDevice, bufferDesc *D3D11_BUFFER_DESC, initialData *D3D11_SUBRESOURCE_DATA, buffer **ID3D11Buffer) HRESULT
	CreateClassLinkage                   uintptr // func(this *IDirect3DDevice, classLinkage **ID3D11ClassLinkage) HRESULT
	CreateComputeShader                  uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, classLinkage *ID3D11ClassLinkage, computeShader **ID3D11ComputeShader) HRESULT
	CreateCounter                        uintptr // func(this *IDirect3DDevice, counterDesc *D3D11_COUNTER_DESC, counter **ID3D11Counter) HRESULT
	CreateDeferredContext                uintptr // func(this *IDirect3DDevice, contextFlags uint32, deferredContext **ID3D11DeviceContext) HRESULT
	CreateDepthStencilState              uintptr // func(this *IDirect3DDevice, depthStencilDesc *D3D11_DEPTH_STENCIL_DESC, depthStencilState **ID3D11DepthStencilState) HRESULT
	CreateDepthStencilView               uintptr // func(this *IDirect3DDevice, resource *ID3D11Resource, depthStencilDesc *D3D11_DEPTH_STENCIL_VIEW_DESC, depthStencilView **ID3D11DepthStencilView) HRESULT
	CreateDomainShader                   uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, classLinkage *ID3D11ClassLinkage, domainShader **ID3D11DomainShader) HRESULT
	CreateGeometryShader                 uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, classLinkage *ID3D11ClassLinkage, geometryShader **ID3D11GeometryShader) HRESULT
	CreateGeometryShaderWithStreamOutput uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, streamOutputDesc *D3D11_SO_DESC, numEntries uint32, bufferStrides *uint32, numStrides uint32, rasterizedStream uint32, classLinkage *ID3D11ClassLinkage, geometryShader **ID3D11GeometryShader) HRESULT
	CreateHullShader                     uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, classLinkage *ID3D11ClassLinkage, hullShader **ID3D11HullShader) HRESULT
	CreateInputLayout                    uintptr // func(this *IDirect3DDevice, inputElementDescs *D3D11_INPUT_ELEMENT_DESC, numElements uint32, shaderBytecodeWithInputSignature *byte, bytecodeLength uint32, inputLayout **ID3D11InputLayout) HRESULT
	CreatePixelShader                    uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, classLinkage *ID3D11ClassLinkage, pixelShader **ID3D11PixelShader) HRESULT
	CreatePredicate                      uintptr // func(this *IDirect3DDevice, predicateDesc *D3D11_PREDICATE_DESC, predicate **ID3D11Predicate) HRESULT
	CreateQuery                          uintptr // func(this *IDirect3DDevice, queryDesc *D3D11_QUERY_DESC, query **ID3D11Query) HRESULT
	CreateRasterizerState                uintptr // func(this *IDirect3DDevice, rasterizerDesc *D3D11_RASTERIZER_DESC, rasterizerState **ID3D11RasterizerState) HRESULT
	CreateRenderTargetView               uintptr // func(this *IDirect3DDevice, resource *ID3D11Resource, renderTargetViewDesc *D3D11_RENDER_TARGET_VIEW_DESC, renderTargetView **ID3D11RenderTargetView) HRESULT
	CreateSamplerState                   uintptr // func(this *IDirect3DDevice, samplerDesc *D3D11_SAMPLER_DESC, samplerState **ID3D11SamplerState) HRESULT
	CreateShaderResourceView             uintptr // func(this *IDirect3DDevice, resource *ID3D11Resource, shaderResourceViewDesc *D3D11_SHADER_RESOURCE_VIEW_DESC, shaderResourceView **ID3D11ShaderResourceView) HRESULT
	CreateTexture1D                      uintptr // func(this *IDirect3DDevice, textureDesc *D3D11_TEXTURE1D_DESC, initialData *D3D11_SUBRESOURCE_DATA, texture **ID3D11Texture1D) HRESULT
	CreateTexture2D                      uintptr // func(this *IDirect3DDevice, textureDesc *D3D11_TEXTURE2D_DESC, initialData *D3D11_SUBRESOURCE_DATA, texture **ID3D11Texture2D) HRESULT
	CreateTexture3D                      uintptr // func(this *IDirect3DDevice, textureDesc *D3D11_TEXTURE3D_DESC, initialData *D3D11_SUBRESOURCE_DATA, texture **ID3D11Texture3D) HRESULT
	CreateUnorderedAccessView            uintptr // func(this *IDirect3DDevice, resource *ID3D11Resource, unorderedAccessViewDesc *D3D11_UNORDERED_ACCESS_VIEW_DESC, unorderedAccessView **ID3D11UnorderedAccessView) HRESULT
	CreateVertexShader                   uintptr // func(this *IDirect3DDevice, shaderBytecode *byte, bytecodeLength uint32, classLinkage *ID3D11ClassLinkage, vertexShader **ID3D11VertexShader) HRESULT
	GetCreationFlags                     uintptr // func(this *IDirect3DDevice) uint32
	GetDeviceRemovedReason               uintptr // func(this *IDirect3DDevice) HRESULT
	GetExceptionMode                     uintptr // func(this *IDirect3DDevice) uint32
	GetFeatureLevel                      uintptr // func(this *IDirect3DDevice) D3D_FEATURE_LEVEL
	GetImmediateContext                  uintptr // func(this *IDirect3DDevice, immediateContext **ID3D11DeviceContext) HRESULT
	GetPrivateData                       uintptr // func(this *IDirect3DDevice, guid *ole.GUID, dataSize *uint32, data *byte) HRESULT
	OpenSharedResource                   uintptr // func(this *IDirect3DDevice, sharedHandle uintptr, riid *ole.GUID, resource **ole.IUnknown) HRESULT
	SetExceptionMode                     uintptr // func(this *IDirect3DDevice, raiseFlags uint32) HRESULT
	SetPrivateData                       uintptr // func(this *IDirect3DDevice, guid *ole.GUID, dataSize uint32, data *byte) HRESULT
	SetPrivateDataInterface              uintptr // func(this *IDirect3DDevice, guid *ole.GUID, unknown *ole.IUnknown) HRESULT
}

func (v *IDirect3DDevice) VTable() *IDirect3DDeviceVtbl {
	return (*IDirect3DDeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDirect3DDevice) CreateTexture2D(desc *dx11.D3D11_TEXTURE2D_DESC, initialData *dx11.D3D11_SUBRESOURCE_DATA) (*ID3D11Texture2D, error) {
	if v == nil {
		return nil, ole.NewError(ole.E_POINTER)
	}

	var texture *ID3D11Texture2D
	r1, _, _ := syscall.SyscallN(v.VTable().CreateTexture2D, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(desc)), uintptr(unsafe.Pointer(initialData)), uintptr(unsafe.Pointer(&texture)))
	if r1 != ole.S_OK {
		return nil, ole.NewError(r1)
	}
	return texture, nil
}
