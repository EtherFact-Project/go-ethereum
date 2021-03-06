// Copyright 2017 The go-etherfact Authors
// This file is part of the go-etherfact library.
//
// The go-etherfact library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etherfact library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etherfact library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8\xcc\xe1\x94\x06!\x94\x83\xe2\x17\u026dG\x9fv\xa9T&\xf6\x89\xefMYQ\x8b\x19\x816X\xf1\xec\xcc\x1aR\x00\x00\xe1\x94M\xbc\xab&*\x16M\xb5\xbd\f\u0242\x8f\xa0-\xd9\xe9\xacN\u018b\x19\x816X\xf1\xec\xcc\x1aR\x00\x00\xe1\x94UJ\x13\xd1I\x81\x9f\x8d^\x18\x1b\xd3H\x82\x91E\xcb\xcb@<\x8b\x19\x816X\xf1\xec\xcc\x1aR\x00\x00\u151d\x17\x03\xc8\xcb\x02/\u01b0\x80\xf9\xc3\x16\xb4B_8B\x9e\xe0\x8b\x19\x816X\xf1\xec\xcc\x1aR\x00\x00\u153c0\xb9\x87p\xe3G\xb4\x94*y\xaf\x10J\xbf\xfa=\xa4\xf2\xbe\x8b\x19\x816X\xf1\xec\xcc\x1aR\x00\x00\xe1\x94\xce:\xb7}*\xf18\xc6_$H\xbbQo\xe1\xed\u008d\xfc\u010b\x19\x816X\xf1\xec\xcc\x1aR\x00\x00"
const testnetAllocData = "\xf8\xcc\xe1\x94\x06!\x94\x83\xe2\x17\u026dG\x9fv\xa9T&\xf6\x89\xefMYQ\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94M\xbc\xab&*\x16M\xb5\xbd\f\u0242\x8f\xa0-\xd9\xe9\xacN\u018b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94UJ\x13\xd1I\x81\x9f\x8d^\x18\x1b\xd3H\x82\x91E\xcb\xcb@<\x8b)[\xe9nd\x06ir\x00\x00\x00\u151d\x17\x03\xc8\xcb\x02/\u01b0\x80\xf9\xc3\x16\xb4B_8B\x9e\xe0\x8b)[\xe9nd\x06ir\x00\x00\x00\u153c0\xb9\x87p\xe3G\xb4\x94*y\xaf\x10J\xbf\xfa=\xa4\xf2\xbe\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94\xce:\xb7}*\xf18\xc6_$H\xbbQo\xe1\xed\u008d\xfc\u010b)[\xe9nd\x06ir\x00\x00\x00"
const rinkebyAllocData = "\xf8\xcc\xe1\x94\x06!\x94\x83\xe2\x17\u026dG\x9fv\xa9T&\xf6\x89\xefMYQ\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94M\xbc\xab&*\x16M\xb5\xbd\f\u0242\x8f\xa0-\xd9\xe9\xacN\u018b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94UJ\x13\xd1I\x81\x9f\x8d^\x18\x1b\xd3H\x82\x91E\xcb\xcb@<\x8b)[\xe9nd\x06ir\x00\x00\x00\u151d\x17\x03\xc8\xcb\x02/\u01b0\x80\xf9\xc3\x16\xb4B_8B\x9e\xe0\x8b)[\xe9nd\x06ir\x00\x00\x00\u153c0\xb9\x87p\xe3G\xb4\x94*y\xaf\x10J\xbf\xfa=\xa4\xf2\xbe\x8b)[\xe9nd\x06ir\x00\x00\x00\xe1\x94\xce:\xb7}*\xf18\xc6_$H\xbbQo\xe1\xed\u008d\xfc\u010b)[\xe9nd\x06ir\x00\x00\x00"