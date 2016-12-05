package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreGraphics -framework AppKit

#import <Foundation/Foundation.h>
#import <CoreGraphics/CoreGraphics.h>
#import <AppKit/NSEvent.h>
#import <AppKit/NSScreen.h>

void click(int x, int y) {
  CGPoint location = CGPointMake(x, y);

  CGEventRef mouseClick = CGEventCreateMouseEvent(NULL, kCGEventLeftMouseDown, location, kCGMouseButtonLeft);
  CGEventSetIntegerValueField(mouseClick, kCGMouseEventClickState, 1);
  CGEventPost(kCGHIDEventTap, mouseClick);
  CGEventSetType(mouseClick, kCGEventLeftMouseUp);
  CGEventPost(kCGHIDEventTap, mouseClick);
  CFRelease(mouseClick);
}

void move(int x, int y) {
  CGEventRef mouseMove = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, CGPointMake(x, y), kCGMouseButtonLeft);
  CGEventPost(kCGHIDEventTap, mouseMove);
  CFRelease(mouseMove);
}

void getMouseLoc(char **ml) {
  *ml = (char*)calloc(32, sizeof(char));
  NSPoint mouseLocation = [NSEvent mouseLocation];
  NSRect screenSize = [[NSScreen mainScreen]frame];
  sprintf(*ml, "xy:%d,%d", (int)mouseLocation.x, (int)(screenSize.size.height-mouseLocation.y));
}
*/
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// Sleep duration of each actions
const sleepD = 20

func main() {
	if len(os.Args) < 2 {
		epe("some arguments required.")
	}

  // Get mouse location mode
	if "loc" == os.Args[1] {
		for {
			getMouseLoc()
			sleepM(3000)
		}
		return
	}
  // Show help
  if "help" == os.Args[1] {
    usage()
    return
  }

  // Actions
	for _, arg := range os.Args[1:] {

		a := strings.Split(arg, ":")
		if len(a) < 2 {
			epe("unexpected arguments.")
		}

		ope, val := a[0], a[1]

		switch {
		case '0' < ope[0] && '0'+10 > ope[0]: // click
			n := parseNum(ope)
			x, y := parseLoc(val)
			move(x, y)
			click(n, x, y)
		case 'm' == ope[0]: // move
			x, y := parseLoc(val)
			move(x, y)
		case 'w' == ope[0]: // wait
			n := parseNum(val)
			sleepM(n)
		default:
			epe("unexpected operator.")
		}
	}
}

func usage() {
  str := 
`Usage: go3click [loc] [num:x,y] [m:x,y] [w:num] [help]
  loc      Print mouse location each 3 seconds endless
  num:x,y  Click specified location x,y num times
  m:x,y    Move mouse to specified location x,y
  w:num    Wait programm num milliseconds
`
  fmt.Fprintf(os.Stderr, "%s", str)
}

// Error print and exit
func epe(s string) {
	fmt.Fprintf(os.Stderr, "error. %s\n", s)
  usage()
	os.Exit(1)
}

// Sleep milliseconds
func sleepM(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// Parse string to int
func parseNum(nums string) int {
	n, err := strconv.Atoi(nums)
	if err != nil {
		epe("could not parse " + nums + " to int.")
	}
	return n
}

// Parse mouse position string to C.int
func parseLoc(locs string) (C.int, C.int) {
	loc := strings.Split(locs, ",")
	if len(loc) < 2 {
		epe("unexpected format " + locs + " .")
	}
	x, err := strconv.Atoi(loc[0])
	if err != nil {
		epe("could not parse " + loc[0] + " to int.")
	}
	y, err := strconv.Atoi(loc[1])
	if err != nil {
		epe("could not parse " + loc[1] + " to int.")
	}
	return C.int(x), C.int(y)
}

// Wrapper of objective-c functions

func click(n int, x, y C.int) {
	for i := 0; i < n; i++ {
		C.click(x, y)
		sleepM(sleepD)
	}
}

func move(x, y C.int) {
	C.move(x, y)
	sleepM(sleepD)
}

func getMouseLoc() {
	var ml *C.char
	defer C.free(unsafe.Pointer(ml))
	C.getMouseLoc(&ml)
	fmt.Println(C.GoString(ml))
}
