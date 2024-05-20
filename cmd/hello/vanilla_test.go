package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a boolean", t, func() {
		var f bool
		Convey("When value is true", func() {
			f = true
			Convey("So result should be true", func() {
				So(f, ShouldBeTrue)
				So(f, ShouldNotEqual, false)
				So(f, ShouldEqual, true)
			})
		})
		Convey("When value is false", func() {
			f = false
			Convey("So result should be false", func() {
				So(f, ShouldBeFalse)
				So(f, ShouldNotEqual, true)
				So(f, ShouldEqual, false)
			})
		})
	})
}
