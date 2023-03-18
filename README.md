# image-generator

<img src="images/gradient-square.png" widht=100>

## Gradient
```bash
$ go run scripts/gradient.go -o output/gradient.png -w 720 -h 100 -angle 135 FFCB5E F7302D DB3251 B63671 8944B8

# Usage of gradient:
#   -w int
#         Width of the image (default 256)
#   -h int
#         Height of the image (default 256)
#   -angle float
#         Angle of rotation in degrees (default 0.0)
#   -o string
#         Output file name (default "output.jpg")
#  color string
#         Hex color codes (required)
```
<img src="images/gradient-header.png" widht=100>


### Samples

```bash
# Gradient Color Icon
go run scripts/gradient.go -o output/gradient-icon.png -angle 45 FFCB5E F7302D DB3251 B63671 8944B8 
# Solid Color Icon
go run scripts/gradient.go -o output/solid-icon.png -w 200 -h 200 076676 076676 
# Wallpaper for PC
go run scripts/gradient.go -o output/gradient-wallpaper-landscape.png -w 1920 -h 1080 -angle 45 FFCB5E F7302D DB3251 B63671 8944B8 
# Wallpaper for Phone
go run scripts/gradient.go -o output/gradient-wallpaper-portrait.png -w 640 -h 940 -angle 45 FFCB5E F7302D DB3251 B63671 8944B8 
```


