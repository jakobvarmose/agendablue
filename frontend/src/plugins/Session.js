import User from './User';

const loadUser = () =>
  User.fromString(localStorage.getItem('cryptouser'));

const saveUser = user =>
  localStorage.setItem('cryptouser', user.toString());

const removeUser = () =>
  localStorage.removeItem('cryptouser');

export default class Session {
  /** */
  static get userLoggedIn() {
    try {
      loadUser();
    } catch (_) {
      return false;
    }
    return true;
  }

  /** */
  static get userServer() {
    try {
      const user = loadUser();
      return user.server;
    } catch (_) {
      return null;
    }
  }

  /** */
  static get userUsername() {
    try {
      const user = loadUser();
      return user.username;
    } catch (_) {
      return null;
    }
  }

  /** */
  static async userGetData() {
    let user;
    try {
      user = loadUser();
    } catch (_) {
      return null;
    }
    const data = await user.getData();
    saveUser(user);
    return data;
  }

  static async userExportData() {
    const user = loadUser();
    return user.exportData();
  }

  static async userImportData(data) {
    const user = loadUser();
    return user.importData(data);
  }

  /** */
  static async userUpdatePassword(oldPassword, newPassword) {
    const user = loadUser();
    user.updatePassword(oldPassword, newPassword);
  }

  /** */
  static async userRegister(server, username, password, data) {
    const user = await User.register(server, username, password, data);
    saveUser(user);
  }

  /** */
  static async userLogIn(server, username, password) {
    const user = await User.logIn(server, username, password);
    saveUser(user);
  }

  /** */
  static async userLogOut() {
    removeUser();
  }

  /** */
  static async userSetData(data) {
    const user = loadUser();
    await user.setData(data);
    saveUser(user);
  }
}
