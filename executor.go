package MidCloudMAPEK

import (
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
)

func Execute(lookupProxy dist.LookupProxy, cloudFunctionName string, chanExecutor chan CloudService) {
	for {
		cloudService := <-chanExecutor
		err := lookupProxy.Bind(cloudFunctionName, cloudService.Aor.ClientProxy)
		lib.FailOnError(err, "Error at lookup.")
	}
}
