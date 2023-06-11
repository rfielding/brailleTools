# You'll need to connect Keybow 2040 to a computer, as you would with a regular
# USB keyboard.
# Drop the keybow2040.py file into your `lib` folder on your `CIRCUITPY` drive.

# NOTE! Requires the adafruit_hid CircuitPython library also!
import board
from keybow2040 import Keybow2040

# Layout: (with usb-c pointing down, thumbs on back, fingers on front):
#
# 8-dot braille (using 7-dots literally) for middle two rows. As seen from the BACK:
# 
# [?     ][d1][d4][tab ]
# [repeat][d2][d5][Fn   ]
# [gui   ][d3][d6][alt ]
# [shift ][d7][d8][ctrl]
#
# Note: gui is the "windows" key
# It's important to push down keys in order.
# the d keys must all come up after other keys
# the tab key sends control on down, so modifiers
# must go down first, etc.

# ex: to navigate shift+tab to tab backwards
# gui+tab to move between apps
# ctrl+tab to move between app tabs
# ctrl+c to send ctrl c.

# For all middle d fingers, no key is emitted until
# all fingers come up.
# When all fingers come up in d rows, the shift,alt,ctrl
# are checked at the time all d fingers come up.
# ie: 'A' is either: d1+d7, or it's d1+shift
#
# Everything above 7-bit ASCII is unreachable directly,
# because keycodes don't consistently map them anyway.
#
# space: d8
# backspace: d7
# enter: d7+d8
#
# Arrows mimic the vi text editor:
#
# left: h+d8 (ie: d1+d2+d5+d8)
# right: l+d8 (ie: d1+d2+d3+d8)
# down: j+d8 (ie: d2+d4+d5+d8)
# up: k+d8 (ie: d1+d3+d8)
#
# home: H+d8 (ie: d1+d2+d5+d8)
# end: L+d8 (ie: d1+d2+d3+d8)
# pagedown: J+d8 (ie: d2+d4+d5+d8)
# pageup: K+d8 (ie: d1+d3+d8)
#
# Function keys F1 .. F12
#
# F1..F9 Fn+[1..9]
# F10,F11,F12 Fn+[0,a,b]
#
# There is a principled way to take the basic 6-dot ASCII braille standard,
# and interpret it as 8dot computer braille. Here is the full map, which
# includes theoretical, but unused, dot patterns as well:
#
#   https://github.com/rfielding/brailleTools/tree/master/cmd/brascii 
#
# Importantly, It shows you how to make raw chars like TAB and ESC, so
# that you can use text editors like vi.
#
#
# This code is hosted here:
#
# https://github.com/rfielding/brailleTools/blob/master/circuitpy/code.py
#
# To get good at this keyboard, practice computer braille.
# Start off typing on qwerty, then mimic it in braille on
# the keybow2040 loaded with this program
#
# https://rfielding.github.io/editor.html

import usb_hid
from adafruit_hid.keyboard import Keyboard
from adafruit_hid.keyboard_layout_us import KeyboardLayoutUS
from adafruit_hid.keycode import Keycode

# Set up Keybow
i2c = board.I2C()
keybow = Keybow2040(i2c)
keys = keybow.keys

isFn = False
isAlt = False
isCtrl = False
isShift = False
isGUI = False
isRepeat = False
keyRepeated = None

kRepeat = 1
kGUI = 2
kShift = 3
kTab = 12
kFn = 13
kAlt = 14
kCtrl = 15

# Set up the keyboard and layout
keyboard = Keyboard(usb_hid.devices)
layout = KeyboardLayoutUS(keyboard)

pink = (255,128,128)
red = (255,0,0)
yellow = (255, 255, 0)
green = (0, 255, 0)
darkGreen = (0, 64, 0)
purple = (255,0,255)
black = (0,0,0)


# keep a key red on init
keys[0].set_led(*red)

# lives from 32..96 in the standard, as all six-dot patterns
braillePattern = [
        0b00000000, 0b00101110, 0b00010000, 0b00111100, 0b00101011, 0b00101001, 0b00101111, 0b00000100, 0b00110111, 0b00111110, 0b00100001, 0b00101100, 0b00100000, 0b00100100, 0b00101000, 0b00001100,
        0b00110100, 0b00000010, 0b00000110, 0b00010010, 0b00110010, 0b00100010, 0b00010110, 0b00110110, 0b00100110, 0b00010100, 0b00110001, 0b00110000, 0b00100011, 0b00111111, 0b00011100, 0b00111001,
        0b00001000, 0b00000001, 0b00000011, 0b00001001, 0b00011001, 0b00010001, 0b00001011, 0b00011011, 0b00010011, 0b00001010, 0b00011010, 0b00000101, 0b00000111, 0b00001101, 0b00011101, 0b00010101,
        0b00001111, 0b00011111, 0b00010111, 0b00001110, 0b00011110, 0b00100101, 0b00100111, 0b00111010, 0b00101101, 0b00111101, 0b00110101, 0b00101010, 0b00110011, 0b00111011, 0b00011000, 0b00111000
]

### Make the 8-dot standard pattern from the 6-dot standard
braillePermutation = [0 for i in range(0,256)]
for i in range(0,32): #ctrl chars
    braillePermutation[i] = braillePattern[i]+64
for i in range(0,32): #punctuation numbers
    braillePermutation[32+i] = braillePattern[i]
for i in range(0,32): #uppercase
    braillePermutation[64+i] = braillePattern[32+i]+64
for i in range(0,32): #lowercase
    braillePermutation[64+32+i] = braillePattern[32+i]
# this odd swap shows up in real terminals
bpt = braillePermutation[0x5F]
braillePermutation[0x5F] = braillePermutation[0x7F]
braillePermutation[0x7F] = bpt
for i in range(0,128): # upper unused half
    braillePermutation[128+i] = braillePermutation[i] + 128
brailleAsciiMap = [0 for i in range(0,256)]
for i in range(0,256):
    brailleAsciiMap[braillePermutation[i]] = i
# for some reason, the stack runs out of space if i dont wipe this
braillePermutation = []

# these are still not right
kcEnter=13

# Map the 0-127 ascii chars of 7-dot braille to keys to send
charToKeycodeMap = [
	[0],[0],[0],[0],[0],[0],[0],[0], [Keycode.BACKSPACE],[Keycode.TAB],[Keycode.ENTER],[0],[0],[Keycode.ENTER],[0],[0],
	[Keycode.DELETE],[0],[0],[0],[0],[0],[0],[0], [0],[0],[0],[Keycode.ESCAPE],[0],[0],[0],[0],
	[Keycode.SPACE],[Keycode.ONE,Keycode.SHIFT],[Keycode.QUOTE,Keycode.SHIFT],[Keycode.THREE,Keycode.SHIFT],[Keycode.FOUR,Keycode.SHIFT],[Keycode.FIVE,Keycode.SHIFT],[Keycode.SEVEN,Keycode.SHIFT],[Keycode.QUOTE], [Keycode.NINE,Keycode.SHIFT],[Keycode.ZERO, Keycode.SHIFT],[Keycode.EIGHT,Keycode.SHIFT],[Keycode.EQUALS,Keycode.SHIFT],[Keycode.COMMA],[Keycode.MINUS],[Keycode.PERIOD],[Keycode.FORWARD_SLASH],
	[Keycode.ZERO],[Keycode.ONE],[Keycode.TWO],[Keycode.THREE],[Keycode.FOUR],[Keycode.FIVE],[Keycode.SIX],[Keycode.SEVEN],[Keycode.EIGHT],[Keycode.NINE],[Keycode.SEMICOLON,Keycode.SHIFT],[Keycode.SEMICOLON],[Keycode.COMMA,Keycode.SHIFT],[Keycode.EQUALS],[Keycode.PERIOD,Keycode.SHIFT],[Keycode.FORWARD_SLASH,Keycode.SHIFT],
	[Keycode.TWO,Keycode.SHIFT],[Keycode.A,Keycode.SHIFT],[Keycode.B,Keycode.SHIFT],[Keycode.C,Keycode.SHIFT],[Keycode.D,Keycode.SHIFT],[Keycode.E,Keycode.SHIFT],[Keycode.F,Keycode.SHIFT],[Keycode.G,Keycode.SHIFT], [Keycode.H,Keycode.SHIFT],[Keycode.I,Keycode.SHIFT],[Keycode.J,Keycode.SHIFT],[Keycode.K,Keycode.SHIFT],[Keycode.L,Keycode.SHIFT],[Keycode.M,Keycode.SHIFT],[Keycode.N,Keycode.SHIFT],[Keycode.O,Keycode.SHIFT],
	[Keycode.P,Keycode.SHIFT],[Keycode.Q,Keycode.SHIFT],[Keycode.R,Keycode.SHIFT],[Keycode.S,Keycode.SHIFT],[Keycode.T,Keycode.SHIFT],[Keycode.U,Keycode.SHIFT],[Keycode.V,Keycode.SHIFT],[Keycode.W,Keycode.SHIFT],[Keycode.X,Keycode.SHIFT],[Keycode.Y,Keycode.SHIFT],[Keycode.Z,Keycode.SHIFT],[Keycode.LEFT_BRACKET],[Keycode.BACKSLASH],[Keycode.RIGHT_BRACKET],[Keycode.SIX,Keycode.SHIFT],[Keycode.MINUS,Keycode.SHIFT],
	[Keycode.GRAVE_ACCENT],[Keycode.A],[Keycode.B],[Keycode.C],[Keycode.D],[Keycode.E],[Keycode.F],[Keycode.G], [Keycode.H],[Keycode.I],[Keycode.J],[Keycode.K],[Keycode.L],[Keycode.M],[Keycode.N],[Keycode.O],
	[Keycode.P],[Keycode.Q],[Keycode.R],[Keycode.S],[Keycode.T],[Keycode.U],[Keycode.V],[Keycode.W], [Keycode.X],[Keycode.Y],[Keycode.Z],[Keycode.LEFT_BRACKET,Keycode.SHIFT],[Keycode.BACKSLASH,Keycode.SHIFT],[Keycode.RIGHT_BRACKET,Keycode.SHIFT],[Keycode.GRAVE_ACCENT,Keycode.SHIFT],[Keycode.DELETE]
]
# 8 dot braille is used
dots = 8
# reverse mapping to find the keys,
keyToDot = [-1 for i in range(0,16)]
keyToHeld = [False for i in range(0,16)]
keyToUsed = [False for i in range(0,16)]
dot2Key = [-1 for i in range(0,8)]

def isBrailleKey(n):
    return keyToDot[n] != -1

# convert dots to binary pattern form
def dots2ord():
    value = 0
    m = 1
    for i in range(0,8):
        n = dot2Key[i]
        h = keyToHeld[n]
        if h:
            value += m
        m *= 2
    return value%256

def charToKeycode(c):
    return charToKeycodeMap[c % 256]

def press(keys):
    keyboard.press(*keys)

def release(keys):
    keyboard.release(*keys)

def handle_down(key):
    global isFn
    global isAlt
    global isCtrl
    global isShift
    global isGUI
    global keyToHeld
    global isRepeat
    global keyRepeated
    n = key.number
    if isBrailleKey(n):
        key.set_led(*green)
        keyToHeld[n] = True
        keyToUsed[n] = True
    else:
        key.set_led(*yellow)
        if key.number == kFn:
            isFn = True
            print("Fn")
            key.set_led(*pink)
        if key.number == kShift:
            isShift = True
            print("shift")
            keyboard.press(Keycode.SHIFT)
        if key.number == kCtrl:
            isCtrl = True
            keyboard.press(Keycode.CONTROL)
            print("ctrl")
        if key.number == kAlt:
            isAlt = True
            keyboard.press(Keycode.ALT)
            print("alt")
        if key.number == kGUI:
            isGUI = True
            keyboard.press(Keycode.GUI)
            print("win")
        if key.number == kTab:
            key.set_led(*red)
            keyboard.press(Keycode.TAB)
            print("tab")
        if key.number == kRepeat:
            key.set_led(*purple)
            isRepeat = True
            print("repeat")
            if keyRepeated:
                keyboard.press(keyRepeated)

def totalUsed():
    global keyToUsed
    t = 0
    for n in range(0,16):
        if keyToUsed[n]:
            t += 1
    return t

def clearDotLEDs():
    global keys
    for n in range(0,16):
        if isBrailleKey(n):
            keys[n].set_led(*black)

def clearDotHeld():
    global keyToHeld
    for n in range(0,16):
        if isBrailleKey(n):
            keyToHeld[n] = False

def handle_up(key):
    global isFn
    global isAlt
    global isCtrl
    global isShift
    global isGUI
    global isRepeat
    global keyRepeated
    n = key.number
    if isBrailleKey(n):
        keyToUsed[n] = False
        t = totalUsed()
        if t == 0:
            o = dots2ord()
            theKeys = [0]
            if o == 64:
                theKeys[0] = Keycode.BACKSPACE
                print("BACKSPACE")
            elif o == 128:
                theKeys[0] = Keycode.SPACE
                print("SPACE")
            elif o == 128+64:
                theKeys[0] = Keycode.ENTER
                print("ENTER")
            elif o >= 128:
                # vi style arrows with dot 8
                if o == 1+2+16+128: # braille h+dot8
                    theKeys[0] = Keycode.LEFT_ARROW
                    print("LEFT_ARROW")
                if o == 1+2+4+128:  # braile l+dot8
                    theKeys[0] = Keycode.RIGHT_ARROW
                    print("RIGHT_ARROW")
                if o == 2+8+16+128: # braille j+dot8 
                    theKeys[0] = Keycode.DOWN_ARROW
                    print("DOWN_ARROW")
                if o == 1+4+128: # braille k+dot8
                    theKeys[0] = Keycode.UP_ARROW
                    print("UP_ARROW")
                if o == 1+2+16+64+128: # braille H+dot8
                    theKeys[0] = Keycode.HOME
                    print("HOME")
                if o == 1+2+4+64+128:  # braille L+dot8
                    theKeys[0] = Keycode.END
                    print("END")
                if o == 2+8+16+64+128:  # braille J+dot8
                    theKeys[0] = Keycode.PAGE_DOWN
                    print("PAGE_DOWN")
                if o == 1+4+64+128:   # braille L+dot8 
                    theKeys[0] = Keycode.PAGE_UP
                    print("PAGE_UP")
            else:
                c = brailleAsciiMap[o%128]
                theKeys = charToKeycodeMap[c%128].copy()
                if isFn: # Fn isn't a real key on a keyboard. But for Fn key combos
                    if c == ord('1'):
                        theKeys = [Keycode.F1]
                        print("F1")    
                    if c == ord('2'):
                        theKeys = [Keycode.F2]    
                        print("F2")    
                    if c == ord('3'):
                        theKeys = [Keycode.F3]    
                        print("F3")    
                    if c == ord('4'):
                        theKeys = [Keycode.F4]    
                        print("F4")    
                    if c == ord('5'):
                        theKeys = [Keycode.F5]    
                        print("F5")    
                    if c == ord('6'):
                        theKeys = [Keycode.F6]    
                        print("F6")    
                    if c == ord('7'):
                        theKeys = [Keycode.F7]    
                        print("F7")    
                    if c == ord('8'):
                        theKeys = [Keycode.F8]    
                        print("F8")    
                    if c == ord('9'):
                        theKeys = [Keycode.F9]    
                        print("F9")    
                    if c == ord('0'):
                        theKeys = [Keycode.F10]    
                        print("F10")    
                    if c == ord('a'):
                        theKeys = [Keycode.F11]    
                        print("F11")    
                    if c == ord('b'):
                        theKeys = [Keycode.F12]    
                        print("F12")    
                else:
                  pass # just do the default braille thing
            press(theKeys)
            clearDotLEDs()
            clearDotHeld()
            if isRepeat:
                keyRepeated = theKeys
            else:
                release(theKeys)
        else:
            pass
    else:
        key.set_led(*black)
        if key.number == kCtrl:
            isCtrl = False
            keyboard.release(Keycode.CONTROL)
        if key.number == kAlt:
            isAlt = False
            keyboard.release(Keycode.ALT)
        if key.number == kFn:
            isFn = False
        if key.number == kShift:
            isShift = False
            keyboard.release(Keycode.SHIFT)
        if key.number == kGUI:
            isGUI = False
            keyboard.release(Keycode.GUI)
        if key.number == kTab:
            keyboard.release(Keycode.TAB)
        if key.number == kRepeat:
            isRepeat = False
            if keyRepeated == None:
                pass
            else:
                release(keyRepeated)
            keyRepeated = None

for key in keys:
    @keybow.on_press(key)
    def press_handler(key):
        handle_down(key)
    @keybow.on_release(key)
    def release_handler(key):
        handle_up(key)

def mapDot(d,n):
    global keyToDot
    global dot2Key
    global keyToUsed
    keyToDot[n] = d
    keyToHeld[n] = False
    keyToUsed[n] = False
    dot2Key[d] = n

# set the braile dots
# otherwise keyToDot[n] == -1
# and dot2key[d] == -1
mapDot(0,4)
mapDot(1,5)
mapDot(2,6)
mapDot(3,8)
mapDot(4,9)
mapDot(5,10)
mapDot(6,7)
mapDot(7,11)

keys[0].set_led(*black)
while True:
    keybow.update()
