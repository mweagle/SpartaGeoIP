package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
	spartaCF "github.com/mweagle/Sparta/aws/cloudformation"
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

func ipGeoLambda(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value(sparta.ContextKeyLogger).(*logrus.Logger)
	lambdaContext, _ := r.Context().Value(sparta.ContextKeyLambdaContext).(*sparta.LambdaContext) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var lambdaEvent sparta.APIGatewayLambdaJSONEvent
	err := decoder.Decode(&lambdaEvent)
	if err != nil {
		logger.Error("Failed to unmarshal event data: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parsedIP := net.ParseIP(lambdaEvent.Context.Identity.SourceIP)
	record, err := dbHandle.City(parsedIP)
	if err != nil {
		logger.Error("Failed to find city: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return the Info
	httpResponse := map[string]interface{}{
		"info": record,
	}
	responseBody, err := json.Marshal(httpResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
	}
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	stage := sparta.NewStage("ipgeo")
	apiGateway := sparta.NewAPIGateway("SpartaGeoIPService", stage)

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFn := sparta.HandleAWSLambda(sparta.LambdaName(ipGeoLambda),
		http.HandlerFunc(ipGeoLambda),
		sparta.IAMRoleDefinition{})
	apiGatewayResource, _ := apiGateway.NewResource("/info", lambdaFn)
	apiGatewayResource.NewMethod("GET", http.StatusOK)
	lambdaFunctions = append(lambdaFunctions, lambdaFn)
	stackName := spartaCF.UserScopedStackName("SpartaGeoIP")

	sparta.Main(stackName,
		"Sparta app supporting ip->geo mapping",
		lambdaFunctions,
		apiGateway,
		nil)
}
