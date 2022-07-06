package resources

import (
	_ "embed"
	"encoding/base64"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkEc2Props struct {
	awscdk.StackProps
	VPC          awsec2.CfnVPC
	PublicSubnet awsec2.CfnSubnet
}

const (
	imageId      string = "ami-004332b441f90509b"
	instanceType string = "t4g.nano"
	keyName      string = "VPNServerKey"
)

//go:embed scripts/user_data.sh
var scripts []byte

func NewEC2(scope constructs.Construct, stack awscdk.Stack, props *CdkEc2Props) {
	vpnServerSG := awsec2.NewCfnSecurityGroup(stack, jsii.String("SecurityGroupEC2"), &awsec2.CfnSecurityGroupProps{
		GroupName:        jsii.String("VPNsg"),
		GroupDescription: jsii.String("VPN SecurityGroup"),
		VpcId:            props.VPC.Ref(),
		SecurityGroupIngress: &[]*awsec2.CfnSecurityGroup_IngressProperty{
			{
				IpProtocol: jsii.String("tcp"),
				CidrIp:     jsii.String("0.0.0.0/0"),
				FromPort:   jsii.Number(22),
				ToPort:     jsii.Number(22),
			},
			{
				IpProtocol: jsii.String("udp"),
				CidrIp:     jsii.String("0.0.0.0/0"),
				FromPort:   jsii.Number(500),
				ToPort:     jsii.Number(500),
			},
			{
				IpProtocol: jsii.String("udp"),
				CidrIp:     jsii.String("0.0.0.0/0"),
				FromPort:   jsii.Number(4500),
				ToPort:     jsii.Number(4500),
			},
			{
				IpProtocol: jsii.String("tcp"),
				CidrIp:     jsii.String("0.0.0.0/0"),
				FromPort:   jsii.Number(1701),
				ToPort:     jsii.Number(1701),
			},
		},
	})

	instance := awsec2.NewCfnInstance(stack, jsii.String("EC2Instance"), &awsec2.CfnInstanceProps{
		ImageId:          jsii.String(imageId),
		InstanceType:     jsii.String(instanceType),
		SubnetId:         props.PublicSubnet.Ref(),
		SecurityGroupIds: jsii.Strings(*vpnServerSG.AttrGroupId()),
		KeyName:          jsii.String(keyName),
		UserData:         jsii.String(base64.StdEncoding.EncodeToString(scripts)),
		Tags:             &[]*awscdk.CfnTag{{Key: jsii.String("Name"), Value: jsii.String("VPNServer")}},
	})

	awsec2.NewCfnEIP(stack, jsii.String("VPNServerEIP"), &awsec2.CfnEIPProps{
		InstanceId: instance.Ref(),
	})
}
