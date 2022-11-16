gen-all:
	make gen-captcha-plugin
	make gen-notification-plugin
	make gen-vcode-plugin
gen-captcha-plugin:
	protoc --go-plugin_out=../captcha --go-plugin_opt=paths=source_relative CaptchaPlugin.proto
gen-notification-plugin:
	protoc --go-plugin_out=../notification --go-plugin_opt=paths=source_relative NotificationPlugin.proto
gen-vcode-plugin:
	protoc --go-plugin_out=../vcode --go-plugin_opt=paths=source_relative VCodePlugin.proto