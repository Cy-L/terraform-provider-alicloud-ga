package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAliCloudGaAcceleratorDdosAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudGaAcceleratorDdosAttachmentCreate,
		Read:   resourceAliCloudGaAcceleratorDdosAttachmentRead,
		Update: resourceAliCloudGaAcceleratorDdosAttachmentUpdate,
		Delete: resourceAliCloudGaAcceleratorDdosAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"accelerator_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ddos_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ddos_region_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudGaAcceleratorDdosAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "AttachDdosToAccelerator"
	request := make(map[string]interface{})
	var err error

	request["ClientToken"] = buildClientToken("AttachDdosToAccelerator")
	request["AcceleratorId"] = d.Get("accelerator_id")
	request["DdosConfigList"] = []map[string]interface{}{
		{
			"DdosId":       d.Get("ddos_id"),
			"DdosRegionId": d.Get("ddos_region_id"),
		},
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = retry.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutCreate)), func() *retry.RetryError {
		response, err = client.RpcPost("Ga", "2019-11-20", action, nil, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"StateError.Accelerator", "NotActive.Accelerator"}) || NeedRetry(err) {
				wait()
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ga_accelerator_ddos_attachment", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["AcceleratorId"], d.Get("ddos_id")))

	return resourceAliCloudGaAcceleratorDdosAttachmentRead(d, meta)
}

func resourceAliCloudGaAcceleratorDdosAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	gaService := GaService{client}
	object, err := gaService.DescribeGaAcceleratorDdosAttachment(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ga_accelerator_ddos_attachment gaService.DescribeGaAcceleratorDdosAttachment Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}

	d.Set("accelerator_id", parts[0])
	d.Set("ddos_id", parts[1])
	d.Set("ddos_region_id", object["DdosRegionId"])
	d.Set("status", object["Status"])

	return nil
}

func resourceAliCloudGaAcceleratorDdosAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println(fmt.Sprintf("[WARNING] The resouce has not update operation."))
	return resourceAliCloudGaAcceleratorDdosAttachmentRead(d, meta)
}

func resourceAliCloudGaAcceleratorDdosAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "DetachDdosFromAccelerator"
	var err error

	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}

	request := map[string]interface{}{
		"RegionId":      client.RegionId,
		"ClientToken":   buildClientToken("DetachDdosFromAccelerator"),
		"AcceleratorId": parts[0],
		"DdosConfigList": []map[string]interface{}{
			{
				"DdosId":       parts[1],
				"DdosRegionId": d.Get("ddos_region_id"),
			},
		},
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = retry.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutDelete)), func() *retry.RetryError {
		response, err = client.RpcPost("Ga", "2019-11-20", action, nil, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"StateError.Accelerator", "NotActive.Accelerator"}) || NeedRetry(err) {
				wait()
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
