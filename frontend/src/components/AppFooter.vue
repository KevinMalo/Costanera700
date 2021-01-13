<template>
  <v-footer dark absolute padless>
    <v-col class="text-center" cols="12">
      {{ new Date().getFullYear() }} — <strong>Costanera700</strong>
    </v-col>
  </v-footer>
</template>

<script>
export default {
  name: "AppFooter",
  data: () => ({
    icons: ["mdi-facebook", "mdi-twitter", "mdi-linkedin", "mdi-instagram"],
  }),
};
</script>

<style></style>


