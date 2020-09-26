# Molly Wallet ($DAG desktop wallet)  
  
![N|Solid](https://i.ibb.co/QXvTVR0/mollywallet.jpg)
  
The Molly Wallet is the official $DAG wallet of the Constellation Network. It'll let users interact with the Hypergraph Network in various ways, not limited to producing $DAG transactions. The wallet is currently under development, the first public release has been shipped with more to come.  
  
### Here's how the Constellation Team describes it on their website  
  
>Created by a Constellation community member, for the best community in crypto, and named after the wallet designer’s daughter, Molly is built with a vision to set future standards in digital commerce.  
  
>Just like your wallet that you use daily to pay for lunch, dinner, clothes, that hold your business cards, ID’s, and that lucky penny, our intention was to design something that people would use everyday. We live in an age where we have anabundance of information, from videos to blogs, that are at our fingertips. Digital commerce is molding how generationsinteract with one another, consume, and educate themselves. The nature of this wallet is digital and thus we wanted to expand the possibilities of what a wallet could be and do because of its multimedia capabilities. We wanted to make it not only user friendly, but create the space for applications to be developed by our vibrant open source community.  
  
>A true articulation of open innovation.  
  
>A cryptocurrency wallet is one of the most cherished pieces of technology to anyone that holds cryptocurrency. Yet many wallets are visually unappealing, have limited functionality and utility, and are treated as an afterthought. We wanted to bring it to the forefront and reimagine the wallet, the community that uses the wallet, and enable a future that will be powered by DAG.  
  
[- Official Source](https://constellationnetwork.io/technology/molly-wallet/)  
  
---  

### Publications  
  
* [Constellation’s Molly a Landmark Moment for the $DAG Ecosystem](https://thedailychain.com/constellations-molly-a-landmark-moment-for-the-dag-ecosystem/) - An article written on The Daily Chain by Anna Larsen.  

---  

### Technologies  
  
Molly Wallet uses a number of open source projects to work properly:  
  
* [VueJS](https://vuejs.org) - The Progressive JavaScript Framework  
* [Go](https://golang.org) - Go is an open source programming language that makes it easy to build simple, reliable, and  
efficient software.  
* [GORM](https://gorm.io) - Object Relational Mapping for Go.  
* [Wails](https://wails.app/) - A framework for building desktop apps using Go & Web Technologies  

---    

### Installation  
  
#### Pre-requisits  
Molly Wallet requires [OpenJDK v9](https://java.com/) and [JRE](https://www.oracle.com/java/technologies/javase-jre8-downloads.html) to run. If installing the MacOS version, use [Unarchiver](https://theunarchiver.com/) to unzip molly_installer.zip
  
#### Download Molly Wallet  
The latest builds can be found under [releases](https://github.com/grvlle/constellation_wallet/releases).  
  
---  
  
### FAQ
The Molly Wallet FAQ can be located under [docs](https://github.com/grvlle/constellation_wallet/blob/develop/docs/faq.md#molly-wallet-faq).

---  

### Development Enviornment  
  
#### 1. Download the Go distribution from the official website.  
  
The Go distribution and tooling is available as an installer for all common operating systems. Visit <https://golang.org/dl/> to download to correct version for your OS. The installation instructions can be found
[here](https://golang.org/doc/install).  
  
#### 2. Download and install NPM.  
  
NPM and Node.js can be downloaded from their [official website](https://nodejs.org/en/download/). Simply select your  
distribution/OS and CPU architecture.  
  
#### 3. Install Wails  
  
The Molly Wallet is build using a light-weight framework for Desktop Applications using Go and VueJS. [Wails](https://github.com/wailsapp/wails) is very similar to Electron but is not packaging the full Chromium web browser as a dependency.  
  
  
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
  
`git clone https://github.com/grvlle/constellation_wallet.git`  

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

### Top Contributers  
* [digitaltwin](https://github.com/digitaltwinnn)  
* [Marcin Wadoń](https://github.com/marcinwadon)
* [junkai121](https://github.com/junkai121)

---
