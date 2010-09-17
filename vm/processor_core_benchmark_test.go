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
		c := make(chan *Vector)
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
    for i := 0; i < b.N; i++ { p.StepForward() }
}

func BenchmarkStepBack(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.StepBack() }
}

func BenchmarkStepping(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ {
		p.StepForward()
		p.StepBack()
	}
}

func BenchmarkStepExecuteRewind(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ {
		p.StepForward()
		p.Execute()
		p.StepBack()
	}
}

func BenchmarkStepExecuteInlineRewind(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ {
		p.StepForward()
		p.ExecuteInline()
		p.StepBack()
	}
}

func BenchmarkJumpTo(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.JumpTo(0) }
}

func BenchmarkJumpRelative(b *testing.B) {
	b.StopTimer()
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(simpleProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.JumpRelative(0) }
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
		p := new(ProcessorCore)
		p.Init(32, nil)
		p.LoadProgram(advancedProgram(p))
	b.StartTimer()
    for i := 0; i < b.N; i++ { p.RunInline(); p.ResetState() }
}
