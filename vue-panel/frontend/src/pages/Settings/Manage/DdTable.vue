<template>
    <card card-body-classes="table-full-width" no-footer-line v-loading="table.isLoading">
        <h4 slot="header" class="card-title">
          Domain detect
          <el-tooltip class="item" effect="dark" content="Create new rule" placement="top">
            <el-button  type="primary" icon="fal fa-plus-circle" size="small" circle
                        class="float-right"
                        @click="create()">
            </el-button>
          </el-tooltip>
          <el-dialog  :title="isCreateModal ? 'Create DD rule' : 'Edit DD rule'"
                      width="450px"
                      :visible.sync="showModal"
          >
            <el-form :model="form" style="margin-top: -30px;">
              <el-form-item label="Group name" style="margin-top: -20px;">
                <el-input v-model="form.group"
                          placeholder="crypto"
                          size="medium"
                ></el-input>
              </el-form-item>
              <el-form-item label="Domains" style="margin-top: -20px;">
                <el-input v-model="form.domains"
                          type="textarea"
                          placeholder="paypal.com,chase,com,..."
                ></el-input>
              </el-form-item>
              <el-form-item label="Color">
                <el-color-picker  v-model="form.color"
                                  size="medium"
                ></el-color-picker>
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
          Create DD rule
        </el-button>
        <div v-else>
          <el-table ref="multipleTable"
                  :data="paginate">
            <el-table-column prop="id" label="#" min-width="50"></el-table-column>
            <el-table-column label="Group" prop="group"></el-table-column>
            <el-table-column label="Domains">
              <template slot-scope="props">
                <span :style="{color: props.row.color}">{{props.row.domains}}</span>
              </template>
            </el-table-column>
            <el-table-column label="Color">
              <template slot-scope="props">
                <span :style="{color: props.row.color}">{{props.row.color}}</span>
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
        isLoading: false,
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
        group: '',
        color: '',
        domains: []
      },
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
    if (!this.$store.getters.DD_DATA.length) {
      await this.$store.dispatch('GET_DD_DATA');
    }
    this.table.data = this.$store.getters.DD_DATA;
    this.table.data.length = this.table.data.length;
  },
  methods: {
    async updateData() {
      this.table.isLoading = true;
      await this.$store.dispatch('GET_DD_DATA');
      this.table.data = this.$store.getters.DD_DATA;
      this.table.data.length = this.table.data.length;
      this.table.isLoading = false;
    },
    async deleteRule(row) {
      this.form.id = row.id;
      this.table.isLoading = true;
      await this.$store.dispatch('DEL_DD_RULE', this.form);
      this.table.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: 'DD rule deleted success!', duration:3000});
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
      this.form.group = row.group;
      this.form.color = row.color;
      this.form.domains = row.domains;
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
        await this.$store.dispatch('CREATE_DD_RULE', this.form);
      else
        await this.$store.dispatch('EDIT_DD_RULE', this.form);
      this.table.isLoading = false;
      const err = this.$store.getters.GET_LAST_ERROR;
      if (!err) {
        this.$notify.success({title: 'Success', message: this.isCreateModal ? 'DD rule created success!':'DD rule updated success!', duration:3000});
      } else {
        this.$notify.error({title: 'Error', message: err, duration:3000});
        this.$store.commit('SET_LAST_ERROR', null);
      }
      await this.updateData();
    },
  }
}
</script>
