package apsarastack

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApsaraStackCommonBandwidthPackageAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApsaraStackCommonBandwidthPackageAttachmentCreate,
		Read:   resourceApsaraStackCommonBandwidthPackageAttachmentRead,
		Delete: resourceApsaraStackCommonBandwidthPackageAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bandwidth_package_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApsaraStackCommonBandwidthPackageAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	vpcService := VpcService{client}
	request := vpc.CreateAddCommonBandwidthPackageIpRequest()
	request.RegionId = client.RegionId
	request.BandwidthPackageId = Trim(d.Get("bandwidth_package_id").(string))
	request.IpInstanceId = Trim(d.Get("instance_id").(string))
	raw, err := client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
		return vpcClient.AddCommonBandwidthPackageIp(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_common_bandwidth_package_attachment", request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	//check the common bandwidth package attachment
	d.SetId(request.BandwidthPackageId + COLON_SEPARATED + request.IpInstanceId)
	if err := vpcService.WaitForCommonBandwidthPackageAttachment(d.Id(), Available, 5*DefaultTimeout); err != nil {
		return WrapError(err)
	}
	return resourceApsaraStackCommonBandwidthPackageAttachmentRead(d, meta)
}

func resourceApsaraStackCommonBandwidthPackageAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	vpcService := VpcService{client}

	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}
	bandwidthPackageId, ipInstanceId := parts[0], parts[1]
	_, err = vpcService.DescribeCommonBandwidthPackageAttachment(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("bandwidth_package_id", bandwidthPackageId)
	d.Set("instance_id", ipInstanceId)
	return nil
}

func resourceApsaraStackCommonBandwidthPackageAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	vpcService := VpcService{client}
	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}
	bandwidthPackageId, ipInstanceId := parts[0], parts[1]

	request := vpc.CreateRemoveCommonBandwidthPackageIpRequest()
	request.BandwidthPackageId = bandwidthPackageId
	request.IpInstanceId = ipInstanceId

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			return vpcClient.RemoveCommonBandwidthPackageIp(request)
		})
		//Waiting for unassociate the common bandwidth package
		if err != nil {
			if IsExpectedErrors(err, []string{"TaskConflict"}) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	return WrapError(vpcService.WaitForCommonBandwidthPackageAttachment(d.Id(), Deleted, DefaultTimeoutMedium))
}
