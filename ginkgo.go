package ginkgo

import (
	"github.com/onsi/ginkgo/config"

	"testing"
	"time"
)

const GINKGO_VERSION = config.VERSION

const defaultTimeout = 1

var globalSuite *suite

func init() {
	config.Flags("ginkgo", true)
	globalSuite = newSuite()
}

func RunSpecs(t *testing.T, description string) {
	RunSpecsWithCustomReporter(t, description, []Reporter{newDefaultReporter(config.DefaultReporterConfig)})
}

func RunSpecsWithCustomReporter(t *testing.T, description string, reporters []Reporter) {
	globalSuite.run(t, description, reporters, config.GinkgoConfig)
}

type Done chan<- interface{}

func Fail(message string, callerSkip ...int) {
	skip := 0
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}
	globalSuite.fail(message, skip)
}

//Describes

func Describe(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypeNone, generateCodeLocation(1))
	return true
}

func FDescribe(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypeFocused, generateCodeLocation(1))
	return true
}

func PDescribe(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypePending, generateCodeLocation(1))
	return true
}

func XDescribe(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypePending, generateCodeLocation(1))
	return true
}

//Context

func Context(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypeNone, generateCodeLocation(1))
	return true
}

func FContext(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypeFocused, generateCodeLocation(1))
	return true
}

func PContext(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypePending, generateCodeLocation(1))
	return true
}

func XContext(text string, body func()) bool {
	globalSuite.pushContainerNode(text, body, flagTypePending, generateCodeLocation(1))
	return true
}

//It

func It(text string, body interface{}, timeout ...float64) bool {
	globalSuite.pushItNode(text, body, flagTypeNone, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

func FIt(text string, body interface{}, timeout ...float64) bool {
	globalSuite.pushItNode(text, body, flagTypeFocused, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

func PIt(text string, body interface{}, timeout ...float64) bool {
	globalSuite.pushItNode(text, body, flagTypePending, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

func XIt(text string, body interface{}, timeout ...float64) bool {
	globalSuite.pushItNode(text, body, flagTypePending, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

//Benchmark

func Measure(text string, body func(Benchmarker), samples int) bool {
	globalSuite.pushMeasureNode(text, body, flagTypeNone, generateCodeLocation(1), samples)
	return true
}

func FMeasure(text string, body func(Benchmarker), samples int) bool {
	globalSuite.pushMeasureNode(text, body, flagTypeFocused, generateCodeLocation(1), samples)
	return true
}

func PMeasure(text string, body func(Benchmarker), samples int) bool {
	globalSuite.pushMeasureNode(text, body, flagTypePending, generateCodeLocation(1), samples)
	return true
}

func XMeasure(text string, body func(Benchmarker), samples int) bool {
	globalSuite.pushMeasureNode(text, body, flagTypePending, generateCodeLocation(1), samples)
	return true
}

//Before, JustBefore, and After

func BeforeEach(body interface{}, timeout ...float64) bool {
	globalSuite.pushBeforeEachNode(body, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

func JustBeforeEach(body interface{}, timeout ...float64) bool {
	globalSuite.pushJustBeforeEachNode(body, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

func AfterEach(body interface{}, timeout ...float64) bool {
	globalSuite.pushAfterEachNode(body, generateCodeLocation(1), parseTimeout(timeout...))
	return true
}

func parseTimeout(timeout ...float64) time.Duration {
	if len(timeout) == 0 {
		return time.Duration(defaultTimeout * int64(time.Second))
	} else {
		return time.Duration(timeout[0] * float64(time.Second))
	}
}
