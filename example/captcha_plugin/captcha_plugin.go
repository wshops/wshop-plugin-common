//go:build tinygo.wasm

package main

import (
	"context"
	"fmt"
	"github.com/knqyf263/go-plugin/types/known/emptypb"
	"github.com/wshops/wshop-plugin-common/pkg/pcaptcha"
	"github.com/wshops/wshop-plugin-common/pkg/wpc"
)

func main() {
	pcaptcha.RegisterCaptcha(HCaptcha{})
}

type HCaptcha struct{}

func (H HCaptcha) GetPluginInfo(ctx context.Context, empty emptypb.Empty) (wpc.PluginInfo, error) {
	fmt.Println("configuring plugin information...")
	pluginConfig := wpc.PluginInfo{
		Id:          "hcaptcha",
		Name:        "HCaptcha",
		Version:     "v1.0.0",
		Author:      "hello",
		Description: "123",
		IconUrl:     "123",
		PluginConstants: map[string]string{
			"foo": "bar",
		},
		UserAttributes: map[string]*wpc.PluginAttribute{
			"api_key": {
				Name:        "API Key",
				Description: "",
				Type:        "text",
				VRule:       "required",
				VMsg:        "required:\"API Key is required\"",
			},
		},
	}

	return pluginConfig, nil
}

func (H HCaptcha) VerifyCaptcha(ctx context.Context, request pcaptcha.VerifyCaptchaRequest) (pcaptcha.VerifyCaptchaResponse, error) {
	fmt.Println("validating captcha...")
	result := pcaptcha.VerifyCaptchaResponse{IsValid: true}
	return result, nil
}

func (H HCaptcha) GetCustomHtmlInputField(ctx context.Context, empty emptypb.Empty) (pcaptcha.GetCustomHtmlInputFieldResponse, error) {
	fmt.Println("getting custom html input field...")
	h := pcaptcha.GetCustomHtmlInputFieldResponse{
		Html: "<input type=\"text\" name=\"captcha\"/>",
	}
	return h, nil
}

func (H HCaptcha) GetCustomHtmlHead(ctx context.Context, empty emptypb.Empty) (pcaptcha.GetCustomHtmlHeadResponse, error) {
	fmt.Println("getting custom html head...")
	h := pcaptcha.GetCustomHtmlHeadResponse{Html: "<test></test>"}
	return h, nil
}

func (H HCaptcha) GetCustomHtmlBodyEnd(ctx context.Context, empty emptypb.Empty) (pcaptcha.GetCustomHtmlBodyEndResponse, error) {
	fmt.Println("getting custom html body end...")
	return pcaptcha.GetCustomHtmlBodyEndResponse{Html: "<div></div>"}, nil
}
