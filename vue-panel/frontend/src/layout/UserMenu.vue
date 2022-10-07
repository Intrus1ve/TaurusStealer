<template>
  <div class="user user__menu">
    <div class="photo">
      <i class="fal fa-user"></i>
    </div>
    <div class="info">
      <a data-toggle="collapse" :aria-expanded="!isClosed" @click.stop="toggleMenu" href="#">
           <span>
             {{title}}
             <b class="caret"></b>
          </span>
      </a>
      <div class="clearfix"></div>
      <div>
        <collapse-transition>
          <ul class="nav user-menu__nav" v-show="!isClosed">
            <slot>
              <li>
                <router-link :to="'/settings/manage'">
                  <span class="sidebar-mini-icon">EU</span>
                  <span class="sidebar-normal">Edit User</span>
                </router-link>
              </li>
              <li>
                <router-link :to="'/login'">
                  <span class="sidebar-mini-icon">L</span>
                  <span class="sidebar-normal">Log out</span>
                </router-link>
              </li>
            </slot>
          </ul>
        </collapse-transition>
      </div>
    </div>
  </div>
</template>
<script>
import { CollapseTransition } from 'vue2-transitions';

export default {
  name: 'user-menu',
  components: {
    CollapseTransition
  },
  props: {
    title: {
      type: String,
      default: 'unknown'
    }
  },
  data() {
    return {
      isClosed: true
    };
  },
  methods: {
    toggleMenu() {
      this.isClosed = !this.isClosed;
    }
  }
};
</script>
<style>
  .user__menu ul.user-menu__nav {
    margin-top: 0;
    padding-top: 20px;
  }
</style>
