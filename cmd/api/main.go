package main

func main() {
	s := newService()

	// Start Workers.
	s.Job.StartWorkers()

	// Listen and serve .
	s.HTTP.ListenAndServe()
	s.HTTP.GracefulShutdown()
}
