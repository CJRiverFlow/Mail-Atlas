<template>
  <div class="bg-red-500 text-white font-bold px-4 py-2">Error</div>
  <div
    class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative"
    role="alert"
    v-if="flashStore.show"
  >
    {{ flashStore.text }}
  </div>
</template>

<script>
import { useFlashStore } from "../stores/flashStore";
export default {
  setup() {
    const flashStore = useFlashStore();
    return { flashStore };
  },
  mounted() {
    setTimeout(() => {
      this.flashStore.hideFlash();
    }, 3000);
  },
};
</script>
