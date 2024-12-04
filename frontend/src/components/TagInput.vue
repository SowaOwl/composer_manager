<template>
  <div>
    <div class="text-with-tag">
      <select v-model="selectedTag">
        <option disabled value="">Select a tag</option>
        <option v-for="tag in tags" :key="tag.id" :value="tag.id">{{ tag.name }}</option>
      </select>
      <input type="text" v-model="newText" placeholder="Enter text" />
      <button type="button" @click="addTextWithTag" class="add-button">+</button>
    </div>
    <ul class="text-tag-list">
      <li v-for="(item, index) in modelValue" :key="item.text + item.tagId">
        {{ item.tag }} - {{ item.text }}
        <button type="button" @click="removeTextWithTag(index)" class="delete-button">✕</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref } from "vue";

export default {
  props: {
    tags: { type: Array, required: true },
    modelValue: { type: Array, default: () => [] },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const newText = ref("");
    const selectedTag = ref("");

    const addTextWithTag = () => {
      if (!newText.value.trim() || !selectedTag.value) return;
      const selectedTagName = props.tags.find(tag => tag.id === parseInt(selectedTag.value))?.name;
      emit("update:modelValue", [...props.modelValue, { text: newText.value, tag: selectedTagName, tagId: selectedTag.value }]);
      newText.value = "";
      selectedTag.value = "";
    };

    const removeTextWithTag = index => {
      const newValue = [...props.modelValue];
      newValue.splice(index, 1);
      emit("update:modelValue", newValue);
    };

    return { newText, selectedTag, addTextWithTag, removeTextWithTag };
  },
};
</script>
<style scoped>
input {
  width: 100%;
  padding: 10px;
  margin: 0;
  border: none;
  border-radius: 4px;
  background: #1e1e1e;
  color: #fff;
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


</style>
