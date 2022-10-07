<template>
  <div class="" :class="{'nav-open': $sidebar.showSidebar}">
    <side-bar>
      <template slot="links">
        <user-menu :title="currentUser"></user-menu>
        <sidebar-item :link="{name: 'Dashboard', icon: 'far fa-chart-scatter', path: '/dashboard'}"/>
        <sidebar-item :link="{name: 'Logs', icon: 'far fa-copy', path: '/logs'}"/>
        <sidebar-item :link="{name: 'Grabber', icon: 'far fa-folder-open', path: '/grabber'}"/>
        <sidebar-item :link="{name: 'Loader', icon: 'fal fa-file-download', path: '/loader'}"/>
        <sidebar-item :link="{name: 'Settings', icon: 'fal fa-cogs'}">
          <sidebar-item :link="{name: 'Config', path: '/settings/config'}"></sidebar-item>
          <sidebar-item :link="{name: 'Manage', path: '/settings/manage'}"></sidebar-item>
        </sidebar-item>
      </template>
    </side-bar>
    <div class="main-panel">
      <top-navbar></top-navbar>
      <router-view name="header"></router-view>
      <div :class="{content: !$route.meta.hideContent}" @click="toggleSidebar">
        <zoom-center-transition :duration="200" mode="out-in">
          <!-- your content here -->
          <router-view></router-view>
        </zoom-center-transition>
      </div>
      <content-footer v-if="!$route.meta.hideFooter"></content-footer>
    </div>
  </div>
</template>

<script>
/* eslint-disable no-new */
import PerfectScrollbar from 'perfect-scrollbar';
import 'perfect-scrollbar/css/perfect-scrollbar.css';

function hasElement(className) {
  return document.getElementsByClassName(className).length > 0;
}

function initScrollbar(className) {
  if (hasElement(className)) {
    new PerfectScrollbar(`.${className}`);
  } else {
    // try to init it later in case this component is loaded async
    setTimeout(() => {
      initScrollbar(className);
    }, 100);
  }
}

import TopNavbar from './TopNavbar.vue';
import ContentFooter from './ContentFooter.vue';
import DashboardContent from './Content.vue';
import UserMenu from './UserMenu.vue';
import { SlideYDownTransition, ZoomCenterTransition } from 'vue2-transitions';
import store from 'src/store'

export default {
  components: {
    TopNavbar,
    ContentFooter,
    DashboardContent,
    UserMenu,
    SlideYDownTransition,
    ZoomCenterTransition
  },
  data() {
    return {
      currentUser: '',
    }
  },
  methods: {
    toggleSidebar() {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.displaySidebar(false);
      }
    }
  },
  mounted() {
    this.currentUser = this.$store.getters.CURRENT_USER;
    let docClasses = document.body.classList;
    let isWindows = navigator.platform.startsWith('Win');
    if (isWindows) {
      initScrollbar('sidebar');
      initScrollbar('sidebar-wrapper');

      docClasses.add('perfect-scrollbar-on');
    } else {
      docClasses.add('perfect-scrollbar-off');
    }
  }
};
</script>

<style lang="scss">
$scaleSize: 0.95;
@keyframes zoomIn95 {
  from {
    opacity: 0;
    transform: scale3d($scaleSize, $scaleSize, $scaleSize);
  }
  to {
    opacity: 1;
  }
}
.main-panel .zoomIn {
  animation-name: zoomIn95;
}
@keyframes zoomOut95 {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
    transform: scale3d($scaleSize, $scaleSize, $scaleSize);
  }
}
.main-panel .zoomOut {
  animation-name: zoomOut95;
}

.main-panel > .content {
    min-height: calc(100vh - 175px);
}
</style>