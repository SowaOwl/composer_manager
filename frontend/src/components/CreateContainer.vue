<template>
  <div class="container">
    <h1 class="title">Create Docker Container</h1>
    <ContainerForm :on-submit="submitForm" :available-tags="availableTags" :container="container"/>
  </div>
</template>

<script>
import {inject, onMounted, ref} from "vue";
import {GetAllTypes, CreateContainer} from "../../wailsjs/go/backend/App.js";

import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import 'codemirror/theme/dracula.css';
import 'codemirror/mode/yaml/yaml';
import ContainerForm from "./ContainerForm.vue";

export default {
  name: "CreateContainer",
  components: {ContainerForm},
  setup() {
    const container = ref({
      name: "",
      body: "",
      textWithTags: [], // Массив для текстов с тегами
    });

    const setPage = inject("setPage")

    const availableTags = ref([]);

    const newText = ref("");
    const selectedTag = ref("");

    const editor = ref(null);
    const cmEditor = ref(null);

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
          type_id: parseInt(item.tagId),
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
        cmEditor.value.setValue("");

        setPage("HomePage");
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
      cmEditor.value = new CodeMirror(editor.value, {
        mode: 'YAML',
        theme: 'dracula',
        lineNumbers: true,
      });

      cmEditor.value.on('change', () => {
        container.value.body = cmEditor.value.getValue();
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
</style>
