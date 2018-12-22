const axios = require('axios');
const nacl = require('tweetnacl');
const scrypt = require('scrypt-async');
const querystring = require('querystring');
require('../borc.min.js');

const cbor = window.Borc;

const computeURL = (server, command) => {
  let protocol;
  if (typeof window === 'undefined') {
    if (/^localhost:[0-9]+$/.test(server)) {
      protocol = 'http:';
    } else {
      protocol = 'https:';
    }
  } else {
    protocol = document.location.protocol;
    if (server.match(/^\[.+\]|^[^:]+/)[0] !== document.location.hostname) {
      protocol = 'https:';
    }
  }
  return `${protocol}//${server}/api/v0/${command}`;
};

const computeMasterKey = async (password, info) => {
  if (info.version !== '1') {
    throw new Error('Unsupported info version');
  }

  return new Promise((resolve, reject) => {
    try {
      scrypt(`cryptouser:${password}`, info.salt, {
        logN: 16,
        r: 8,
        p: 1,
        dkLen: 64,
        encoding:
        'binary',
      }, resolve);
    } catch (err) {
      reject(err);
    }
  });
};

const computeAccessKey = masterKey =>
  masterKey.slice(32, 64);
const computeEncryptionKey = masterKey =>
  masterKey.slice(0, 32);


const api = async (key, server, action, obj) => {
  if (!key) {
    // eslint-disable-next-line
    key = nacl.sign.keyPair.fromSeed(new Buffer(32));
  }

  if (!server) {
    // eslint-disable-next-line
    server = document.location.host;
  }

  // eslint-disable-next-line
  obj.time = new Date().toISOString();
  // eslint-disable-next-line
  obj.action = action;

  const message = cbor.encode(obj);
  const signedMessage = nacl.sign(message, key.secretKey);

  const res = await axios.post(computeURL(server, 'signed'), querystring.stringify({
    data: cbor.encode({
      signedMessage,
      publicKey: key.publicKey,
    }).toString('base64')
      .replace(/\+/g, '-')
      .replace(/\//g, '_')
      .replace(/=/g, ''),
  }));

  const val = cbor.decode(Buffer.from(res.data, 'base64'));

  return val;
};

const encrypt = (data, encryptionKey) => {
  const plaintext = cbor.encode(data);
  const nonce = nacl.randomBytes(nacl.secretbox.nonceLength);
  const ciphertext = Buffer.concat([
    Buffer.from(nonce),
    Buffer.from(nacl.secretbox(plaintext, nonce, encryptionKey)),
  ]);
  return ciphertext;
};

const decrypt = (ciphertext, encryptionKey) => {
  const plaintext = nacl.secretbox.open(
    ciphertext.slice(24),
    ciphertext.slice(0, 24),
    encryptionKey,
  );
  return cbor.decode(plaintext);
};

const priv = new WeakMap();

/**
 * Represents a user.
 */
class User {
  /**
   * New instances should be created by calling register() or logIn(),
   * calling this constructor won't work.
   */
  constructor() {
    throw new Error();
  }

  get server() {
    const user = priv.get(this);
    return user.server;
  }

  get username() {
    const user = priv.get(this);
    return user.username;
  }

  /**
   * Get the user data.
   * @return {string}
   */
  async getData() {
    const user = priv.get(this);

    await this.refreshData();

    return user.content;
  }

  /**
   * Refresh the use data.
   * @return {string}
   */
  async refreshData() {
    const user = priv.get(this);

    const contentKey = nacl.sign.keyPair.fromSeed(user.contentKeySeed);

    const res = await api(contentKey, user.server, 'readUserContent', {
      username: user.username,
    });

    user.content = decrypt(res.content, user.contentEncryptionKey);
  }

  /**
   * Update the user data.
   * @param {string} data The new user data.
   */
  async setData(data) {
    const user = priv.get(this);

    const contentKey = nacl.sign.keyPair.fromSeed(user.contentKeySeed);

    await api(contentKey, user.server, 'updateUserContent', {
      username: user.username,
      content: encrypt(data, user.contentEncryptionKey),
    });

    user.content = data;
  }

  /**
   * Convert user to a string.
   */
  toString() {
    return this.toBuffer().toString('base64');
  }

  toBuffer() {
    const user = priv.get(this);
    return cbor.encode({
      server: user.server,
      username: user.username,
      contentKeySeed: user.contentKeySeed,
      contentEncryptionKey: user.contentEncryptionKey,
    });
  }

  /**
   * Restore a user from a string.
   * @param {string} str
   */
  static fromString(str) {
    return this.fromBuffer(Buffer.from(str, 'base64'));
  }

  static fromBuffer(buf) {
    let obj;
    try {
      obj = cbor.decode(buf);
    } catch (ex) {
      return null;
    }
    if (!obj) {
      return null;
    }
    const user = Object.create(User.prototype, {});
    priv.set(user, obj);
    return user;
  }

  async exportData() {
    const data = await this.getData();
    if (!data) {
      throw new Error('no data');
    }
    return cbor.encode(data);
  }

  async importData(data) {
    const data2 = cbor.decode(data);
    if (data2.version !== '1') {
      throw new Error('Invalid version');
    }
    await this.setData(data2);
  }

  /**
   * Log in as an existing user.
   * @param {string} server
   * @param {string} username
   * @param {string} password
   */
  static async logIn(server, username, password) {
    const res1 = await api(null, server, 'readUserInfo', {
      username,
    });

    const info = cbor.decode(res1.info);

    const masterKey = await computeMasterKey(password, info);
    const accessKey = nacl.sign.keyPair.fromSeed(computeAccessKey(masterKey));

    const res2 = await api(accessKey, server, 'readUserBootstrap', {
      username,
    });

    const encryptionKey = computeEncryptionKey(masterKey);

    const bootstrap = decrypt(res2.bootstrap, encryptionKey);

    const content = decrypt(res2.content, bootstrap.contentEncryptionKey);

    const user = Object.create(User.prototype, {});
    priv.set(user, {
      server,
      username,
      contentKeySeed: bootstrap.contentKeySeed,
      contentEncryptionKey: bootstrap.contentEncryptionKey,
      content,
    });
    return user;
  }

  async updatePassword(oldPassword, newPassword) {
    const user = priv.get(this);

    const res1 = await api(null, user.server, 'readUserInfo', {
      username: user.username,
    });
    const oldInfo = cbor.decode(res1.info);
    const oldMasterKey = await computeMasterKey(oldPassword, oldInfo);
    const oldAccessKey = nacl.sign.keyPair.fromSeed(computeAccessKey(oldMasterKey));

    const newInfo = {
      salt: nacl.randomBytes(16),
      version: '1',
    };
    const newBootstrap = {
      contentKeySeed: user.contentKeySeed,
      contentEncryptionKey: user.contentEncryptionKey,
    };
    const newMasterKey = await computeMasterKey(newPassword, newInfo);
    const newAccessKey = nacl.sign.keyPair.fromSeed(computeAccessKey(newMasterKey));
    const newEncryptionKey = computeEncryptionKey(newMasterKey);

    await api(oldAccessKey, user.server, 'updateUserBootstrap', {
      username: user.username,
      accessKey: newAccessKey.publicKey,
      info: cbor.encode(newInfo),
      bootstrap: encrypt(newBootstrap, newEncryptionKey),
    });
  }

  /**
   * Register a new user.
   * @param {string} server
   * @param {string} username
   * @param {string} password
   * @param {string} data
   */
  static async register(server, username, password, content) {
    const info = {
      salt: nacl.randomBytes(16),
      version: '1',
    };
    const masterKey = await computeMasterKey(password, info);
    const accessKey = nacl.sign.keyPair.fromSeed(computeAccessKey(masterKey));
    const encryptionKey = computeEncryptionKey(masterKey);
    const contentKeySeed = nacl.randomBytes(nacl.sign.seedLength);
    const contentKey = nacl.sign.keyPair.fromSeed(contentKeySeed);
    const contentEncryptionKey = nacl.randomBytes(nacl.secretbox.keyLength);
    const bootstrap = {
      contentKeySeed,
      contentEncryptionKey,
    };
    const ciphertext1 = encrypt(bootstrap, encryptionKey);
    const ciphertext2 = encrypt(content, contentEncryptionKey);

    await api(null, server, 'createUser', {
      username,
      accessKey: accessKey.publicKey,
      info: cbor.encode(info),
      bootstrap: ciphertext1,
      contentKey: contentKey.publicKey,
      content: ciphertext2,
    });

    const user = Object.create(User.prototype, {});
    priv.set(user, {
      server,
      username,
      contentKeySeed,
      contentEncryptionKey,
      content,
    });
    return user;
  }
}

export default User;
