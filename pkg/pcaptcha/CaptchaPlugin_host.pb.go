//go:build !tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.7
// source: CaptchaPlugin.proto

package pcaptcha

import (
	context "context"
	errors "errors"
	fmt "fmt"
	emptypb "github.com/knqyf263/go-plugin/types/known/emptypb"
	wazero "github.com/tetratelabs/wazero"
	api "github.com/tetratelabs/wazero/api"
	wasi_snapshot_preview1 "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	sys "github.com/tetratelabs/wazero/sys"
	wpc "github.com/wshops/wshop-plugin-common/pkg/wpc"
	io "io"
	fs "io/fs"
	os "os"
)

const CaptchaPluginAPIVersion = 1

type CaptchaPluginOption struct {
	Stdout io.Writer
	Stderr io.Writer
	FS     fs.FS
}

type CaptchaPlugin struct {
	runtime wazero.Runtime
	config  wazero.ModuleConfig
}

func NewCaptchaPlugin(ctx context.Context, opt CaptchaPluginOption) (*CaptchaPlugin, error) {

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)

	// Combine the above into our baseline config, overriding defaults.
	config := wazero.NewModuleConfig().
		// By default, I/O streams are discarded and there's no file system.
		WithStdout(opt.Stdout).WithStderr(opt.Stderr).WithFS(opt.FS)

	return &CaptchaPlugin{
		runtime: r,
		config:  config,
	}, nil
}

func (p *CaptchaPlugin) Close(ctx context.Context) (err error) {
	if r := p.runtime; r != nil {
		err = r.Close(ctx)
	}
	return
}

func (p *CaptchaPlugin) Load(ctx context.Context, pluginPath string) (Captcha, error) {
	b, err := os.ReadFile(pluginPath)
	if err != nil {
		return nil, err
	}

	// Create an empty namespace so that multiple modules will not conflict
	ns := p.runtime.NewNamespace(ctx)

	if _, err = wasi_snapshot_preview1.NewBuilder(p.runtime).Instantiate(ctx, ns); err != nil {
		return nil, err
	}

	// Compile the WebAssembly module using the default configuration.
	code, err := p.runtime.CompileModule(ctx, b)
	if err != nil {
		return nil, err
	}

	// InstantiateModule runs the "_start" function, WASI's "main".
	module, err := ns.InstantiateModule(ctx, code, p.config)
	if err != nil {
		// Note: Most compilers do not exit the module after running "_start",
		// unless there was an Error. This allows you to call exported functions.
		if exitErr, ok := err.(*sys.ExitError); ok && exitErr.ExitCode() != 0 {
			return nil, fmt.Errorf("unexpected exit_code: %d", exitErr.ExitCode())
		} else if !ok {
			return nil, err
		}
	}

	// Compare API versions with the loading plugin
	apiVersion := module.ExportedFunction("captcha_api_version")
	if apiVersion == nil {
		return nil, errors.New("captcha_api_version is not exported")
	}
	results, err := apiVersion.Call(ctx)
	if err != nil {
		return nil, err
	} else if len(results) != 1 {
		return nil, errors.New("invalid captcha_api_version signature")
	}
	if results[0] != CaptchaPluginAPIVersion {
		return nil, fmt.Errorf("API version mismatch, host: %d, plugin: %d", CaptchaPluginAPIVersion, results[0])
	}

	configplugininfo := module.ExportedFunction("captcha_config_plugin_info")
	if configplugininfo == nil {
		return nil, errors.New("captcha_config_plugin_info is not exported")
	}
	verifycaptcha := module.ExportedFunction("captcha_verify_captcha")
	if verifycaptcha == nil {
		return nil, errors.New("captcha_verify_captcha is not exported")
	}
	getcustomhtmlinputfield := module.ExportedFunction("captcha_get_custom_html_input_field")
	if getcustomhtmlinputfield == nil {
		return nil, errors.New("captcha_get_custom_html_input_field is not exported")
	}
	getcustomhtmlhead := module.ExportedFunction("captcha_get_custom_html_head")
	if getcustomhtmlhead == nil {
		return nil, errors.New("captcha_get_custom_html_head is not exported")
	}
	getcustomhtmlbodyend := module.ExportedFunction("captcha_get_custom_html_body_end")
	if getcustomhtmlbodyend == nil {
		return nil, errors.New("captcha_get_custom_html_body_end is not exported")
	}

	malloc := module.ExportedFunction("malloc")
	if malloc == nil {
		return nil, errors.New("malloc is not exported")
	}

	free := module.ExportedFunction("free")
	if free == nil {
		return nil, errors.New("free is not exported")
	}
	return &captchaPlugin{module: module,
		malloc:                  malloc,
		free:                    free,
		configplugininfo:        configplugininfo,
		verifycaptcha:           verifycaptcha,
		getcustomhtmlinputfield: getcustomhtmlinputfield,
		getcustomhtmlhead:       getcustomhtmlhead,
		getcustomhtmlbodyend:    getcustomhtmlbodyend,
	}, nil
}

type captchaPlugin struct {
	module                  api.Module
	malloc                  api.Function
	free                    api.Function
	configplugininfo        api.Function
	verifycaptcha           api.Function
	getcustomhtmlinputfield api.Function
	getcustomhtmlhead       api.Function
	getcustomhtmlbodyend    api.Function
}

func (p *captchaPlugin) ConfigPluginInfo(ctx context.Context, request emptypb.Empty) (response wpc.PluginInfo, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))
	if dataSize == 0 {
		return response, nil
	}
	results, err := p.malloc.Call(ctx, dataSize)
	if err != nil {
		return response, err
	}

	dataPtr := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer p.free.Call(ctx, dataPtr)

	// The pointer is a linear memory offset, which is where we write the name.
	if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
		return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
	}
	ptrSize, err := p.configplugininfo.Call(ctx, dataPtr, dataSize)
	if err != nil {
		return response, err
	}

	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	resPtr := uint32(ptrSize[0] >> 32)
	resSize := uint32(ptrSize[0])

	// The pointer is a linear memory offset, which is where we write the name.
	bytes, ok := p.module.Memory().Read(ctx, resPtr, resSize)
	if !ok {
		return response, fmt.Errorf("Memory.Read(%d, %d) out of range of memory size %d",
			resPtr, resSize, p.module.Memory().Size(ctx))
	}

	if err = response.UnmarshalVT(bytes); err != nil {
		return response, err
	}

	return response, nil
}
func (p *captchaPlugin) VerifyCaptcha(ctx context.Context, request VerifyCaptchaRequest) (response VerifyCaptchaResponse, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))
	if dataSize == 0 {
		return response, nil
	}
	results, err := p.malloc.Call(ctx, dataSize)
	if err != nil {
		return response, err
	}

	dataPtr := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer p.free.Call(ctx, dataPtr)

	// The pointer is a linear memory offset, which is where we write the name.
	if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
		return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
	}
	ptrSize, err := p.verifycaptcha.Call(ctx, dataPtr, dataSize)
	if err != nil {
		return response, err
	}

	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	resPtr := uint32(ptrSize[0] >> 32)
	resSize := uint32(ptrSize[0])

	// The pointer is a linear memory offset, which is where we write the name.
	bytes, ok := p.module.Memory().Read(ctx, resPtr, resSize)
	if !ok {
		return response, fmt.Errorf("Memory.Read(%d, %d) out of range of memory size %d",
			resPtr, resSize, p.module.Memory().Size(ctx))
	}

	if err = response.UnmarshalVT(bytes); err != nil {
		return response, err
	}

	return response, nil
}
func (p *captchaPlugin) GetCustomHtmlInputField(ctx context.Context, request GetCustomHtmlInputFieldRequest) (response GetCustomHtmlInputFieldResponse, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))
	if dataSize == 0 {
		return response, nil
	}
	results, err := p.malloc.Call(ctx, dataSize)
	if err != nil {
		return response, err
	}

	dataPtr := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer p.free.Call(ctx, dataPtr)

	// The pointer is a linear memory offset, which is where we write the name.
	if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
		return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
	}
	ptrSize, err := p.getcustomhtmlinputfield.Call(ctx, dataPtr, dataSize)
	if err != nil {
		return response, err
	}

	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	resPtr := uint32(ptrSize[0] >> 32)
	resSize := uint32(ptrSize[0])

	// The pointer is a linear memory offset, which is where we write the name.
	bytes, ok := p.module.Memory().Read(ctx, resPtr, resSize)
	if !ok {
		return response, fmt.Errorf("Memory.Read(%d, %d) out of range of memory size %d",
			resPtr, resSize, p.module.Memory().Size(ctx))
	}

	if err = response.UnmarshalVT(bytes); err != nil {
		return response, err
	}

	return response, nil
}
func (p *captchaPlugin) GetCustomHtmlHead(ctx context.Context, request GetCustomHtmlHeadRequest) (response GetCustomHtmlHeadResponse, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))
	if dataSize == 0 {
		return response, nil
	}
	results, err := p.malloc.Call(ctx, dataSize)
	if err != nil {
		return response, err
	}

	dataPtr := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer p.free.Call(ctx, dataPtr)

	// The pointer is a linear memory offset, which is where we write the name.
	if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
		return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
	}
	ptrSize, err := p.getcustomhtmlhead.Call(ctx, dataPtr, dataSize)
	if err != nil {
		return response, err
	}

	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	resPtr := uint32(ptrSize[0] >> 32)
	resSize := uint32(ptrSize[0])

	// The pointer is a linear memory offset, which is where we write the name.
	bytes, ok := p.module.Memory().Read(ctx, resPtr, resSize)
	if !ok {
		return response, fmt.Errorf("Memory.Read(%d, %d) out of range of memory size %d",
			resPtr, resSize, p.module.Memory().Size(ctx))
	}

	if err = response.UnmarshalVT(bytes); err != nil {
		return response, err
	}

	return response, nil
}
func (p *captchaPlugin) GetCustomHtmlBodyEnd(ctx context.Context, request GetCustomHtmlBodyEndRequest) (response GetCustomHtmlBodyEndResponse, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))
	if dataSize == 0 {
		return response, nil
	}
	results, err := p.malloc.Call(ctx, dataSize)
	if err != nil {
		return response, err
	}

	dataPtr := results[0]
	// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
	// So, we have to free it when finished
	defer p.free.Call(ctx, dataPtr)

	// The pointer is a linear memory offset, which is where we write the name.
	if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
		return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
	}
	ptrSize, err := p.getcustomhtmlbodyend.Call(ctx, dataPtr, dataSize)
	if err != nil {
		return response, err
	}

	// Note: This pointer is still owned by TinyGo, so don't try to free it!
	resPtr := uint32(ptrSize[0] >> 32)
	resSize := uint32(ptrSize[0])

	// The pointer is a linear memory offset, which is where we write the name.
	bytes, ok := p.module.Memory().Read(ctx, resPtr, resSize)
	if !ok {
		return response, fmt.Errorf("Memory.Read(%d, %d) out of range of memory size %d",
			resPtr, resSize, p.module.Memory().Size(ctx))
	}

	if err = response.UnmarshalVT(bytes); err != nil {
		return response, err
	}

	return response, nil
}
