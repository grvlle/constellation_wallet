const { Crypto } = require("@peculiar/webcrypto");

const crypto = new Crypto();

delete self.crypto;
self.crypto = crypto;

delete self.localStorage;
