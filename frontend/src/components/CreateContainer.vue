<template>
  <div class="container">
    <h1 class="title">Create Docker Container</h1>
    <form @submit.prevent="submitForm">
      <!-- Поле для имени -->
      <div class="form-group">
        <label for="name">Name:</label>
        <input
            type="text"
            id="name"
            v-model="container.name"
            placeholder="Enter container name"
            required
        />
      </div>

      <!-- Поле для тела -->
      <!--      <div class="form-group">-->
      <!--        <label for="body">Body:</label>-->
      <!--        <textarea-->
      <!--            id="body"-->
      <!--            v-model="container.body"-->
      <!--            placeholder="Enter container description"-->
      <!--            required-->
      <!--        ></textarea>-->
      <!--      </div>-->

      <div class="form-group">
        <label for="body">Body:</label>
        <div ref="editor" class="code-editor"></div>
      </div>

      <!-- Поле для добавления текстов с тегами -->
      <div class="form-group">
        <label>Add Text with Tag:</label>
        <div class="text-with-tag">

          <select v-model="selectedTag">
            <option disabled value="">Select a tag</option>
            <option v-for="tag in availableTags" :key="tag.id" :value="tag.id">
              {{ tag.name }}
            </option>
          </select>
          <input
              type="text"
              v-model="newText"
              placeholder="Enter text"
          />
          <button
              type="button"
              class="add-button"
              @click="addTextWithTag"
          >
            +
          </button>
        </div>
        <ul class="text-tag-list">
          <transition-group name="list" tag="div">
            <li v-for="(item, index) in container.textWithTags" :key="item.text + item.tag">
              {{ item.tag }} - {{ item.text }}
              <button
                  type="button"
                  class="delete-button"
                  @click="removeTextWithTag(index)"
              >
                ✕
              </button>
            </li>
          </transition-group>
        </ul>
      </div>

      <!-- Кнопка создания контейнера -->
      <button id="create-btn" type="submit">Create Container</button>
    </form>
  </div>
</template>

<script>
import {onMounted, ref} from "vue";
import {GetAllTypes, CreateContainer} from "../../wailsjs/go/backend/App.js";

import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import 'codemirror/theme/dracula.css';
import 'codemirror/mode/yaml/yaml';

export default {
  name: "CreateContainer",
  setup() {
    const container = ref({
      name: "",
      body: "",
      textWithTags: [], // Массив для текстов с тегами
    });

    const availableTags = ref([]);

    const newText = ref("");
    const selectedTag = ref("");

    const editor = ref(null);

    const addTextWithTag = () => {
      if (!newText.value.trim() || !selectedTag.value) {
        alert("Please enter text and select a tag.");
        return;
      }

      const selectedTagName = availableTags.value.find(
          (tag) => tag.id === parseInt(selectedTag.value)
      )?.name;

      container.value.textWithTags.push({
        text: newText.value,
        tag: selectedTagName,
        tagId: selectedTag.value,
      });

      newText.value = "";
      selectedTag.value = "";
    };

    const removeTextWithTag = (index) => {
      container.value.textWithTags.splice(index, 1);
    };

    const submitForm = async () => {
      try {
        const publicTypes = container.value.textWithTags.map((item) => ({
          type_id: parseInt(item.tagId), // Используем ID тега
          data: item.text,
        }));

        const requestData = {
          name: container.value.name,
          body: container.value.body,
          public_types: publicTypes
        }

        await CreateContainer(requestData)
        alert(`Container "${container.value.name}" created successfully!`);

        container.value = {name: "", body: "", textWithTags: []};
      } catch (error) {
        console.error("Failed to create container:", error);
        alert("Error creating container. Please try again.");
      }
    };

    const fetchTags = async () => {
      try {
        const tags = await GetAllTypes();
        availableTags.value = tags.map((tag) => ({
          id: tag.ID,
          name: tag.Name,
        }));
      } catch (error) {
        console.error("Failed to fetch tags: ", error)
      }
    }

    onMounted(() => {
      fetchTags();
      new CodeMirror(editor.value, {
        mode: 'YAML',
        theme: 'dracula',
        lineNumbers: true,
      });

      editor.on('change', () => {
        container.value.body = editor.getValue();
      });
    })

    return {
      container,
      availableTags,
      newText,
      selectedTag,
      addTextWithTag,
      removeTextWithTag,
      submitForm,
      editor,
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
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input,
textarea,
select,
.code-editor {
  width: 100%;
  padding: 10px;
  margin: 0;
  border: none;
  border-radius: 4px;
  background: #1e1e1e;
  color: #fff;
}

textarea {
  height: 30vh;
  resize: none;
}

.text-with-tag {
  display: flex;
  gap: 10px;
}

.text-with-tag input {
  flex: 2;
}

.text-with-tag select {
  flex: 1;
  color: #1b1513;
}

.add-button {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #4caf50;
  color: #fff;
  border: none;
  font-size: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.3s;
}

.add-button:hover {
  background: #45a049;
}

.text-tag-list {
  list-style-type: none;
  padding: 0;
  margin: 10px 0 0;
}

.text-tag-list li {
  margin-bottom: 5px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.delete-button {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #e74c3c;
  color: #fff;
  border: none;
  padding: 5px 10px;
  cursor: pointer;
  font-size: 14px;
  transition: background 0.3s;
}

.delete-button:hover {
  background: #c0392b;
}

/* Плавная анимация добавления */
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from {
  opacity: 0;
  transform: translateY(20px); /* Появление снизу */
}

.list-leave-to {
  opacity: 0;
  transform: translateY(-20px); /* Уход вверх */
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

#create-btn {
  width: 40vh;  /* Убираем 100% ширину */
  padding: 8px 16px;  /* Меньше отступы */
  background: #4caf50;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s;
  margin: 0 auto;  /* Центрируем кнопку */
  display: block;  /* Включаем блочный режим, чтобы использовать margin: auto */
}

#create-btn:hover {
  background: #45a049;
}
</style>
