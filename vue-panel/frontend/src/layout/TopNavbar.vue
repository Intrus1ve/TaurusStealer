<template>
  <navbar :show-navbar="showNavbar">
    <div class="navbar-wrapper">
      <div class="navbar-toggle" :class="{toggled: $sidebar.showSidebar}">
        <navbar-toggle-button @click.native="toggleSidebar">
        </navbar-toggle-button>
      </div>
      <span class="navbar-brand">
        {{$route.name}}
      </span>
    </div>

    <template slot="navbar-menu">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link" style="font-size: 13px;" href="#">
            <i class="fal fa-angle-right"> Taurus v1.5</i>
          </a>
        </li>
      </ul>
    </template>
  </navbar>
</template>

<script>
import {Navbar, NavbarToggleButton} from 'src/components';

export default {
  components: {
    Navbar,
    NavbarToggleButton
  },
  computed: {
    routeName() {
      const { name } = this.$route;
      return name.charAt(0).toUpperCase() + name.slice(1);
    }
  },
  data() {
    return {
      showNavbar: false
    };
  },
  methods: {
    toggleSidebar() {
      this.$sidebar.displaySidebar(!this.$sidebar.showSidebar);
    },
  }
};
</script>

<style>
</style>
