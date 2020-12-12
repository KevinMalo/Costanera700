import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Uploads from "../views/Uploads.vue";
import Buyers from "../views/Buyers.vue";
import Search from "../views/Search.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/uploads",
    name: "Uploads",
    component: Uploads
  },
  {
    path: "/buyers",
    name: "Buyers",
    component: Buyers
  },
  {
    path: "/Search",
    name: "Search",
    component: Search
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
