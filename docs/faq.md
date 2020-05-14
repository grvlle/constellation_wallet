# Molly Wallet FAQ

This is an attempt to address the most Frequently Asked Questions around the Molly Wallet in hopes of easing the community into leveraging the Desktop Wallet for $DAG to securely store their tokens.

The Molly Wallet is still in the early phases of development with the first stable, mainnet connected build released to the general public mere days ago (at the time of writing this). This means that the way we're distributing the application, and some of the manual steps required will be improved upon with time. 

## General Questions

#### Q: Where can I download the Molly Wallet?
https://github.com/grvlle/constellation_wallet/releases

#### Q: Why can't I login to my wallet?
You can always reset the Molly Wallet, simply by closing the application, removing your `.dag` folder, starting the application and re-importing your key. See the below sections under your OS of choice for how to locate the `.dag` folder.

#### Q: What does the IMPORT feature do on the login screen?

The Import feature exists to make it possible to access your wallet key from different computers. This only has to be done once on any computer, and it'll sync your transactions with the mainnet, and update the wallet with your previous transactions. 

Once imported, you will access your funds by simply login into the wallet.

#### Q: How do I create a new wallet?

Select the create wallet button on the login screen and it'll let you browse to a location in which you wish to store your keyfile.p12. Remember to give it a name, and select save.

After that you need to provide the keyfile with authentication, populate the remaining fields and click create. Then the keyfile.p12 will be saved to the location you specified, and you can use that to access your funds through Molly Wallet.


I will organize the questions based on the OS the wallet is installed upon.

## Windows

#### Q: Why am I getting an error when sending transaction?

**A:** This is most likely because you've used an older testnet build of the wallet with testnet artifacts interfering with your mainnet wallet. In order to fix this you need to remove the `.dag` folder that is located in your `C:/Users/<username>/` directory. Once done, feel free to import your wallet again and you should be good to go.

#### Q: I am getting `Unable to detect your Java path, make sure Java has been installed` when trying to login or create/import a wallet.

***A:*** This is either because [OpenJDK v9](https://java.com/) and [JRE](https://www.oracle.com/java/technologies/javase-jre8-downloads.html) hasn't been installed, or because the enviornment paths haven't been set up correctly.

For Molly Wallet to be able to detect the installation directory of Java, the `JAVA_HOME` enviornmental variable needs t  
o be set. The Java path also has to be included in the `Path` variable.  
  
Search for enviornment variables in windows search and select *Edit the system enviornment variables*, then set the JAVA  
paths like in the below image.  
  
[![env](https://i.ibb.co/Br1M31s/envvars.png)](https://constellationnetwork.io/technology/molly-wallet/) 

#### Q: I am running Windows 7 (or earlier) - can I run Molly Wallet?

**A:** The answer is maybe. It's not officially supported, so probably not.

## MacOS

#### Q: Why am I getting an error when sending transaction?

**A:** This is most likely because you've used an older testnet build of the wallet with testnet artifacts interfering with your mainnet wallet. In order to fix this you need to remove the `.dag` folder that is located in your `$HOME` (if you don't know what this is, see [this article](https://www.cnet.com/how-to/how-to-find-your-macs-home-folder-and-add-it-to-finder/)) path. Once done, feel free to import your wallet again and you should be good to go.

#### Q: Why am I getting `The application 'Molly - Constellation Desktop Wallet` can't be opened?

![env](https://i.ibb.co/VWw30HN/a123555f-0881-4ae8-9b1d-7dd36d4d6802.jpg)


**A:** This is because the application is compressed and archived. You need to use a software called [Unarchiver](https://theunarchiver.com/). 

After you've downloaded that, right click on the .zip file and select "open with" and then choose "Unarchiver" instead of the default program. It'll unzip and you'll be able to run it.

#### Q: Why is my Molly Wallet download not downloaded as a .zip file?

**A:** On some systems, Safari will automatically unzip the contents. We do not want that, we want to leverage the [Unarchiver](https://theunarchiver.com/). Here's a guide [how to disable automatic unzipping](https://www.addictivetips.com/mac-os/stop-automatically-unzipping-files-in-safari/) of files in Safari. Once that has been disabled, redownload Molly Wallet from the [official website](https://constellationnetwork.io/technology/molly-wallet/) and use [Unarchiver](https://theunarchiver.com/) to extract the contents.

## Linux
#### Q: Why does it say `Alias not found` when trying to import/login/create wallet?

**A:** This is probably due to Java missing on the system. For now, I've only implemented means of detecting that on Windows. So please, if you're using a debian based distribution, download `openjdk-8` from aptitude.

#### Q: Why am I getting an error when sending transaction?

**A:** This is most likely because you've used an older testnet build of the wallet with testnet artifacts interfering with your mainnet wallet. In order to fix this you need to remove the `.dag` folder that is located in your `$HOME` path. Once done, feel free to import your wallet again and you should be good to go.

## Java

#### Q: Why is Java needed?

**A:** The wallet binary that I'm integrating Molly Wallet against has been built by the Constellation Core team in a programming language called *Scala*. This is a functional programming language that runs on JVM (Java virtual machine). Thus Java becomes a dependency.

