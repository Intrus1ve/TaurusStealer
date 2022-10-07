<template>
    <card card-body-classes="table-full-width"
          no-footer-line
          v-loading="table.isLoading"
          v-loading.lock.fullscreen="isBackupCreating">
        <h4 slot="header" class="card-title">
          Backups
          <el-tooltip class="item" effect="dark" content="Create new rule" placement="top">
            <el-button  type="primary" icon="fal fa-plus-circle" size="small" circle
                        class="float-right"
                        @click="create()">
            </el-button>
          </el-tooltip>
          <el-dialog  title="Create backup"
                      width="450px"
                      :visible.sync="showModal"
          >
            <el-form :model="form" style="margin-top: -30px;">
              <el-form-item label="Comment" style="margin-top: -20px;">
                <el-input v-model="form.comment"
                          size="medium"
                          placeholder="comment..."
                ></el-input>
              </el-form-item>
              <el-form-item label="Delete all logs after create backup" style="margin-top: -20px;">
                <el-switch v-model="form.delete"/>
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
                    type="text"
                    @click="create()"
        >
          Create backup
        </el-button>
        <div v-else>
          <el-table ref="multipleTable"
                  :data="paginate">
            <el-table-column prop="id" label="#" min-width="50"></el-table-column>
            <el-table-column prop="date" label="Date"></el-table-column>
            <el-table-column prop="size" label="Size"></el-table-column>
            <el-table-column prop="comment" label="Comment"></el-table-column>
            <el-table-column label="Actions">
              <template slot-scope="props">
                <el-tooltip class="item" effect="dark" content="Download" placement="top">
                  <el-button  icon="fal fa-file-download"
                              circle
                              size="medium"
                              @click="downloadBackup(props.row)"
                  ></el-button>
                </el-tooltip>
                <el-tooltip class="item" effect="dark" content="Delete" placement="top">
                  <el-button  icon="fal fa-trash"
                              type="danger"
                              circle
                              size="medium"
                              @click="deleteRule(props.row)"
                  ></el-button>
                </el-tooltip>
              </template>
            </el-table-column>
          </el-table>
          <div slot="footer" class="col-12 d-flex justify-content-center justify-content-sm-between flex-wrap mt-3">
            <p class="text-muted">Total {{table.data.length}} backups</p>
            <pagination class="pagination-primary"
                          v-model="pagination.currentPage"
                          :per-page="pagination.perPage"
                          :total="table.data.length"/>
            <p class="text-muted">Showing {{from}}-{{to}} backups</p>
          </div>
        </div>
      </card>
</template>

<script>
import { Pagination, Card} from 'src/components';
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
      isBackupCreating: false,
      pagination: {
        perPage: 5,
        currentPage: 1,
      },
      showModal: false,
      form: {},
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
  async mounted() {
    this.table.isLoading = true;
    if (!this.$store.getters.BACKUPS_DATA.length) {
      await this.$store.dispatch('GET_BACKUPS_DATA');
    }
    this.table.data = this.$store.getters.BACKUPS_DATA;
    this.table.isLoading = false;
  },
  methods: {
    async updateData() {
      this.table.isLoading = true;
      await this.$store.dispatch('GET_BACKUPS_DATA');
      this.table.data = this.$store.getters.BACKUPS_DATA;
      this.table.isLoading = false;
    },
    create() {
      this.form = {};
      this.showModal = true;
    },
    async deleteRule(row) {
      this.form.id = row.id;
      this.form.date = row.date;
      this.table.isLoading = true;
      await this.$store.dispatch('DEL_BACKUP', this.form);
      this.table.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Backup deleted success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }

      const index = this.table.data.indexOf(row);
      if (index != -1) {
        this.table.data.splice(index, 1);
      }
    },
    async modalSubmit() {
      this.showModal = false;
      this.isBackupCreating= true;
      await this.$store.dispatch('CREATE_BACKUP', this.form);
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'Backup created success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
      await this.updateData();
      this.isBackupCreating = false;
    },
    async downloadBackup(row) {
      this.form.date = row.date;
      this.$store.dispatch('DOWNLOAD_BACKUP', this.form);

      this.$notify.success({title: 'Success', message: 'Downloading started in background, please wait', duration:3000});
    },
  }
}
</script>
