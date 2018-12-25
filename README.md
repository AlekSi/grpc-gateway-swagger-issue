# grpc-gateway-swagger-issue

```
$ dep status

PROJECT                                 CONSTRAINT     VERSION        REVISION  LATEST   PKGS USED
github.com/PuerkitoBio/purell           v1.1.0         v1.1.0         0bcb03f   v1.1.0   1
github.com/PuerkitoBio/urlesc           branch master  branch master  de5bf2a   de5bf2a  1
github.com/asaskevich/govalidator       v9             v9             ccb8e96   v9       1
github.com/fsnotify/fsnotify            v1.4.7         v1.4.7         c282820   v1.4.7   1
github.com/ghodss/yaml                  *                             0ca9ea5            1
github.com/globalsign/mgo               branch master  branch master  eeefdec   eeefdec  2
github.com/go-openapi/analysis          branch master  branch master  e2f3fdb   e2f3fdb  2
github.com/go-openapi/errors            branch master  branch master  7a7ff1b   7a7ff1b  1
github.com/go-openapi/inflect           branch master  branch master  16a898e   16a898e  1
github.com/go-openapi/jsonpointer       v0.18.0        v0.18.0        ef5f0af   v0.18.0  1
github.com/go-openapi/jsonreference     v0.18.0        v0.18.0        8483a88   v0.18.0  1
github.com/go-openapi/loads             branch master  branch master  7462858   7462858  2
github.com/go-openapi/runtime           branch master  branch master  41e24cc   41e24cc  8
github.com/go-openapi/spec              branch master  branch master  5b6cdde   5b6cdde  1
github.com/go-openapi/strfmt            branch master  branch master  e471370   e471370  1
github.com/go-openapi/swag              branch master  branch master  1d29f06   1d29f06  1
github.com/go-openapi/validate          branch master  branch master  d2eab7d   d2eab7d  1
github.com/go-swagger/go-swagger        v0.18.0        v0.18.0        6b23bb6   v0.18.0  6
github.com/golang/glog                  *                             23def4e            1
github.com/golang/protobuf              v1.2.0         v1.2.0         aa810b6   v1.2.0   14
github.com/gorilla/handlers             v1.4.0         v1.4.0         7e0847f   v1.4.0   1
github.com/grpc-ecosystem/grpc-gateway  v1.6.3         v1.6.3         719aaad   v1.6.3   12
github.com/hashicorp/hcl                v1.0.0         v1.0.0         8cb6e5b   v1.0.0   10
github.com/jessevdk/go-flags            v1.4.0         v1.4.0         c6ca198   v1.4.0   1
github.com/kr/pretty                    branch master  branch master  73f6ac0   73f6ac0  1
github.com/kr/text                      v0.1.0         v0.1.0         e2ffdb1   v0.1.0   1
github.com/magiconair/properties        v1.8.0         v1.8.0         c235336   v1.8.0   1
github.com/mailru/easyjson              branch master  branch master  60711f1   60711f1  3
github.com/mitchellh/mapstructure       v1.1.2         v1.1.2         3536a92   v1.1.2   1
github.com/pelletier/go-toml            v1.2.0         v1.2.0         c01d127   v1.2.0   1
github.com/spf13/afero                  v1.2.0         v1.2.0         a5d6946   v1.2.0   2
github.com/spf13/cast                   v1.3.0         v1.3.0         8c9545a   v1.3.0   1
github.com/spf13/jwalterweatherman      v1.0.0         v1.0.0         4a4406e   v1.0.0   1
github.com/spf13/pflag                  v1.0.3         v1.0.3         298182f   v1.0.3   1
github.com/spf13/viper                  branch master  branch master  6d33b5a   6d33b5a  1
github.com/toqueteos/webbrowser         v1.1           v1.1           3232c91   v1.1     1
golang.org/x/net                        branch master  branch master  927f977   927f977  7
golang.org/x/sys                        branch master  branch master  b4a75ba   b4a75ba  1
golang.org/x/text                       v0.3.0         v0.3.0         f21a4df   v0.3.0   15
golang.org/x/tools                      branch master  branch master  d00ac6d   d00ac6d  13
google.golang.org/genproto              383e8b2                       383e8b2            3
google.golang.org/grpc                  v1.17.0        v1.17.0        df01485   v1.17.0  31
gopkg.in/yaml.v2                        branch v2      branch v2      51d6538   51d6538  1
```

```
$ make

rm -f api/*.pb.go api/*.pb.gw.go api/*.swagger.json api/swagger.json
go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go
go install -v ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go install -v ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go install -v ./vendor/github.com/go-swagger/go-swagger/cmd/swagger
protoc --proto_path=. \
			--proto_path=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			api/*.proto --go_out=plugins=grpc:.
protoc --proto_path=. \
			--proto_path=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			api/*.proto --grpc-gateway_out=logtostderr=true:.
protoc --proto_path=. \
			--proto_path=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			api/*.proto --swagger_out=logtostderr=true:.
swagger mixin api/api.json api/*.swagger.json > api/swagger.json
2018/12/25 18:02:22 args[0] = api/api.json
2018/12/25 18:02:22 args[1:] = [api/bar.swagger.json api/baz.swagger.json api/common.swagger.json]
2018/12/25 18:02:22 definitions entry 'apiCommon' already exists in primary or higher priority mixin, skipping

make: *** [all] Error 1
```
