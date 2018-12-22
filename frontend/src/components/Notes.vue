<template>
  <div v-if="$global.loggedIn">
    <v-toolbar app flat
      color="white"
    >
      <h1 class="title ds-light-forecolor">Notes</h1>
      <span style="display: none;"><c-title>Notes - Agenda.blue</c-title></span>
      <v-spacer></v-spacer>
      <c-menu>Notes</c-menu>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <v-btn @click=createNoteDialog_open>New note</v-btn>
      </v-container>
    </v-content>
    <v-dialog v-model="createNoteDialog">
      <v-card>
        <v-card-title>
          <span class="headline">New note</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-text-field
              label="Title"
              v-model="createNoteDialog_title"
            ></v-text-field>
            <v-textarea
              label="Note"
              v-model="createNoteDialog_note"
            ></v-textarea>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn @click="createNoteDialog_save">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
  <c-login v-else></c-login>
</template>
<script>
  import CTitle from '@/components/CTitle';
  import CMenu from '@/components/CMenu';

  export default {
    components: {
      CTitle,
      CMenu,
    },
    data() {
      return {
        notes: [],
        createNoteDialog: false,
      };
    },
    methods: {
      createNoteDialog_open() {
        this.createNoteDialog = true;
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
