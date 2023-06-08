<template>
  <div class="spinner-container">
    <p-progressSpinner v-if="loading"></p-progressSpinner>
  </div>
</template>

<script>
import ProgressSpinner from "primevue/progressspinner"; // Corrected import
export default {
  name: "Spinner",
  components: {
    "p-progressSpinner": ProgressSpinner,
  },
  props: {
    loading: {
      type: Boolean,
      default: false,
    },
  },
};
</script>

<style scoped>
.spinner-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>
