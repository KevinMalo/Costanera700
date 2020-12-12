<template>
<v-container>
<v-card>
    <v-card-title>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :items="buyers"
      :search="search"
    ></v-data-table>
  </v-card>
</v-container>
</template>

<script>

import api from "@/services/api"

export default {
  name: "Buyers",
  components: { },
  data () {
      return {
        search: '',
        headers: [
          {
            text: 'ID',
            align: 'start',
            filterable: true,
            value: 'id',
          },
          { text: 'NAME', filterable: false, value: 'name' },
          { text: 'AGE', filterable: false, value: 'age' },
        ],
        buyers: []
      }
    },
    created() {
      api.getBuyers()
      .then( buyers => this.buyers = buyers )
    }
};
</script>