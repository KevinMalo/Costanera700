<template>
  <v-container>
    <v-col cols="12">
      <v-container class="grey lighten-5">
        <v-row no-gutters>
          <v-col>
            <v-card class="pa-2" outlined tile>
              <v-text-field label="USER ID" v-model="userId"></v-text-field>
            </v-card>
            <div class="top">
              <v-btn
                color="primary"
                type="submit"
                v-on:click="
                  getProducts(userId);
                  getBuyersIP(userId);
                  getBestSellers();
                "
              >
                Submit
              </v-btn>
            </div>
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
                <th class="text-left">Buyer id</th>
                <th class="text-left">Name</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, i) in dataProducts" :key="i">
                <td>{{ item.product[0].product_id }}</td>
                <td>{{ item.product[0].name }}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </div>

    <div v-if="dataBuyers.length > 0">
      <h1>Other buyer with same IP</h1>
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
              <tr v-for="(b, i) in dataBuyers" :key="i">
                <template v-if="b.buyers[0] != undefined">
                  <td>{{ b.buyers[0].id }}</td>
                  <td>{{ b.buyers[0].name }}</td>
                </template>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-col>
    </div>

    <div v-if="dataBestSellers.length > 0">
      <h1>Best Sellers</h1>
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
              <tr v-for="(bs, i) in dataBestSellers" :key="i">
                <template v-if="bs.product != undefined">
                  <td>{{ bs.product[0].product_id }}</td>
                  <td>{{ bs.product[0].name }}</td>
                </template>
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
import getBuyersIpsService from "@/services/getBuyersIpsService";
import getBestSellersService from "@/services/getBestSellersService";

export default {
  name: "Search",
  components: {},
  data() {
    return {
      userId: "9cc96330",
      dataProducts: [],
      dataBuyers: [],
      dataBestSellers: [],
    };
  },
  methods: {
    getProducts: function (userId) {
      getUserTransactionsService
        .getUserTransactions(userId)
        .then((dataProducts) => (this.dataProducts = dataProducts));
    },
    getBuyersIP: function (userId) {
      getBuyersIpsService
        .getBuyersIps(userId)
        .then((dataBuyers) => (this.dataBuyers = dataBuyers));
    },
    getBestSellers: function () {
      getBestSellersService
        .getBestSellers()
        .then((dataBestSellers) => (this.dataBestSellers = dataBestSellers));
    },
  },
};
</script>

<style>
.top {
  margin-top: 8px;
}
</style>