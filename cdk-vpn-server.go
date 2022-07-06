package main

import (
	"cdk-vpn-server/stack"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := awscdk.NewApp(nil)

	stack.NewVPC(app, &awscdk.StackProps{})

	app.Synth(nil)
}
