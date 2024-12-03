<template>
  <div class="container">
    <h1 class="title">Docker Containers</h1>

    <button @click="generateDockerCompose" class="generate-button">Generate docker-compose.yml</button>

    <div v-for="container in containers" :key="container.id" class="card">
      <div class="info">
        <h2>{{ container.name }}</h2>
        <p>{{ container.description }}</p>
        <p>Status: <span :class="container.status">{{ container.status }}</span></p>
      </div>
      <div class="actions">
        <button @click="toggleContainer(container)">
          {{ container.status === 'active' ? 'Deactivate' : 'Activate' }}
        </button>
        <button @click="editContainer(container)">Edit</button>
        <button @click="deleteContainer(container)">Delete</button>
      </div>
    </div>
  </div>
</template>

<script>
import {onMounted, ref} from "vue";
import {GetAllContainers, DeleteContainer, SwitchContainerActive, GenerateDockerCompose} from "../../wailsjs/go/backend/App.js";

export default {
  name: "DockerContainers",
  setup() {
    const containers = ref([]);

    const fetchContainers = async () => {
      try {
        const response = await GetAllContainers();
        console.log(response)
        containers.value = response.map((container) => ({
          id: container.ID,
          name: container.Name,
          body: "Test",
          status: container.IsActive ? "active" : "disabled",
        }))
      } catch (error) {
        console.error("Error with fetch containers: ", error)
      }
    }

    const generateDockerCompose = async () => {
      await GenerateDockerCompose()
      alert('Successful generationk')
    }

    const toggleContainer = async (container) => {
      const containerStatus = await SwitchContainerActive(container.id);
      container.status = containerStatus ? "active" : "disabled";
    };

    const editContainer = (container) => {
      alert(`Edit functionality for ${container.name}`);
    };

    const deleteContainer = async (container) => {
      try {
        await DeleteContainer(container.id)
        alert(`Container deleted successful ${container.name}`)
        await fetchContainers()
      } catch (error) {
        console.error("Error with delete container: ", error)
      }
    }

    onMounted(() => {
      fetchContainers()
    })

    return {
      containers,
      toggleContainer,
      editContainer,
      deleteContainer,
      generateDockerCompose,
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

.info .active {
  color: #4caf50;
}

.info .disabled {
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

.generate-button {
  background: #dc6b26;
  color: #fff;
  border: none;
  padding: 10px 15px;
  margin: 10px 0;
  border-radius: 4px;
  cursor: pointer;
  display: block;
  transition: background 0.3s;
}

.generate-button:hover {
  background: #aa531e;
}

</style>
