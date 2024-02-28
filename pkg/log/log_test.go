package log_test

import (
	"os"

	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/raharper/stacker/pkg/log"
)

func TestLog(t *testing.T) {
	Convey("With timestamps", t, func() {
		handler := log.NewTextHandler(os.Stderr, true)
		So(handler, ShouldNotBeNil)

		So(func() { log.Debugf("debug msg") }, ShouldNotPanic)
		So(func() { log.Infof("info msg") }, ShouldNotPanic)
		So(func() { log.Errorf("error msg") }, ShouldNotPanic)

		So(func() { log.FilterNonStackerLogs(handler, 1) }, ShouldNotPanic)

		So(func() { log.Debugf("debug msg") }, ShouldNotPanic)
		So(func() { log.Infof("info msg") }, ShouldNotPanic)
		So(func() { log.Errorf("error msg") }, ShouldNotPanic)
	})

	Convey("Without timestamps", t, func() {
		handler := log.NewTextHandler(os.Stderr, false)
		So(handler, ShouldNotBeNil)

		So(func() { log.Debugf("debug msg") }, ShouldNotPanic)
		So(func() { log.Infof("info msg") }, ShouldNotPanic)
		So(func() { log.Errorf("error msg") }, ShouldNotPanic)

		So(func() { log.FilterNonStackerLogs(handler, 1) }, ShouldNotPanic)

		So(func() { log.Debugf("debug msg") }, ShouldNotPanic)
		So(func() { log.Infof("info msg") }, ShouldNotPanic)
		So(func() { log.Errorf("error msg") }, ShouldNotPanic)
	})
}
