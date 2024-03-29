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
	for i := range stages {
		out = next(stages[i](out), done)
	}
	return out
}

func next(in, done In) Out {
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
