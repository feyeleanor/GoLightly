//	TODO:	check both synchronous and asynchronous channels are supported
//	TODO:	Read() and Write() interface support
//	TODO:	write tests

package vm

import . "golightly/storage"

type IOController []chan IntBuffer

func (ioc *IOController) realloc(length int) (c []chan IntBuffer) {
	c = make([]chan IntBuffer, length)
	copy(c, *ioc)
	*ioc = c
	return
}

/*
func (ioc IOController) Open(c chan IntBuffer) {
	starting_length := ioc.Len()
	ioc.realloc(starting_length + 1)
	ioc.Set(starting_length, c)
}
*/

func (ioc IOController) Clone() IOController {
	c := make(IOController, len(ioc))
	copy(c, ioc)
	return c
}

//func (ioc *IOController) Init()									{ ioc.realloc(0) }
func (ioc IOController) Close(i int)							{ close(ioc[i]) }
func (ioc IOController) CloseAll()								{ for i, _ := range ioc { ioc.Close(i) } }
func (ioc IOController) Send(i int, x IntBuffer)				{ go func() { ioc[i] <- x.Clone() }() }
func (ioc IOController) Receive(i int) IntBuffer				{ return <-ioc[i] }
func (ioc IOController) Copy(i, j int)							{ ioc[i]<- <-ioc[j] }

//func (ioc *IOController) Do(f func(x IntVector))				{ for _, v := range *ioc { f(<-v) } }
