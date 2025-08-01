<template>
  <el-form ref="loginForm" :model="loginForm" label-width="80px">
    <el-form-item label="用户名" prop="username">
      <el-input v-model="loginForm.username"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input type="password" v-model="loginForm.password"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="handleLogin">登录</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import request from "@/utils/request";

const loginForm = ref({
  username: "",
  password: "",
});
const router = useRouter();

const handleLogin = async () => {
  try {
    const res = await request.post("/auth/login", loginForm.value);
    localStorage.setItem("token", res.data.token);
    router.push("/");
  } catch (error) {
    console.error(error);
  }
};
</script>

<style scoped></style>
