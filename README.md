# VPN Server with CDK Go

Create VPN Server with CDK Go

## resources

- VPC
  - 1 public subnet
- EC2
  - t4g.nano instance
- EIP
- IGW

## to be prepared in advance

1. Create KeyPair
   - <https://ap-northeast-1.console.aws.amazon.com/ec2/v2/home?region=ap-northeast-1#KeyPairs:>
2. AWS CLI credentials configure
   - <https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-configure-profiles.html>
3. run `make bootstrap PROFILE=your_profile_name REGION=your_region ACCOUNT_ID=your_aws_account_id`
4. edit `PSK` and `USERS` in `resources/scripts/user_data.sh`
   - <https://hub.docker.com/r/siomiz/softethervpn/>

## usage

VPN Setting
<https://support.apple.com/ja-jp/guide/mac-help/mchlp2963/12.0/mac/12.0>

### diff

`make diff PROFILE=your_profile_name REGION=your_region`

### deploy

`make deploy PROFILE=your_profile_name REGION=your_region`
