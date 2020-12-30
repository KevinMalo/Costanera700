<template>
  <v-app>
    <v-main>
    <v-container fluid>
      <navbar></navbar>
      <router-view></router-view>
    </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from "@/components/Navbar";

export default {
  name: "App",
  components: {
    Navbar
  }
};
</script>
