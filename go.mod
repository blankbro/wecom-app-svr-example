module wecom-app-to-dify

go 1.22

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/blankbro/wecom-app-svr v0.0.0-20241108123316-602b923ff5e3
	github.com/langgenius/dify-sdk-go v0.0.0-20241031143354-972c4addddf6
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/lestrrat-go/strftime v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sbzhu/weworkapi_golang v0.0.0-20210525081115-1799804a7c8d // indirect
	golang.org/x/sys v0.27.0 // indirect
)

replace github.com/langgenius/dify-sdk-go => /Users/zexin.li/projects/github/blankbro/dify-sdk-go

replace github.com/blankbro/wecom-app-svr => /Users/zexin.li/projects/github/blankbro/wecom-app-svr
