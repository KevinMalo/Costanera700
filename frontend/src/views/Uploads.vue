<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-form method="post" :action="`${url}${date}`">
        <v-col cols="12">
          <v-container class="grey lighten-5">
            <v-row no-gutters>
              <v-col>
                <v-layout row wrap>
                  <v-text-field
                    label="Select a date"
                    readonly
                    :value="date"
                  ></v-text-field>
                </v-layout>
                <v-date-picker v-model="date" no-title></v-date-picker>
                <div class="top">
                  <v-btn color="primary" type="submit"> Submit </v-btn>
                </div>
              </v-col>
            </v-row>
          </v-container>
        </v-col>
      </v-form>
    </v-row>
  </v-container>
</template>  

<script>
export default {
  name: "Uploads",
  data() {
    return {
      fromDateVal: null,
      date: new Date().toISOString().substr(0, 10),
      url: "http://localhost:9000/uploads/",
    };
  },
};
</script>

<style>
.top {
  margin-top: 8px;
}
</style>