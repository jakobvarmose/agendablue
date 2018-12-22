<template>
  <div>
    <v-toolbar app flat
      color="white"
    >
      <h1 class="title ds-light-forecolor"><c-title>Agenda.blue</c-title></h1>
      <v-spacer></v-spacer>
      <c-menu>Home</c-menu>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <p>Agenda.blue is an encrypted calendar.</p>
        <p>
          <v-btn color="primary" @click="$router.push('/register');">Sign up</v-btn>
          <v-btn @click="$router.push('/login');">Log in</v-btn>
        </p>
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
    created() {
      if (this.$global.loggedIn) {
        this.$router.push('/calendar');
      }
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
