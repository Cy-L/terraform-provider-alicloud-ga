---
subcategory: "Global Accelerator (GA)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ga_accelerator_ddos_attachment"
sidebar_current: "docs-alicloud-resource-ga-accelerator-ddos-attachment"
description: |-
  Provides a Alicloud Global Accelerator (GA) Accelerator Ddos Attachment resource.
---

# alicloud_ga_accelerator_ddos_attachment

Provides a Global Accelerator (GA) Accelerator Ddos Attachment resource.

For information about Global Accelerator (GA) Accelerator Ddos Attachment and how to use it, see [What is Accelerator Ddos Attachment](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-attachddostoaccelerator).

-> **NOTE:** Available since v1.240.0.

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/terraform?resource=alicloud_ga_accelerator_ddos_attachment&exampleId=63ba4e5a-0874-fd03-ee1b-3964e0a191ca3b17a5f5&activeTab=example&spm=docs.r.ga_accelerator_ddos_attachment.0.63ba4e5a08&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
resource "alicloud_ga_accelerator" "default" {
  payment_type           = "PayAsYouGo"
  bandwidth_billing_type = "CDT"
}

resource "alicloud_ga_accelerator_ddos_attachment" "default" {
  accelerator_id = alicloud_ga_accelerator.default.id
  ddos_id        = "ddoscoo-cn-xxx"
  ddos_region_id = "cn-hangzhou"
}
```

## Argument Reference

The following arguments are supported:

* `accelerator_id` - (Required, ForceNew) The ID of the Global Accelerator instance.
* `ddos_id` - (Required, ForceNew) The ID of the DDoS protection instance.
* `ddos_region_id` - (Required, ForceNew) The region ID of the DDoS protection instance. Valid values: `cn-hangzhou`, `ap-southeast-1`.
* `dry_run` - (Optional, Bool) Specifies whether to only precheck the request.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID of Accelerator Ddos Attachment. The value formats as `<accelerator_id>:<ddos_id>`.
* `status` - The status of the Accelerator Ddos Attachment.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 mins) Used when create the Accelerator Ddos Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Accelerator Ddos Attachment.

## Import

Global Accelerator (GA) Accelerator Ddos Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_ga_accelerator_ddos_attachment.example <accelerator_id>:<ddos_id>
```
