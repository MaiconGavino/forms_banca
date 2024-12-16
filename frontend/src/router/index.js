import { createRouter, createWebHistory } from "vue-router";
import UserLogin from "@/components/UserLogin.vue";
import AdminDash from "@/components/AdminDash.vue";
import StartupEvaluation from "@/components/StartupEvaluation.vue";
import StartupResult from "@/components/StartupResult.vue";

const routes = [
  { path: "/", name: "UserLogin", component: UserLogin },
  { path: "/dash", name: "AdminDash", component: AdminDash },
  { path: "/evaluation", name: "StartupEvaluation", component: StartupEvaluation },
  { path: "/result", name: "StartupResult", component: StartupResult },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
