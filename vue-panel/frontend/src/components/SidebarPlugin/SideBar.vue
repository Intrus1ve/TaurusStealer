<template>
  <div class="sidebar"
       :data-color="backgroundColor">

    <div class="logo">
      <router-link to="/dashboard" class="simple-text logo-mini">
        <div class="logo-image">
          <img :src="logo">
        </div>
      </router-link>

      <router-link to="/dashboard" class="simple-text logo-normal" style="outline: none !important;">
        {{title}}
      </router-link>
      <div class="navbar-minimize">
        <el-button  icon="fal fa-bars"
                    class="minimize-btn"
                    circle
                    @click="minimizeSidebar"
        ></el-button>
      </div>
    </div>
    <div class="sidebar-wrapper" ref="sidebarScrollArea">
      <slot></slot>
      <ul class="nav">
        <slot name="links">
          <sidebar-item v-for="(link, index) in sidebarLinks"
                        :key="link.name + index"
                        :link="link">

            <sidebar-item v-for="(subLink, index) in link.children"
                          :key="subLink.name + index"
                          :link="subLink">
            </sidebar-item>
          </sidebar-item>
        </slot>

      </ul>
    </div>
  </div>
</template>
<script>

export default {
  name: 'sidebar',
  props: {
    title: {
      type: String,
      default: 'Taurus'
    },
    backgroundColor: {
      type: String,
      default: 'black',
      validator: value => {
        let acceptedValues = [
          '',
          'blue',
          'azure',
          'green',
          'orange',
          'red',
          'purple',
          'black'
        ];
        return acceptedValues.indexOf(value) !== -1;
      }
    },
    logo: {
      type: String,
      default: './img/taurus.png'
    },
    sidebarLinks: {
      type: Array,
      default: () => []
    },
    autoClose: {
      type: Boolean,
      default: true
    }
  },
  provide() {
    return {
      autoClose: this.autoClose
    };
  },
  methods: {
    minimizeSidebar() {
      if (this.$sidebar) {
        this.$sidebar.toggleMinimize();
      }
    }
  },
  beforeDestroy() {
    if (this.$sidebar.showSidebar) {
      this.$sidebar.showSidebar = false;
    }
  }
};
</script>

<style scoped>
@media (min-width: 992px) {
  .navbar-search-form-mobile,
  .nav-mobile-menu {
    display: none;
  }
}

.minimize-btn {
  color: #FFF;
  background-color: rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.1);
  outline: none !important;
}
</style>
