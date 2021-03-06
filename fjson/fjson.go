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

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_e "encoding/json";_da "github.com/unidoc/unipdf/v3/core";_ad "github.com/unidoc/unipdf/v3/model";_a "io";_dc "os";);

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_f ,_fb :=_dc .Open (filePath );if _fb !=nil {return nil ,_fb ;};defer _f .Close ();return LoadFromJSON (_f );};

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_ab ,_dbf :=_dc .Open (filePath );if _dbf !=nil {return nil ,_dbf ;};defer _ab .Close ();return LoadFromPDF (_ab );};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _a .Reader )(*FieldData ,error ){var _b FieldData ;_c :=_e .NewDecoder (r ).Decode (&_b ._af );if _c !=nil {return nil ,_c ;};return &_b ,nil ;};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_af []fieldValue };type fieldValue struct{Name string `json:"name"`;Value string `json:"value"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};

// FieldValues implements model.FieldValueProvider interface.
func (_gb *FieldData )FieldValues ()(map[string ]_da .PdfObject ,error ){_abb :=make (map[string ]_da .PdfObject );for _ ,_dec :=range _gb ._af {if len (_dec .Value )> 0{_abb [_dec .Name ]=_da .MakeString (_dec .Value );};};return _abb ,nil ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _a .ReadSeeker )(*FieldData ,error ){_ec ,_bf :=_ad .NewPdfReader (rs );if _bf !=nil {return nil ,_bf ;};if _ec .AcroForm ==nil {return nil ,nil ;};var _ca []fieldValue ;_fba :=_ec .AcroForm .AllFields ();for _ ,_cg :=range _fba {var _afa []string ;_fc :=make (map[string ]struct{});_g ,_ee :=_cg .FullName ();if _ee !=nil {return nil ,_ee ;};if _de ,_adg :=_cg .V .(*_da .PdfObjectString );_adg {_ca =append (_ca ,fieldValue {Name :_g ,Value :_de .Decoded ()});continue ;};var _fd string ;for _ ,_ef :=range _cg .Annotations {_bd ,_fg :=_da .GetName (_ef .AS );if _fg {_fd =_bd .String ();};_ed ,_bg :=_da .GetDict (_ef .AP );if !_bg {continue ;};_adf ,_ :=_da .GetDict (_ed .Get ("\u004e"));for _ ,_bb :=range _adf .Keys (){_caa :=_bb .String ();if _ ,_dga :=_fc [_caa ];!_dga {_afa =append (_afa ,_caa );_fc [_caa ]=struct{}{};};};_fdc ,_ :=_da .GetDict (_ed .Get ("\u0044"));for _ ,_ga :=range _fdc .Keys (){_eg :=_ga .String ();if _ ,_ae :=_fc [_eg ];!_ae {_afa =append (_afa ,_eg );_fc [_eg ]=struct{}{};};};};_gg :=fieldValue {Name :_g ,Value :_fd ,Options :_afa };_ca =append (_ca ,_gg );};_ac :=FieldData {_af :_ca };return &_ac ,nil ;};

// JSON returns the field data as a string in JSON format.
func (_aa FieldData )JSON ()(string ,error ){_ege ,_be :=_e .MarshalIndent (_aa ._af ,"","\u0020\u0020\u0020\u0020");return string (_ege ),_be ;};