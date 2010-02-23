package vm
import "testing"
import "os"

func defaultSynchronousChannel() chan *Stream {
	return make(chan *Stream)
}

func defaultAsynchronousChannel() chan *Stream {
	return make(chan *Stream, 256)
}

func defaultIOController() *IOController {
	i := new(IOController)
	i.Init()
	i.Open(defaultSynchronousChannel())
	i.Open(defaultAsynchronousChannel())
	return i
}

func checkChannelsAssigned(i *IOController, t *testing.T, channels int) {
	compareValues(i, t, i.Len(), channels)
	compareValues(i, t, i.Cap(), channels)
}

func TestCreateIOController(t *testing.T) {
	os.Stdout.WriteString("IOController Creation\n")
	i := new(IOController)
	i.Init()
	checkChannelsAssigned(i, t, 0)
	i.Open(defaultSynchronousChannel())
	checkChannelsAssigned(i, t, 1)
	i.Open(defaultAsynchronousChannel())
	checkChannelsAssigned(i, t, 2)
	ioc := i.Clone()
	checkChannelsAssigned(ioc, t, 2)
}

func TestIOControllerTraffic(t *testing.T) {
	os.Stdout.WriteString("IOController Traffic\n")
	i := defaultIOController()
	i.Send(0, defaultStream())
	checkDefaultStream(i.Receive(0), t, true)
}
