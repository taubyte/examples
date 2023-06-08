import { createRouter, createWebHistory } from "vue-router";
import Login from "../components/Login.vue";
import GenderSelection from "../components/GenderSelection.vue";
import Spinner from "../components/Spinner.vue";
import Avatar from "../components/Avatar.vue";

const routes = [
  { path: "/", component: Login },
  { path: "/gender-selection", component: GenderSelection },
  { path: "/spinner", component: Spinner },
  { path: "/avatar", component: Avatar },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
