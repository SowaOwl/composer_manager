<template>
  <div class="container">
    <h1 class="title">Docker Containers</h1>
    <div v-for="container in containers" :key="container.id" class="card">
      <div class="info">
        <h2>{{ container.name }}</h2>
        <p>{{ container.description }}</p>
        <p>Status: <span :class="container.status">{{ container.status }}</span></p>
      </div>
      <div class="actions">
        <button @click="toggleContainer(container)">
          {{ container.status === 'running' ? 'Stop' : 'Start' }}
        </button>
        <button @click="editContainer(container)">Edit</button>
      </div>
    </div>
  </div>
</template>

<script>
import {onMounted, ref} from "vue";
import {GetAllContainers} from "../../wailsjs/go/backend/App.js";

export default {
  name: "DockerContainers",
  setup() {
    const containers = ref([]);

    const fetchContainers = async () => {
      try {
        const response = await GetAllContainers();
        containers.value = response.map((container) => ({
          id: container.ID,
          name: container.Name,
          body: "Test",
          status: "stopped",
        }))
      } catch (error) {
        console.error("Error with fetch containers: ", error)
      }
    }

    const toggleContainer = (container) => {
      container.status = container.status === "running" ? "stopped" : "running";
    };

    const editContainer = (container) => {
      alert(`Edit functionality for ${container.name}`);
    };

    onMounted(() => {
      fetchContainers()
    })

    return {
      containers,
      toggleContainer,
      editContainer,
    };
  },
};
</script>

<style scoped>
body {
  margin: 0;
  font-family: Arial, sans-serif;
  background-color: #121212;
  color: #fff;
}

.container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.title {
  text-align: center;
  margin-bottom: 20px;
}

.card {
  background: #1e1e1e;
  padding: 20px;
  margin-bottom: 15px;
  border-radius: 8px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info h2 {
  margin: 0 0 5px;
  font-size: 18px;
}

.info p {
  margin: 2px 0;
}

.info .running {
  color: #4caf50;
}

.info .stopped {
  color: #f44336;
}

.actions button {
  background: #333;
  color: #fff;
  border: none;
  padding: 8px 12px;
  margin-left: 10px;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s;
}

.actions button:hover {
  background: #444;
}
</style>
