# Molly Wallet ($DAG desktop wallet)

[![N|Solid](https://i.ibb.co/pKC9WMs/2020-02-06-1649x916-scrot.png)](https://constellationnetwork.io/technology/molly-wallet/)

The Molly Wallet is the official $DAG wallet of the Constellation Network. It'll let users interact with the Hypergraph Network in various ways, not limited to producing $DAG transactions. The wallet is currently under development, the first public release has been shipped with more to come.

### Here's how the Constellation Team describes it on their website

>Created by a Constellation community member, for the best community in crypto, and named after the wallet designer’s daughter, Molly is built with a vision to set >future standards in digital commerce.

>Just like your wallet that you use daily to pay for lunch, dinner, clothes, that hold your business cards, ID’s, and that lucky penny, our intention was to design something that people would use everyday. We live in an age where we have an abundance of information, from videos to blogs, that are at our fingertips. Digital commerce is molding how generations interact with one another, consume, and educate themselves. The nature of this wallet is digital and thus we wanted to expand the possibilities of what a wallet could be and do because of its multimedia capabilities. We wanted to make it not only user friendly, but create the space for applications to be developed by our vibrant open source community.

>A true articulation of open innovation. 

>A cryptocurrency wallet is one of the most cherished pieces of technology to anyone that holds cryptocurrency. Yet many wallets are visually unappealing, have limited functionality and utility, and are treated as an afterthought. We wanted to bring it to the forefront and reimagine the wallet, the community that uses the wallet, and enable a future that will be powered by DAG.

[- Official Source](https://constellationnetwork.io/technology/molly-wallet/)

### Technologies

Molly Wallet uses a number of open source projects to work properly:

* [VueJS](https://vuejs.org) - The Progressive JavaScript Framework
* [Go](https://golang.org) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [GORM](https://gorm.io) - Object Relational Mapping for Go.
* [Wails](https://wails.app/) - A framework for building desktop apps using Go & Web Technologies

### Installation

Molly Wallet requires [Java](https://java.com/) to run.

Simply download from the [official website](https://constellationnetwork.io/technology/molly-wallet/) for the selected OS.

### Development Enviornment

##### 1. Download the Go distribution from the official website.

   The Go distribution and tooling is available as an installer for all common operating systems.
   Visit <https://golang.org/dl/> to download to correct version for your OS. The installation instructions can be found [here](https://golang.org/doc/install).

##### 2. Download and install NPM.

   NPM and Node.js can be downloaded from their [official website](https://nodejs.org/en/download/). Simply select your distribution/OS and CPU architecture. 

##### 3. Install Wails
  
   The Molly Wallet is build using a light-weight framework for Desktop Applications using Go and VueJS. Wails is very similar to Electron but is not packaging the full Chromium web browser as a dependency.


 ###### MacOS

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Make sure you have the xcode command line tools installed. This can be done by running:
   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `xcode-select --`
   
   ##### Linux

   ###### Debian/Ubuntu

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   `sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   _Debian: 8, 9, 10_

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   _Ubuntu: 16.04, 18.04, 19.04_

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   _Also succesfully tested on: Zorin 15, Parrot 4.7, Linuxmint 19, Elementary 5, Kali, Neon_

   ###### Arch Linux

   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`sudo pacman -S webkit2gtk gtk3`

   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;_Also succesfully test on: Manjaro & ArcoLinux_

   ###### Centos

   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`sudo yum install webkitgtk3-devel gtk3-devel`

   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;_CentOS 6, 7_

   ###### Fedora

   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`sudo yum install webkit2gtk3-devel gtk3-devel`

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; _Fedora 29, 30_
 
   ###### VoidLinux & VoidLinux-musl

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `xbps-install gtk+3-devel webkit2gtk-devel`

   ###### Gentoo

 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  `sudo emerge gtk+:3 webkit-gtk`

   ##### Windows

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Windows requires gcc and related tooling. The recommended download is from [http://tdm-gcc.tdragon.net/download](http://tdm-gcc.tdragon.net/download). Once this is installed, you are good to go.

   ##### Installation

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; **Ensure Go modules are enabled: GO111MODULE=on and go/bin is in your PATH variable.**

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; Installation is as simple as running the following command:
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ```
   go get -u github.com/wailsapp/wails/cmd/wails
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ```
---

##### 4. Clone this repository into your GOPATH

   `git clone https://github.com/grvlle/constellation_wallet.git`

### Development

Want to contribute? Great!

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

If you wish to compile the wallet, simply run:
```sh
wails build
```
or for the debug version, run:
```sh
wails build -d
```
