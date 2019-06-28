package MidCloudMAPEK

import (
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
)

type CloudService struct {
	Aor     common.NamingRecord
	Price   float64
	Status  bool
	Removed bool
	Rank    int
}

// Calls each of the cloud services using a generic proxy to get data to make the Rank based on a predefined sort method
func (cs *CloudService) RefreshPrice() {
	// Todo implement call to each of the cloud services using a generic proxy to get current price
	cfp := *NewCloudFunctionsProxy(cs.Aor.ClientProxy.Ip, cs.Aor.ClientProxy.Port, cs.Aor.ClientProxy.ObjectId)
	price, err := cfp.Price(100)
	if err != nil {
		lib.PrintlnError("Error while getting price from", cs.Aor.ServiceName, ". Error:", err)
	}
	err = cfp.Close()
	if err != nil {
		lib.PrintlnError("Error while closing CloudFunctionsProxy.Price from", cs.Aor.ServiceName, ". Error:", err)
	}

	cs.Price = price
}

// Calls each of the cloud services using a generic proxy to get status of them to make the Rank based on a predefined sort method
func (cs *CloudService) RefreshStatus() {
	// Todo implement call to each of the cloud services using a generic proxy to get status
	cfp := *NewCloudFunctionsProxy(cs.Aor.ClientProxy.Ip, cs.Aor.ClientProxy.Port, cs.Aor.ClientProxy.ObjectId)
	status, err := cfp.Availability()
	if err != nil {
		lib.PrintlnError("Error while getting availability from", cs.Aor.ServiceName, ". Error:", err)
	}
	err = cfp.Close()
	if err != nil {
		lib.PrintlnError("Error while closing CloudFunctionsProxy.Status from", cs.Aor.ServiceName, ". Error:", err)
	}

	cs.Status = status
}
