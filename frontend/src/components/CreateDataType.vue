<template>
  <div class="container">
    <h1 class="title">Create New Data Type</h1>
    <form @submit.prevent="submitForm">
      <!-- Поле для имени -->
      <div class="form-group">
        <label for="name">Name:</label>
        <input
            type="text"
            id="name"
            v-model="dataType.name"
            placeholder="Enter unique name"
            required
        />
      </div>

      <!-- Поле для ManyName -->
      <div class="form-group">
        <label for="manyName">Many Name:</label>
        <input
            type="text"
            id="manyName"
            v-model="dataType.manyName"
            placeholder="Enter many name"
        />
      </div>

      <!-- Переключатель для IsTabulate -->
      <div class="form-group switch-group">
        <label for="isTabulate">Is Tabulate:</label>
        <label class="switch">
          <input
              type="checkbox"
              id="isTabulate"
              v-model="dataType.isTabulate"
          />
          <span class="slider"></span>
        </label>
      </div>

      <!-- Кнопка создания нового типа данных -->
      <button type="submit">Create Data Type</button>
    </form>
  </div>
</template>

<script>
import {ref} from "vue";
import {CreateDataType} from "../../wailsjs/go/backend/App.js"; // Предполагаемая функция для API

export default {
  name: "CreateDataType",
  setup() {
    const dataType = ref({
      name: "",
      manyName: "",
      isTabulate: false,
    });

    const submitForm = async () => {
      try {
        await CreateDataType({
          name: dataType.value.name,
          many_name: dataType.value.manyName,
          is_tabulate: dataType.value.isTabulate,
        });
        alert(`Data type "${dataType.value.name}" created successfully!`);
        dataType.value = {name: "", manyName: "", isTabulate: false};
      } catch (error) {
        console.error("Failed to create data type:", error);
        alert("Error creating data type. Please try again.");
      }
    };

    return {
      dataType,
      submitForm,
    };
  },
};
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 600px;
  margin: 0 auto;
}

.title {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  font-weight: bold;
  display: block;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 10px;
  margin: 0;
  border: none;
  border-radius: 4px;
  background: #1e1e1e;
  color: #fff;
}

.switch-group {
  display: flex;
  align-items: center;
}

.switch-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
}

.label-text {
  font-weight: bold;
  color: #fff;
}

.switch {
  position: relative;
  display: inline-block;
  width: 40px;
  height: 20px;
  margin-left: 10px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: 0.4s;
  border-radius: 20px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 14px;
  width: 14px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: #4caf50;
}

input:checked + .slider:before {
  transform: translateX(20px);
}

button {
  width: 100%;
  padding: 10px;
  background: #4caf50;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s;
}

button:hover {
  background: #45a049;
}
</style>
