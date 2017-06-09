package librato

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/henrikhodne/go-librato/librato"
)

func TestAccLibratoAlert_Minimal(t *testing.T) {
	var alert librato.Alert
	name := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLibratoAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLibratoAlertConfig_minimal(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					testAccCheckLibratoAlertName(&alert, name),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", name),
				),
			},
		},
	})
}

func TestAccLibratoAlert_Basic(t *testing.T) {
	var alert librato.Alert
	name := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLibratoAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLibratoAlertConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					testAccCheckLibratoAlertName(&alert, name),
					testAccCheckLibratoAlertDescription(&alert, "A Test Alert"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", name),
				),
			},
		},
	})
}

func TestAccLibratoAlert_Full(t *testing.T) {
	var alert librato.Alert
	name := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLibratoAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLibratoAlertConfig_full(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					testAccCheckLibratoAlertName(&alert, name),
					testAccCheckLibratoAlertDescription(&alert, "A Test Alert"),
					testAccCheckLibratoAlertTags(&alert, "some_tag", []*string{librato.String("value1")}),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.836525194.metric_name", "librato.cpu.percent.idle"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.836525194.type", "above"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.836525194.threshold", "10"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.836525194.duration", "600"),
				),
			},
		},
	})
}

func TestAccLibratoAlert_Updated(t *testing.T) {
	var alert librato.Alert
	name := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLibratoAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLibratoAlertConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					testAccCheckLibratoAlertDescription(&alert, "A Test Alert"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", name),
				),
			},
			{
				Config: testAccCheckLibratoAlertConfig_new_value(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					testAccCheckLibratoAlertDescription(&alert, "A modified Test Alert"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "description", "A modified Test Alert"),
				),
			},
		},
	})
}

func TestAccLibratoAlert_Rename(t *testing.T) {
	var alert librato.Alert
	name := acctest.RandString(10)
	newName := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLibratoAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLibratoAlertConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", name),
				),
			},
			{
				Config: testAccCheckLibratoAlertConfig_basic(newName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", newName),
				),
			},
		},
	})
}

func TestAccLibratoAlert_FullUpdate(t *testing.T) {
	var alert librato.Alert
	name := acctest.RandString(10)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLibratoAlertDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLibratoAlertConfig_full_update(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLibratoAlertExists("librato_alert.foobar", &alert),
					testAccCheckLibratoAlertName(&alert, name),
					testAccCheckLibratoAlertDescription(&alert, "A Test Alert"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "rearm_seconds", "1200"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.2524844643.metric_name", "librato.cpu.percent.idle"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.2524844643.type", "above"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.2524844643.threshold", "10"),
					resource.TestCheckResourceAttr(
						"librato_alert.foobar", "condition.2524844643.duration", "60"),
				),
			},
		},
	})
}

func testAccCheckLibratoAlertDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*librato.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "librato_alert" {
			continue
		}

		id, err := strconv.ParseUint(rs.Primary.ID, 10, 0)
		if err != nil {
			return fmt.Errorf("ID not a number")
		}

		_, _, err = client.Alerts.Get(uint(id))

		if err == nil {
			return fmt.Errorf("Alert still exists")
		}
	}

	return nil
}

func testAccCheckLibratoAlertName(alert *librato.Alert, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if alert.Name == nil || *alert.Name != name {
			return fmt.Errorf("Bad name: %s", *alert.Name)
		}

		return nil
	}
}

func testAccCheckLibratoAlertDescription(alert *librato.Alert, description string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if alert.Description == nil || *alert.Description != description {
			return fmt.Errorf("Bad description: %s", *alert.Description)
		}

		return nil
	}
}

func testAccCheckLibratoAlertExists(n string, alert *librato.Alert) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Alert ID is set")
		}

		client := testAccProvider.Meta().(*librato.Client)

		id, err := strconv.ParseUint(rs.Primary.ID, 10, 0)
		if err != nil {
			return fmt.Errorf("ID not a number")
		}

		foundAlert, _, err := client.Alerts.Get(uint(id))

		if err != nil {
			return err
		}

		if foundAlert.ID == nil || *foundAlert.ID != uint(id) {
			return fmt.Errorf("Alert not found")
		}

		*alert = *foundAlert

		return nil
	}
}
func CompareSlicesOfStringPointers(a, b []*string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if *v != *b[i] {
			return false
		}
	}

	return true
}

func testAccCheckLibratoAlertTags(alert *librato.Alert, key string, values []*string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if alert.Conditions == nil || len(alert.Conditions) != 1 {
			return fmt.Errorf("No condition found: %s", *alert.Name)
		}

		tagSet := alert.Conditions[0].Tags[0]

		if *tagSet.Name != key {
			return fmt.Errorf("Incorrect key for tag %s in alert %s", *tagSet.Name, *alert.Name)
		}
		//reflect.DeepEqual(tagSet.Values, values)

		if !CompareSlicesOfStringPointers(tagSet.Values, values) {
			return fmt.Errorf("Incorrect value for tag %s in alert %s", *tagSet.Name, *alert.Name)
		}

		return nil
	}
}

func testAccCheckLibratoAlertConfig_minimal(name string) string {
	return fmt.Sprintf(`
resource "librato_alert" "foobar" {
    name = "%s"
}`, name)
}

func testAccCheckLibratoAlertConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "librato_alert" "foobar" {
    name = "%s"
    description = "A Test Alert"
}`, name)
}

func testAccCheckLibratoAlertConfig_new_value(name string) string {
	return fmt.Sprintf(`
resource "librato_alert" "foobar" {
    name = "%s"
    description = "A modified Test Alert"
}`, name)
}

func testAccCheckLibratoAlertConfig_full(name string) string {
	return fmt.Sprintf(`
resource "librato_service" "foobar" {
    title = "Foo Bar"
    type = "mail"
    settings = <<EOF
{
  "addresses": "admin@example.com"
}
EOF
}

resource "librato_alert" "foobar" {
    name = "%s"
    description = "A Test Alert"
    services = [ "${librato_service.foobar.id}" ]
    condition {
      type = "above"
      threshold = 10
      duration = 600
      metric_name = "librato.cpu.percent.idle"
			tag {
				name = "some_tag"
				values = [ "value1" ]
			}
    }
    attributes {
      runbook_url = "https://www.youtube.com/watch?v=oHg5SJYRHA0"
    }
    active = false
    rearm_seconds = 300
}`, name)
}

func testAccCheckLibratoAlertConfig_full_update(name string) string {
	return fmt.Sprintf(`
resource "librato_service" "foobar" {
    title = "Foo Bar"
    type = "mail"
    settings = <<EOF
{
  "addresses": "admin@example.com"
}
EOF
}

resource "librato_alert" "foobar" {
    name = "%s"
    description = "A Test Alert"
    services = [ "${librato_service.foobar.id}" ]
    condition {
      type = "above"
      threshold = 10
      duration = 60
      metric_name = "librato.cpu.percent.idle"
    }
    attributes {
      runbook_url = "https://www.youtube.com/watch?v=oHg5SJYRHA0"
    }
    active = false
    rearm_seconds = 1200
}`, name)
}
