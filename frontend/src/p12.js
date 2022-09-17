import forge from 'node-forge'

class KeyStoreFile {
    constructor() {
        this.privateKeyHexPrefix = "30818d020100301006072a8648ce3d020106052b8104000a047630740201010420";
        this.privateKeyHexPrefixLength = this.privateKeyHexPrefix.length;
        this.secp256k1StringIdentifier = 'a00706052b8104000aa144034200';
    }
    readP12(file, pass) {
        const p12Asn1 = forge.asn1.fromDer(file);
        const p12 = forge.pkcs12.pkcs12FromAsn1(p12Asn1, pass);
        const keyData = p12.getBags({ bagType: forge.pki.oids.pkcs8ShroudedKeyBag });
        const bag1 = keyData[forge.pki.oids.pkcs8ShroudedKeyBag][0];
        const hex = forge.asn1.toDer(bag1.asn1).toHex();
        const keyMinusPrefix = hex.substring(this.privateKeyHexPrefixLength, hex.length);
        const keyParts = keyMinusPrefix.split(this.secp256k1StringIdentifier);
        const privateKey = keyParts[0];
        const publicKey = keyParts[1];
        return { privateKey, publicKey };
    }
}

export const keyStoreFile = new KeyStoreFile();