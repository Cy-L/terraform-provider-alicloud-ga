package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAliCloudGaAcceleratorWafAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudGaAcceleratorWafAttachmentCreate,
		Read:   resourceAliCloudGaAcceleratorWafAttachmentRead,
		Update: resourceAliCloudGaAcceleratorWafAttachmentUpdate,
		Delete: resourceAliCloudGaAcceleratorWafAttachmentDelete,
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
			"associated_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"associated_resource_region_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"associated_resource_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "WAF",
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"WAF"}, false),
			},
			"associated_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Associated",
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"Associated", "Managed"}, false),
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudGaAcceleratorWafAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "AssociateResources"
	request := make(map[string]interface{})
	var err error

	request["ClientToken"] = buildClientToken("AssociateResources")
	request["AcceleratorId"] = d.Get("accelerator_id")
	request["AssociatedResourceType"] = d.Get("associated_resource_type")
	request["AssociatedResourceId"] = d.Get("associated_resource_id")
	request["AssociatedResourceRegionId"] = d.Get("associated_resource_region_id")

	if v, ok := d.GetOk("associated_mode"); ok {
		request["AssociatedMode"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ga_accelerator_waf_attachment", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["AcceleratorId"], d.Get("associated_resource_id")))

	return resourceAliCloudGaAcceleratorWafAttachmentRead(d, meta)
}

func resourceAliCloudGaAcceleratorWafAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	gaService := GaService{client}
	object, err := gaService.DescribeGaAcceleratorWafAttachment(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ga_accelerator_waf_attachment gaService.DescribeGaAcceleratorWafAttachment Failed!!! %s", err)
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
	d.Set("associated_resource_id", parts[1])
	d.Set("associated_resource_region_id", object["AssociatedResourceRegionId"])
	d.Set("associated_resource_type", object["AssociatedResourceType"])
	d.Set("associated_mode", object["AssociatedMode"])
	d.Set("state", object["State"])

	return nil
}

func resourceAliCloudGaAcceleratorWafAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println(fmt.Sprintf("[WARNING] The resouce has not update operation."))
	return resourceAliCloudGaAcceleratorWafAttachmentRead(d, meta)
}

func resourceAliCloudGaAcceleratorWafAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "DisassociateResources"
	var err error

	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}

	request := map[string]interface{}{
		"RegionId":                 client.RegionId,
		"ClientToken":              buildClientToken("DisassociateResources"),
		"AcceleratorId":            parts[0],
		"AssociatedResourceType":   d.Get("associated_resource_type"),
		"AssociatedResourceId":     parts[1],
		"AssociatedResourceRegionId": d.Get("associated_resource_region_id"),
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
