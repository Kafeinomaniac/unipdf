//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package endian ;import (_a "encoding/binary";_ae "unsafe";);func IsLittle ()bool {return !_e };func init (){const _aa =int (_ae .Sizeof (0));_ee :=1;_be :=(*[_aa ]byte )(_ae .Pointer (&_ee ));if _be [0]==0{_e =true ;ByteOrder =_a .BigEndian ;}else {ByteOrder =_a .LittleEndian ;};};func IsBig ()bool {return _e };var (ByteOrder _a .ByteOrder ;_e bool ;);