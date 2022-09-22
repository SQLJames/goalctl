// Code generated by 'ccgo stdlib\gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o stdlib\stdlib_windows_386.go -pkgname stdlib', DO NOT EDIT.

package stdlib

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	CHAR_BIT                                        = 8
	CHAR_MAX                                        = 127
	CHAR_MIN                                        = -128
	DUMMYSTRUCTNAME                                 = 0
	DUMMYSTRUCTNAME1                                = 0
	DUMMYSTRUCTNAME2                                = 0
	DUMMYSTRUCTNAME3                                = 0
	DUMMYSTRUCTNAME4                                = 0
	DUMMYSTRUCTNAME5                                = 0
	DUMMYUNIONNAME                                  = 0
	DUMMYUNIONNAME1                                 = 0
	DUMMYUNIONNAME2                                 = 0
	DUMMYUNIONNAME3                                 = 0
	DUMMYUNIONNAME4                                 = 0
	DUMMYUNIONNAME5                                 = 0
	DUMMYUNIONNAME6                                 = 0
	DUMMYUNIONNAME7                                 = 0
	DUMMYUNIONNAME8                                 = 0
	DUMMYUNIONNAME9                                 = 0
	EXIT_FAILURE                                    = 1
	EXIT_SUCCESS                                    = 0
	INT_MAX                                         = 2147483647
	INT_MIN                                         = -2147483648
	LLONG_MAX                                       = 9223372036854775807
	LLONG_MIN                                       = -9223372036854775808
	LONG_LONG_MAX                                   = 9223372036854775807
	LONG_LONG_MIN                                   = -9223372036854775808
	LONG_MAX                                        = 2147483647
	LONG_MIN                                        = -2147483648
	MB_LEN_MAX                                      = 5
	MINGW_DDK_H                                     = 0
	MINGW_DDRAW_VERSION                             = 7
	MINGW_HAS_DDK_H                                 = 1
	MINGW_HAS_DDRAW_H                               = 1
	MINGW_HAS_SECURE_API                            = 1
	MINGW_SDK_INIT                                  = 0
	PATH_MAX                                        = 260
	RAND_MAX                                        = 0x7fff
	SCHAR_MAX                                       = 127
	SCHAR_MIN                                       = -128
	SHRT_MAX                                        = 32767
	SHRT_MIN                                        = -32768
	SIZE_MAX                                        = 4294967295
	SSIZE_MAX                                       = 2147483647
	UCHAR_MAX                                       = 255
	UINT_MAX                                        = 4294967295
	ULLONG_MAX                                      = 18446744073709551615
	ULONG_LONG_MAX                                  = 18446744073709551615
	ULONG_MAX                                       = 4294967295
	UNALIGNED                                       = 0
	USE___UUIDOF                                    = 0
	USHRT_MAX                                       = 65535
	WIN32                                           = 1
	WINNT                                           = 1
	X_AGLOBAL                                       = 0
	X_ALLOCA_S_HEAP_MARKER                          = 0xDDDD
	X_ALLOCA_S_MARKER_SIZE                          = 8
	X_ALLOCA_S_STACK_MARKER                         = 0xCCCC
	X_ALLOCA_S_THRESHOLD                            = 1024
	X_ANONYMOUS_STRUCT                              = 0
	X_ANONYMOUS_UNION                               = 0
	X_ARGMAX                                        = 100
	X_CALL_REPORTFAULT                              = 0x2
	X_CONST_RETURN                                  = 0
	X_CRTNOALIAS                                    = 0
	X_CRTRESTRICT                                   = 0
	X_CRT_ABS_DEFINED                               = 0
	X_CRT_ALGO_DEFINED                              = 0
	X_CRT_ALLOCATION_DEFINED                        = 0
	X_CRT_ALTERNATIVE_IMPORTED                      = 0
	X_CRT_ATOF_DEFINED                              = 0
	X_CRT_DOUBLE_DEC                                = 0
	X_CRT_ERRNO_DEFINED                             = 0
	X_CRT_MANAGED_HEAP_DEPRECATE                    = 0
	X_CRT_PACKING                                   = 8
	X_CRT_PERROR_DEFINED                            = 0
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES          = 0
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES_MEMORY   = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES        = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT  = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_MEMORY = 0
	X_CRT_SWAB_DEFINED                              = 0
	X_CRT_SYSTEM_DEFINED                            = 0
	X_CRT_TERMINATE_DEFINED                         = 0
	X_CRT_WPERROR_DEFINED                           = 0
	X_CRT_WSYSTEM_DEFINED                           = 0
	X_CVTBUFSIZE                                    = 349
	X_DIV_T_DEFINED                                 = 0
	X_DLL                                           = 0
	X_ERRCODE_DEFINED                               = 0
	X_FILE_OFFSET_BITS                              = 64
	X_FREEA_INLINE                                  = 0
	X_FREEENTRY                                     = 0
	X_GCC_LIMITS_H_                                 = 0
	X_HEAPBADBEGIN                                  = -3
	X_HEAPBADNODE                                   = -4
	X_HEAPBADPTR                                    = -6
	X_HEAPEMPTY                                     = -1
	X_HEAPEND                                       = -5
	X_HEAPINFO_DEFINED                              = 0
	X_HEAPOK                                        = -2
	X_HEAP_MAXREQ                                   = 0xFFFFFFE0
	X_I16_MAX                                       = 32767
	X_I16_MIN                                       = -32768
	X_I32_MAX                                       = 2147483647
	X_I32_MIN                                       = -2147483648
	X_I64_MAX                                       = 9223372036854775807
	X_I64_MIN                                       = -9223372036854775808
	X_I8_MAX                                        = 127
	X_I8_MIN                                        = -128
	X_INC_CRTDEFS                                   = 0
	X_INC_CRTDEFS_MACRO                             = 0
	X_INC_LIMITS                                    = 0
	X_INC_MINGW_SECAPI                              = 0
	X_INC_STDLIB                                    = 0
	X_INC_STDLIB_S                                  = 0
	X_INC_VADEFS                                    = 0
	X_INC__MINGW_H                                  = 0
	X_INT128_DEFINED                                = 0
	X_INTEGRAL_MAX_BITS                             = 64
	X_INTPTR_T_DEFINED                              = 0
	X_LIMITS_H___                                   = 0
	X_MALLOC_H_                                     = 0
	X_MAX_DIR                                       = 256
	X_MAX_DRIVE                                     = 3
	X_MAX_ENV                                       = 32767
	X_MAX_EXT                                       = 256
	X_MAX_FNAME                                     = 256
	X_MAX_PATH                                      = 260
	X_MAX_WAIT_MALLOC_CRT                           = 60000
	X_MM_MALLOC_H_INCLUDED                          = 0
	X_MT                                            = 0
	X_M_IX86                                        = 600
	X_ONEXIT_T_DEFINED                              = 0
	X_OUT_TO_DEFAULT                                = 0
	X_OUT_TO_MSGBOX                                 = 2
	X_OUT_TO_STDERR                                 = 1
	X_PGLOBAL                                       = 0
	X_PTRDIFF_T_                                    = 0
	X_PTRDIFF_T_DEFINED                             = 0
	X_QSORT_S_DEFINED                               = 0
	X_REENTRANT                                     = 1
	X_REPORT_ERRMODE                                = 3
	X_RSIZE_T_DEFINED                               = 0
	X_SECURECRT_FILL_BUFFER_PATTERN                 = 0xFD
	X_SIZE_T_DEFINED                                = 0
	X_SSIZE_T_DEFINED                               = 0
	X_TAGLC_ID_DEFINED                              = 0
	X_THREADLOCALEINFO                              = 0
	X_TIME32_T_DEFINED                              = 0
	X_TIME64_T_DEFINED                              = 0
	X_TIME_T_DEFINED                                = 0
	X_UI16_MAX                                      = 0xffff
	X_UI32_MAX                                      = 0xffffffff
	X_UI64_MAX                                      = 0xffffffffffffffff
	X_UI8_MAX                                       = 0xff
	X_UINTPTR_T_DEFINED                             = 0
	X_USEDENTRY                                     = 1
	X_USE_32BIT_TIME_T                              = 0
	X_VA_LIST_DEFINED                               = 0
	X_W64                                           = 0
	X_WCHAR_T_DEFINED                               = 0
	X_WCTYPE_T_DEFINED                              = 0
	X_WIN32                                         = 1
	X_WIN32_WINNT                                   = 0x502
	X_WINT_T                                        = 0
	X_WRITE_ABORT_MSG                               = 0x1
	X_WSTDLIBP_DEFINED                              = 0
	X_WSTDLIBP_S_DEFINED                            = 0
	X_WSTDLIB_DEFINED                               = 0
	X_WSTDLIB_S_DEFINED                             = 0
	X_X86_                                          = 1
	I386                                            = 1
)

type Ptrdiff_t = int32 /* <builtin>:3:26 */

type Size_t = uint32 /* <builtin>:9:23 */

type Wchar_t = uint16 /* <builtin>:15:24 */

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

type Va_list = X__builtin_va_list /* <builtin>:50:27 */

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// This macro holds an monotonic increasing value, which indicates
//    a specific fix/patch is present on trunk.  This value isn't related to
//    minor/major version-macros.  It is increased on demand, if a big
//    fix was applied to trunk.  This macro gets just increased on trunk.  For
//    other branches its value won't be modified.

// mingw.org's version macros: these make gcc to define
//    MINGW32_SUPPORTS_MT_EH and to use the _CRT_MT global
//    and the __mingwthr_key_dtor() function from the MinGW
//    CRT in its private gthr-win32.h header.

// Set VC specific compiler target macros.

// For x86 we have always to prefix by underscore.

// Special case nameless struct/union.

// MinGW-w64 has some additional C99 printf/scanf feature support.
//    So we add some helper macros to ease recognition of them.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// http://msdn.microsoft.com/en-us/library/ms175759%28v=VS.100%29.aspx
// Templates won't work in C, will break if secure API is not enabled, disabled

// https://blogs.msdn.com/b/sdl/archive/2010/02/16/vc-2010-and-memcpy.aspx?Redirected=true
// fallback on default implementation if we can't know the size of the destination

// Include _cygwin.h if we're building a Cygwin application.

// Target specific macro replacement for type "long".  In the Windows API,
//    the type long is always 32 bit, even if the target is 64 bit (LLP64).
//    On 64 bit Cygwin, the type long is 64 bit (LP64).  So, to get the right
//    sized definitions and declarations, all usage of type long in the Windows
//    headers have to be replaced by the below defined macro __LONG32.

// C/C++ specific language defines.

// Note the extern. This is needed to work around GCC's
// limitations in handling dllimport attribute.

// Attribute `nonnull' was valid as of gcc 3.3.  We don't use GCC's
//    variadiac macro facility, because variadic macros cause syntax
//    errors with  --traditional-cpp.

//  High byte is the major version, low byte is the minor.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

type X__gnuc_va_list = X__builtin_va_list /* vadefs.h:24:29 */

type Ssize_t = int32 /* crtdefs.h:47:13 */

type Rsize_t = Size_t /* crtdefs.h:52:16 */

type Intptr_t = int32 /* crtdefs.h:64:13 */

type Uintptr_t = uint32 /* crtdefs.h:77:22 */

type Wint_t = uint16   /* crtdefs.h:106:24 */
type Wctype_t = uint16 /* crtdefs.h:107:24 */

type Errno_t = int32 /* crtdefs.h:113:13 */

type X__time32_t = int32 /* crtdefs.h:118:14 */

type X__time64_t = int64 /* crtdefs.h:123:35 */

type Time_t = X__time32_t /* crtdefs.h:136:20 */

type Threadlocaleinfostruct = struct {
	Frefcount      int32
	Flc_codepage   uint32
	Flc_collate_cp uint32
	Flc_handle     [6]uint32
	Flc_id         [6]LC_ID
	Flc_category   [6]struct {
		Flocale    uintptr
		Fwlocale   uintptr
		Frefcount  uintptr
		Fwrefcount uintptr
	}
	Flc_clike            int32
	Fmb_cur_max          int32
	Flconv_intl_refcount uintptr
	Flconv_num_refcount  uintptr
	Flconv_mon_refcount  uintptr
	Flconv               uintptr
	Fctype1_refcount     uintptr
	Fctype1              uintptr
	Fpctype              uintptr
	Fpclmap              uintptr
	Fpcumap              uintptr
	Flc_time_curr        uintptr
} /* crtdefs.h:422:1 */

type Pthreadlocinfo = uintptr /* crtdefs.h:424:39 */
type Pthreadmbcinfo = uintptr /* crtdefs.h:425:36 */

type Localeinfo_struct = struct {
	Flocinfo Pthreadlocinfo
	Fmbcinfo Pthreadmbcinfo
} /* crtdefs.h:428:9 */

type X_locale_tstruct = Localeinfo_struct /* crtdefs.h:431:3 */
type X_locale_t = uintptr                 /* crtdefs.h:431:19 */

type TagLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
} /* crtdefs.h:422:1 */

type LC_ID = TagLC_ID  /* crtdefs.h:439:3 */
type LPLC_ID = uintptr /* crtdefs.h:439:9 */

type Threadlocinfo = Threadlocaleinfostruct /* crtdefs.h:468:3 */

// Copyright (C) 1992-2018 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free
// Software Foundation; either version 3, or (at your option) any later
// version.
//
// GCC is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
// for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// This administrivia gets added to the beginning of limits.h
//    if the system has its own version of limits.h.

// We use _GCC_LIMITS_H_ because we want this not to match
//    any macros that the system's limits.h uses for its own purposes.

// Use "..." so that we find syslimits.h only in this same directory.
// syslimits.h stands for the system's own limits.h file.
//    If we can use it ok unmodified, then we install this text.
//    If fixincludes fixes it, then the fixed version is installed
//    instead of this text.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.
// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// File system limits
//
// NOTE: Apparently the actual size of PATH_MAX is 260, but a space is
//       required for the NUL. TODO: Test?
// NOTE: PATH_MAX is the POSIX equivalent for Microsoft's MAX_PATH; the two
//       are semantically identical, with a limit of 259 characters for the
//       path name, plus one for a terminating NUL, for a total of 260.

// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free
// Software Foundation; either version 3, or (at your option) any later
// version.
//
// GCC is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
// for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// Number of bits in a `char'.

// Maximum length of a multibyte character.

// Minimum and maximum values a `signed char' can hold.

// Maximum value an `unsigned char' can hold.  (Minimum is 0).

// Minimum and maximum values a `char' can hold.

// Minimum and maximum values a `signed short int' can hold.

// Maximum value an `unsigned short int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed int' can hold.

// Maximum value an `unsigned int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed long int' can hold.
//    (Same as `int').

// Maximum value an `unsigned long int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed long long int' can hold.

// Maximum value an `unsigned long long int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed long long int' can hold.

// Maximum value an `unsigned long long int' can hold.  (Minimum is 0).

// This administrivia gets added to the end of limits.h
//    if the system has its own version of limits.h.

type X_onexit_t = uintptr /* stdlib.h:49:15 */

type X_div_t = struct {
	Fquot int32
	Frem  int32
} /* stdlib.h:59:11 */

type Div_t = X_div_t /* stdlib.h:62:5 */

type X_ldiv_t = struct {
	Fquot int32
	Frem  int32
} /* stdlib.h:64:11 */

type Ldiv_t = X_ldiv_t /* stdlib.h:67:5 */

type X_LDOUBLE = struct{ Fld [10]uint8 } /* stdlib.h:76:5 */

type X_CRT_DOUBLE = struct{ Fx float64 } /* stdlib.h:83:5 */

type X_CRT_FLOAT = struct{ Ff float32 } /* stdlib.h:87:5 */

type X_LONGDOUBLE = struct{ Fx float64 } /* stdlib.h:94:5 */

type X_LDBL12 = struct{ Fld12 [12]uint8 } /* stdlib.h:101:5 */

type X_purecall_handler = uintptr /* stdlib.h:142:16 */

type X_invalid_parameter_handler = uintptr /* stdlib.h:147:16 */

type Lldiv_t = struct {
	Fquot int64
	Frem  int64
} /* stdlib.h:699:61 */

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// Return codes for _heapwalk()

// Values for _heapinfo.useflag

// The structure used to walk through the heap with _heapwalk.
type X_heapinfo = struct {
	F_pentry  uintptr
	F_size    Size_t
	F_useflag int32
} /* malloc.h:46:11 */

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// Return codes for _heapwalk()

// Values for _heapinfo.useflag

// The structure used to walk through the heap with _heapwalk.
type X_HEAPINFO = X_heapinfo /* malloc.h:50:5 */

var _ int8 /* gen.c:2:13: */
