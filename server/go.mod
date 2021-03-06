module github.com/fieldkit/cloud/server

go 1.12

require (
	github.com/O-C-R/singlepage v0.0.0-20170327184421-bbbe2159ecec
	github.com/PyYoshi/goa-logging-zap v0.2.3
	github.com/ajg/form v1.5.1
	github.com/armon/go-metrics v0.0.0-20180713145231-3c58d8115a78
	github.com/aws/aws-lambda-go v0.0.0-20180413184133-ea03c2814414
	github.com/aws/aws-sdk-go v0.0.0-20170317202926-5b99715ae294
	github.com/cavaliercoder/grab v2.0.0+incompatible
	github.com/cenkalti/backoff v0.0.0-20170309153948-3db60c813733
	github.com/conservify/gonaturalist v0.0.0-20190530183130-1509fd074b2c
	github.com/conservify/protobuf-tools v0.0.0-20180715164506-43b897198d14
	github.com/conservify/sqlxcache v0.0.0-20190613231538-4c025edbc64e
	github.com/conservify/tooling v0.0.0-20190530175209-bf2b6e69e188 // indirect
	github.com/creack/goselect v0.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dghubble/go-twitter v0.0.0-20170219215544-fc93bb35701b
	github.com/dghubble/oauth1 v0.0.0-20170219195226-3c7784d12fed
	github.com/dghubble/sling v0.0.0-20170219194632-91b015f8c5e2
	github.com/dgrijalva/jwt-go v0.0.0-20170201225849-2268707a8f08
	github.com/dimfeld/httppath v0.0.0-20170720192232-ee938bf73598
	github.com/dimfeld/httptreemux v5.0.1+incompatible
	github.com/fieldkit/data-protocol v0.0.0-20180327223702-03a931c33d8e
	github.com/go-ini/ini v1.26.0
	github.com/go-kit/kit v0.0.0-20190326234956-40f35b54ddf5
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-stack/stack v1.8.0
	github.com/goadesign/goa v1.4.0
	github.com/gogo/protobuf v0.0.0-20190324160722-382325bbbb4d
	github.com/golang/freetype v0.0.0-20161208064710-d9be45aaf745
	github.com/golang/protobuf v1.2.0
	github.com/google/go-querystring v0.0.0-20170111101155-53e6ce116135
	github.com/google/uuid v1.1.1
	github.com/hashicorp/go-immutable-radix v0.0.0-20180129170900-7f3cd4390caa
	github.com/hashicorp/golang-lru v0.0.0-20180201235237-0fb14efe8c47
	github.com/inconshreveable/log15 v0.0.0-20180818164646-67afb5ed74ec
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/jmespath/go-jmespath v0.0.0-20160803190731-bd40a432e4c7
	github.com/jmoiron/sqlx v0.0.0-20171020205521-3379e5993990
	github.com/kelseyhightower/envconfig v0.0.0-20170206223400-8bf4bbfc795e
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
	github.com/lib/pq v0.0.0-20171022192043-b609790bd85e
	github.com/llgcode/draw2d v0.0.0-20161104081029-1286d3b2030a
	github.com/lucasb-eyer/go-colorful v0.0.0-20170223221042-c900de9dbbc7
	github.com/manveru/faker v0.0.0-20171103152722-9fbc68a78c4d
	github.com/mattn/go-colorable v0.1.1
	github.com/mattn/go-isatty v0.0.7
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/paulmach/go.geo v0.0.0-20170321183534-b160a6efed6c
	github.com/paulmach/go.geojson v0.0.0-20170327170536-40612a87147b
	github.com/pkg/errors v0.0.0-20190227000051-27936f6d90f9
	github.com/robinpowered/go-proto v0.0.0-20160614142341-85ea3e1f1d3d
	github.com/satori/go.uuid v0.0.0-20180103174451-36e9d2ebbde5
	github.com/segmentio/ksuid v1.0.1
	github.com/sirupsen/logrus v0.0.0-20190331131941-a6c0064cfaf9
	github.com/sohlich/go-dbscan v0.0.0-20161128164835-242a0c72bf77
	github.com/spf13/cobra v0.0.0-20170314171253-7be4beda01ec
	github.com/spf13/pflag v0.0.0-20170130214245-9ff6c6923cff
	github.com/ugorji/go v0.0.0-20190320090025-2dc34c0b8780
	github.com/xeipuuv/gojsonpointer v0.0.0-20170225233418-6fe8760cad35
	github.com/xeipuuv/gojsonreference v0.0.0-20150808065054-e02fc20de94c
	github.com/xeipuuv/gojsonschema v0.0.0-20170324002221-702b404897d4
	github.com/zach-klippenstein/goregen v0.0.0-20160303162051-795b5e3961ea
	go.bug.st/serial.v1 v0.0.0-20180827123349-5f7892a7bb45 // indirect
	go.uber.org/atomic v1.3.1
	go.uber.org/multierr v0.0.0-20180122172545-ddea229ff1df
	go.uber.org/zap v1.8.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/image v0.0.0-20170322222000-c0851fbc5b92
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
	golang.org/x/oauth2 v0.0.0-20190523182746-aaccbc9213b0
	golang.org/x/sys v0.0.0-20190329044733-9eb1bfa1ce65
	golang.org/x/tools v0.0.0-20190401163957-4fc9f0bfa59a
	google.golang.org/appengine v1.4.0
	gopkg.in/yaml.v2 v2.2.2
)
