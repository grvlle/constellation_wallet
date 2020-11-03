package app

import (
	"errors"
	"io/ioutil"
	"os"
    "path/filepath"
    "runtime"
    "strconv"
    "strings"
)

func (a *WalletApplication) LoginJsonWallet(keystorePath, password string) (string, error) {

    if keystorePath == "" {
        a.LoginError("Please provide a path to the KeyStore file.")
        return "", errors.New("Please provide a path to the KeyStore file")
    }

    return a.loadJsonPrivateKey(keystorePath, password)
}

func (a *WalletApplication) MigrateWallet(keystorePath, keystorePassword, keyPassword, alias string) (string, error) {

    alias = strings.ToLower(alias)

    if runtime.GOOS == "windows" && !a.javaInstalled() {
        a.LoginError("Unable to detect your Java path. Please make sure that Java has been installed.")
        return "", errors.New("Unable to detect your Java path. Please make sure that Java has been installed")
    }

    if keystorePath == "" {
        a.LoginError("Please provide a path to the KeyStore file.")
        return "", errors.New("Please provide a path to the KeyStore file")
    }

    if !a.passwordsProvided(keystorePassword, keyPassword, alias) {
        a.log.Warnln("One or more passwords were not provided.")
        return "", errors.New("Unable to detect your Java path. Please make sure that Java has been installed")
    }

    os.Setenv("CL_STOREPASS", keystorePassword)
    os.Setenv("CL_KEYPASS", keyPassword)

    var pKeyFilePath string

    if runtime.GOOS == "windows"  {
        wKeyFilePath, err := os.Getwd()
        if err != nil {
            return "", errors.New("Unable to access the file system")
        }
        pKeyFilePath = filepath.Join(wKeyFilePath, "id_ecdsa.hex");
    } else {
        pKeyFilePath = filepath.Join(filepath.Dir(keystorePath), "id_ecdsa.hex")
    }

    a.log.Infoln("Migrate working directory - " + filepath.Dir(pKeyFilePath))

    if a.fileExists(pKeyFilePath) {
        err := os.Remove(pKeyFilePath)
        if err != nil {
            return "", errors.New("Unable to access the file system")
        }
    }

    _, err :=  a.producePrivateKeyMigrateV2(keystorePath, alias)
    if err != nil {
        return "", err
    }

    content, err := ioutil.ReadFile(pKeyFilePath)
    if err != nil {
        return "", errors.New("Unable to extract private key")
    }

    _ = os.Remove(pKeyFilePath)

    jsonKey, err := a.generateKeyStore(string(content), keystorePassword)
    if err != nil {
        return "", err
    }

    jsonKey, err = prettyPrint(jsonKey)
    if err != nil {
        return "", err
    }

    return a.saveMigrateKeyStoreFile(keystorePath, jsonKey)
}

func (a *WalletApplication) saveMigrateKeyStoreFile(p12FilePath string, jsonKey []byte) (string, error) {

    fullFilePath := filepath.Dir(p12FilePath) + string(os.PathSeparator) + strings.TrimSuffix(filepath.Base(p12FilePath),  filepath.Ext(p12FilePath)) + ".json"

    a.log.Info("Saving Migrate KeyStore File to: " + fullFilePath)

    if a.fileExists(fullFilePath) {
        a.log.Errorln("Private Key file already exists: ", fullFilePath)
        return "", errors.New("Private Key file already exists: " + fullFilePath)
    }

	err := ioutil.WriteFile(fullFilePath, jsonKey, 0644)
	if err != nil {
        a.log.Errorln("Unable to write file. Reason: ", err)
        return "", errors.New("Unable to write file: " + fullFilePath)
	}

    return fullFilePath, nil
}

// java -jar ~/.dag/cl-keytool.jar generate-wallet --keystore testA.p12 --alias alias --storepass test1 --keypass test2
// java -jar ~/.dag/cl-keytool.jar migrate-to-store-password-only --keystore testA.p12 --alias alias --storepass test1 --keypass test2
func (a *WalletApplication) producePrivateKeyMigrateV2(keystorePath, alias string) (bool, error) {

	err := a.runWalletCMD("keytool", "export-private-key-hex", "--keystore="+keystorePath, "--alias="+alias, "--env_args=true")
	if err != nil {
		s := err.Error()
		if strings.Contains(s, "java.io.IOException: keystore password was incorrect") {
			a.log.Errorln("Password is incorrect. Reason: ", err)
			return false, errors.New("Possibly wrong KeyStore Password")
		}
		if strings.Contains(s, "java.security.UnrecoverableKeyException:") {
			a.log.Errorln("Password is incorrect. Reason: ", err)
			return false, errors.New("Possibly wrong Key Password")
		}
		if strings.Contains(s, "java.lang.NullPointerException\n	at org.constellation.keytool.KeyStoreUtils$.$anonfun$unlockKeyPair$1") {
			a.log.Errorln("Alias is incorrect. Reason: ", err)
			return false, errors.New("Unable to find alias")
		}
		if strings.Contains(s, "java.io.IOException: toDerInputStream rejects tag type") {
			a.log.Errorln("Private key file type is incorrect.", err)
			return false, errors.New("Possibly wrong Private key file type")
		}
		a.log.Errorln("Unable to migrate Keystore file. Reason: ", err)
		return false, errors.New("Unable to migrate Keystore file to V2. Please report this issue")
	}

	return true, nil
}

func (a *WalletApplication) CreateKeyStoreFile(fileName, password string) (string, error) {

    var fullFilePath = a.paths.HomeDir + string(os.PathSeparator)  + fileName;

    // Check if a file with the same name exists
    i := 0
    suffix := ""
    for {
        if (i != 0) {
            suffix = strconv.Itoa(i)
        }
        if _, err := os.Stat(fullFilePath + suffix + ".json"); os.IsNotExist(err) {
            break
        }
        i ++
    }

    fullFilePath += suffix + ".json"

    jsonKey, err := a.generateNewKeyStore(password)
    if err != nil {
        return "", err
    }

    jsonKey, err = prettyPrint(jsonKey)
    if err != nil {
        return "", err
    }

    err = WriteToFile(fullFilePath, jsonKey)
    if err != nil {
        a.log.Errorln("Unable to write file. Reason: ", err)
        return "", errors.New("Unable to write file")
    }

    if (suffix != "") {
        fullFilePath += "."
    }

    return fullFilePath, nil
}
