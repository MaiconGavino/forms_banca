<template>
  <v-app>
    <Navbar />
    <v-main class="app-background">
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import Navbar from "@/components/NavBar.vue";

export default {
  name: "App",
  components: {
    Navbar,
  },
};
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  background-color: #f5f5f5; /* Cinza claro */
  font-family: 'Roboto', Arial, sans-serif;
}
.v-application {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>

