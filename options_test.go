package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionBlank(t *testing.T) {
	var opts options
	assert.Equal(t, "", opts.encode())
}

func TestOptionSetVal(t *testing.T) {
	opts := options{}
	opts.set("foo", "bar")
	assert.Equal(t, "foo=bar", opts.encode())
}

func TestOptionSetTwoKeys(t *testing.T) {
	opts := options{}
	opts.set("foo", "bar")
	opts.set("baz", "quux")
	assert.Equal(t, "baz=quux&foo=bar", opts.encode())
}

func TestOptionSetBool(t *testing.T) {
	opts := options{}
	opts.setBool("foo", true)
	assert.Equal(t, "foo=true", opts.encode())
}

func TestOptionReplaceVal(t *testing.T) {
	opts := options{}
	opts.set("foo", "bar")
	opts.set("foo", "baz")
	assert.Equal(t, "foo=baz", opts.encode())
}

func TestOptionAddVal(t *testing.T) {
	opts := options{}
	opts.add("foo", "bar")
	assert.Equal(t, "foo=bar", opts.encode())
}

func TestOptionAddTwoVals(t *testing.T) {
	opts := options{}
	opts.add("foo", "bar")
	opts.add("foo", "baz")
	assert.Equal(t, "foo=bar;baz", opts.encode())
}

func TestOptionAddTwoValsAsVariadic(t *testing.T) {
	opts := options{}
	opts.add("foo", "bar", "baz")
	assert.Equal(t, "foo=bar;baz", opts.encode())
}

func TestOptionSetKeyAndAddTwoVals(t *testing.T) {
	opts := options{}
	opts.set("foo", "bar")
	opts.add("baz", "quux")
	opts.add("baz", "zuko")
	assert.Equal(t, "baz=quux;zuko&foo=bar", opts.encode())
}

func TestOptionAddTwoValsAndSetKey(t *testing.T) {
	opts := options{}
	opts.add("foo", "bar")
	opts.add("foo", "baz")
	opts.set("quux", "zuko")
	assert.Equal(t, "foo=bar;baz&quux=zuko", opts.encode())
}

func TestOptionAddIntVal(t *testing.T) {
	opts := options{}
	opts.addInt("foo", 1)
	assert.Equal(t, "foo=1", opts.encode())
}

func TestOptionAddIntValsAsVariadic(t *testing.T) {
	opts := options{}
	opts.addInt("foo", 1, 2)
	assert.Equal(t, "foo=1;2", opts.encode())
}

func TestOptionsAddInt64Val(t *testing.T) {
	opts := options{}
	opts.addInt64("foo", int64(1))
	assert.Equal(t, "foo=1", opts.encode())
}

func TestOptionsAddFloatVal(t *testing.T) {
	opts := options{}
	opts.addFloat("foo", 0.1231)
	assert.Equal(t, "foo=0.1231", opts.encode())
}

func TestOptionsAddFloatValsAsVariadic(t *testing.T) {
	opts := options{}
	opts.addFloat("foo", 1.1231312, 2.1233)
	assert.Equal(t, "foo=1.1231312;2.1233", opts.encode())
}
