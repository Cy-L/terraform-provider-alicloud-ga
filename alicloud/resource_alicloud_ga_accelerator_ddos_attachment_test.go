package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAliCloudGaAcceleratorDdosAttachment_basic0(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.GaSupportRegions)
	resourceId := "alicloud_ga_accelerator_ddos_attachment.default"
	ra := resourceAttrInit(resourceId, AliCloudGaAcceleratorDdosAttachmentMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GaService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGaAcceleratorDdosAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sGaAcceleratorDdosAttachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudGaAcceleratorDdosAttachmentBasicDependence0)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPrePaidResources(t)
			testAccPreCheckWithRegions(t, true, connectivity.DdoscooSupportedRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"accelerator_id": "${alicloud_ga_accelerator.default.id}",
					"ddos_id":        "${data.alicloud_ddoscoo_instances.default.ids.0}",
					"ddos_region_id": "${data.alicloud_regions.default.regions.0.id}",
					"dry_run":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": CHECKSET,
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

func TestAccAliCloudGaAcceleratorDdosAttachment_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.GaSupportRegions)
	resourceId := "alicloud_ga_accelerator_ddos_attachment.default"
	ra := resourceAttrInit(resourceId, AliCloudGaAcceleratorDdosAttachmentMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GaService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGaAcceleratorDdosAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sGaAcceleratorDdosAttachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudGaAcceleratorDdosAttachmentBasicDependence0)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckPrePaidResources(t)
			testAccPreCheckWithRegions(t, true, connectivity.DdoscooSupportedRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"accelerator_id": "${alicloud_ga_accelerator.default.id}",
					"ddos_id":        "${data.alicloud_ddoscoo_instances.default.ids.0}",
					"ddos_region_id": "${data.alicloud_regions.default.regions.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": CHECKSET,
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

var AliCloudGaAcceleratorDdosAttachmentMap0 = map[string]string{
	"status": CHECKSET,
}

func AliCloudGaAcceleratorDdosAttachmentBasicDependence0(name string) string {
	return fmt.Sprintf(`
data "alicloud_regions" "default" {
  current = true
}

data "alicloud_ddoscoo_instances" "default" {
}

resource "alicloud_ga_accelerator" "default" {
  payment_type           = "PayAsYouGo"
  bandwidth_billing_type = "CDT"
}
`, name)
}
