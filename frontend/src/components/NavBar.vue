<template>
  <v-app-bar app color="primary" dark>
    <v-toolbar-title>Forms Banca</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn text to="/" v-if="!authenticated">Login</v-btn>
    <v-btn text to="/dash" v-if="isAdmin">Dashboard</v-btn>
    <v-btn text to="/evaluation" v-if="authenticated">Avaliar Startups</v-btn>
    <v-btn text to="/result" v-if="authenticated">Resultados</v-btn>
    <v-btn text @click="logout" v-if="authenticated">Logout</v-btn>
  </v-app-bar>
</template>

<script>
export default {
  computed: {
    authenticated() {
      return !!localStorage.getItem("role");
    },
    isAdmin() {
      return localStorage.getItem("role") === "admin";
    },
  },
  methods: {
    logout() {
      localStorage.clear();
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.v-app-bar {
  position: fixed;
  top: 0;
  z-index: 1000;
}
</style>
