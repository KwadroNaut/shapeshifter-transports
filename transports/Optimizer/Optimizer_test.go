package Optimizer

import (
	"github.com/OperatorFoundation/shapeshifter-transports/transports/shadow"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	transport := shadow.NewShadowServer("orange", "aes-128-ctr")
	listener := transport.Listen("127.0.0.1:1234")
	go acceptConnections(listener)

	os.Exit(m.Run())
}


func acceptConnections(listener net.Listener) {
	for {
		_, err := listener.Accept()
		if err != nil {
			return
		}
	}
}


func TestShadowDial1(t *testing.T) {
	shadowTransport := ShadowTransport{"orange", "aes-128-ctr", "127.0.0.1:1234"}
	conn := shadowTransport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestShadowDial2(t *testing.T) {
	shadowTransport := ShadowTransport{"banana", "aes-192-ctr", "127.0.0.1:1234"}
	conn := shadowTransport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestOptimizerShadowDial1(t *testing.T) {
	shadowTransport := ShadowTransport{"orange", "aes-128-ctr", "127.0.0.1:1234"}
	transports := []Transport{shadowTransport}
	strategy := FirstStrategy{}
	optimizerTransport := NewOptimizerClient(transports, strategy)
	conn := optimizerTransport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestOptimizerShadowDial2(t *testing.T) {
	shadowTransport := ShadowTransport{"banana", "aes-192-ctr", "127.0.0.1:1234"}
	transports := []Transport{shadowTransport}
	strategy := FirstStrategy{}
	optimizerTransport := NewOptimizerClient(transports, strategy)
	conn := optimizerTransport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestObfs4Transport_Dial1(t *testing.T) {
	Obfs4Transport := Obfs4Transport{"UsuF7oN4KNKviZP54JOyTCoCphrdM5gwZK4vT8GnCAcmqLUJEJxyw1dpko9a/ii6He4iZg", 0, "77.81.104.251:443"}
	conn := Obfs4Transport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestObfs4Transport_Dial2(t *testing.T) {
	Obfs4Transport := Obfs4Transport{"BBKeJPokZXigyKpn+E/iKim/BwNEiIdifbHfaXQmyu1GpSHtNlruAIWebci9m8Yb0tGUOw", 0, "5.253.87.21:443"}
	conn := Obfs4Transport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestOptimizerObfs4Transport_Dial1(t *testing.T) {
	Obfs4Transport := Obfs4Transport{"UsuF7oN4KNKviZP54JOyTCoCphrdM5gwZK4vT8GnCAcmqLUJEJxyw1dpko9a/ii6He4iZg", 0, "77.81.104.251:443"}
	transports := []Transport{Obfs4Transport}
	strategy := FirstStrategy{}
	optimizerTransport := NewOptimizerClient(transports, strategy)
	conn := optimizerTransport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestOptimizerObfs4Transport_Dial2(t *testing.T) {
	Obfs4Transport := Obfs4Transport{"BBKeJPokZXigyKpn+E/iKim/BwNEiIdifbHfaXQmyu1GpSHtNlruAIWebci9m8Yb0tGUOw", 0, "5.253.87.21:443"}
	transports := []Transport{Obfs4Transport}
	strategy := FirstStrategy{}
	optimizerTransport := NewOptimizerClient(transports, strategy)
	conn := optimizerTransport.Dial()
	if conn == nil {
		t.Fail()
	}
}

func TestOptimizerTransportFirstDial(t *testing.T) {
	obfs4Transport := Obfs4Transport{"UsuF7oN4KNKviZP54JOyTCoCphrdM5gwZK4vT8GnCAcmqLUJEJxyw1dpko9a/ii6He4iZg", 0, "77.81.104.251:443"}
	shadowTransport := ShadowTransport{"orange", "aes-128-ctr", "127.0.0.1:1234"}
	transports := []Transport{obfs4Transport, shadowTransport}
	optimizerTransport := NewOptimizerClient(transports, FirstStrategy{})
	for i := 1; i <= 3; i++ {
		conn := optimizerTransport.Dial()
		if conn == nil {
			t.Fail()
		}
	}
}

func TestOptimizerTransportRandomDial(t *testing.T) {
	obfs4Transport := Obfs4Transport{"UsuF7oN4KNKviZP54JOyTCoCphrdM5gwZK4vT8GnCAcmqLUJEJxyw1dpko9a/ii6He4iZg", 0, "77.81.104.251:443"}
	shadowTransport := ShadowTransport{"orange", "aes-128-ctr", "127.0.0.1:1234"}
	transports := []Transport{obfs4Transport, shadowTransport}
	optimizerTransport := NewOptimizerClient(transports, RandomStrategy{})

	for i := 1; i <= 3; i++ {
		conn := optimizerTransport.Dial()
		if conn == nil {
			t.Fail()
		}
	}
}

func TestOptimizerTransportRotateDial(t *testing.T) {
	obfs4Transport := Obfs4Transport{"UsuF7oN4KNKviZP54JOyTCoCphrdM5gwZK4vT8GnCAcmqLUJEJxyw1dpko9a/ii6He4iZg", 0, "77.81.104.251:443"}
	shadowTransport := ShadowTransport{"orange", "aes-128-ctr", "127.0.0.1:1234"}
	transports := []Transport{obfs4Transport, shadowTransport}
	optimizerTransport := NewOptimizerClient(transports, RotateStrategy{})

	for i := 1; i <= 3; i++ {
		conn := optimizerTransport.Dial()
		if conn == nil {
			t.Fail()
		}
	}
}