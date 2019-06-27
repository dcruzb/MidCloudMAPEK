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
	lookup dist.LookupProxy
	//	cloudFunctionName     string
	cloudFunctionsPattern string
}

// Starts monitoring the clouds available and caches the rank of their prices
//
// lookupProxy its the NameServer that will be used for the monitoring
// cloudFunctionName stands for the name that will be binded in the nameServer
// cloudFunctionsPattern stands for de pattern of the services that will be monitored. Any service that contains this pattern in the Lookup Proxy will be monitored
//
//	go monitor(lp, "cloudFunctions", "CloudFunctions")
func (mon Monitor) Start(ip string, port int, cloudFunctionsPattern string, chanAnalyzer chan []CloudService) {

	//mon.cloudFunctionName = cloudFunctionName
	mon.cloudFunctionsPattern = cloudFunctionsPattern

	for {
		mon.lookup = *dist.NewLookupProxy(ip, port)

		mon.refreshCloudServices()

		err := mon.lookup.Close()
		if err != nil {
			lib.PrintlnError("Error at closing lookup. Error:", err)
		}

		for i := range mon.cloudServices {
			mon.cloudServices[i].RefreshPrice()
			mon.cloudServices[i].RefreshStatus()
		}

		if len(mon.cloudServices) > 0 {
			chanAnalyzer <- mon.cloudServices
		}

		time.Sleep(5 * time.Second)
	}
}

// Get the list of cloud services based on name server list of binded servers
func (mon *Monitor) refreshCloudServices() {
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

	for _, cloudService := range mon.cloudServices {
		cloudService.Removed = true
	}

	for _, service := range services {
		// If the service registred in NameServer is a CloudFunctions server
		if strings.Contains(service.ServiceName, mon.cloudFunctionsPattern) {
			found := false
			for _, cloudService := range mon.cloudServices {
				if cloudService.Aor.ServiceName == service.ServiceName {
					found = true
					cloudService.Removed = false
				}
			}
			if !found {
				newCloudService := CloudService{}
				newCloudService.Aor = service
				newCloudService.Removed = false
				mon.cloudServices = append(mon.cloudServices, newCloudService)
			}
		}
	}
}
