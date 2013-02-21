//	TODO:	check both synchronous and asynchronous channels are supported
//	TODO:	Read() and Write() interface support
//	TODO:	write tests

package vm

import "github.com/feyeleanor/slices"

type IOController []chan slices.ISlice

func (ioc *IOController) realloc(length int) (c []chan slices.ISlice) {
	c = make([]chan slices.ISlice, length)
	copy(c, *ioc)
	*ioc = c
	return
}

/*
func (ioc IOController) Open(c chan slices.ISlice) {
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
func (ioc IOController) Send(i int, x slices.ISlice) {
	go func() {
		c := make(slices.ISlice, 0, x.Len())
		copy(c, x)
		ioc[i] <- c
	}()
}
func (ioc IOController) Receive(i int) slices.ISlice				{ return <-ioc[i] }
func (ioc IOController) Copy(i, j int)							{ ioc[i]<- <-ioc[j] }

//func (ioc *IOController) Do(f func(x slices.ISlice))				{ for _, v := range *ioc { f(<-v) } }
