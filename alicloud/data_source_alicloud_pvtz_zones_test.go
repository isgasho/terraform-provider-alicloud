package alicloud

import (
	"testing"

	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAlicloudPvtzZonesDataSource_keyword(t *testing.T) {
	var pvtzZone pvtz.DescribeZoneInfoResponse
	rand := acctest.RandIntRange(10000, 999999)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccAlicloudPvtzZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAlicloudPvtzZoneDataSource_keyword(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccAlicloudPvtzZoneExists("alicloud_pvtz_zone.basic", &pvtzZone),
					testAccCheckAlicloudDataSourceID("data.alicloud_pvtz_zones.keyword"),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.#", "1"),
					resource.TestCheckResourceAttrSet("data.alicloud_pvtz_zones.keyword", "zones.0.id"),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.0.name", fmt.Sprintf("tf-testacc%d.test.com", rand)),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.0.remark", ""),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.0.record_count", "0"),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.0.is_ptr", "false"),
					resource.TestCheckResourceAttrSet("data.alicloud_pvtz_zones.keyword", "zones.0.creation_time"),
					resource.TestCheckResourceAttrSet("data.alicloud_pvtz_zones.keyword", "zones.0.update_time"),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.0.bind_vpcs.#", "0"),
				),
			},
			{
				Config: testAccCheckAlicloudPvtzZoneDataSource_keyword_empty(rand),
				Check: resource.ComposeTestCheckFunc(
					testAccAlicloudPvtzZoneExists("alicloud_pvtz_zone.basic", &pvtzZone),
					testAccCheckAlicloudDataSourceID("data.alicloud_pvtz_zones.keyword"),
					resource.TestCheckResourceAttr("data.alicloud_pvtz_zones.keyword", "zones.#", "0"),
				),
			},
		},
	})
}

func testAccCheckAlicloudPvtzZoneDataSource_keyword(rand int) string {
	return fmt.Sprintf(`
	resource "alicloud_pvtz_zone" "basic" {
		name = "tf-testacc%d.test.com"
	}
	data "alicloud_pvtz_zones" "keyword" {
		keyword = "${alicloud_pvtz_zone.basic.name}"
	}
	`, rand)
}

func testAccCheckAlicloudPvtzZoneDataSource_keyword_empty(rand int) string {
	return fmt.Sprintf(`
	resource "alicloud_pvtz_zone" "basic" {
		name = "tf-testacc%d.test.com"
	}
	data "alicloud_pvtz_zones" "keyword" {
		keyword = "${alicloud_pvtz_zone.basic.name}-fake"
	}
	`, rand)
}
