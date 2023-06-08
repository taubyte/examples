<template>
  <div>
    <Message v-if="showMessage" severity="success"> Avatar Minted </Message>
    <img v-if="avatar" :src="avatar" alt="Avatar" />
    <div>
      <p-button v-if="avatar" @click="mint">Mint</p-button>
    </div>
    <Spinner ref="spinnerRef" :loading="loading" />
  </div>
</template>

<script>
import axios from "axios";
import { MintContract, provider } from "../ethereum";
import Spinner from "./Spinner.vue";
import Button from "primevue/button";
import Message from "primevue/message";

export default {
  name: "Avatar",
  data() {
    return {
      avatar: null,
      cid: null,
      loading: true,
      showMessage: false,
    };
  },
  components: {
    Spinner,
    "p-button": Button,
    Message,
  },
  created() {
    window.socket.onmessage = (event) => {
      console.log("GOT MESSAGE");
      let dec = new TextDecoder("utf-8");
      let decoded = dec.decode(event.data);
      const data = JSON.parse(decoded);
      if (data.error) {
        console.error(data.error);
      } else {
        console.log("DATA::", data);
        if (data.Cid) {
          this.cid = data.Cid;
          axios
            .get(`https://ipfs.tau.link/get?cid=${data.Cid}`, {
              responseType: "blob",
            })
            .then((res) => {
              let blob = new Blob([res.data], { type: "image/png" });
              let image = URL.createObjectURL(blob);
              this.loading = false;
              this.avatar = image;
            });
        }

        if (data.TokenId && data.Sender && data.TokenURI) {
          this.loading = false;
          this.showMessage = true;
          window.socket.onmessage = null
        }
      }
    };
  },
  methods: {
    async mint() {
      this.loading = true;
      const accounts = await provider.listAccounts();
      const account = accounts[0];

      await MintContract.mintTauNFT(account, `ipfs://${this.cid}`);
    },
  },
};
</script>
