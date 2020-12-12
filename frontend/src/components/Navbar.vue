<template>
  <v-toolbar class="hidden-xs-and-down">
    <v-toolbar-title>
        <router-link to="/" tag="span" style="cursor: pointer">
          Costanera700
        </router-link>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items>
        <v-toolbar-items class="hidden-xs-only">
        <v-btn
          flat
          v-for="item in menuItems"
          :key="item.title"
          :to="item.path">
          {{ item.title }}
        </v-btn>
      </v-toolbar-items>
    </v-toolbar-items>
  </v-toolbar>
</template>

<script>
export default {
  name: "Navbar",
  data(){
    return {
      appTitle: 'Awesome App',
      sidebar: false,
      menuItems: [
          { title: 'Home', path: '/'},
          { title: 'Uploads', path: '/uploads'},
          { title: 'Buyers', path: '/buyers'},
          { title: 'Search Buyers', path: '/search'},
     ]
    }
  },
};
</script>

<style></style>


