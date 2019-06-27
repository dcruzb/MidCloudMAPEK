package MidCloudMAPEK

import (
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
)

type CloudFunctionsProxy struct {
	host      string
	port      int
	ObjectId  int
	requestor dist.Requestor
}

func NewCloudFunctionsProxy(host string, port int, objectId int) *CloudFunctionsProxy {
	return &CloudFunctionsProxy{host, port, objectId, dist.NewRequestorImpl(host, port)}
}

func (cfp CloudFunctionsProxy) Price(size float64) (price float64, err error) {
	inv := *dist.NewInvocation(cfp.ObjectId, cfp.host, cfp.port, lib.FunctionName(), []interface{}{size})
	termination, err := cfp.requestor.Invoke(inv)
	if err != nil {
		return price, err
	}

	price = termination.Result.([]interface{})[0].(float64)
	//	err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return price, err
	}

	return price, nil
}

func (cfp CloudFunctionsProxy) Availability() (available bool, err error) {
	inv := *dist.NewInvocation(cfp.ObjectId, cfp.host, cfp.port, lib.FunctionName(), []interface{}{})
	termination, err := cfp.requestor.Invoke(inv)
	if err != nil {
		return available, err
	}

	available = termination.Result.([]interface{})[0].(bool)
	//	err = termination.Result.([]interface{})[1].(error)
	if err != nil {
		return available, err
	}

	return available, nil
}

func (cfp CloudFunctionsProxy) Close() error {
	return cfp.requestor.Close()
}
