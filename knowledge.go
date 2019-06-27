package MidCloudMAPEK

import "github.com/dcbCIn/MidCloud/services/common"

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
}

// Calls each of the cloud services using a generic proxy to get status of them to make the Rank based on a predefined sort method
func (cs *CloudService) RefreshStatus() {
	// Todo implement call to each of the cloud services using a generic proxy to get status
}
