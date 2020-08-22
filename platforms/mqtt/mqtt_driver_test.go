package mqtt

import (
	"AIrpi/framework"
	"AIrpi/test"
	"strings"
	"testing"
)

var _ framework.Driver = (*Driver)(nil)

func TestMqttDriver(t *testing.T) {
	d := NewDriver(initTestMqttAdaptor(), "/test/topic")

	test.Assert(t, strings.HasPrefix(d.Name(), "MQTT"), true)
	test.Assert(t, strings.HasPrefix(d.Connection().Name(), "MQTT"), true)

	test.Assert(t, d.Start(), nil)
	test.Assert(t, d.Halt(), nil)
}

func TestMqttDriverName(t *testing.T) {
	d := NewDriver(initTestMqttAdaptor(), "/test/topic")
	test.Assert(t, strings.HasPrefix(d.Name(), "MQTT"), true)
	d.SetName("NewName")
	test.Assert(t, d.Name(), "NewName")
}

func TestMqttDriverTopic(t *testing.T) {
	d := NewDriver(initTestMqttAdaptor(), "/test/topic")
	test.Assert(t, d.Topic(), "/test/topic")
	d.SetTopic("/test/newtopic")
	test.Assert(t, d.Topic(), "/test/newtopic")
}

func TestMqttDriverPublish(t *testing.T) {
	a := initTestMqttAdaptor()
	d := NewDriver(a, "/test/topic")
	a.Connect()
	d.Start()
	defer d.Halt()
	test.Assert(t, d.Publish([]byte{0x01, 0x02, 0x03}), true)
}

func TestMqttDriverPublishError(t *testing.T) {
	a := initTestMqttAdaptor()
	d := NewDriver(a, "/test/topic")
	d.Start()
	defer d.Halt()
	test.Assert(t, d.Publish([]byte{0x01, 0x02, 0x03}), false)
}
