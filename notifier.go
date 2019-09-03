package notifier

// Notifier defines an interface for send info and errors to
// the busha team. can be via slack, email, sms e.t.c.
type Notifier interface {
	Error(msg string, optMsg ...string)
	Info(msg string, optMsg ...string)
}

// EmptyNotifier does nothing
type EmptyNotifier struct{}

// Error do nothing
func (EmptyNotifier) Error(msg string, optMsg ...string) {}

// Info do nothing
func (EmptyNotifier) Info(msg string, optMsg ...string) {}
