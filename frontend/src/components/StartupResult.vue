<template>
  <v-container>
    <h3>Resultados</h3>
    <v-data-table
        :headers="headers"
        :items="results"
        class="elevation-1"
    ></v-data-table>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      headers: [
        { text: "Startup", value: "startup_name" },
        { text: "Média de Pontuação", value: "average_score" },
      ],
      results: [],
    };
  },
  created() {
    this.fetchResults();
  },
  methods: {
    async fetchResults() {
      try {
        const response = await axios.get("http://localhost:8080/result");
        this.results = response.data;
      } catch (error) {
        alert("Erro ao carregar resultados.");
      }
    },
  },
};
</script>

<style scoped>
.v-container {
  margin-top: 20px;
}
</style>
