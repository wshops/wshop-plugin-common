package main

import (
	"context"
	"fmt"
	"github.com/knqyf263/go-plugin/types/known/emptypb"
	"github.com/wshops/wshop-plugin-common/pkg/pcaptcha"
	"os"
)

func main() {
	ctx := context.Background()
	p, err := pcaptcha.NewCaptchaPlugin(ctx, pcaptcha.CaptchaPluginOption{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		panic(err)
	}

	pluginInstance, err := p.Load(ctx, "./example/hcaptcha.wasm")
	if err != nil {
		panic(err)
	}

	pi, err := pluginInstance.GetPluginInfo(ctx, emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println("plugin version: ", pi.GetVersion())

	res, err := pluginInstance.VerifyCaptcha(ctx, pcaptcha.VerifyCaptchaRequest{Captcha: "123"})
	if err != nil {
		panic(err)
	}
	fmt.Println("verify result: ", res.GetIsValid())

	htmlInput, err := pluginInstance.GetCustomHtmlInputField(ctx, emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println("input html: ", htmlInput.GetHtml())

	htmlHead, err := pluginInstance.GetCustomHtmlHead(ctx, emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println("head html: ", htmlHead.GetHtml())

	htmlBody, err := pluginInstance.GetCustomHtmlBodyEnd(ctx, emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println("body html: ", htmlBody.GetHtml())
}
