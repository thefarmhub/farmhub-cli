//go:build !windows

/*
 * This file is part of go-win32-utils.
 *
 * Copyright 2018-2023 ARDUINO SA (http://www.arduino.cc/)
 *
 * go-win32-utils is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 */

package win32

import (
	"fmt"
	"runtime"
)

// The functions defined below allow compile on non-Windows OS. The caller
// may choose to not call those functions based on runtime.GOOS value.

// GetDocumentsFolder returns the Document folder
func GetDocumentsFolder() (string, error) {
	return "", fmt.Errorf("operating system not supported: %s", runtime.GOOS)
}

// GetLocalAppDataFolder returns the LocalAppData folder
func GetLocalAppDataFolder() (string, error) {
	return "", fmt.Errorf("operating system not supported: %s", runtime.GOOS)
}

// GetRoamingAppDataFolder returns the AppData folder
func GetRoamingAppDataFolder() (string, error) {
	return "", fmt.Errorf("operating system not supported: %s", runtime.GOOS)
}
