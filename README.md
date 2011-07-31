Introduction
============

This package provides a simple interface to user extended attributes on Linux and OSX. Support for xattrs is filesystem dependant, so not a given even if you are running one of those operating systems.

Linux notes
-----------

Support for extended attributes is available on ext3/ext4 but generally not enabled by default on most distributions. Make sure that you add the 

    user_xattr

flag to /etc/fstab for the filesystem you want to use.

Installation
============

    goinstall github.com/davecheney/xattr


Documentation
=============

    godoc github.com/davecheney/xattr


Usage
=====

A example program is provided with the source. The simplest way to compile and install it is

    make -C $GOROOT/src/pkg/github.com/davecheney/xattr/example clean install

This will install it to your $GOBIN directory. If you have trouble running this example, make sure there isn't another xattr somewhere higher in your $PATH.

Before you start
----------------

All extended attributes need a file to be associated with. In this example I'm going to create an empty file in my home directory (see notes in the installation section)

    touch ~/testfile

Setting an attribute
--------------------

    % $GOBIN/xattr -w username dave ~/testfile

Listing known attributes
--------------------
  
    % $GOBIN/xattr ~/testfile
    username

Printing attribute values
-------------------------
 
    % $GOBIN/xattr -p username ~/testfile
    dave

Listing names and values
------------------------

     % $GOBIN/xattr -l ~/testfile
     username: dave



