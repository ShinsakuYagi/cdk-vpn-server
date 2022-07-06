package resources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewVPC(scope constructs.Construct, stack awscdk.Stack) (awsec2.CfnVPC, awsec2.CfnSubnet) {
	// Vpc
	vpc := awsec2.NewCfnVPC(stack, jsii.String("Vpc"), &awsec2.CfnVPCProps{
		CidrBlock: jsii.String("10.0.0.0/16"),
	})

	// PublicSubnet
	publicSubnet := awsec2.NewCfnSubnet(stack, jsii.String("PublicSubnet"), &awsec2.CfnSubnetProps{
		VpcId:               vpc.Ref(),
		CidrBlock:           jsii.String("10.0.0.0/24"),
		MapPublicIpOnLaunch: true,
	})

	// InternetGateway
	igw := awsec2.NewCfnInternetGateway(stack, jsii.String("InternetGateway"), &awsec2.CfnInternetGatewayProps{
		Tags: &[]*awscdk.CfnTag{{
			Key:   jsii.String("Name"),
			Value: jsii.String("VPNIG")}},
	})

	awsec2.NewCfnVPCGatewayAttachment(stack, jsii.String("VPCGatewayAttachment"), &awsec2.CfnVPCGatewayAttachmentProps{
		VpcId:             vpc.Ref(),
		InternetGatewayId: igw.Ref(),
	})

	// RouteTable
	routeTable := awsec2.NewCfnRouteTable(stack, jsii.String("RouteTable"), &awsec2.CfnRouteTableProps{
		VpcId: vpc.Ref(),
	})

	awsec2.NewCfnRoute(stack, jsii.String("Route"), &awsec2.CfnRouteProps{
		RouteTableId:         routeTable.Ref(),
		DestinationCidrBlock: jsii.String("0.0.0.0/0"),
		GatewayId:            igw.Ref(),
	})

	awsec2.NewCfnSubnetRouteTableAssociation(stack, jsii.String("RouteTableAssociation"), &awsec2.CfnSubnetRouteTableAssociationProps{
		RouteTableId: routeTable.Ref(),
		SubnetId:     publicSubnet.Ref(),
	})

	return vpc, publicSubnet
}
