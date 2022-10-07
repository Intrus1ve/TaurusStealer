<template>
 <div class="container" v-loading.fullscreen.lock="isLoading" style="text-align: center;">
   <div class="login-card">
     <card class="card">
      <h4 slot="header" class="card-title" style="text-align: center;">Login</h4>
      <div class="row">
        <div class="col">
          <el-input v-model="data.username"
                      prefix-icon="fal fa-user"
                      placeholder="login"
            >
            </el-input>
        </div>
      </div>
      <div class="row">
        <div class="col">
            <el-input v-model="data.password"
                      prefix-icon="fal fa-key"
                      placeholder="password"
                      show-password
            >
            </el-input>
        </div>
      </div>
      <div slot="footer" style="text-align: center;">
          <el-button v-on:click="login()">login</el-button>
      </div>
    </card>
   </div>
 </div>
</template>

<script>
import {Card} from 'src/components'
import store from 'src/store'

export default {
  components: {
    Card
  },
  data() {
    return {
      isLoading: false,
      data: {
        username: '',
        password: ''
      }
    };
  },
  mounted() {
   this.$store.dispatch('LOGOUT');
  },
  methods: {
    async login() {
      this.isLoading = true;
      await this.$store.dispatch('LOGIN', this.data);
      this.isLoading = false;
      if (this.$store.getters.CURRENT_USER) {
        this.$notify.success({title: "Success", message: "Redirecting to dashboard...", duration:1500});
        setTimeout(() => {
          this.$router.push('dashboard');
        }, 1500);

      } else {
        this.$notify.error({title: "Error", message: "Invalid login/pass!", duration:2000});
      }
    }
  }
};
</script>

<style>
.login-card {
  width: 310px !important;
  margin: 0 auto !important;
  margin-top: 20% !important;
}
.row {
    margin-bottom: 20px;
  }
</style>
