<template>
  <!-- Container Principal -->
  <v-container class="fill-height login-container">
    <v-row justify="center" align="center" class="h-100">
      <v-col cols="12" sm="8" md="4">
        <!-- Card Centralizado -->
        <v-card class="elevation-12 rounded-card pa-4">
          <!-- Título do Card -->
          <v-card-title class="text-center">
            <span class="text-h5 font-weight-bold primary--text">Bem-vindo!</span>
          </v-card-title>

          <!-- Formulário de Login -->
          <v-card-text>
            <v-form @submit.prevent="login">
              <!-- Campo de Usuário -->
              <v-text-field
                  v-model="name"
                  label="Usuário"
                  outlined
                  color="primary"
                  prepend-icon="mdi-account"
                  placeholder="Digite seu usuário"
                  required
                  class="mb-3"
              ></v-text-field>

              <!-- Campo de Senha -->
              <v-text-field
                  v-model="password"
                  label="Senha"
                  type="password"
                  outlined
                  color="primary"
                  prepend-icon="mdi-lock"
                  placeholder="Digite sua senha"
                  required
              ></v-text-field>

              <!-- Botão de Login -->
              <v-btn
                  block
                  color="primary"
                  class="mt-4 rounded-pill"
                  dark
                  :loading="loading"
                  type="submit"
              >
                Entrar
              </v-btn>
            </v-form>
          </v-card-text>

          <!-- Mensagem de Erro (Snackbar) -->
          <v-snackbar v-model="error" color="red" bottom>
            {{ errorMessage }}
          </v-snackbar>
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
      name: "",
      password: "",
      loading: false,
      error: false,
      errorMessage: "",
    };
  },
  methods: {
    async login() {
      this.loading = true;
      this.error = false;

      try {
        const response = await axios.post("http://localhost:8080/login", {
          name: this.name,
          password: this.password,
        });

        // Armazena informações do usuário
        localStorage.setItem("role", response.data.role);
        this.$router.push(response.data.role === "admin" ? "/dash" : "/evaluation");
      } catch {
        this.error = true;
        this.errorMessage = "Usuário ou senha incorretos.";
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
/* Fundo da Página com Gradiente */
.login-container {
  background: linear-gradient(to bottom right, #f5f5f5, #e8eaf6);
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* Card Customizado */
.rounded-card {
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
</style>
