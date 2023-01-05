<template>
  <div class="app">
    <h1 class="header">Mail Atlas &#x1F30D;</h1>

    <form @submit.prevent="submitForm" class="mx-8 my-8">
      <div class="relative">
        <div class="form-container">
          <svg
            aria-hidden="true"
            class="w-5 h-5 text-gray-500 dark:text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
        </div>
        <input
          class="form-input"
          type="search"
          v-model="term"
          placeholder="Enter a search term"
        />
        <button type="submit" class="form-button">Search</button>
      </div>
    </form>
    <MailGrid :mails="mails" />
  </div>
</template>

<script>
import MailGrid from "./components/MailGrid.vue";

export default {
  data() {
    return {
      term: "",
      mails: [],
    };
  },
  components: {
    MailGrid,
  },
  methods: {
    async submitForm() {
      const body = { term: this.term };
      const response = await fetch("http://localhost:3000/api/mails", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });
      if (response.ok) {
        const docs = await response.json();
        if ("Documents" in docs) {
          this.mails = docs.Documents;
        } else {
          console.error(
            "Error: API response does not contain a 'Documents' key"
          );
          this.mails = [];
        }
      } else {
        console.error(`Error: ${response.status} - ${response.statusText}`);
        this.mails = [];
      }
    },
  },
};
</script>
