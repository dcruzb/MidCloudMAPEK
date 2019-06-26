package MidCloudMAPEK

import (
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"strings"
	"time"
)

type Monitor struct {
	cloudServices []CloudService
	//nameServerIp          string
	//nameServerPort        int
	lookup                dist.LookupProxy
	cloudFunctionName     string
	cloudFunctionsPattern string
}

// Starts monitoring the clouds available and caches the rank of their prices
//
// lookupProxy its the NameServer that will be used for the monitoring
// cloudFunctionName stands for the name that will be binded in the nameServer
// cloudFunctionsPattern stands for de pattern of the services that will be monitored. Any service that contains this pattern in the Lookup Proxy will be monitored
//
//	go monitor(lp, "cloudFunctions", "CloudFunctions")
func (mon Monitor) Start(lookupProxy dist.LookupProxy, cloudFunctionName string, cloudFunctionsPattern string, chanAnalyzer chan []CloudService) {
	mon.lookup = lookupProxy
	mon.cloudFunctionName = cloudFunctionName
	mon.cloudFunctionsPattern = cloudFunctionsPattern

	for {
		mon.refreshCloudServices()

		for _, service := range mon.cloudServices {
			service.RefreshPrice()
			service.RefreshStatus()
		}

		chanAnalyzer <- mon.cloudServices

		time.Sleep(30 * time.Second)
	}
}

// Get the list of cloud services based on name server list of binded servers
func (mon Monitor) refreshCloudServices() {
	//lp := dist.NewLookupProxy(mon.nameServerIp, mon.nameServerPort)
	services, err := mon.lookup.List()
	if err != nil {
		lib.PrintlnError("Error at lookup. Error:", err)
	}

	// Todo monitor não está fechando e reabrindo a conexão para o lookup, pois é acessado diversas vezes depois (a cada intervalo de tempo pré-definido). Isso ocasiona a quebra do sistema de monitoramento caso o servidor de nomes seja reiniciado.
	//err = mon.lookup.Close()
	//if err != nil {
	//	lib.PrintlnError("Error at closing lookup. Error:", err)
	//}

	for _, service := range services {
		// If the service registred in NameServer is a CloudFunctions server
		if strings.Contains(service.ServiceName, mon.cloudFunctionsPattern) {
			found := false
			for _, cloudService := range mon.cloudServices {
				if cloudService.Aor.ServiceName == service.ServiceName {
					found = true
				}
			}
			if !found {
				newCloudService := CloudService{}
				newCloudService.Aor = service
				mon.cloudServices = append(mon.cloudServices, newCloudService)
			}
		}
	}
}
