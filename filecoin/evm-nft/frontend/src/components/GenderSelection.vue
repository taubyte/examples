<template>
  <div class="gender-selection">
    <div
      class="gender-option"
      v-for="(option, index) in genderOptions"
      :key="index"
    >
      <p-radioButton v-model="gender" :value="option.value"></p-radioButton>
      <label>{{ option.label }}</label>
    </div>
    <p-button
      class="generate-avatar-button"
      @click="submit"
      label="Generate Avatar"
    ></p-button>
    <Spinner ref="spinnerRef" :loading="loading" />
  </div>
</template>

<script>
import { ComputeContract, provider } from "../ethereum";
import router from "../router";
import Button from "primevue/button";
import RadioButton from "primevue/radiobutton";
import Spinner from "./Spinner.vue";

export default {
  name: "GenderSelection",
  components: {
    Spinner,
    "p-button": Button,
    "p-radioButton": RadioButton,
  },
  data() {
    return {
      loading: false,
      gender: null,
      genderOptions: [
        { label: "Male", value: "Male" },
        { label: "Female", value: "Female" },
        { label: "Non-Binary", value: "Non-Binary" },
      ],
    };
  },
  methods: {
    async submit() {
    this.loading = true 
      const accounts = await provider.listAccounts();
      const account = accounts[0];

      await ComputeContract.submitCompute(this.gender, 1, {
        from: account,
        value: 1,
      });
      this.loading = false 
      router.push("avatar");
    },
  },
};
</script>

<style scoped>
.gender-selection {
  text-align: left;
}

.gender-option {
  margin-bottom: 10px;
}

.generate-avatar-button {
  margin-top: 20px;
}
</style>
