package Shapeshifter_obfs4_OpenVPN_Transport_Plugin
import "C"
import (
	"github.com/OperatorFoundation/shapeshifter-transports/transports/obfs4"
	"github.com/OperatorFoundation/shapeshifter-transports/transports/base"
)

var Obfs4_c_client *obfs4.Obfs4Transport = nil
var obfs4_transport_connection base.TransportConn = nil
var obfs4_transport_listener base.TransportListener = nil

func Initialize_obfs4_c_client(certString *C.char, iatMode int) {

	goCertString:= C.GoString(certString)
	Obfs4_c_client = obfs4.NewObfs4Client(goCertString, iatMode)
}

func Obfs4_listen(address_string *C.char) {

	goAddressString := C.GoString(address_string)
	obfs4_transport_listener = Obfs4_c_client.Listen(goAddressString)
}

func Obfs4_dial(address_string *C.char) {

	goAddressString := C.GoString(address_string)
	obfs4_transport_connection = Obfs4_c_client.Dial(goAddressString)
}