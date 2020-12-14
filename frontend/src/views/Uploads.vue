<template>
<v-container>
<v-form method="post" :action="`${url}${date}`">

  <v-col cols="6">
   <v-container class="grey lighten-5">
    <v-row no-gutters>
      <v-col>
        <v-card
          class="pa-2"
          outlined
          tile> 
          
          <v-text-field
            label="Date" v-model="date"
          ></v-text-field>
        </v-card>
          <v-btn
            color="primary"
            text
            type="submit"
          >
            Submit
          </v-btn>
      </v-col>
      <v-col order="12">
        <v-card
          class="pa-2"
          outlined
          tile
        >
      <v-date-picker
      v-model="picker"
      color="green lighten-1"
    ></v-date-picker>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
  
  </v-col>



</v-form>
</v-container>
</template>  

<script>
  export default {
    name: "Uploads",
    data () {
      return {
        picker: new Date().toISOString().substr(0, 10),
        picker2: "",
        date: Math.floor(Date.now() / 1000),
        url: "http://localhost:9000/uploads/"
      }
    },
  }
</script>

<style scope>
</style>