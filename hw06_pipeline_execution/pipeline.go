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
	out := in
	for _, stage := range stages {
		out = next(done, stage(out))
	}
	return out
}

func next(done, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case data, open := <-in:
				if !open {
					return
				}
				out <- data
			}
		}
	}()
	return out
}
