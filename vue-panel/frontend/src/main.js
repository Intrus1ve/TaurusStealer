import Vue from 'vue';
import VueRouter from 'vue-router';
Vue.use(VueRouter);

import Components from './plugins';
Vue.use(Components);

import App from './App.vue';

import LoginLayout from 'src/pages/Login/LoginLayout.vue';
import DashboardLayout from 'src/layout/DashboardLayout.vue';

import store from 'src/store'

const router = new VueRouter({
  routes:
  [
    {
      path: '/login',
      name: 'Login',
      meta: {layout: LoginLayout},
      component: () => import('src/pages/Login/Login.vue')
    },
    {
      path: '/',
      component: DashboardLayout,
      redirect: '/dashboard',
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          meta: {auth: true, isRoot: true},
          components: {
            default: () => import('src/pages/Dashboard/Dashboard.vue'),
            header: () => import('src/pages/Dashboard/DashboardHeader.vue')
          }
        },
        {
          path: 'logs',
          name: 'Logs',
          meta: {auth: true, isRoot: true},
          components: {
            default: () => import('src/pages/Logs/Logs.vue'),
            header: () => import('src/layout/DefaultHeader.vue')
          }
        },
        {
          path: 'grabber',
          name: 'Grabber',
          meta: {auth: true, isRoot: true},
          components: {
            default: () => import('src/pages/Grabber/Grabber.vue'),
            header: () => import('src/layout/DefaultHeader.vue')
          }
        },
        {
          path: 'loader',
          name: 'Loader',
          meta: {auth: true, isRoot: true},
          components: {
            default: () => import('src/pages/Loader/Loader.vue'),
            header: () => import('src/layout/DefaultHeader.vue')
          }
        },
        {
          path: 'settings/manage',
          name: 'Manage',
          meta: {auth: true, isRoot: true},
          components: {
            default: () => import('src/pages/Settings/Manage/Manage.vue'),
            header: () => import('src/layout/DefaultHeader.vue')
          }
        },
        {
          path: 'settings/config',
          name: 'Config',
          meta: {auth: true, isRoot: true},
          components: {
            default: () => import('src/pages/Settings/Config/Config.vue'),
            header: () => import('src/layout/DefaultHeader.vue')
          }
        },
      ]
    },
    {
      path: '/user/logs',
      name: 'UserLogs',
      meta: {layout: DashboardLayout, auth: true, isRoot: false},
      component: () => import('src/pages/UserLogs/UserLogs.vue')
    },
    { path: '*' }
  ],
  linkActiveClass: 'active'
});

router.beforeEach(async function (to, from, next) {
  const currentUser = store.getters.CURRENT_USER;
  const requireAuth = to.matched.some(record => record.meta.auth);
  if (requireAuth && !currentUser) {
    next('/login');
  } else {
    next();
  }

  if (currentUser)
  {
    const isRoot = store.getters.IS_ROOT;
    const requireRoot = to.matched.some(record => record.meta.isRoot);
    if (!isRoot && requireRoot) {
      next('/user/logs');
    } else {
      next();
    }
  }
})

new Vue({
  el: '#app',
  render: h => h(App),
  router,
  store
});
