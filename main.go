package main

import (
	"context"
	"net"
	"net/http"

	sparta "github.com/mweagle/Sparta"
	spartaAPIGateway "github.com/mweagle/Sparta/aws/apigateway"
	spartaCF "github.com/mweagle/Sparta/aws/cloudformation"
	spartaAWSEvents "github.com/mweagle/Sparta/aws/events"
	"github.com/mweagle/SpartaGeoIP/constants"
	"github.com/oschwald/geoip2-golang"
)

var dbHandle *geoip2.Reader

// One time load of the database
func init() {
	openHandle, err := geoip2.FromBytes(constants.FSMustByte(false, "/GeoLite2-Country.mmdb"))
	if err != nil {
		panic(err.Error())
	}
	dbHandle = openHandle
}

//go:generate mkdir -pv ./constants
//go:generate rm -f ./constants/CONSTANTS.go
//go:generate esc -o ./constants/CONSTANTS.go -pkg constants GeoLite2-Country.mmdb
////////////////////////////////////////////////////////////////////////////////
// IP->Geo results
//

func ipGeoLambda(ctx context.Context,
	apiRequest spartaAWSEvents.APIGatewayRequest) (*spartaAPIGateway.Response, error) {
	parsedIP := net.ParseIP(apiRequest.Context.Identity.SourceIP)
	record, err := dbHandle.City(parsedIP)
	if err != nil {
		return nil, err
	}
	requestResponse := map[string]interface{}{
		"ip":     parsedIP,
		"record": record,
	}
	return spartaAPIGateway.NewResponse(http.StatusOK, requestResponse), nil
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	stage := sparta.NewStage("ipgeo")
	apiGateway := sparta.NewAPIGateway("SpartaGeoIPService", stage)

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFn, _ := sparta.NewAWSLambda(sparta.LambdaName(ipGeoLambda),
		ipGeoLambda,
		sparta.IAMRoleDefinition{})
	apiGatewayResource, _ := apiGateway.NewResource("/info", lambdaFn)
	apiMethod, _ := apiGatewayResource.NewMethod("GET", http.StatusOK, http.StatusOK)
	apiMethod.SupportedRequestContentTypes = []string{"application/json"}

	lambdaFunctions = append(lambdaFunctions, lambdaFn)
	stackName := spartaCF.UserScopedStackName("SpartaGeoIP")

	sparta.Main(stackName,
		"Sparta app supporting ip->geo mapping",
		lambdaFunctions,
		apiGateway,
		nil)
}
