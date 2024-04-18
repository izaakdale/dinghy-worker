package consensus

import "log"

type snapshotSink struct{}

func (s *snapshotSink) Write([]byte) (int, error) {
	log.Printf("write in snapshot sink\n")
	return 0, nil
}
func (s *snapshotSink) Close() error {
	log.Printf("closing\n")
	return nil
}
func (s *snapshotSink) ID() string {
	log.Printf("ID\n")
	return ""
}
func (s *snapshotSink) Cancel() error {
	log.Printf("cancel\n")
	return nil
}
