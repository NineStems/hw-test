package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return nil
	}
	innerIn := in
	for i := range stages {
		innerIn = next(innerIn, done, stages[i])
	}
	return innerIn
}

func next(in, done In, f Stage) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for data := range f(in) {
			select {
			case <-done:
				return
			default:
				out <- data
			}
		}
	}()
	return out
}
