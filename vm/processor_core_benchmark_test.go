package vm

import . "golightly/storage"
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
		c := make(chan IntBuffer)
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

func BenchmarkLoadProgram(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.LoadProgram(simpleProgram(p)) }
}

func BenchmarkStepForward(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.PC++ }
}

func BenchmarkStepBack(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.PC++ }
}

func BenchmarkStepping(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ {
		p.PC++
		p.PC++
	}
}

func BenchmarkStepExecuteRewind(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ {
		p.PC++
		p.Execute()
		p.PC++
	}
}

func BenchmarkStepExecuteInlineRewind(b *testing.B) {
	b.StopTimer()
		p := new(InlinedProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ {
		p.PC++
		p.Execute()
		p.PC++
	}
}

func BenchmarkJumpTo(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.PC = 0 }
}

func BenchmarkJumpRelative(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.PC += 0 }
}

func BenchmarkProgramRun(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(advancedProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.Run(); p.ResetState() }
}

func BenchmarkProgramRunInline(b *testing.B) {
	b.StopTimer()
		p := new(InlinedProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(advancedProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.Run(); p.ResetState() }
}
