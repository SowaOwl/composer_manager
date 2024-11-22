<template>
  <div class="app">
    <header>
      <h1>Docker Management</h1>
      <nav>
        <button @click="setPage('home')">Home</button>
        <button @click="setPage('create')">Create New Container</button>
      </nav>
    </header>
    <main>
      <!-- Динамическое отображение компонентов -->
      <HomePage v-if="currentPage === 'home'" />
      <CreateContainerPage v-if="currentPage === 'create'" />
    </main>
  </div>
</template>

<script>
import { ref } from "vue";
import HomePage from "./components/Main.vue";
import CreateContainerPage from "./components/CreateContainer.vue";

export default {
  components: {
    HomePage,
    CreateContainerPage,
  },
  setup() {
    const currentPage = ref("home");

    const setPage = (page) => {
      currentPage.value = page;
    };

    return {
      currentPage,
      setPage,
    };
  },
};
</script>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background: #333;
  color: white;
}

header h1 {
  margin: 0;
}

nav button {
  background: #4caf50;
  color: white;
  border: none;
  padding: 10px 15px;
  margin-left: 10px;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s;
}

nav button:hover {
  background: #45a049;
}

main {
  padding: 20px;
}
</style>
