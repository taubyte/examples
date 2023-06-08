<template>
  <div>
      <p-button @click="connect">Connect Wallet</p-button>
      <Spinner ref="spinnerRef" :loading="loading" />
  </div>
</template>

<script>
import axios from "axios";
import { provider } from "../ethereum";
import Spinner from "./Spinner.vue";
import router from "../router";
import Button from 'primevue/button';

export default {
  name: "Login",
  data() {
    return {
      socket: null,
      loading: false,
      api: "https://3kgmdeza0.gtau.link"
    };
  },
  components: {
    Spinner,
    'p-button': Button,
  },
  methods: {
    async connect() {
      this.loading = true;
      await provider.send("eth_requestAccounts");
      const accounts = await provider.listAccounts();
      const account = accounts[0];

      try {
        await axios.post(
            `${this.api}/listen/compute`,{}
        )

        await axios.post(
            `${this.api}/listen/mint`,{}
        )

        const res = await axios.get(
          `${this.api}/notification/url?wallet=${account}`
        );
        window.socket = new WebSocket(`wss://3kgmdeza0.gtau.link/${res.data}`);
        window.socket.binaryType = "arraybuffer"
        window.socket.onopen = (event) => {
          this.loading = false;
          router.push("gender-selection");
          console.log("notifications websocket connected");
        };
      } catch (e) {
        console.log(e);
      }
    },
  },
};
</script>
