<template>
  <div ref="editor" class="code-editor"></div>
</template>

<script>
import { ref, onMounted, watch } from "vue";
import CodeMirror from "codemirror";
import "codemirror/lib/codemirror.css";
import "codemirror/theme/dracula.css";
import "codemirror/mode/yaml/yaml";

export default {
  props: {
    modelValue: { type: String, default: "" },
  },
  emits: ["update:modelValue"],
  setup(props, { emit }) {
    const editor = ref(null);
    let cmInstance = null;

    onMounted(() => {
      cmInstance = new CodeMirror(editor.value, {
        mode: "yaml",
        theme: "dracula",
        lineNumbers: true,
      });

      cmInstance.on("change", () => {
        emit("update:modelValue", cmInstance.getValue());
      });

      cmInstance.setValue(props.modelValue);
    });

    watch(() => props.modelValue, (newValue) => {
      if (cmInstance && cmInstance.getValue() !== newValue) {
        cmInstance.setValue(newValue);
      }
    });

    return { editor };
  },
};
</script>

<style scoped>
.code-editor {
  width: 100%;
  padding: 10px;
  margin: 0;
  border: none;
  border-radius: 4px;
  background: #1e1e1e;
  color: #fff;
}
</style>