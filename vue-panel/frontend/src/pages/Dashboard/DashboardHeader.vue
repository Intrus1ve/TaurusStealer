<template>
  <div class="panel-header panel-header-lg" v-loading="isLoading">
    <header-chart v-if="!isLoading"
                  :height="255"
                  :data="chartData"
                  :labels="chartLabels">
    </header-chart>
  </div>
</template>

<script>
import HeaderChart from './HeaderChart';
import store from 'src/store'

export default {
  name: 'overview-header',
  components: {
    HeaderChart
  },
  data() {
    return {
      isLoading: true,
      chartData: [0, 0, 0, 0, 0, 0, 0],
      chartLabels: [
        '6 day ago',
        '5 day ago',
        '4 day ago',
        '3 day ago',
        '2 day ago',
        '1 day ago',
        'today'
      ]
    };
  },
  async mounted() {
    this.isLoading = true;
    if (!this.$store.getters.DAYS_DATA.length) {
      await this.$store.dispatch('GET_DAYS_DATA');
    }
    this.chartData = this.$store.getters.DAYS_DATA;
    this.isLoading = false;
  }
};
</script>