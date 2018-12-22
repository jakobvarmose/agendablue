<template>
  <v-menu>
    <v-btn flat slot="activator">
      <slot></slot>
    </v-btn>
    <v-list v-if="$global.loggedIn">
      <v-list-tile @click="$router.push('/calendar')">
        <v-list-tile-content>
          <v-list-tile-title>
            Calendar
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <!--<v-list-tile @click="$router.push('/notes')">
        <v-list-tile-content>
          <v-list-tile-title>
            Notes
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>-->
      <v-list-tile @click="$router.push('/settings')">
        <v-list-tile-content>
          <v-list-tile-title>
            Settings
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile @click="$global.logOut">
        <v-list-tile-content>
          <v-list-tile-title>
            Log out
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
    <v-list v-else>
      <v-list-tile @click="$router.push('/')">
        <v-list-tile-content>
          <v-list-tile-title>
            Home
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile @click="$router.push('/login')">
        <v-list-tile-content>
          <v-list-tile-title>
            Log in
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile @click="$router.push('/register')">
        <v-list-tile-content>
          <v-list-tile-title>
            Sign up
          </v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </v-menu>
</template>
<script>
  export default {};
</script>
