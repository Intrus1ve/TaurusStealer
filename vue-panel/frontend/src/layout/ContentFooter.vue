<template>
  <footer class="footer">
    <div class="container-fluid d-flex flex-wrap justify-content-between">
      <nav>
        <ul class="footer-menu">
          <li>
            <router-link :to="'/dashboard'">Go to Dashboard</router-link>
          </li>
        </ul>
      </nav>
      <div class="copyright">
        &copy; {{year}}, Taurus Project
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  data() {
    return {
      year: new Date().getFullYear()
    };
  }
};
</script>
