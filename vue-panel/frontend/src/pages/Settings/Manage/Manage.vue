<template>
  <div class="">
    <div class="row">
      <div class="col-md-6">
        <UsersTable/>
      </div>
      <div class="col-md-6">
        <BackupsTable/>
      </div>
    </div>
    <div class="row">
      <div class="col-md-6">
        <DdTable/>
      </div>
      <div class="col-md-6">
        <BannedTable/>
      </div>
    </div>
  </div>
</template>

<script>
import UsersTable from './UsersTable'
import BackupsTable from './BackupsTable'
import DdTable from './DdTable'
import BannedTable from './BannedTable'

export default {
  components: {
    UsersTable,
    BackupsTable,
    DdTable,
    BannedTable,
  }
};
</script>

<style>
.text-muted {
    color: #70657b !important;
    font-size: 12px;
}
</style>
