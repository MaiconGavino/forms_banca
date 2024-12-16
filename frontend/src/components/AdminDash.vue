<template>
  <v-container class="fill-height d-flex align-center justify-center">
    <v-card max-width="600" class="pa-4" elevation="10">
      <v-card-title class="text-h5 text-center primary--text">Dashboard - Admin</v-card-title>
      <v-card-text>
        <h4 class="secondary--text mb-2">Cadastro de Usuário</h4>
        <v-form>
          <v-text-field
              v-model="userName"
              label="Nome do Usuário"
              outlined
              color="primary"
          ></v-text-field>
          <v-text-field
              v-model="userPassword"
              label="Senha"
              type="password"
              outlined
              color="primary"
          ></v-text-field>
          <v-btn block color="primary" class="mt-3" dark @click="registerUser">Cadastrar Usuário</v-btn>
        </v-form>

        <v-divider class="my-5"></v-divider>

        <h4 class="secondary--text mb-2">Cadastro de Startup</h4>
        <v-form>
          <v-text-field
              v-model="startupName"
              label="Nome da Startup"
              outlined
              color="secondary"
          ></v-text-field>
          <v-btn block color="secondary" class="mt-3" dark @click="registerStartup">Cadastrar Startup</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      userName: "",
      userPassword: "",
      startupName: "",
    };
  },
  methods: {
    async registerUser() {
      try {
        await axios.post("http://localhost:8080/register", {
          name: this.userName,
          password: this.userPassword,
        });
        alert("Usuário cadastrado com sucesso!");
      } catch {
        alert("Erro ao cadastrar usuário.");
      }
    },
    async registerStartup() {
      try {
        await axios.post("http://localhost:8080/startup/register", {
          name: this.startupName,
        });
        alert("Startup cadastrada com sucesso!");
      } catch {
        alert("Erro ao cadastrar startup.");
      }
    },
  },
};
</script>

<style scoped>
.v-container {
  background-color: #f5f5f5;
}
</style>
