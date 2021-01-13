<template>
  <v-container>
    <v-parallax
      dark
      src="https://images.pexels.com/photos/3298057/pexels-photo-3298057.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=750&w=1260"
    >
      <v-row align="center" justify="center">
        <v-col class="text-center" cols="12">
          <v-layout justify-center>
            <v-img
              class="justify-center"
              lazy-src="https://sushiswapclassic.org/static/media/chef.ec78e8e8.png"
              max-height="120"
              max-width="120"
              src="https://sushiswapclassic.org/static/media/chef.ec78e8e8.png"
            ></v-img>
          </v-layout>
          <h1 class="display-3 mb-1">WELCOME</h1>
          <h4 class="subheading">MasterChef is Ready!</h4>
        </v-col>
      </v-row>
    </v-parallax>
  </v-container>
</template>

<script>
export default {
  name: "Home",
  components: {},
};
</script>

<style>
h1 {
  color: #f80a18;
}
</style>
