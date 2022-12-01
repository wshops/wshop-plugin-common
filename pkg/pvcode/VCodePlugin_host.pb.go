//go:build !tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.9
// source: VCodePlugin.proto

package pvcode

import (
	context "context"
	errors "errors"
	fmt "fmt"
	emptypb "github.com/knqyf263/go-plugin/types/known/emptypb"
	wasm "github.com/knqyf263/go-plugin/wasm"
	wazero "github.com/tetratelabs/wazero"
	api "github.com/tetratelabs/wazero/api"
	wasi_snapshot_preview1 "github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	sys "github.com/tetratelabs/wazero/sys"
	wpc "github.com/wshops/wshop-plugin-common/pkg/wpc"
	io "io"
	fs "io/fs"
	os "os"
)

const (
	i32 = api.ValueTypeI32
	i64 = api.ValueTypeI64
)

type _hostFunctions struct {
	HostFunctions
}

// Instantiate a Go-defined module named "env" that exports host functions.
func (h _hostFunctions) Instantiate(ctx context.Context, r wazero.Runtime, ns wazero.Namespace) error {
	envBuilder := r.NewHostModuleBuilder("env")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._LogInfo), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("log_info")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._LogError), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("log_error")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._LogDebug), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("log_debug")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._LogWarn), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("log_warn")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._LogFatal), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("log_fatal")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._LogPanic), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("log_panic")

	envBuilder.NewFunctionBuilder().
		WithGoModuleFunction(api.GoModuleFunc(h._HttpRequest), []api.ValueType{i32, i32}, []api.ValueType{i64}).
		WithParameterNames("offset", "size").
		Export("http_request")

	_, err := envBuilder.Instantiate(ctx, ns)
	return err
}

func (h _hostFunctions) _LogInfo(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncLogRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.LogInfo(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

func (h _hostFunctions) _LogError(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncLogRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.LogError(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

func (h _hostFunctions) _LogDebug(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncLogRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.LogDebug(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

func (h _hostFunctions) _LogWarn(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncLogRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.LogWarn(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

func (h _hostFunctions) _LogFatal(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncLogRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.LogFatal(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

func (h _hostFunctions) _LogPanic(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncLogRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.LogPanic(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

func (h _hostFunctions) _HttpRequest(ctx context.Context, m api.Module, params []uint64) []uint64 {
	offset, size := uint32(params[0]), uint32(params[1])
	buf, err := wasm.ReadMemory(ctx, m, offset, size)
	if err != nil {
		panic(err)
	}
	var request wpc.HFuncHttpRequest
	err = request.UnmarshalVT(buf)
	if err != nil {
		panic(err)
	}
	resp, err := h.HttpRequest(ctx, request)
	if err != nil {
		panic(err)
	}
	buf, err = resp.MarshalVT()
	if err != nil {
		panic(err)
	}
	ptr, err := wasm.WriteMemory(ctx, m, buf)
	if err != nil {
		panic(err)
	}
	ptrLen := (ptr << uint64(32)) | uint64(len(buf))
	return []uint64{ptrLen}
}

const VCodePluginAPIVersion = 1

type VCodePluginOption struct {
	Stdout io.Writer
	Stderr io.Writer
	FS     fs.FS
}

type VCodePlugin struct {
	runtime wazero.Runtime
	config  wazero.ModuleConfig
}

func NewVCodePlugin(ctx context.Context, opt VCodePluginOption) (*VCodePlugin, error) {

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)

	// Combine the above into our baseline config, overriding defaults.
	config := wazero.NewModuleConfig().
		// By default, I/O streams are discarded and there's no file system.
		WithStdout(opt.Stdout).WithStderr(opt.Stderr).WithFS(opt.FS)

	return &VCodePlugin{
		runtime: r,
		config:  config,
	}, nil
}

func (p *VCodePlugin) Close(ctx context.Context) (err error) {
	if r := p.runtime; r != nil {
		err = r.Close(ctx)
	}
	return
}

func (p *VCodePlugin) Load(ctx context.Context, pluginPath string, hostFunctions HostFunctions) (VCode, error) {
	b, err := os.ReadFile(pluginPath)
	if err != nil {
		return nil, err
	}

	// Create an empty namespace so that multiple modules will not conflict
	ns := p.runtime.NewNamespace(ctx)

	h := _hostFunctions{hostFunctions}

	if err := h.Instantiate(ctx, p.runtime, ns); err != nil {
		return nil, err
	}

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
	apiVersion := module.ExportedFunction("v_code_api_version")
	if apiVersion == nil {
		return nil, errors.New("v_code_api_version is not exported")
	}
	results, err := apiVersion.Call(ctx)
	if err != nil {
		return nil, err
	} else if len(results) != 1 {
		return nil, errors.New("invalid v_code_api_version signature")
	}
	if results[0] != VCodePluginAPIVersion {
		return nil, fmt.Errorf("API version mismatch, host: %d, plugin: %d", VCodePluginAPIVersion, results[0])
	}

	getplugininfo := module.ExportedFunction("v_code_get_plugin_info")
	if getplugininfo == nil {
		return nil, errors.New("v_code_get_plugin_info is not exported")
	}
	sendverificationcode := module.ExportedFunction("v_code_send_verification_code")
	if sendverificationcode == nil {
		return nil, errors.New("v_code_send_verification_code is not exported")
	}

	malloc := module.ExportedFunction("malloc")
	if malloc == nil {
		return nil, errors.New("malloc is not exported")
	}

	free := module.ExportedFunction("free")
	if free == nil {
		return nil, errors.New("free is not exported")
	}
	return &vCodePlugin{module: module,
		malloc:               malloc,
		free:                 free,
		getplugininfo:        getplugininfo,
		sendverificationcode: sendverificationcode,
	}, nil
}

type vCodePlugin struct {
	module               api.Module
	malloc               api.Function
	free                 api.Function
	getplugininfo        api.Function
	sendverificationcode api.Function
}

func (p *vCodePlugin) GetPluginInfo(ctx context.Context, request emptypb.Empty) (response wpc.PluginInfo, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))

	var dataPtr uint64
	// If the input data is not empty, we must allocate the in-Wasm memory to store it, and pass to the plugin.
	if dataSize != 0 {
		results, err := p.malloc.Call(ctx, dataSize)
		if err != nil {
			return response, err
		}
		dataPtr = results[0]
		// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
		// So, we have to free it when finished
		defer p.free.Call(ctx, dataPtr)

		// The pointer is a linear memory offset, which is where we write the name.
		if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
			return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
		}
	}

	ptrSize, err := p.getplugininfo.Call(ctx, dataPtr, dataSize)
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
func (p *vCodePlugin) SendVerificationCode(ctx context.Context, request SendVerificationCodeRequest) (response SendVerificationCodeResponse, err error) {
	data, err := request.MarshalVT()
	if err != nil {
		return response, err
	}
	dataSize := uint64(len(data))

	var dataPtr uint64
	// If the input data is not empty, we must allocate the in-Wasm memory to store it, and pass to the plugin.
	if dataSize != 0 {
		results, err := p.malloc.Call(ctx, dataSize)
		if err != nil {
			return response, err
		}
		dataPtr = results[0]
		// This pointer is managed by TinyGo, but TinyGo is unaware of external usage.
		// So, we have to free it when finished
		defer p.free.Call(ctx, dataPtr)

		// The pointer is a linear memory offset, which is where we write the name.
		if !p.module.Memory().Write(ctx, uint32(dataPtr), data) {
			return response, fmt.Errorf("Memory.Write(%d, %d) out of range of memory size %d", dataPtr, dataSize, p.module.Memory().Size(ctx))
		}
	}

	ptrSize, err := p.sendverificationcode.Call(ctx, dataPtr, dataSize)
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
