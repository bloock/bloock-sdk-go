package exception

type WaitAnchorTimeoutException struct {
}

func NewWaitAnchorTimeoutException() WaitAnchorTimeoutException{
	return WaitAnchorTimeoutException{
	}
}

func(h WaitAnchorTimeoutException) Error() string {
	return "Timeout exceeded while waiting for anchor"
}