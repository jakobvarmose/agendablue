import Vue from 'vue';
import Router from 'vue-router';
import Index from '@/components/Index';
import Calendar from '@/components/Calendar';
import Settings from '@/components/Settings';
import Register from '@/components/Register';
import CLogin from '@/components/CLogin';
import Notes from '@/components/Notes';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      component: Index,
    },
    {
      path: '/calendar',
      component: Calendar,
    },
    {
      path: '/notes',
      component: Notes,
    },
    {
      path: '/settings',
      component: Settings,
    },
    {
      path: '/register',
      component: Register,
    },
    {
      path: '/login',
      component: CLogin,
    },
  ],
});
