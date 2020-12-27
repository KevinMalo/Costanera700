<template>
  <v-container>
    <v-col cols="12">
      <v-container class="grey lighten-5">
        <v-row no-gutters>
          <v-col>
            <v-card class="pa-2" outlined tile>
              <v-text-field label="userId" v-model="userId"></v-text-field>
            </v-card>
            <v-btn color="primary" type="submit" v-on:click="getData(userId)">
              Submit
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-col>

    <div v-if="dataProducts.length > 0">
      <h1>Product History</h1>
      <v-col cols="12">
        <v-simple-table>
          <template v-slot:default>
            <thead>
              <tr>
                <th class="text-left">Product id</th>
                <th class="text-left">Name</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="item in dataProducts"
                :key="item.product[0].product_id"
              >
                <td>{{ item.product[0].product_id }}</td>
                <td>{{ item.product[0].name }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </div>
  </v-container>
</template>

<script>
import getUserTransactionsService from "@/services/getUserTransactionsService";

export default {
  name: "Search",
  components: {},
  data() {
    return {
      userId: "14f1317e",
      dataProducts: [],
    };
  },
  methods: {
    getData: function (userId) {
      getUserTransactionsService
        .getUserTransactions(userId)
        .then((dataProducts) => (this.dataProducts = dataProducts));
    },
  },
};
</script>