# Keylogger
Read the keycodes from a device and print to stdout

# Usage
```sh
keylogger $file # where $file corresponds to the keyboard, e.g. /dev/input/by-id/usb-OLKB_Planck_0-event-kbd
```

# Example
```sh
keylogger /dev/input/by-id/usb-OLKB_Planck_0-event-kbd > /tmp/keylog
# do some typing...
cat /tmp/keylog | sort | uniq -c | sort -rh
```
