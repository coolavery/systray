// +build linux
//go:build linux

package systray

/*
#cgo linux pkg-config: appindicator3-0.1
#cgo linux CFLAGS: -DUSE_LEGACY_APPINDICATOR

#include "systray.h"
*/
import "C"
