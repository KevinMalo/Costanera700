<template>
  <v-app>
    <v-main>
      <navbar></navbar>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
      <app-footer></app-footer>
    </v-main>
  </v-app>
</template>

<script>
import Navbar from "@/components/Navbar";
import AppFooter from "@/components/AppFooter";

export default {
  name: "App",
  components: {
    Navbar,
    AppFooter,
  },
};
</script>
