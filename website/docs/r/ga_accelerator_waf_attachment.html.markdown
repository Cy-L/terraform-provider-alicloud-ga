---
subcategory: "Global Accelerator (GA)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ga_accelerator_waf_attachment"
sidebar_current: "docs-alicloud-resource-ga-accelerator-waf-attachment"
description: |-
  Provides a Alicloud Global Accelerator (GA) Accelerator Waf Attachment resource.
---

# alicloud_ga_accelerator_waf_attachment

Provides a Global Accelerator (GA) Accelerator Waf Attachment resource.

For information about Global Accelerator (GA) Accelerator Waf Attachment and how to use it, see [What is Accelerator Waf Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-associateresources).

-> **NOTE:** Available since v1.240.0.

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/terraform?resource=alicloud_ga_accelerator_waf_attachment&exampleId=63ba4e5a-0874-fd03-ee1b-3964e0a191ca3b17a5f5&activeTab=example&spm=docs.r.ga_accelerator_waf_attachment.0.63ba4e5a08&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
resource "alicloud_ga_accelerator" "test" {
  payment_type           = "PayAsYouGo"
  bandwidth_billing_type = "CDT"
}

resource "alicloud_ga_accelerator_waf_attachment" "test" {
  accelerator_id                = alicloud_ga_accelerator.test.id
  associated_resource_id        = "waf_v2_public_cn-xxxx"
  associated_resource_region_id = "cn-hangzhou"
}
```

## Argument Reference

The following arguments are supported:

* `accelerator_id` - (Required, ForceNew) The ID of the Global Accelerator instance.
* `associated_resource_id` - (Required, ForceNew) The ID of the WAF instance.
* `associated_resource_region_id` - (Required, ForceNew) The region ID of the WAF instance. Valid values: `cn-hangzhou`, `ap-southeast-1`.
* `associated_resource_type` - (Optional, ForceNew) The type of the associated resource. Valid values: `WAF`. Default value: `WAF`.
* `associated_mode` - (Optional, ForceNew) The association mode. Valid values: `Associated`, `Managed`. Default value: `Associated`.
* `dry_run` - (Optional, Bool) Specifies whether to only precheck the request.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID of Accelerator Waf Attachment. The value formats as `<accelerator_id>:<associated_resource_id>`.
* `state` - The state of the Accelerator Waf Attachment.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 mins) Used when create the Accelerator Waf Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Accelerator Waf Attachment.

## Import

Global Accelerator (GA) Accelerator Waf Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ga_accelerator_waf_attachment.example <accelerator_id>:<associated_resource_id>
```
