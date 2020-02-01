## The following instructions will let you compile the Wallet from source

### 1. Download the Go distribution from the official website.

   The Go distribution and tooling is available as an installer for all common operating systems.
   Visit <https://golang.org/dl/> to download to correct version for your OS. The installation instructions can be found [here](https://golang.org/doc/install).

### 2. Download and install NPM.

   NPM and Node.js can be downloaded from their [official website](https://nodejs.org/en/download/). Simply select your distribution/OS and CPU architecture. 

### 3. Install Wails
  
   The Molly Wallet is build using a light-weight framework for Desktop Applications using Go and VueJS. Wails is very similar to Electron but is not packaging the full Chromium web browser as a dependency.

   To install Wails, follow the instructions for your OS below.

   #### MacOS

   Make sure you have the xcode command line tools installed. This can be done by running:

   `xcode-select --install`

   #### Linux

   ##### Debian/Ubuntu

   `sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev`

   _Debian: 8, 9, 10_

   _Ubuntu: 16.04, 18.04, 19.04_

   _Also succesfully tested on: Zorin 15, Parrot 4.7, Linuxmint 19, Elementary 5, Kali, Neon_

   ##### Arch Linux

   `sudo pacman -S webkit2gtk gtk3`

   _Also succesfully test on: Manjaro & ArcoLinux_

   ##### Centos

   `sudo yum install webkitgtk3-devel gtk3-devel`

   _CentOS 6, 7_

   ##### Fedora

   `sudo yum install webkit2gtk3-devel gtk3-devel`

   _Fedora 29, 30_
 
   ##### VoidLinux & VoidLinux-musl

   `xbps-install gtk+3-devel webkit2gtk-devel`

   ##### Gentoo

   `sudo emerge gtk+:3 webkit-gtk`

   #### Windows

   Windows requires gcc and related tooling. The recommended download is from [http://tdm-gcc.tdragon.net/download](http://tdm-gcc.tdragon.net/download). Once this is installed, you are good to go.

   ### Installation

   **Ensure Go modules are enabled: GO111MODULE=on and go/bin is in your PATH variable.**

   Installation is as simple as running the following command:

   <pre style='color:white'>
   go get -u github.com/wailsapp/wails/cmd/wails
   </pre>

4. Clone this repository into your GOPATH

   `https://github.com/grvlle/constellation_wallet.git`

5. Change to 'develop' branch

   `git checkout develop`

6. Compile the wallet using Wails CLI tool

   `wails build -d`
