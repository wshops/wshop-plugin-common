gen-all:
	make gen-common
	make gen-captcha-plugin
	make gen-notification-plugin
	make gen-vcode-plugin
gen-captcha-plugin:
	protoc --go-plugin_out=../pcaptcha --go-plugin_opt=paths=source_relative CaptchaPlugin.proto
gen-notification-plugin:
	protoc --go-plugin_out=../pnotify --go-plugin_opt=paths=source_relative NotificationPlugin.proto
gen-vcode-plugin:
	protoc --go-plugin_out=../pvcode --go-plugin_opt=paths=source_relative VCodePlugin.proto
gen-common:
	protoc --go-plugin_out=../wpc --go-plugin_opt=paths=source_relative common.proto