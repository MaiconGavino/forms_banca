<template>
  <v-container>
    <v-forms>
      <v-text-field label="Nome" v-model="name"></v-text-field>
      <v-text-field label="Senha" v-model="password" type="password"></v-text-field>
      <v-btn @click="login">Login</v-btn>
    </v-forms>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data(){
    return{
      name:'',
      password:'',
    };
  },
  methods:{
    async login(){
      try {
        await axios.post('http://localhost:8080/login', {
          name: this.name,
          password: this.password
        });
        alert('Login successful!');
        this.$router.push('/dash');
      }catch(err){
        alert('Error no login');
      }
    }
  }
};
</script>
