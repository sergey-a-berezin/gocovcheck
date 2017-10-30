// Copyright 2016 sergey@sergeyberezin.com

package coverage

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	t.Parallel()
	Convey("parseLine works", t, func() {
		Convey("for mode: set", func() {
			total, covered, err := parseLine("mode: set")
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 0)
			So(covered, ShouldEqual, 0)
		})

		Convey("for mode: unknown", func() {
			_, _, err := parseLine("mode: unknown")
			So(err, ShouldNotBeNil)
		})

		Convey("for invalid string", func() {
			_, _, err := parseLine("totallyInvalid")
			So(err, ShouldNotBeNil)
		})

		Convey("for a valid coverage entry", func() {
			total, covered, err := parseLine("file:1.12,9.25 100 80")
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 100)
			So(covered, ShouldEqual, 100)
		})

		Convey("for an invalid coverage entry (total is bad)", func() {
			_, _, err := parseLine("file:1.12,9.25 bad 80")
			So(err, ShouldNotBeNil)
		})

		Convey("for an invalid coverage entry (covered is bad)", func() {
			_, _, err := parseLine("file:1.12,9.25 100 bad")
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Extract works", t, func() {
		Convey("For correct input", func() {
			total, covered, err := Extract(strings.NewReader(
				"mode: set\n" +
					"file 20 1\n" +
					"file 30 0\n"))
			So(err, ShouldBeNil)
			So(total, ShouldEqual, 50)
			So(covered, ShouldEqual, 20)
		})

		Convey("For malformed input", func() {
			_, _, err := Extract(strings.NewReader(
				"mode: set\n" +
					"bad-robot\n" +
					"file 30 15\n"))
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Percentage works", t, func() {
		Convey("for normal values", func() {
			So(Percentage(100, 50), ShouldEqual, 50.0)
		})

		Convey("for zero total", func() {
			So(Percentage(0, 50), ShouldEqual, 0.0)
		})
	})
}
