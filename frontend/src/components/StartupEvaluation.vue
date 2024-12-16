<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-card class="pa-5" color="surface" outlined>
          <v-card-title class="text-h5 text-center">Avaliação de Startups</v-card-title>
          <v-card-text>
            <v-select
                v-model="selectedStartup"
                :items="startups"
                item-text="name"
                item-value="id"
                label="Selecione uma Startup"
                outlined
                color="primary"
            ></v-select>

            <div v-if="selectedStartup" class="mt-5">
              <v-row v-for="(label, criterion) in criterionLabels" :key="criterion">
                <v-col cols="12" md="6">
                  <p class="text-secondary">{{ label }}</p>
                  <v-rating v-model="scores[criterion]" length="5" outlined></v-rating>
                </v-col>
              </v-row>
              <v-btn block color="accent" class="mt-5" dark @click="submitEvaluation">
                Enviar Avaliação
              </v-btn>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      startups: [],
      selectedStartup: null,
      scores: {
        innovation: 0,
        impact: 0,
        economicFeasibility: 0,
        technicalFeasibility: 0,
        pitch: 0,
      },
      criterionLabels: {
        innovation: "Grau de Inovação",
        impact: "Impacto Socioambiental",
        economicFeasibility: "Viabilidade Econômica",
        technicalFeasibility: "Viabilidade Técnica",
        pitch: "Pitch",
      },
    };
  },
  created() {
    this.fetchStartups();
  },
  methods: {
    async fetchStartups() {
      try {
        const response = await axios.get("http://localhost:8080/startup/list");
        this.startups = response.data;
      } catch {
        this.error = true;
        this.errorMessage = "Erro ao carregar startups. Verifique o servidor.";
      }
    },
    async submitEvaluation() {
      try {
        await axios.post("http://localhost:8080/startup/evaluate", {
          startup_id: this.selectedStartup,
          mentor_id: 1,
          ...this.scores,
        });
        alert("Avaliação enviada com sucesso!");
        this.$router.push("/result");
      } catch {
        alert("Erro ao enviar avaliação.");
      }
    },
  },
};
</script>

<style scoped>
.v-container {
  min-height: 100vh;
}
</style>
