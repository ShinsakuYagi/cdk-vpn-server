package stack

import (
	"cdk-vpn-server/resources"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewVPC(scope constructs.Construct, props *awscdk.StackProps) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = *props
	}
	stack := awscdk.NewStack(scope, jsii.String("VPNServerStack"), &sprops)

	vpc, publicSubnet := resources.NewVPC(scope, stack)
	resources.NewEC2(scope, stack, &resources.CdkEc2Props{
		VPC:          vpc,
		PublicSubnet: publicSubnet,
	})
}
