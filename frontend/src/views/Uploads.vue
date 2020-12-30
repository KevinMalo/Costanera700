<template>
  <v-container>
    <v-form method="post" :action="`${url}${date}`">
      <v-col cols="6">
        <v-container class="grey lighten-5">
          <v-row no-gutters>
            <v-col>
              <v-card class="pa-2" outlined tile>
                <v-text-field label="Date" v-model="date"></v-text-field>
              </v-card>

              <v-layout row wrap>
                <!-- other code -->
                <v-text-field
                  label="Date"
                  readonly
                  :value="fromDateDisp"
                ></v-text-field>
              </v-layout>

              <v-date-picker
                locale="en-in"
                v-model="fromDateVal"
                no-title
              ></v-date-picker>

              <v-btn color="primary" text type="submit"> Submit </v-btn>
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
  data() {
    return {
      fromDateVal: null,
      date: Math.floor(Date.now() / 1000),
      url: "http://localhost:9000/uploads/",
    };
  },
  computed: {
      fromDateDisp() {
        return this.fromDateVal;
      },
    },
};
</script>

<style scope>
</style>