// Code generated by 'ccgo unistd\gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o unistd\unistd_windows_arm64.go -pkgname unistd', DO NOT EDIT.

package unistd

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
	FTRUNCATE_DEFINED                               = 0
	F_OK                                            = 0
	MINGW_DDK_H                                     = 0
	MINGW_HAS_DDK_H                                 = 1
	MINGW_HAS_SECURE_API                            = 1
	MINGW_SDK_INIT                                  = 0
	OLD_P_OVERLAY                                   = 2
	P_DETACH                                        = 4
	P_NOWAIT                                        = 1
	P_NOWAITO                                       = 3
	P_OVERLAY                                       = 2
	P_WAIT                                          = 0
	R_OK                                            = 4
	SEEK_CUR                                        = 1
	SEEK_END                                        = 2
	SEEK_SET                                        = 0
	STDERR_FILENO                                   = 2
	STDIN_FILENO                                    = 0
	STDOUT_FILENO                                   = 1
	UNALIGNED                                       = 0
	USE___UUIDOF                                    = 0
	WAIT_CHILD                                      = 0
	WAIT_GRANDCHILD                                 = 1
	WIN32                                           = 1
	WIN64                                           = 1
	WINNT                                           = 1
	WIN_PTHREADS_UNISTD_H                           = 0
	W_OK                                            = 2
	X_OK                                            = 1
	X_AGLOBAL                                       = 0
	X_ANONYMOUS_STRUCT                              = 0
	X_ANONYMOUS_UNION                               = 0
	X_ARGMAX                                        = 100
	X_ARM64_                                        = 1
	X_A_ARCH                                        = 0x20
	X_A_HIDDEN                                      = 0x02
	X_A_NORMAL                                      = 0x00
	X_A_RDONLY                                      = 0x01
	X_A_SUBDIR                                      = 0x10
	X_A_SYSTEM                                      = 0x04
	X_CONST_RETURN                                  = 0
	X_CRTNOALIAS                                    = 0
	X_CRTRESTRICT                                   = 0
	X_CRT_ALTERNATIVE_IMPORTED                      = 0
	X_CRT_DIRECTORY_DEFINED                         = 0
	X_CRT_GETPID_DEFINED                            = 0
	X_CRT_MANAGED_HEAP_DEPRECATE                    = 0
	X_CRT_MEMORY_DEFINED                            = 0
	X_CRT_PACKING                                   = 8
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES          = 0
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES_MEMORY   = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES        = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT  = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_MEMORY = 0
	X_CRT_SWAB_DEFINED                              = 0
	X_CRT_SYSTEM_DEFINED                            = 0
	X_CRT_TERMINATE_DEFINED                         = 0
	X_CRT_USE_WINAPI_FAMILY_DESKTOP_APP             = 0
	X_CRT_WSYSTEM_DEFINED                           = 0
	X_DEV_T_DEFINED                                 = 0
	X_DLL                                           = 0
	X_ERRCODE_DEFINED                               = 0
	X_FILE_OFFSET_BITS                              = 64
	X_FILE_OFFSET_BITS_SET_FTRUNCATE                = 0
	X_FILE_OFFSET_BITS_SET_LSEEK                    = 0
	X_FILE_OFFSET_BITS_SET_OFFT                     = 0
	X_FINDDATA_T_DEFINED                            = 0
	X_FSIZE_T_DEFINED                               = 0
	X_INC_CORECRT                                   = 0
	X_INC_CORECRT_STARTUP                           = 0
	X_INC_CRTDEFS                                   = 0
	X_INC_CRTDEFS_MACRO                             = 0
	X_INC_MINGW_SECAPI                              = 0
	X_INC_PROCESS                                   = 0
	X_INC_STRING                                    = 0
	X_INC_STRING_S                                  = 0
	X_INC_TYPES                                     = 0
	X_INC_VADEFS                                    = 0
	X_INC__MINGW_H                                  = 0
	X_INO_T_DEFINED                                 = 0
	X_INT128_DEFINED                                = 0
	X_INTPTR_T_DEFINED                              = 0
	X_IO_H_                                         = 0
	X_MODE_T_                                       = 0
	X_MT                                            = 0
	X_M_ARM64                                       = 1
	X_NLSCMPERROR                                   = 2147483647
	X_NLSCMP_DEFINED                                = 0
	X_OFF64_T_DEFINED                               = 0
	X_OFF_T_                                        = 0
	X_OFF_T_DEFINED                                 = 0
	X_OLD_P_OVERLAY                                 = 2
	X_PGLOBAL                                       = 0
	X_PID_T_                                        = 0
	X_POSIX_BARRIERS                                = 200112
	X_POSIX_CLOCK_SELECTION                         = 200112
	X_POSIX_READER_WRITER_LOCKS                     = 200112
	X_POSIX_SEMAPHORES                              = 200112
	X_POSIX_SPIN_LOCKS                              = 200112
	X_POSIX_THREADS                                 = 200112
	X_POSIX_TIMEOUTS                                = 200112
	X_PTRDIFF_T_                                    = 0
	X_PTRDIFF_T_DEFINED                             = 0
	X_P_DETACH                                      = 4
	X_P_NOWAIT                                      = 1
	X_P_NOWAITO                                     = 3
	X_P_OVERLAY                                     = 2
	X_P_WAIT                                        = 0
	X_RSIZE_T_DEFINED                               = 0
	X_SECURECRT_FILL_BUFFER_PATTERN                 = 0xFD
	X_SIGSET_T_                                     = 0
	X_SIZE_T_DEFINED                                = 0
	X_SPAWNV_DEFINED                                = 0
	X_SSIZE_T_DEFINED                               = 0
	X_TAGLC_ID_DEFINED                              = 0
	X_THREADLOCALEINFO                              = 0
	X_TIME32_T_DEFINED                              = 0
	X_TIME64_T_DEFINED                              = 0
	X_TIMESPEC_DEFINED                              = 0
	X_TIME_T_DEFINED                                = 0
	X_UCRT                                          = 0
	X_UINTPTR_T_DEFINED                             = 0
	X_UNISTD_H                                      = 0
	X_VA_LIST_DEFINED                               = 0
	X_W64                                           = 0
	X_WAIT_CHILD                                    = 0
	X_WAIT_GRANDCHILD                               = 1
	X_WCHAR_T_DEFINED                               = 0
	X_WCTYPE_T_DEFINED                              = 0
	X_WConst_return                                 = 0
	X_WEXEC_DEFINED                                 = 0
	X_WFINDDATA_T_DEFINED                           = 0
	X_WIN32                                         = 1
	X_WIN32_WINNT                                   = 0x601
	X_WIN64                                         = 1
	X_WINT_T                                        = 0
	X_WIO_DEFINED                                   = 0
	X_WSPAWN_DEFINED                                = 0
	X_WSTRING_DEFINED                               = 0
	X_WSTRING_S_DEFINED                             = 0
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = uint16 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

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

// MS does not prefix symbols by underscores for 64-bit.
// As we have to support older gcc version, which are using underscores
//       as symbol prefix for x64, we have to check here for the user label
//       prefix defined by gcc.

// Special case nameless struct/union.

// MinGW-w64 has some additional C99 printf/scanf feature support.
//    So we add some helper macros to ease recognition of them.

// If _FORTIFY_SOURCE is enabled, some inline functions may use
//    __builtin_va_arg_pack().  GCC may report an error if the address
//    of such a function is used.  Set _FORTIFY_VA_ARG=0 in this case.

// Enable workaround for ABI incompatibility on affected platforms

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

// Allow both 0x1400 and 0xE00 to identify UCRT

// ===-------- vadefs.h ---------------------------------------------------===
//
// Part of the LLVM Project, under the Apache License v2.0 with LLVM Exceptions.
// See https://llvm.org/LICENSE.txt for license information.
// SPDX-License-Identifier: Apache-2.0 WITH LLVM-exception
//
//===-----------------------------------------------------------------------===

// Only include this if we are aiming for MSVC compatibility.
// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// for backward compatibility

type X__gnuc_va_list = X__builtin_va_list /* vadefs.h:24:29 */

type Ssize_t = int64 /* corecrt.h:45:35 */

type Rsize_t = Size_t /* corecrt.h:52:16 */

type Intptr_t = int64 /* corecrt.h:62:35 */

type Uintptr_t = uint64 /* corecrt.h:75:44 */

type Wint_t = uint16   /* corecrt.h:106:24 */
type Wctype_t = uint16 /* corecrt.h:107:24 */

type Errno_t = int32 /* corecrt.h:113:13 */

type X__time32_t = int32 /* corecrt.h:118:14 */

type X__time64_t = int64 /* corecrt.h:123:35 */

type Time_t = X__time64_t /* corecrt.h:138:20 */

type Threadlocaleinfostruct = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
} /* corecrt.h:430:1 */

type Pthreadlocinfo = uintptr /* corecrt.h:432:39 */
type Pthreadmbcinfo = uintptr /* corecrt.h:433:36 */

type Localeinfo_struct = struct {
	Flocinfo Pthreadlocinfo
	Fmbcinfo Pthreadmbcinfo
} /* corecrt.h:436:9 */

type X_locale_tstruct = Localeinfo_struct /* corecrt.h:439:3 */
type X_locale_t = uintptr                 /* corecrt.h:439:19 */

type TagLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
} /* corecrt.h:443:9 */

type LC_ID = TagLC_ID  /* corecrt.h:447:3 */
type LPLC_ID = uintptr /* corecrt.h:447:9 */

type Threadlocinfo = Threadlocaleinfostruct /* corecrt.h:482:3 */
type X_fsize_t = uint32                     /* io.h:29:25 */

type X_finddata32_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        X_fsize_t
	Fname        [260]int8
} /* io.h:35:3 */

type X_finddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        int64
	Fname        [260]int8
	F__ccgo_pad1 [4]byte
} /* io.h:44:3 */

type X_finddata64i32_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        X_fsize_t
	Fname        [260]int8
} /* io.h:53:3 */

type X__finddata64_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        int64
	Fname        [260]int8
	F__ccgo_pad2 [4]byte
} /* io.h:62:3 */

type X_wfinddata32_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        X_fsize_t
	Fname        [260]Wchar_t
} /* io.h:94:3 */

type X_wfinddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create X__time32_t
	Ftime_access X__time32_t
	Ftime_write  X__time32_t
	Fsize        int64
	Fname        [260]Wchar_t
} /* io.h:103:3 */

type X_wfinddata64i32_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        X_fsize_t
	Fname        [260]Wchar_t
	F__ccgo_pad2 [4]byte
} /* io.h:112:3 */

type X_wfinddata64_t = struct {
	Fattrib      uint32
	F__ccgo_pad1 [4]byte
	Ftime_create X__time64_t
	Ftime_access X__time64_t
	Ftime_write  X__time64_t
	Fsize        int64
	Fname        [260]Wchar_t
} /* io.h:121:3 */

type X_off_t = int32 /* _mingw_off_t.h:5:16 */
type Off32_t = int32 /* _mingw_off_t.h:7:16 */

type X_off64_t = int64 /* _mingw_off_t.h:13:39 */
type Off64_t = int64   /* _mingw_off_t.h:15:39 */

type Off_t = Off64_t /* _mingw_off_t.h:24:17 */

type X_PVFV = uintptr /* corecrt_startup.h:20:14 */
type X_PIFV = uintptr /* corecrt_startup.h:21:13 */
type X_PVFI = uintptr /* corecrt_startup.h:22:14 */

type X_onexit_table_t1 = struct {
	F_first uintptr
	F_last  uintptr
	F_end   uintptr
} /* corecrt_startup.h:24:9 */

type X_onexit_table_t = X_onexit_table_t1 /* corecrt_startup.h:28:3 */

type X_onexit_t = uintptr /* corecrt_startup.h:30:13 */

// Includes a definition of _pid_t and pid_t
// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

type X_ino_t = uint16 /* types.h:43:24 */
type Ino_t = uint16   /* types.h:45:24 */

type X_dev_t = uint32 /* types.h:51:22 */
type Dev_t = uint32   /* types.h:53:22 */

type X_pid_t = int64 /* types.h:63:17 */

type Pid_t = X_pid_t /* types.h:68:16 */

type X_mode_t = uint16 /* types.h:74:24 */

type Mode_t = X_mode_t /* types.h:77:17 */

type Useconds_t = uint32 /* types.h:84:22 */

type Timespec = struct {
	Ftv_sec      Time_t
	Ftv_nsec     int32
	F__ccgo_pad1 [4]byte
} /* types.h:89:1 */

type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec      Time_t
		Ftv_nsec     int32
		F__ccgo_pad1 [4]byte
	}
	Fit_value struct {
		Ftv_sec      Time_t
		Ftv_nsec     int32
		F__ccgo_pad1 [4]byte
	}
} /* types.h:94:1 */

type X_sigset_t = uint64 /* types.h:104:28 */

type X_beginthread_proc_type = uintptr   /* process.h:32:16 */
type X_beginthreadex_proc_type = uintptr /* process.h:33:20 */

type X_tls_callback_type = uintptr /* process.h:64:16 */

//
//    Copyright (c) 2011-2016  mingw-w64 project
//
//    Permission is hereby granted, free of charge, to any person obtaining a
//    copy of this software and associated documentation files (the "Software"),
//    to deal in the Software without restriction, including without limitation
//    the rights to use, copy, modify, merge, publish, distribute, sublicense,
//    and/or sell copies of the Software, and to permit persons to whom the
//    Software is furnished to do so, subject to the following conditions:
//
//    The above copyright notice and this permission notice shall be included in
//    all copies or substantial portions of the Software.
//
//    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//    FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
//    DEALINGS IN THE SOFTWARE.

// Set defines described by the POSIX Threads Extension (1003.1c-1995)
// _SC_THREADS
//   Basic support for POSIX threads is available. The functions
//
//   pthread_atfork(),
//   pthread_attr_destroy(),
//   pthread_attr_getdetachstate(),
//   pthread_attr_getschedparam(),
//   pthread_attr_init(),
//   pthread_attr_setdetachstate(),
//   pthread_attr_setschedparam(),
//   pthread_cancel(),
//   pthread_cleanup_push(),
//   pthread_cleanup_pop(),
//   pthread_cond_broadcast(),
//   pthread_cond_destroy(),
//   pthread_cond_init(),
//   pthread_cond_signal(),
//   pthread_cond_timedwait(),
//   pthread_cond_wait(),
//   pthread_condattr_destroy(),
//   pthread_condattr_init(),
//   pthread_create(),
//   pthread_detach(),
//   pthread_equal(),
//   pthread_exit(),
//   pthread_getspecific(),
//   pthread_join(,
//   pthread_key_create(),
//   pthread_key_delete(),
//   pthread_mutex_destroy(),
//   pthread_mutex_init(),
//   pthread_mutex_lock(),
//   pthread_mutex_trylock(),
//   pthread_mutex_unlock(),
//   pthread_mutexattr_destroy(),
//   pthread_mutexattr_init(),
//   pthread_once(),
//   pthread_rwlock_destroy(),
//   pthread_rwlock_init(),
//   pthread_rwlock_rdlock(),
//   pthread_rwlock_tryrdlock(),
//   pthread_rwlock_trywrlock(),
//   pthread_rwlock_unlock(),
//   pthread_rwlock_wrlock(),
//   pthread_rwlockattr_destroy(),
//   pthread_rwlockattr_init(),
//   pthread_self(),
//   pthread_setcancelstate(),
//   pthread_setcanceltype(),
//   pthread_setspecific(),
//   pthread_testcancel()
//
//   are present.

// _SC_READER_WRITER_LOCKS
//   This option implies the _POSIX_THREADS option. Conversely, under
//   POSIX 1003.1-2001 the _POSIX_THREADS option implies this option.
//
//   The functions
//   pthread_rwlock_destroy(),
//   pthread_rwlock_init(),
//   pthread_rwlock_rdlock(),
//   pthread_rwlock_tryrdlock(),
//   pthread_rwlock_trywrlock(),
//   pthread_rwlock_unlock(),
//   pthread_rwlock_wrlock(),
//   pthread_rwlockattr_destroy(),
//   pthread_rwlockattr_init()
//
//   are present.

// _SC_SPIN_LOCKS
//   This option implies the _POSIX_THREADS and _POSIX_THREAD_SAFE_FUNCTIONS
//   options. The functions
//
//   pthread_spin_destroy(),
//   pthread_spin_init(),
//   pthread_spin_lock(),
//   pthread_spin_trylock(),
//   pthread_spin_unlock()
//
//   are present.

// _SC_BARRIERS
//   This option implies the _POSIX_THREADS and _POSIX_THREAD_SAFE_FUNCTIONS
//   options. The functions
//
//   pthread_barrier_destroy(),
//   pthread_barrier_init(),
//   pthread_barrier_wait(),
//   pthread_barrierattr_destroy(),
//   pthread_barrierattr_init()
//
//   are present.

// _SC_TIMEOUTS
//   The functions
//
//   mq_timedreceive(), - not supported
//   mq_timedsend(), - not supported
//   posix_trace_timedgetnext_event(), - not supported
//   pthread_mutex_timedlock(),
//   pthread_rwlock_timedrdlock(),
//   pthread_rwlock_timedwrlock(),
//   sem_timedwait(),
//
//   are present.

// _SC_TIMERS - not supported
//   The functions
//
//   clock_getres(),
//   clock_gettime(),
//   clock_settime(),
//   nanosleep(),
//   timer_create(),
//   timer_delete(),
//   timer_gettime(),
//   timer_getoverrun(),
//   timer_settime()
//
//   are present.
// #undef _POSIX_TIMERS

// _SC_CLOCK_SELECTION
//    This option implies the _POSIX_TIMERS option. The functions
//
//    pthread_condattr_getclock(),
//    pthread_condattr_setclock(),
//    clock_nanosleep()
//
//    are present.

// _SC_SEMAPHORES
//   The include file <semaphore.h> is present. The functions
//
//   sem_close(),
//   sem_destroy(),
//   sem_getvalue(),
//   sem_init(),
//   sem_open(),
//   sem_post(),
//   sem_trywait(),
//   sem_unlink(),
//   sem_wait()
//
//   are present.

var _ int8 /* gen.c:2:13: */
