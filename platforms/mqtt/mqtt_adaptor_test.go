package mqtt

import (
	"AIrpi/framework"
	"AIrpi/test"
	"errors"
	"fmt"
	multierror "github.com/hashicorp/go-multierror"
	"strings"
	"testing"
)

var _ framework.Adaptor = (*Adaptor)(nil)

func initTestMqttAdaptor() *Adaptor {
	return NewAdaptor("tcp://localhost:1883", "client")
}

func TestMqttAdaptorName(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, strings.HasPrefix(a.Name(), "MQTT"), true)
	a.SetName("NewName")
	test.Assert(t, a.Name(), "NewName")
}

func TestMqttAdaptorPort(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.Port(), "tcp://localhost:1883")
}

func TestMqttAdaptorAutoReconnect(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.AutoReconnect(), false)
	a.SetAutoReconnect(true)
	test.Assert(t, a.AutoReconnect(), true)
}

func TestMqttAdaptorCleanSession(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.CleanSession(), true)
	a.SetCleanSession(false)
	test.Assert(t, a.CleanSession(), false)
}

func TestMqttAdaptorUseSSL(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.UseSSL(), false)
	a.SetUseSSL(true)
	test.Assert(t, a.UseSSL(), true)
}

func TestMqttAdaptorUseServerCert(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.ServerCert(), "")
	a.SetServerCert("/path/to/server.cert")
	test.Assert(t, a.ServerCert(), "/path/to/server.cert")
}

func TestMqttAdaptorUseClientCert(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.ClientCert(), "")
	a.SetClientCert("/path/to/client.cert")
	test.Assert(t, a.ClientCert(), "/path/to/client.cert")
}

func TestMqttAdaptorUseClientKey(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.ClientKey(), "")
	a.SetClientKey("/path/to/client.key")
	test.Assert(t, a.ClientKey(), "/path/to/client.key")
}

func TestMqttAdaptorConnectError(t *testing.T) {
	a := NewAdaptor("tcp://localhost:1884", "client")

	err := a.Connect()
	test.Assert(t, strings.Contains(err.Error(), "connection refused"), true)
}

func TestMqttAdaptorConnectSSLError(t *testing.T) {
	a := NewAdaptor("tcp://localhost:1884", "client")
	a.SetUseSSL(true)
	err := a.Connect()
	test.Assert(t, strings.Contains(err.Error(), "connection refused"), true)
}

func TestMqttAdaptorConnectWithAuthError(t *testing.T) {
	a := NewAdaptorWithAuth("xyz://localhost:1883", "client", "user", "pass")
	var expected error
	expected = multierror.Append(expected, errors.New("Network Error : Unknown protocol"))

	test.Assert(t, a.Connect(), expected)
}

func TestMqttAdaptorFinalize(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.Finalize(), nil)
}

func TestMqttAdaptorCannotPublishUnlessConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	data := []byte("o")
	test.Assert(t, a.Publish("test", data), false)
}

func TestMqttAdaptorPublishWhenConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	a.Connect()
	data := []byte("o")
	test.Assert(t, a.Publish("test", data), true)
}

func TestMqttAdaptorCannotOnUnlessConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	test.Assert(t, a.On("hola", func(msg Message) {
		fmt.Println("hola")
	}), false)
}

func TestMqttAdaptorOnWhenConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	a.Connect()
	test.Assert(t, a.On("hola", func(msg Message) {
		fmt.Println("hola")
	}), true)
}

func TestMqttAdaptorQoS(t *testing.T) {
	a := initTestMqttAdaptor()
	a.SetQoS(1)
	test.Assert(t, 1, a.qos)
}
