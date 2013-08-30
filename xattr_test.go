package xattr_test

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"."
	. "launchpad.net/gocheck"
)

func TestXattr(t *testing.T) { TestingT(t) }

type F struct {
	f    string
	attr string
}

var _ = Suite(&F{})

func (f *F) SetUpTest(c *C) {
	file, err := ioutil.TempFile("", "test_xattr_")
	c.Assert(err, IsNil)
	err = file.Close()
	c.Assert(err, IsNil)
	f.f = file.Name()
	f.attr = "test xattr"
}

func (f *F) TearDownTest(c *C) {
	if !c.Failed() {
		err := os.Remove(f.f)
		c.Assert(err, IsNil)
	}
}

func (f *F) TestFlow(c *C) {
	data := []byte("test xattr data")
	attr2 := "text xattr 2"

	attrs, err := xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{})

	err = xattr.Set(f.f, f.attr, data)
	c.Check(err, IsNil)

	attrs, err = xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{f.attr})

	err = xattr.Set(f.f, attr2, data)
	c.Check(err, IsNil)

	attrs, err = xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{f.attr, attr2})

	data1, err := xattr.Get(f.f, f.attr)
	c.Check(err, IsNil)
	c.Check(data1, DeepEquals, data)

	data1, err = xattr.Get(f.f, "test other xattr")
	if runtime.GOOS == "linux" {
		c.Check(err, ErrorMatches, "*. no data available")
	} else {
		c.Check(err, ErrorMatches, "*. attribute not found")
	}
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})
	c.Check(xattr.IsNotExist(err), Equals, true)
	c.Check(data1, IsNil)

	err = xattr.Remove(f.f, f.attr)
	c.Check(err, IsNil)

	attrs, err = xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{attr2})

	err = xattr.Remove(f.f, attr2)
	c.Check(err, IsNil)

	attrs, err = xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{})
}

func (f *F) TestEmptyAttr(c *C) {
	data := []byte{}

	err := xattr.Set(f.f, f.attr, data)
	c.Check(err, IsNil)

	attrs, err := xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{f.attr})

	data1, err := xattr.Get(f.f, f.attr)
	c.Check(err, IsNil)
	c.Check(data1, DeepEquals, data)

	err = xattr.Remove(f.f, f.attr)
	c.Check(err, IsNil)

	attrs, err = xattr.List(f.f)
	c.Check(err, IsNil)
	c.Check(attrs, DeepEquals, []string{})
}

func (f *F) TestNoFile(c *C) {
	fn := "no-such-file"
	data := []byte("test_xattr data")

	attrs, err := xattr.List(fn)
	c.Check(err, ErrorMatches, "*. no such file or directory")
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})
	c.Check(attrs, IsNil)

	err = xattr.Set(fn, f.attr, data)
	c.Check(err, ErrorMatches, "*. no such file or directory")
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})

	attrs, err = xattr.List(fn)
	c.Check(err, ErrorMatches, "*. no such file or directory")
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})
	c.Check(attrs, IsNil)

	data1, err := xattr.Get(fn, f.attr)
	c.Check(err, ErrorMatches, "*. no such file or directory")
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})
	c.Check(data1, IsNil)

	err = xattr.Remove(fn, f.attr)
	c.Check(err, ErrorMatches, "*. no such file or directory")
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})

	attrs, err = xattr.List(fn)
	c.Check(err, ErrorMatches, "*. no such file or directory")
	c.Check(err, FitsTypeOf, &xattr.XAttrError{})
	c.Check(attrs, IsNil)
}
