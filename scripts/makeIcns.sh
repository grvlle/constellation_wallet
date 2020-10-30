#!/bin/bash

# convert appicon.png -resize '128x128!' icon-128.png
# convert appicon.png -resize '256x256!' icon-256.png
# convert icontray.png -resize '16x16!' icon-16.png

mkdir MyIcon.iconset
sips -z 16 16     appicon.png --out MyIcon.iconset/icon_16x16.png
sips -z 32 32     appicon.png --out MyIcon.iconset/icon_16x16@2x.png
sips -z 32 32     appicon.png --out MyIcon.iconset/icon_32x32.png
sips -z 64 64     appicon.png --out MyIcon.iconset/icon_32x32@2x.png
sips -z 128 128   appicon.png --out MyIcon.iconset/icon_128x128.png
sips -z 256 256   appicon.png --out MyIcon.iconset/icon_128x128@2x.png
sips -z 256 256   appicon.png --out MyIcon.iconset/icon_256x256.png
sips -z 512 512   appicon.png --out MyIcon.iconset/icon_256x256@2x.png
sips -z 512 512   appicon.png --out MyIcon.iconset/icon_512x512.png
sips -z 1024 1024   appicon.png --out MyIcon.iconset/icon_512x512@2x.png
# cp appicon.png MyIcon.iconset/icon_512x512@2x.png
iconutil -c icns MyIcon.iconset
rm -R MyIcon.iconset
mv MyIcon.icns iconfile.icns
