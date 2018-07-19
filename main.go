package main
import "C"
import (
	"github.com/OperatorFoundation/shapeshifter-transports/transports/obfs4"
	"github.com/OperatorFoundation/shapeshifter-transports/transports/base"
	"unsafe"
)

var obfs4Clients = map[int]*obfs4.Obfs4Transport{}
var obfs4Connections = map[int]base.TransportConn{}
var nextID = 0
var obfs4_transport_listener base.TransportListener = nil

//export Initialize_obfs4_c_client
func Initialize_obfs4_c_client(certString *C.char, iatMode int) (clientKey int) {

	goCertString := C.GoString(certString)
	var obfs4Client *obfs4.Obfs4Transport = obfs4.NewObfs4Client(goCertString, iatMode)
	obfs4Clients[nextID] = obfs4Client

	// This is the return value
	clientKey = nextID

	nextID += 1
	return
}

//export Obfs4_listen
func Obfs4_listen(address_string *C.char) {

	//goAddressString := C.GoString(address_string)
	//obfs4_transport_listener = Obfs4_c_client.Listen(goAddressString)
}

//export Obfs4_dial
func Obfs4_dial(client_id int, address_string *C.char) int {

	goAddressString := C.GoString(address_string)

	var transport = obfs4Clients[client_id]
	var obfs4_transport_connection = transport.Dial(goAddressString)

	if obfs4_transport_connection == nil {
		return 1
	} else {
		obfs4Connections[client_id] = obfs4_transport_connection
		return 0
	}
}

//export Obfs4_write
func Obfs4_write(client_id int, buffer unsafe.Pointer, buffer_length C.int) int {
	var connection = obfs4Connections[client_id]
	var bytesBuffer = C.GoBytes(buffer, buffer_length)
	numberOfBytesWritten, error := connection.Write(bytesBuffer)

	if error != nil {
		return -1
	} else {
		return numberOfBytesWritten
	}
}

//export Obfs4_read
func Obfs4_read(client_id int, buffer unsafe.Pointer, buffer_length C.int) int {

	var connection = obfs4Connections[client_id]
	var bytesBuffer = C.GoBytes(buffer, buffer_length)

	numberOfBytesRead, error := connection.Read(bytesBuffer)

	if error != nil {
		return -1
	} else {
		return numberOfBytesRead
	}
}

//export Obfs4_close_connection
func Obfs4_close_connection(client_id int) {

	var connection = obfs4Connections[client_id]
	connection.Close()
	delete(obfs4Connections, client_id)
	delete(obfs4Clients, client_id)
}

func main() {}