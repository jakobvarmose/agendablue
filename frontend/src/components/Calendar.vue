<template>
  <ds-calendar-app
    v-if="calendar"
    :calendar="calendar"
    @change="saveState"
    ref="cal"
  >
    <template slot="title">
      Agenda.blue
    </template>
    <template slot="menuRight">
      <span style="display: none;"><c-title>Calendar - Agenda.blue</c-title></span>
      <c-menu>Calendar</c-menu>
    </template>
    <template slot="eventPopover" slot-scope="slotData">
      <ds-calendar-event-popover
        v-bind="slotData"
        @finish="saveState"
      ></ds-calendar-event-popover>
    </template>
    <template slot="eventCreatePopover" slot-scope="{placeholder, calendar, close}">
      <ds-calendar-event-create-popover 
        :calendar-event="placeholder"
        :calendar="calendar"
        :close="$refs.cal.$refs.calendar.clearPlaceholder"
        @create-edit="$refs.cal.editPlaceholder"
        @create-popover-closed="saveState"
        >
      </ds-calendar-event-create-popover>
    </template>
  </ds-calendar-app>
  <div v-else-if="$global.loggedIn"></div>
  <c-login v-else></c-login>
</template>
<script>
  import CLogin from '@/components/CLogin';
  import CMenu from '@/components/CMenu';
  import CTitle from '@/components/CTitle';
  import { Calendar } from 'dayspan';

  const { Session } = require('@/plugins/cryptouser');

  export default {
    components: {
      CLogin,
      CMenu,
      CTitle,
    },
    data() {
      let calendar = null;
      if (this.$global.data) {
        this.$dayspan.setLocale(this.$global.data.locale);
        calendar = Calendar.months();
        calendar.addEvents(this.$global.data.events);
      }
      return {
        calendar,
      };
    },
    methods: {
      async saveState() {
        const state = this.calendar.toInput(true);
        this.$global.data.events = state.events;
        await Session.userSetData(this.$global.data);
      },
    },
    watch: {
      // eslint-disable-next-line
      '$global.data'(val) {
        if (val) {
          this.$dayspan.setLocale(val.locale);
          this.calendar = Calendar.months();
          this.calendar.addEvents(val.events);
        } else {
          this.calendar = null;
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
