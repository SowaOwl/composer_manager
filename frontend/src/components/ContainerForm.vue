<template>
  <form @submit.prevent="submitForm">
    <div class="form-group">
      <label for="name">Name:</label>
      <input type="text" id="name" v-model="container.name" placeholder="Enter container name" required />
    </div>

    <CodeEditor v-model="container.body" />

    <TagInput :tags="availableTags" v-model="container.textWithTags" class="tag-input"/>

    <button type="submit">{{ submitButtonText }}</button>
  </form>
</template>

<script>
import TagInput from "./TagInput.vue";
import CodeEditor from "./CodeEditor.vue";
export default {
  components: {TagInput, CodeEditor},
  props: {
    container: { type: Object, required: true },
    availableTags: { type: Array, required: true },
    submitButtonText: { type: String, default: "Create Container" },
    onSubmit: { type: Function, required: true },
  },
  setup(props) {
    const submitForm = () => props.onSubmit();
    return { submitForm };
  },
};
</script>

<style scoped>
.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

input,
select {
  width: 100%;
  padding: 10px;
  margin: 0;
  border: none;
  border-radius: 4px;
  background: #1e1e1e;
  color: #fff;
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

.tag-input {
  padding-top: 20px;
}
</style>
