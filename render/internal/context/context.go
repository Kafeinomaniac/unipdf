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

package context ;import (_f "errors";_eg "github.com/golang/freetype/truetype";_bg "github.com/unidoc/unipdf/v3/core";_dd "github.com/unidoc/unipdf/v3/internal/textencoding";_e "github.com/unidoc/unipdf/v3/internal/transform";_fg "github.com/unidoc/unipdf/v3/model";_bc "golang.org/x/image/font";_d "image";_b "image/color";);type TextFont struct{Font *_fg .PdfFont ;Face _bc .Face ;Size float64 ;_dfb *_eg .Font ;_bdg *_fg .PdfFont ;};type Context interface{Push ();Pop ();Matrix ()_e .Matrix ;SetMatrix (_bb _e .Matrix );Translate (_dc ,_da float64 );Scale (_fe ,_cg float64 );Rotate (_cb float64 );MoveTo (_fce ,_ac float64 );LineTo (_g ,_be float64 );CubicTo (_fgf ,_ba ,_eb ,_gb ,_daa ,_ege float64 );QuadraticTo (_cc ,_df ,_bd ,_af float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();ResetClip ();LineWidth ()float64 ;SetLineWidth (_fb float64 );SetLineCap (_db LineCap );SetLineJoin (_dac LineJoin );SetDash (_de ...float64 );SetDashOffset (_ed float64 );Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_cga ,_ag ,_bbd ,_aa float64 );SetFillRGBA (_fgd ,_aad ,_gd ,_gg float64 );SetFillStyle (_bdc Pattern );SetFillRule (_bf FillRule );SetStrokeRGBA (_cgg ,_bdf ,_bea ,_cgc float64 );SetStrokeStyle (_ea Pattern );TextState ()*TextState ;DrawString (_ae string ,_fceb ,_ggb float64 );MeasureString (_aga string )(_fcf ,_gdd float64 );DrawRectangle (_fee ,_cf ,_fcg ,_ef float64 );DrawImage (_afd _d .Image ,_bdd ,_egd int );DrawImageAnchored (_efe _d .Image ,_bfb ,_cba int ,_afb ,_dea float64 );Height ()int ;Width ()int ;};func (_gda *TextState )ProcTStar (){_gda .ProcTd (0,-_gda .Tl )};func (_egdc *TextState )ProcTd (tx ,ty float64 ){_egdc .Tlm .Concat (_e .TranslationMatrix (tx ,-ty ));_egdc .Tm =_egdc .Tlm .Clone ();};func NewTextFont (font *_fg .PdfFont ,size float64 )(*TextFont ,error ){_dbc :=font .FontDescriptor ();if _dbc ==nil {return nil ,_f .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");};_dbf ,_cgd :=_bg .GetStream (_dbc .FontFile2 );if !_cgd {return nil ,_f .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};_dg ,_ad :=_bg .DecodeStream (_dbf );if _ad !=nil {return nil ,_ad ;};_fgdg ,_ad :=_eg .Parse (_dg );if _ad !=nil {return nil ,_ad ;};if size <=1{size =10;};return &TextFont {Font :font ,Face :_eg .NewFace (_fgdg ,&_eg .Options {Size :size }),Size :size ,_dfb :_fgdg },nil ;};func (_bcc *TextFont )CharcodesToUnicode (charcodes []_dd .CharCode )[]rune {if _bcc ._bdg !=nil {return _bcc ._bdg .CharcodesToUnicode (charcodes );};return _bcc .Font .CharcodesToUnicode (charcodes );};func (_edb *TextState )ProcQ (data []byte ,ctx Context ){_edb .ProcTStar ();_edb .ProcTj (data ,ctx )};func (_bdde *TextFont )GetRuneMetrics (r rune )(float64 ,float64 ,bool ){if _cbd ,_edc :=_bdde .Font .GetRuneMetrics (r );_edc &&_cbd .Wx !=0{return _cbd .Wx ,_cbd .Wy ,_edc ;};if _bdde ._bdg ==nil {return 0,0,false ;};_fcb ,_feb :=_bdde ._bdg .GetRuneMetrics (r );return _fcb .Wx ,_fcb .Wy ,_feb &&_fcb .Wx !=0;};func (_bga *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_bga .Tw =aw ;_bga .Tc =ac ;_bga .ProcQ (data ,ctx );};type LineCap int ;func (_gba *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_gba .Tm =_e .NewMatrix (a ,b ,c ,d ,e ,-f );_gba .Tlm =_gba .Tm .Clone ();};func (_ab *TextState )ProcTD (tx ,ty float64 ){_ab .Tl =-ty ;_ab .ProcTd (tx ,ty )};func (_add *TextState )Reset (){_add .Tm =_e .IdentityMatrix ();_add .Tlm =_e .IdentityMatrix ()};func (_age *TextState )ProcTj (data []byte ,ctx Context ){_dec :=_age .Tf .Size ;_cgb :=_age .Th /100.0;_abe :=_e .NewMatrix (_dec *_cgb ,0,0,_dec ,0,_age .Ts );_abf :=_age .Tf .CharcodesToUnicode (_age .Tf .BytesToCharcodes (data ));for _ ,_gf :=range _abf {if _gf =='\x00'{continue ;};_fceg :=_age .Tm .Clone ();_age .Tm .Concat (_abe );_cbg ,_agg :=_age .Tm .Transform (0,0);ctx .Scale (1,-1);ctx .DrawString (string (_gf ),_cbg ,_agg );ctx .Scale (1,-1);_ggd :=0.0;if _gf ==' '{_ggd =_age .Tw ;};var _dee float64 ;if _ff ,_ ,_egc :=_age .Tf .GetRuneMetrics (_gf );_egc {_dee =_ff *0.001*_dec ;}else {_dee ,_ =ctx .MeasureString (string (_gf ));};_abg :=(_dee +_age .Tc +_ggd )*_cgb ;_age .Tm =_e .TranslationMatrix (_abg ,0).Mult (_fceg );};};func NewTextState ()*TextState {return &TextState {Th :100,Tm :_e .IdentityMatrix (),Tlm :_e .IdentityMatrix ()};};type LineJoin int ;const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);type FillRule int ;func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_dae ,_egdb :=_fg .NewPdfFontFromTTFFile (filePath );if _egdb !=nil {return nil ,_egdb ;};return NewTextFont (_dae ,size );};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;);type Pattern interface{ColorAt (_bgc ,_egf int )_b .Color ;};func (_efg *TextFont )GetCharMetrics (code _dd .CharCode )(float64 ,float64 ,bool ){if _ga ,_agb :=_efg .Font .GetCharMetrics (code );_agb &&_ga .Wx !=0{return _ga .Wx ,_ga .Wy ,_agb ;};if _efg ._bdg ==nil {return 0,0,false ;};_cggd ,_adc :=_efg ._bdg .GetCharMetrics (code );return _cggd .Wx ,_cggd .Wy ,_adc &&_cggd .Wx !=0;};const (LineCapRound LineCap =iota ;LineCapButt ;LineCapSquare ;);func (_ge *TextFont )BytesToCharcodes (data []byte )[]_dd .CharCode {if _ge ._bdg !=nil {return _ge ._bdg .BytesToCharcodes (data );};return _ge .Font .BytesToCharcodes (data );};func (_bfd *TextState )Translate (tx ,ty float64 ){_bfd .Tm =_e .TranslationMatrix (tx ,ty ).Mult (_bfd .Tm );};type TextState struct{Tc float64 ;Tw float64 ;Th float64 ;Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _e .Matrix ;Tlm _e .Matrix ;};func (_gge *TextFont )WithSize (size float64 ,originalFont *_fg .PdfFont )*TextFont {if size <=1{size =10;};return &TextFont {Font :_gge .Font ,Face :_eg .NewFace (_gge ._dfb ,&_eg .Options {Size :size }),Size :size ,_dfb :_gge ._dfb ,_bdg :originalFont };};func (_fcec *TextState )ProcTf (font *TextFont ){_fcec .Tf =font };type Gradient interface{Pattern ;AddColorStop (_a float64 ,_fc _b .Color );};