<template>
  <div class="mx-8 my-8">
    <div class="overflow-x-auto relative shadow-md sm:rounded-lg">
      <table class="table-class">
        <thead class="table-head">
          <tr>
            <th scope="col" class="py-3 px-6">From</th>
            <th scope="col" class="py-3 px-5">To</th>
            <th scope="col" class="py-3 px-6">Subject</th>
          </tr>
        </thead>
        <tbody v-if="mails.length > 0">
          <tr
            v-for="mail in visibleMails"
            :key="mail._id"
            @click="this.selectedMail = mail"
            class="table-row"
          >
            <td class="py-4 px-6">{{ mail._source.From }}</td>
            <td class="py-4 px-6">{{ mail._source.To }}</td>
            <td class="py-4 px-6">{{ mail._source.Subject }}</td>
          </tr>
        </tbody>
        <tbody v-else>
          <tr>
            <td class="py-4 px-6" colspan="3">No mails found</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="flex justify-between py-4">
      <button
        class="text-blue-500 hover:text-blue-800"
        @click="prevPage"
        :disabled="currentPage === 1"
      >
        Previous
      </button>
      <p v-if="totalPages > 0">Page {{ currentPage }} of {{ totalPages }}</p>
      <p v-else>Page 1 of 1</p>
      <button
        class="text-blue-500 hover:text-blue-800"
        @click="nextPage"
        :disabled="currentPage === totalPages"
      >
        Next
      </button>
    </div>

    <div>
      <h2><strong>Content</strong></h2>
      <p v-if="selectedMail">
        <MailContent :mail="selectedMail" />
      </p>
      <p v-else>No mail data</p>
    </div>
  </div>
</template>

<script>
import MailContent from "./MailContent.vue";
export default {
  props: {
    mails: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      selectedMail: null,
      rowsPerPage: 5,
      currentPage: 1,
    };
  },
  components: {
    MailContent,
  },
  watch: {
    mails: {
      inmediate: true,
      handler() {
        if (this.mails.length === 0) {
          this.selectedMail = null;
        }
        this.currentPage = 1;
      },
    },
  },
  computed: {
    totalPages() {
      return Math.ceil(this.mails.length / this.rowsPerPage);
    },
    visibleMails() {
      const startIndex = (this.currentPage - 1) * this.rowsPerPage;
      return this.mails.slice(startIndex, startIndex + this.rowsPerPage);
    },
  },
  methods: {
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }
    },
  },
};
</script>
