<template>
  <v-container>
    <v-form method="post" :action="`${url}${date}`">
      <v-col cols="6">
        <v-container class="grey lighten-5">
          <v-row no-gutters>
            <v-col>
              <v-layout row wrap>
                <!-- other code -->
                <v-text-field
                  label="Select a date"
                  readonly
                  :value="date"
                ></v-text-field>
              </v-layout>

              <v-date-picker v-model="date" no-title></v-date-picker>

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
      date: new Date().toISOString().substr(0, 10),
      url: "http://localhost:9000/uploads/",
    };
  }
};
</script>

<style scope>
</style>