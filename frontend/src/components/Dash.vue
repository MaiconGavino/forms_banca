<template>
  <v-container>
    <v-form>
      <h3>Cadastro de Usuário</h3>
      <v-text-field label="Nome" v-model="userName"></v-text-field>
      <v-text-field label="Senha" v-model="userPassword" type="password"></v-text-field>
      <v-btn @click="registerUser">Cadastrar Usuário</v-btn>
      <h3>Cadastro de Startup</h3>
      <v-text-field label="Nome da Startup" v-model="startupName"></v-text-field>
      <v-btn @click="registerStartup">Cadastrar Startup</v-btn>
    </v-form>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      userName: '',
      userPassword: '',
      startupName: ''
    };
  },
  methods: {
    async registerUser() {
      try {
        await axios.post('http://localhost:8080/register', {
          name: this.userName,
          password: this.userPassword
        });
        alert('Usuário cadastrado!');
      } catch (error) {
        alert('Erro ao cadastrar usuário.');
      }
    },
    async registerStartup() {
      try {
        await axios.post('http://localhost:8080/startup/register', {
          name: this.startupName
        });
        alert('Startup cadastrada!');
      } catch (error) {
        alert('Erro ao cadastrar startup.');
      }
    }
  }
};
</script>
