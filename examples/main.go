package main

import (
	"context"
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
	"github.com/samjegal/go-fincloud-helpers/authentication"
	"github.com/samjegal/go-fincloud-helpers/sender"
)

func authorize(sender autorest.Sender) autorest.Authorizer {
	httpMethod := "GET"
	requestUrl := "/vpc/v2/getVpcList?regionCode=FKR&responseFormatType=json"
	auth := authentication.Builder{
		AccessKeyId: "xx",
		SecretKey:   "xx",
		HttpMethod:  httpMethod,
		RequestURL:  requestUrl,
	}
	env, err := authentication.DetermineEnvironment("FINCLOUD")
	if err != nil {
		panic("Authentication determine set environment failed...")
	}
	config, err := auth.Build()
	if err != nil {
		panic("Authentication build failed...")
	}
	a, err := config.GetAuthorizationToken(sender, env.ResourceManagerEndpoint)
	if err != nil {
		panic("Get Authorization token failed...")
	}
	return a
}
func main() {
	sender := sender.BuildSender("FINCLOUD")
	client := vpc.NewClientWithBaseURI("https://fin-ncloud.apigw.fin-ntruss.com")
	client.Authorizer = authorize(sender)
	client.Sender = sender
	ctx := context.Background()
	resp, err := client.GetList(ctx, "json", "", "", "", "")
	if err != nil {
		panic("Send request faild...")
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(*resp.ReturnMessage)
}
