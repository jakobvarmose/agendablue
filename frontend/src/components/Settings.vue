<template>
  <div v-if="$global.loggedIn">
    <v-toolbar app flat
      color="white"
    >
      <h1 class="title ds-light-forecolor">Settings</h1>
      <span style="display: none;"><c-title>Settings - Agenda.blue</c-title></span>
      <v-spacer></v-spacer>
      <c-menu>Settings</c-menu>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <v-text-field
          label="Username"
          v-model="username"
          readonly
        ></v-text-field>
        <v-select
          label="Locale"
          v-model="locale"
          :items="locales"
        ></v-select>
        <v-btn @click="downloadBackup">Download backup</v-btn>
        <v-btn @click="restoreFromBackup">Restore from backup</v-btn>
        <v-btn @click="passwordDialog_open">Change password</v-btn>
      </v-container>
    </v-content>
    <v-dialog v-model="importDialog">
      <v-card>
        <v-card-text>
          <v-container>
            <v-textarea
              label="Data"
              v-model="importDialog_data"
              ref="importDialog_data"
            ></v-textarea>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="importDialog_import">Import</v-btn>
          <v-btn @click="importDialog_cancel">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="exportDialog">
      <v-card>
        <v-card-text>
          <v-container>
            <v-textarea
              label="Data"
              readonly
              v-model="exportDialog_data"
              ref="exportDialog_data"
            ></v-textarea>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="exportDialog_close">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="passwordDialog">
      <v-card>
        <v-card-text>
          <v-container>
            <v-text-field
              label="Old password"
              type="password"
              v-model="passwordDialog_oldPassword"
              ref="passwordDialog_oldPassword"
            ></v-text-field>
            <v-text-field
              label="New password"
              type="password"
              v-model="passwordDialog_password"
            ></v-text-field>
            <v-text-field
              label="Repeat new password"
              type="password"
              v-model="passwordDialog_passwordCheck"
            ></v-text-field>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="passwordDialog_save" color="primary">Save</v-btn>
          <v-btn @click="passwordDialog_cancel">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
  <c-login v-else></c-login>
</template>
<script>
  import CLogin from '@/components/CLogin';
  import CMenu from '@/components/CMenu';
  import CTitle from '@/components/CTitle';

  require('js-file-manager/dist/jsfilemanager.min.js');

  const JSFileManager = window.JSFileManager;
  const JSFile = window.JSFile;

  const { Session } = require('@/plugins/cryptouser');

  export default {
    components: {
      CLogin,
      CMenu,
      CTitle,
    },
    data() {
      return {
        locale: this.$global.data ? String(this.$global.data.locale) : '',
        locales: ['', ...Object.keys(this.$dayspan.locales)],
        username: this.$global.username,
        password: '',
        newPassword: '',
        newPasswordCheck: '',

        importDialog: false,
        importDialog_data: '',
        exportDialog: false,
        exportDialog_data: '',
        passwordDialog: false,
        passwordDialog_oldPassword: '',
        passwordDialog_password: '',
        passwordDialog_passwordCheck: '',

        data: null,
      };
    },
    methods: {
      async importDialog_open() {
        this.importDialog = true;
        this.$nextTick(() => {
          this.$refs.importDialog_data.focus();
        });
      },
      async importDialog_import() {
        this.importDialog = false;
        this.$global.importData(Buffer.from(this.importDialog_data, 'base64'));
        this.importDialog_data = '';
      },
      async importDialog_cancel() {
        this.importDialog = false;
        this.importDialog_data = '';
      },
      async exportDialog_open() {
        this.exportDialog_data = (await this.$global.exportData()).toString('base64');
        this.exportDialog = true;
        this.$nextTick(() => {
          this.$refs.exportDialog_data.focus();
        });
      },
      async exportDialog_close() {
        this.exportDialog = false;
        this.exportDialog_data = '';
      },
      async passwordDialog_open() {
        this.passwordDialog = true;
        this.$nextTick(() => {
          this.$refs.passwordDialog_oldPassword.focus();
        });
      },
      async passwordDialog_save() {
        if (this.passwordDialog_passwordCheck !== this.passwordDialog_password) {
          alert('passwords don\'t match');
          return;
        }
        await this.$global.updatePassword(
          this.passwordDialog_oldPassword, this.passwordDialog_password);

        this.passwordDialog = false;
        this.passwordDialog_oldPassword = '';
        this.passwordDialog_password = '';
        this.passwordDialog_passwordCheck = '';
      },
      async passwordDialog_cancel() {
        this.passwordDialog = false;
        this.passwordDialog_oldPassword = '';
        this.passwordDialog_password = '';
        this.passwordDialog_passwordCheck = '';
      },
      async downloadBackup() {
        const exportData = await this.$global.exportData();
        const filename = `${this.$global.username}-${new Date().toISOString().replace(/:/g, '-')}.cbor`;
        const file = new JSFile(exportData, filename);
        file.save();
      },
      async restoreFromBackup() {
        const file = await JSFileManager.pick({});
        try {
          const buf = await file.getArrayBuffer();
          await this.$global.importData(Buffer.from(buf));
        } catch (err) {
          alert(err);
        }
      },
    },
    watch: {
      async locale(val) {
        this.$dayspan.setLocale(val);
        this.$global.data.locale = val;
        await Session.userSetData(this.$global.data);
      },
      // eslint-disable-next-line
      '$global.data'(val) {
        if (val) {
          this.locale = String(val.locale);
        } else {
          this.locale = '';
        }
      },
    },
  };
</script>
<style>
  .v-btn--flat, .v-text-field--solo .v-input__slot {
    background-color: #f5f5f5 !important;
    margin-bottom: 8px !important;
  }
</style>
