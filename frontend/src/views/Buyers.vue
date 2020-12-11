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
        buyers: [
          {"id":"78feb807","name":"Laughry","age":71},{"id":"99a18da9","name":"Dora","age":55},{"id":"da4f6e2a","name":"Charbonnier","age":20},{"id":"65feeeee","name":"Cas","age":72},{"id":"4fa997b","name":"Gabey","age":51},{"id":"b2818207","name":"Julius","age":55},{"id":"6cc3fecc","name":"Ehman","age":82},{"id":"6e7cb370","name":"Lemon","age":26},{"id":"d0f94b88","name":"Rudman","age":24},{"id":"8865d7d7","name":"Elisha","age":43},{"id":"96111d4f","name":"Hawger","age":31},{"id":"68108d58","name":"Katherina","age":20},{"id":"6720715c","name":"Dougherty","age":76},{"id":"1611e146","name":"Garber","age":20},{"id":"9e55e903","name":"Shaylynn","age":75},{"id":"4e26ced1","name":"Dynah","age":27},{"id":"b8048387","name":"Calley","age":55},{"id":"b959a7fb","name":"Auguste","age":73},{"id":"d96e6f83","name":"Cynarra","age":55},{"id":"99555287","name":"Fillbert","age":81},{"id":"730361e3","name":"Robbert","age":73},{"id":"9b998516","name":"Bamford","age":44},{"id":"e187b07","name":"Tabib","age":59},{"id":"cd1ac90c","name":"Key","age":60},{"id":"1f556a10","name":"MacLean","age":34},{"id":"33bdc5c3","name":"Terryl","age":63},{"id":"332e5f8a","name":"Tega","age":26},{"id":"c068e19b","name":"Conley","age":57},{"id":"79b46a9e","name":"Bergeman","age":36},{"id":"5f0b4f45","name":"Sugden","age":35},{"id":"b9b7702","name":"Tor","age":38},
        ],
      }
    },
};
</script>