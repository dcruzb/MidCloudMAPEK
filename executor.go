package MidCloudMAPEK

import (
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
)

func Execute(ip string, port int, cloudFunctionName string, chanExecutor chan CloudService) {
	for {
		cloudService := <-chanExecutor

		lp := *dist.NewLookupProxy(ip, port)
		err := lp.Bind(cloudFunctionName, cloudService.Aor.ClientProxy)
		lib.FailOnError(err, "Error at lookup.bind.")
		lp.Close()
	}
}
