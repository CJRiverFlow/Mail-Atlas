import { defineStore } from "pinia";

export const useFlashStore = defineStore("flashStore", {
  state: () => ({
    text: "",
    show: false,
  }),
  actions: {
    showFlash(message) {
      this.text = message;
      this.show = true;
    },
    hideFlash() {
      this.text = "";
      this.show = false;
    },
  },
});
