<template>
  <div>
    <v-toolbar app flat
      color="white"
    >
      <h1 class="title ds-light-forecolor"><c-title>Log in</c-title></h1>
      <v-spacer></v-spacer>
      <c-menu>Log in</c-menu>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <v-form @submit.prevent="logIn">
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
          <v-btn type="submit">Log in</v-btn>
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

        usernameError: '',
        passwordError: '',
      };
    },
    created() {
      if (this.$global.loggedIn) {
        this.$router.push('/calendar');
      }
    },
    methods: {
      async logIn() {
        this.usernameError =
          this.username === '' ?
          'Please enter your username' : '';
        this.passwordError =
          this.password === '' ?
          'Please enter your username' : '';
        if (this.usernameError || this.passwordError || this.passwordCheckError) {
          return;
        }
        await this.$global.logIn(this.username, this.password);
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
