<template>
  <div class="app">
    <header>
      <h1>Docker Management</h1>
      <nav>
        <button @click="setPage('HomePage')">Home</button>
        <button @click="setPage('CreateContainerPage')">Create New Container</button>
      </nav>
    </header>
    <main>
      <!-- Динамическое отображение компонентов с анимацией -->
      <transition name="fade" mode="out-in">
        <component :is="currentPage" />
      </transition>
    </main>
  </div>
</template>

<script>
import {ref} from "vue";
import HomePage from "./components/Main.vue";
import CreateContainerPage from "./components/CreateContainer.vue";

export default {
  components: {
    HomePage,
    CreateContainerPage,
  },
  setup() {
    // Текущее состояние страницы
    const currentPage = ref("HomePage"); // Указываем имя компонента

    // Функция для смены страницы
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
/* Стили для Header */
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

/* Стили для Main */
main {
  padding: 20px;
}

/* Анимация переходов */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
