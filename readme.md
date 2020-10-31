# Molly Wallet  
  
![N|Solid](https://i.ibb.co/qRK8Cj9/mollywallet.png)
  
The Molly Wallet is the official $DAG desktop wallet of the Constellation Network. The wallet is supported and developed by the Stardust Collective.
  
---  

### Technologies  
  
Molly Wallet uses a number of open source projects:  
  
* [VueJS](https://vuejs.org) - The Progressive JavaScript Framework  
* [Go](https://golang.org) - Go is an open source programming language that makes it easy to build simple, reliable, and  
efficient software.  
* [GORM](https://gorm.io) - Object Relational Mapping for Go.  
* [Wails](https://wails.app/) - A framework for building desktop apps using Go & Web Technologies  

---    

### Installation  
  
#### Download Molly Wallet  
The latest builds can be found under [releases](https://github.com/StardustCollective/molly_wallet/releases).  
  
---  
  
### Support
Need any help? You can visit the Stardust Support channel in Telegram [here](https://t.me/StardustSupport).

---  

### Development Enviornment  
  
#### 1. Download the Go distribution from the official website.  
  
The Go distribution and tooling is available as an installer for all common operating systems. Visit <https://golang.org/dl/> to download to correct version for your OS. The installation instructions can be found
[here](https://golang.org/doc/install).  
  
#### 2. Download and install NPM.  
  
NPM and Node.js can be downloaded from their [official website](https://nodejs.org/en/download/). Simply select your  
distribution/OS and CPU architecture.  
  
#### 3. Install Wails  
  
The Molly Wallet is built using a light-weight framework for Desktop Applications using Go and VueJS. [Wails](https://github.com/wailsapp/wails) is very similar to Electron but does not packaging the full Chromium web browser as a dependency.  
  
  
###### MacOS  
  
Make sure you have the xcode command line tools installed. This can be done by running:  
`xcode-select --`  
  

###### Debian/Ubuntu  
  
`sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev`  
  
_Debian: 8, 9, 10_  
  
_Ubuntu: 16.04, 18.04, 19.04_  
  
_Also succesfully tested on: Zorin 15, Parrot 4.7, Linuxmint 19, Elementary 5, Kali, Neon_  
  
###### Arch Linux  
  
`sudo pacman -S webkit2gtk gtk3`  
  
_Also succesfully test on: Manjaro & ArcoLinux_  
  
###### Centos  
  
`sudo yum install webkitgtk3-devel gtk3-devel`  
  
_CentOS 6, 7_  
  
###### Fedora  
  
`sudo yum install webkit2gtk3-devel gtk3-devel`  
  
_Fedora 29, 30_  
  
###### VoidLinux & VoidLinux-musl  
  
`xbps-install gtk+3-devel webkit2gtk-devel`  
  
###### Gentoo  
  
`sudo emerge gtk+:3 webkit-gtk`  
  
###### Windows  
  
Windows requires gcc and related tooling. The recommended download is from [http://tdm-gcc.tdragon.net/download](http://tdm-gcc.tdragon.net/download). Once this is installed, you are good to go.  
  

  
**Ensure Go modules are enabled: GO111MODULE=on and go/bin is in your PATH variable.**  
  
Installation is as simple as running the following command:  
  
`go get -u github.com/wailsapp/wails/cmd/wails`  
  
  
##### 4. Clone this repository into your GOPATH  
  
`git clone git@github.com:StardustCollective/molly_wallet.git`  

---

### Want to contribute? Great!  
  
  
Molly Wallet uses Wails + Webpack for fast frontend development.  
Make a change in your file and instantaneously see your updates!  
  
Open your favorite Terminal and run these commands.  
  
In the constellation_wallet directory, run:  
```sh  
$ wails serve  
```  
  
In the frontend directory, run:  
```sh  
$ npm run serve  
```  
  
##### Alternatively:  
If you wish to compile the wallet, simply run:  
```sh  
wails build  
```  
or for the debug version, run:  
```sh  
wails build -d  
```  

---
