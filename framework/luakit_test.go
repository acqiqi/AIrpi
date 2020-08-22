package framework

import (
	"testing"
	"time"

	"AIrpi/test"
)

func TestRobotConnectionEach(t *testing.T) {
	r := newTestRobot("Robot1")

	i := 0
	r.Connections().Each(func(conn Connection) {
		i++
	})
	test.Assert(t, r.Connections().Len(), i)
}

func TestRobotToJSON(t *testing.T) {
	r := newTestRobot("Robot99")
	r.AddCommand("test_function", func(params map[string]interface{}) interface{} {
		return nil
	})
	json := NewJSONRobot(r)
	test.Assert(t, len(json.Devices), r.Devices().Len())
	test.Assert(t, len(json.Commands), len(r.Commands()))
}

func TestRobotDevicesToJSON(t *testing.T) {
	r := newTestRobot("Robot99")
	json := NewJSONRobot(r)
	test.Assert(t, len(json.Devices), r.Devices().Len())
	test.Assert(t, json.Devices[0].Name, "Device1")
	test.Assert(t, json.Devices[0].Driver, "*gobot.testDriver")
	test.Assert(t, json.Devices[0].Connection, "Connection1")
	test.Assert(t, len(json.Devices[0].Commands), 1)
}

func TestRobotStart(t *testing.T) {
	r := newTestRobot("Robot99")
	test.Assert(t, r.Start(), nil)
	test.Assert(t, r.Stop(), nil)
	test.Assert(t, r.Running(), false)
}

func TestRobotStartAutoRun(t *testing.T) {
	adaptor1 := newTestAdaptor("Connection1", "/dev/null")
	driver1 := newTestDriver(adaptor1, "Device1", "0")
	//work := func() {}
	r := NewRobot("autorun",
		[]Connection{adaptor1},
		[]Device{driver1},
		//work,
	)

	go func() {
		test.Assert(t, r.Start(), nil)
	}()

	time.Sleep(10 * time.Millisecond)
	test.Assert(t, r.Running(), true)

	// stop it
	test.Assert(t, r.Stop(), nil)
	test.Assert(t, r.Running(), false)
}
