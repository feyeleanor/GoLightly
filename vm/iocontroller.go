//	TODO:	check both synchronous and asynchronous channels are supported
//	TODO:	Read() and Write() interface support
//	TODO:	write tests

package vm

type IOController []chan int

func (ioc *IOController) realloc(length int) (c []chan int) {
	c = make([]chan int, length)
	copy(c, *ioc)
	*ioc = c
	return
}

func (ioc *IOController) Open(c chan int) {
	starting_length := ioc.Len()
	ioc.realloc(starting_length + 1)
	ioc.Set(starting_length, c)
}

func (ioc *IOController) iterate(c chan<- int) {
	for _, v := range *ioc { c<- <-v }
	close(c)
}

func (ioc *IOController) Iter() <-chan int {
	c := make(chan int)
	go ioc.iterate(c)
	return c
}

func (ioc *IOController) Clone() *IOController {
	c := new(IOController)
	c.realloc(ioc.Len())
	copy(*c, *ioc)
	return c
}

func (ioc *IOController) Init()									{ ioc.realloc(0) }
func (ioc *IOController) Len() int								{ return len(*ioc) }
func (ioc *IOController) Cap() int								{ return cap(*ioc) }
func (ioc *IOController) At(i int) chan int						{ return (*ioc)[i] }
func (ioc *IOController) Set(i int, c chan int)					{ (*ioc)[i] = c }
func (ioc *IOController) Close(i int)							{ close(ioc.At(i)) }
func (ioc *IOController) CloseAll()								{ for i, _ := range *ioc { ioc.Close(i) } }
func (ioc *IOController) Send(i, x int)							{ ioc.At(i) <- x }
func (ioc *IOController) Receive(i int) int						{ return <-ioc.At(i) }
func (ioc *IOController) Copy(i, j int)							{ ioc.At(i)<- <-ioc.At(j) }
func (ioc *IOController) Do(f func(elem interface{}))			{ for _, e := range *ioc { f(e) } }
