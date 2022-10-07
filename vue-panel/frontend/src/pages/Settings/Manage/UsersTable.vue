<template>
    <card card-body-classes="table-full-width" no-footer-line v-loading="table.isLoading">
        <h4 slot="header" class="card-title">
          Users
          <el-tooltip class="item" effect="dark" content="Create new rule" placement="top">
            <el-button  type="primary" icon="fal fa-plus-circle" size="small" circle
                        class="float-right"
                        @click="create()">
            </el-button>
          </el-tooltip>
          <el-dialog  :title="isCreateModal ? 'Create new user' : 'Edit user'"
                      width="450px"
                      :visible.sync="showModal"
          >
            <el-form :model="form" style="margin-top: -30px;">
              <el-form-item label="Username" style="margin-top: -20px;">
                <el-input v-model="form.username"
                          size="medium"
                ></el-input>
              </el-form-item>
              <el-form-item label="Password" style="margin-top: -20px;">
                <el-input v-model="form.password"
                          show-password
                          size="medium"
                ></el-input>
              </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
              <el-button size="small" @click="showModal = false">Cancel</el-button>
              <el-button size="small" type="primary" @click="modalSubmit">Submit</el-button>
            </span>
          </el-dialog>
        </h4>
        <el-button  v-if="table.data.length < 1"
                    class="col d-flex justify-content-center"
                    type="text" @click="create()"
        >
          Create new user
        </el-button>
        <div v-else>
          <el-table ref="multipleTable"
                    :data="paginate">
            <el-table-column prop="id" label="#" min-width="50"></el-table-column>
            <el-table-column label="Username">
              <template slot-scope="props">
                <el-tooltip class="item" effect="dark" content="Root user" placement="top">
                  <span style="color: #F56C6C;" v-if="props.row.access">{{props.row.username}}</span>
                </el-tooltip>
                <el-tooltip class="item" effect="dark" content="Not root user" placement="top">
                  <span style="color: #1989fa;" v-if="!props.row.access">{{props.row.username}}</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column prop="last_online" label="Online">
              <template slot-scope="props">
                <span>{{props.row.last_online | timeago}}</span>
              </template>
            </el-table-column>
            <el-table-column label="Actions">
              <template slot-scope="props">
                <el-tooltip class="item" effect="dark" content="Edit" placement="top">
                  <el-button  icon="fal fa-edit"
                              circle
                              @click="edit(props.row)"
                              size="medium"
                  ></el-button>
                </el-tooltip>
                <el-tooltip class="item" effect="dark" content="Delete" placement="top">
                  <el-button  icon="fal fa-trash"
                              type="danger"
                              circle
                              @click="deleteRule(props.row)"
                              size="medium"
                  ></el-button>
                </el-tooltip>
              </template>
            </el-table-column>
          </el-table>
          <div slot="footer" class="col-12 d-flex justify-content-center justify-content-sm-between flex-wrap mt-3">
            <p class="text-muted">Total {{table.data.length}} users</p>
            <pagination class="pagination-primary"
                          v-model="pagination.currentPage"
                          :per-page="pagination.perPage"
                          :total="table.data.length"/>
            <p class="text-muted">Showing {{from}}-{{to}} users</p>
          </div>
        </div>
      </card>
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
      table: {
        isLoading: true,
        data: [],
        size: 0,
      },
      pagination: {
        perPage: 5,
        currentPage: 1,
      },
      showModal: false,
      isCreateModal: true,
      form: {
        id: 0,
        username: '',
      },
    }
  },
  filters: {
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
    }
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
    }
  },
   async created() {
    this.table.isLoading = true;
    if (!this.$store.getters.USERS_DATA.length) {
      await this.$store.dispatch('GET_USERS_DATA');
    }
    this.table.data = this.$store.getters.USERS_DATA;
    this.table.isLoading = false;
  },
  methods: {
    async updateData() {
      this.table.isLoading = true;
      await this.$store.dispatch('GET_USERS_DATA');
      this.table.data = this.$store.getters.USERS_DATA;

      this.table.isLoading = false;
    },
    async deleteRule(row) {
      this.form.id = row.id;
      this.form.username = row.username;
      this.table.isLoading = true;
      await this.$store.dispatch('DEL_USER', this.form);
      this.table.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'User deleted success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }

      const index = this.table.data.indexOf(row);
      if (index != -1) {
        this.table.data.splice(index, 1);
      }
    },
    edit(row) {
      this.form.id = row.id;
      this.form.username = row.username;
      this.isCreateModal = false;
      this.showModal = true;
    },
    create() {
      this.form = {};
      this.isCreateModal = true;
      this.showModal = true;
    },
    async modalSubmit() {
      this.showModal = false;
      this.table.isLoading = true;
      if (this.isCreateModal)
        await this.$store.dispatch('CREATE_USER', this.form);
      else
        await this.$store.dispatch('EDIT_USER', this.form);
      this.table.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: this.isCreateModal ? 'User created success!':'User updated success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
      await this.updateData();
    },
  }
}
</script>
