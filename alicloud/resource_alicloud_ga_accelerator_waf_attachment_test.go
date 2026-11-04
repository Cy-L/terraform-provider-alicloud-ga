package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAliCloudGaAcceleratorWafAttachment_basic0(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.GaSupportRegions)
	resourceId := "alicloud_ga_accelerator_waf_attachment.default"
	ra := resourceAttrInit(resourceId, AliCloudGaAcceleratorWafAttachmentMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GaService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGaAcceleratorWafAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sGaAcceleratorWafAttachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudGaAcceleratorWafAttachmentBasicDependence0)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPrePaidResources(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"accelerator_id":                "${alicloud_ga_accelerator.default.id}",
					"associated_resource_id":        "${data.alicloud_wafv3_instances.default.ids.0}",
					"associated_resource_region_id": "${data.alicloud_regions.default.regions.0.id}",
					"associated_resource_type":      "WAF",
					"associated_mode":               "Associated",
					"dry_run":                       "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"state": CHECKSET,
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudGaAcceleratorWafAttachment_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.GaSupportRegions)
	resourceId := "alicloud_ga_accelerator_waf_attachment.default"
	ra := resourceAttrInit(resourceId, AliCloudGaAcceleratorWafAttachmentMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GaService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGaAcceleratorWafAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sGaAcceleratorWafAttachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudGaAcceleratorWafAttachmentBasicDependence0)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPrePaidResources(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"accelerator_id":                "${alicloud_ga_accelerator.default.id}",
					"associated_resource_id":        "${data.alicloud_wafv3_instances.default.ids.0}",
					"associated_resource_region_id": "${data.alicloud_regions.default.regions.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"state": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

var AliCloudGaAcceleratorWafAttachmentMap0 = map[string]string{
	"state": CHECKSET,
}

func AliCloudGaAcceleratorWafAttachmentBasicDependence0(name string) string {
	return fmt.Sprintf(`
data "alicloud_regions" "default" {
  current = true
}

data "alicloud_wafv3_instances" "default" {
}

resource "alicloud_ga_accelerator" "default" {
  payment_type           = "PayAsYouGo"
  bandwidth_billing_type = "CDT"
}
`, name)
}
