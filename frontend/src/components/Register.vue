<template>
  <div>
    <v-toolbar app flat
      color="white"
    >
      <h1 class="title ds-light-forecolor">Sign up</h1>
      <span style="display: none;"><c-title>Sign up - Agenda.blue</c-title></span>
      <v-spacer></v-spacer>
      <c-menu>Sign up</c-menu>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <v-form @submit.prevent="register">
          <v-text-field
            label="Username"
            v-model="username"
            :error-messages="usernameError"
          ></v-text-field>
          <v-text-field
            label="Password"
            type="password"
            v-model="password"
            :error-messages="passwordError"
          ></v-text-field>
          <v-text-field
            label="Repeat password"
            type="password"
            v-model="passwordCheck"
            :error-messages="passwordCheckError"
          ></v-text-field>
          <v-btn type="submit">Sign up</v-btn>
        </v-form>
      </v-container>
    </v-content>
  </div>
</template>
<script>
  import CMenu from '@/components/CMenu';
  import CTitle from '@/components/CTitle';

  export default {
    components: {
      CMenu,
      CTitle,
    },
    data() {
      return {
        username: '',
        password: '',
        passwordCheck: '',

        usernameError: '',
        passwordError: '',
        passwordCheckError: '',
      };
    },
    created() {
      if (this.$global.loggedIn) {
        this.$router.push('/calendar');
      }
    },
    methods: {
      async register() {
        this.usernameError =
          this.username === '' ?
          'Please enter a username' : '';
        this.passwordError =
          this.password === '' ?
          'Please enter a username' : '';
        this.passwordCheckError =
          this.passwordCheck !== this.password ?
          'Passwords do not match' : '';
        if (this.usernameError || this.passwordError || this.passwordCheckError) {
          return;
        }
        await this.$global.register(this.username, this.password);
      },
    },
    watch: {
      // eslint-disable-next-line
      '$global.loggedIn'(val) {
        if (val) {
          this.$router.push('/calendar');
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
  .error--text .v-label {
    animation: initial !important;
  }
</style>
