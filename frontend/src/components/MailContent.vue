<template>
  <div>
    <p class="whitespace-pre-wrap leading-none text-gray-600 text-sm">
      {{ mail._source.Content }}
    </p>
  </div>
</template>

<script>
export default {
  props: {
    mail: {
      type: Object,
      required: true,
    },
  },
};
</script>
