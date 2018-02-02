package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionSetVal(t *testing.T) {
	opts := Options{}
	opts.Set("foo", "bar")
	assert.Equal(t, "foo=bar", opts.Encode())
}

func TestOptionSetTwoKeys(t *testing.T) {
	opts := Options{}
	opts.Set("foo", "bar")
	opts.Set("baz", "quux")
	assert.Equal(t, "baz=quux&foo=bar", opts.Encode())
}

func TestOptionSetBool(t *testing.T) {
	opts := Options{}
	opts.SetBool("foo", true)
	assert.Equal(t, "foo=true", opts.Encode())
}

func TestOptionReplaceVal(t *testing.T) {
	opts := Options{}
	opts.Set("foo", "bar")
	opts.Set("foo", "baz")
	assert.Equal(t, "foo=baz", opts.Encode())
}

func TestOptionAddVal(t *testing.T) {
	opts := Options{}
	opts.Add("foo", "bar")
	assert.Equal(t, "foo=bar", opts.Encode())
}

func TestOptionAddTwoVals(t *testing.T) {
	opts := Options{}
	opts.Add("foo", "bar")
	opts.Add("foo", "baz")
	assert.Equal(t, "foo=bar;baz", opts.Encode())
}

func TestOptionAddTwoValsAsVariadic(t *testing.T) {
	opts := Options{}
	opts.Add("foo", "bar", "baz")
	assert.Equal(t, "foo=bar;baz", opts.Encode())
}

func TestOptionSetKeyAndAddTwoVals(t *testing.T) {
	opts := Options{}
	opts.Set("foo", "bar")
	opts.Add("baz", "quux")
	opts.Add("baz", "zuko")
	assert.Equal(t, "baz=quux;zuko&foo=bar", opts.Encode())
}

func TestOptionAddTwoValsAndSetKey(t *testing.T) {
	opts := Options{}
	opts.Add("foo", "bar")
	opts.Add("foo", "baz")
	opts.Set("quux", "zuko")
	assert.Equal(t, "foo=bar;baz&quux=zuko", opts.Encode())
}

func TestOptionAddIntVal(t *testing.T) {
	opts := Options{}
	opts.AddInt("foo", 1)
	assert.Equal(t, "foo=1", opts.Encode())
}

func TestOptionAddIntValsAsVariadic(t *testing.T) {
	opts := Options{}
	opts.AddInt("foo", 1, 2)
	assert.Equal(t, "foo=1;2", opts.Encode())
}

func TestOptionsAddInt64Val(t *testing.T) {
	opts := Options{}
	opts.AddInt64("foo", int64(1))
	assert.Equal(t, "foo=1", opts.Encode())
}

func TestOptionsAddFloatVal(t *testing.T) {
	opts := Options{}
	opts.AddFloat("foo", 0.1231)
	assert.Equal(t, "foo=0.1231", opts.Encode())
}

func TestOptionsAddFloatValsAsVariadic(t *testing.T) {
	opts := Options{}
	opts.AddFloat("foo", 1.1231312, 2.1233)
	assert.Equal(t, "foo=1.1231312;2.1233", opts.Encode())
}
