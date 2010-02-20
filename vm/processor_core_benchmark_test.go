package vm
import "testing"

func BenchmarkCreateProcessorCore(b *testing.B) {
    for i := 0; i < b.N; i++ {
		p := new(ProcessorCore)
		p.Init(32, nil)
	}
}

func BenchmarkCloneProcessorCore(b *testing.B) {
	b.StopTimer()
	p := new(ProcessorCore)
	p.Init(32, nil)
	c := make(chan *Buffer)
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.Clone(c) }
}

func BenchmarkValidPC(b *testing.B) {
	b.StopTimer()
	p := new(ProcessorCore)
	p.Init(32, nil)
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.ValidPC() }
}

func BenchmarkResetState(b *testing.B) {
	b.StopTimer()
	p := new(ProcessorCore)
	p.Init(32, nil)
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.ResetState() }
}
