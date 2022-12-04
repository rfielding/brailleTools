import storage, usb_cdc
import board, digitalio

print("boot.py is running")

button = digitalio.DigitalInOut(board.SW0)
button.pull = digitalio.Pull.UP

if button.value:
  print("disabling usb cdc")
  usb_cdc.disable()
  print("disabling usb drive and cdc")
  storage.disable_usb_drive()
else:
  print("press SW0 to disable usb drive and cdc")

